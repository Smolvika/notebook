package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createNote(c *gin.Context) {
	id, _ := c.Get(userCxt)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllNotes(c *gin.Context) {

}

func (h *Handler) getNoteById(c *gin.Context) {

}

func (h *Handler) updateNote(c *gin.Context) {

}

func (h *Handler) deleteNote(c *gin.Context) {

}
