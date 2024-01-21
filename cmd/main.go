package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
	"twitch-htmx-server/state"
	"twitch-htmx-server/templates"
)

func main() {
	err := templates.RegisterThemes("public/themes", true)
	templates.ActivateTheme("casper")
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("*", pageHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// TODO: implement template helper function "asset",
// Example use case: <link rel="preload" as="script" href="{{asset "built/casper.js"}}" />

func pageHandler(c echo.Context) error {
	fmt.Println(c.Request().URL.Path)
	path := strings.TrimPrefix(c.Request().URL.Path, "/")

	tplData := state.TplData{
		Site: state.Site{
			Locale: "TR-tr",
		},
	}

	rendered, err := templates.RenderPage(path, tplData)
	if err != nil {
		return c.HTML(http.StatusInternalServerError,
			"<html><head><title>Internal Server Error</title></head><body>"+err.Error()+"</body></html>")
	}

	return c.HTML(http.StatusOK, rendered)
}
