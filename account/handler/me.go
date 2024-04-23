package handler

import (
	"log"
	"net/http"

	"github.com/ArtamonovDen/memo/model"
	"github.com/ArtamonovDen/memo/model/apperrors"
	"github.com/gin-gonic/gin"
)

// Me handler calls services for getting user's details
func (h *AccountHandler) Me(c *gin.Context) {

	// Expect User to be added by middleware
	user, exists := c.Get("user")

	// double check
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	uid := user.(*model.User).UID

	// TODO: do not really understand, why we need call to service
	// if we have it from middleware
	u, err := h.UserService.Get(c, uid)

	if err != nil {
		log.Printf("Failed to find user: %v: %v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})

}
