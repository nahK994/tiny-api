# tiny-api

A minimal and lightweight web framework in Go — inspired by Gin, built for simplicity.


## Features

- Tiny and fast HTTP router
- Route handlers with method + path matching
- Request body parsing (JSON)
- Simple and expressive API like
- Build in documentation support


## Roadmap

* [x] JSON request parsing
* [x] JSON response rendering
* [x] Basic routing
* [x] Dynamic routing
* [x] Route groups
* [ ] Template engine support
* [ ] Middleware support
* [ ] Swagger documentation support
* [ ] Unit tests
* [ ] Static file serving


## Installation

```bash
go get github.com/nahK994/tiny-api
````

## Usage

```go
package main

import (
	"fmt"

	"github.com/nahK994/tiny-api/tiny"
)

type NameInput struct {
	Name string `json:"name"`
}

func main() {
	app := tiny.New()

	app.POST("/hello", func(c *tiny.Context) {
		var input NameInput
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, map[string]string{"error": "Invalid JSON"})
			return
		}
		c.JSON(200, fmt.Sprintf("Hello, %s!", input.Name))
	})

	app.GET("/students/:studentId/courses/:courseId/", func(c *tiny.Context) {
		var resp struct {
			StudentId int `json:"studentId"`
			CourseId  int `json:"courseId"`
		}

		resp.StudentId = c.PathParam["studentId"].(int)
		resp.CourseId = c.PathParam["courseId"].(int)
		c.JSON(200, resp)
	})

	if err := app.Run("127.0.0.1:8000"); err != nil {
		panic(err)
	}
}
```

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.
