package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"book-api/internal/usecase"
)

type Handler struct {
	createBook *usecase.CreateBookUsecase
}

func NewHandler(c *usecase.CreateBookUsecase) *Handler {
	return &Handler{createBook: c}
}

type createBookReq struct {
	Title string `json:"title"`
}

func (h *Handler) CreateBook(c *gin.Context) {
	var req createBookReq
	if err := c.ShouldBindJSON(&req); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

book, err := h.createBook.Execute(req.Title)
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

c.JSON(http.StatusCreated, gin.H{
	"id": book.ID,
	"title": book.Title,
	})
}