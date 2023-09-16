package main

import (
	"goclizidian/data"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// init router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Static("/assets", "./vol/assets")
	router.LoadHTMLGlob("vol/templates/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "Select a menu item hereabove to get started",
		})
	})
	router.GET("/size", dicsize)
	router.GET("/listpy/:py", listpy)
	router.GET("/listzi/:zi", listzi)
	router.GET("/zistring", buildZiString)
	router.GET("/zistring/:py/:zistr", addZi)
	router.GET("/quiz", quiz)
	router.Run(":8081")
}

func dicsize(c *gin.Context) {
	len, time := data.Dicsize()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "The dictionary presently contains " + len + " entries ; last updated on " + time,
	})
}

func listpy(c *gin.Context) {
	py := c.Param("py")
	// log.Println("résultat =", data.Listforpy(py))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": data.Listforpy(py),
	})
}

func listzi(c *gin.Context) {
	zi := c.Param("zi")
	// log.Println("résultat =", data.Listforzi(zi))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": data.Listforzi(zi),
	})
}

func buildZiString(c *gin.Context) {
	c.HTML(http.StatusOK, "zilist.html", gin.H{
		"zistring": template.HTML("<input type='text' id='zistring' size='40'>"),
	})
}

func addZi(c *gin.Context) {
	py := c.Param("py")
	st := c.Param("zistr")
	if st == "v" {
		st = ""
	}
	// log.Println("résultat =", data.Listforpy(py))
	c.HTML(http.StatusOK, "zilist.html", gin.H{
		"content":  data.GetZiList(py),
		"zistring": template.HTML("<input type='text' id='zistring' size='40' value='" + st + "'>"),
	})
}

func quiz(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz.html", gin.H{
		"content": data.GetQuizZi(),
	})
}
