package main

import (
	"fmt"

	"github.com/nahK994/tiny-api/tiny"
)

type NameInput struct {
	Name string `json:"name"`
}

func LoggingMiddleware(next tiny.HandlerFunc) tiny.HandlerFunc {
	return func(c *tiny.Context) {
		fmt.Println("Before handler execution")
		next(c)
		fmt.Println("After handler execution")
	}
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
	}, LoggingMiddleware)

	app.GET("/students/:studentId/courses/:courseId/", func(c *tiny.Context) {
		var resp struct {
			StudentId int `json:"studentId"`
			CourseId  int `json:"courseId"`
		}

		resp.StudentId = c.PathParam["studentId"].(int)
		resp.CourseId = c.PathParam["courseId"].(int)
		c.JSON(200, resp)
	})

	subjectGroup := app.Group("/subjects")
	subjectGroup.Use(LoggingMiddleware)
	subjectGroup.GET("/:subjectId/", func(ctx *tiny.Context) {
		subjectId := ctx.PathParam["subjectId"].(int)
		response := map[string]any{
			"subjectId":   subjectId,
			"name":        "Sample Subject",
			"description": "This is a sample subject description.",
		}
		ctx.JSON(200, response)
	})
	subjectGroup.GET("/", func(ctx *tiny.Context) {
		courses := []string{"Math", "Science", "History"}
		ctx.JSON(200, courses)
	})

	if err := app.Run("127.0.0.1:8000"); err != nil {
		panic(err)
	}
}
