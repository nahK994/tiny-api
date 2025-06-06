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
