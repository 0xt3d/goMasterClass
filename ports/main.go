package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

var (
	logger          = log.New(os.Stdout, "", log.LstdFlags)
	dockerClient    *client.Client
	currentContainerID = os.Getenv("HOSTNAME")
	hostname         = os.Getenv("HOST_HOSTNAME")
	version          = os.Getenv("VERSION")
)

func init() {
	if os.Getenv("VERBOSE") != "" {
		logger.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	var err error
	dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	logger.Printf("Running as container ID: %s on external host %s", currentContainerID, hostname)
}

func checkPortProtocol(hostname, port string) string {
	for _, protocol := range []string{"http", "https"} {
		url := fmt.Sprintf("%s://%s:%s", protocol, hostname, port)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}
		req.Header.Set("x-whatsrunning-probe", "true")

		client := &http.Client{
			Timeout: 2 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			logger.Printf("Error checking %s: %v", url, err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return protocol
		}
	}
	return ""
}

func processContainer(container types.Container, hostname, currentContainerID string) *ContainerData {
	logger.Printf("Processing container %s", container.Names[0])

	if currentContainerID != "" && strings.HasPrefix(container.ID, currentContainerID) {
		logger.Printf("Skipping (current) container %s", container.Names[0])
		return nil
	}

	ports := []Port{}
	inspect, err := dockerClient.ContainerInspect(context.Background(), container.ID)
	if err != nil {
		logger.Printf("Failed to inspect container %s: %v", container.ID, err)
		return nil
	}

	if inspect.NetworkSettings != nil && inspect.NetworkSettings.Ports != nil {
		for name, value := range inspect.NetworkSettings.Ports {
			if !strings.HasSuffix(string(name), "/tcp") {
				continue
			}
			if len(value) == 0 {
				continue
			}
			candidatePorts := make(map[string]struct{})
			for _, v := range value {
				if v.HostPort != "" {
					candidatePorts[v.HostPort] = struct{}{}
				}
			}

			logger.Printf("Container %s has published ports %v", container.Names[0], candidatePorts)

			for port := range candidatePorts {
				protocol := checkPortProtocol(hostname, port)
				if protocol != "" {
					ports = append(ports, Port{Protocol: protocol, Port: port})
				}
			}
		}
	}

	if len(ports) > 0 {
		return &ContainerData{Name: container.Names[0], Ports: ports}
	}

	return nil
}

func processContainers(containers []types.Container, hostname, currentContainerID string) []*ContainerData {
	var wg sync.WaitGroup
	var mu sync.Mutex
	containerData := []*ContainerData{}

	for _, container := range containers {
		wg.Add(1)
		go func(container types.Container) {
			defer wg.Done()
			data := processContainer(container, hostname, currentContainerID)
			if data != nil {
				mu.Lock()
				containerData = append(containerData, data)
				mu.Unlock()
			}
		}(container)
	}

	wg.Wait()
	return containerData
}

func about(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Version: %s", version))
}

func listPorts(c *gin.Context) {
	if c.GetHeader("x-whatsrunning-probe") != "" {
		logger.Println("Ignoring probe request")
		c.String(http.StatusOK, "Alive")
		return
	}

	var containers, err := dockerClient.ContainerList(context.Background(), types.ContainerListOptions(All: true))
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to list containers: %v", err))
		return
	}

	sort.Slice(containers, func(i, j int) bool {
		return containers[i].Names[0] < containers[j].Names[0]
	})

	containerData := processContainers(containers, hostname, currentContainerID)

	type TemplateData struct {
		Containers []*ContainerData
		Hostname   string
		AppVersion string
	}

	data := TemplateData{
		Containers: containerData,
		Hostname:   hostname,
		AppVersion: version,
	}

	c.HTML(http.StatusOK, "template", gin.H{"data": data})
}

type ContainerData struct {
	Name  string
	Ports []Port
}

type Port struct {
	Protocol string
	Port     string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/about", about)
	r.GET("/", listPorts)

	port := os.Getenv("FLASK_PORT")
	if port == "" {
		port = "5000"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
