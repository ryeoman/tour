package handler

import "github.com/gin-gonic/gin"

// Note: May need to abstract away gin library usage or just create adapter so we could swap gin with general http handler in golang.

type TourPlanHandler interface {
	GetByID(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
}
