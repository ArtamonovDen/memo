package handler

import (
	"github.com/ArtamonovDen/memo/model"
	"github.com/gin-gonic/gin"
)

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type RouteConfig struct {
	R           *gin.Engine
	UserService model.UserService
}

// NewAccountHandler initializes account group handlers and inject them to app
// by direct reference to the gin Engine
func NewAccountHandler(c *RouteConfig) {

	h := &AccountHandler{UserService: c.UserService}

	g := c.R.Group("/api/account")

	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

}
