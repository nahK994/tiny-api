# tiny-api

A minimal and lightweight web framework in Go — inspired by Gin, built for simplicity.


## Features

- Tiny and fast HTTP router
- Route handlers with method + path matching
- Request body parsing (JSON)
- Simple and expressive API like
- Build in documentation support


## Roadmap

* [x] Basic routing
* [x] JSON request parsing
* [x] JSON response rendering
* [ ] Swagger documentation support
* [ ] Middleware support
* [ ] Route groups
* [ ] Unit tests
* [ ] Static file serving
* [ ] Template engine support


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

	if err := app.Run("127.0.0.1:8000"); err != nil {
		panic(err)
	}
}
```


## Example Request

```bash
curl -X POST http://localhost:8080/hello \
     -H "Content-Type: application/json" \
     -d '{"name": "Shomi"}'
```

**Response:**

```json
{
  "message": "Hello, Shomi!"
}
```


## Folder Structure

```
tiny-api/
│
├── tiny/               # Core framework code
│   ├── router.go
│   ├── context.go
│   └── server.go
│
├── main.go             # Example app using tiny-api
└── README.md
```



## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.
