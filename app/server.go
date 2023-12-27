package main

import(
	"net/http"
	"github.com/labstack/echo/v4"
)



func main(){
	
	e := echo.New()
	name := []string{"John", "Andreas", "Doe", "Peter", "Griffin", "Brian", "Stewie"}

	e.GET("/", func(c echo.Context) error{
		return c.JSON(http.StatusOK, map[string][]string{"names": name})
	})
	e.Logger.Fatal(e.Start(":1323"))
}