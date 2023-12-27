package main

import(
	"net/http"
	"github.com/labstack/echo/v4"
	n "github.com/riveong/echo-api-go/names"
)



func main(){
	
	e := echo.New()
	

	e.GET("/", func(c echo.Context) error{
		return c.JSON(http.StatusOK, map[string][]string{"names": n.persons()})
	})
	e.Logger.Fatal(e.Start(":1323"))
}