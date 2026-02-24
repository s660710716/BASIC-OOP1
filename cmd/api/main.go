package main

import (
	"github.com/gin-gonic/gin"

	httpHandler "book-api/internal/delivery/http"
	"book-api/internal/infrastructure/memory"
	"book-api/internal/usecase"
)

func main() {
	repo := memory.NewBookRepo()
	uc := usecase.NewCreateBookUsecase(repo)
	h := httpHandler.NewHandler(uc)

	r := gin.Default()
	r.POST("/books", h.CreateBook)
	r.Run(":8080")
}