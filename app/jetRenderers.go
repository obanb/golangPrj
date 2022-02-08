package app

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gin-gonic/gin"
	"log"
)

func renderPage(c *gin.Context, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(c.Writer, data, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func home(c *gin.Context) {
	err := renderPage(c, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func register(c *gin.Context) {
	err := renderPage(c, "register.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func login(c *gin.Context) {
	err := renderPage(c, "login.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func registerActiveEndpoint(c *gin.Context) {
	err := renderPage(c, "registerActiveEndpoint.jet", nil)
	if err != nil {
		log.Println(err)
	}
}
