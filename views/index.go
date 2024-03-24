package views

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(
			c.Writer,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		http.Error(
			c.Writer,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}
