package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:username", func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello, "+username)
	})

	e.GET("/users/:username/followers", func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello, followers "+username)
	})

	e.GET("/users/:username/following", func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello, following "+username)
	})

	e.GET("/courses/:slug/episodes/:index", func(c echo.Context) error {
		slug := c.Param("slug")
		index := c.Param("index")
		return c.String(http.StatusOK, "Hello, following "+slug+", index "+index)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
