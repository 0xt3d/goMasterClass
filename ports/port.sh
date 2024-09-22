#!/bin/bash

# Get a list of all open TCP ports
open_ports=$(netstat -ntl | awk 'NR>2 {print $4}' | awk -F: '{print $NF}' | sort -u)

# Loop through each open port
for port in $open_ports; do
    # Get the container ID associated with the port
    container_id=$(docker ps --format "{{.ID}}:{{.Ports}}" | grep ":$port/" | awk -F: '{print $1}')

    # If a container ID is found, get the container name
    if [ -n "$container_id" ]; then
        container_name=$(docker inspect --format '{{.Name}}' $container_id | sed 's/\///')
        echo "Port $port is associated with container $container_name (ID: $container_id)"
    else
        echo "Port $port is not associated with any container"
    fi
done
