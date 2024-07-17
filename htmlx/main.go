package main

import(
	"github.com/gofiber/fiber/v2"
	"github.com/chasefleming/elem-go/attrs"
    "github.com/chasefleming/elem-go/styles"
    "github.com/chasefleming/elem-go/htmx"
)

bodyStyle := styles.Props{  
    styles.BackgroundColor: "#f4f4f4",
    styles.FontFamily:      "Arial, sans-serif",
    styles.Height:          "100vh",
    styles.Display:         "flex",
    styles.FlexDirection:   "column",
    styles.AlignItems:      "center",
    styles.JustifyContent:  "center",
}

buttonStyle := styles.Props{
    styles.Padding:       "10px 20px",
    styles.BackgroundColor: "#007BFF",
    styles.Color:         "#fff",
    styles.BorderColor:   "#007BFF",
    styles.BorderRadius:  "5px",
    styles.Margin:        "10px",
    styles.Cursor:        "pointer",
}

body := elem.Body(
    attrs.Props{
        attrs.Style: bodyStyle.ToInline(),
    },
    elem.H1(nil, elem.Text("Counter App")),
    elem.Div(attrs.Props{attrs.ID: "count"}, elem.Text("0")),
    elem.Button(
        attrs.Props{
            htmx.HXPost:   "/increment",
            htmx.HXTarget: "#count",
            attrs.Style:   buttonStyle.ToInline(),
        }, 
        elem.Text("+")
    ),
    elem.Button(
        attrs.Props{
            htmx.HXPost:   "/decrement",
            htmx.HXTarget: "#count",
            attrs.Style:   buttonStyle.ToInline(),
        }, 
        elem.Text("-")
    ),
)

func main()  {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	c.Type("html")
	// 	return c.SendString("Hello, World")
	// })
	app.Listen(":3030")

	app.Get("/", func(c *fiber.Ctx) error {
		head := elem.Head(nil, elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@1.9.6"}))
		//...[rest of the code]...
	})
}