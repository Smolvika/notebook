package handler

import (
	"github.com/Smolvika/notebook.git/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}

}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		notes := api.Group("/notes")
		{
			notes.POST("/", h.createNote)
			notes.GET("/", h.getAllNotes)
			notes.GET("/:id", h.getNoteById)
			notes.PUT("/:id", h.updateNote)
			notes.DELETE("/:id", h.deleteNote)
		}
	}
	return router
}
