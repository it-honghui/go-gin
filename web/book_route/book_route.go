package book_route

import (
	"github.com/gin-gonic/gin"
	"go-gin/domain"
	"go-gin/domain/entity"
	"go-gin/service/book_service"
)

func getBooks(c *gin.Context) {
	books, err := book_service.FindBooks()
	if err != nil {
		domain.Panic(domain.ERROR, err.Error())
	}
	domain.Ok(c, books)
}

func getBook(c *gin.Context) {
	domain.Panic(domain.NOT_FOUND, "book")
	domain.Ok(c, entity.Book{Id: 1, Name: "西游记", Price: 1})
}

func Setup(e *gin.Engine) {
	g := e.Group("/book")
	{
		g.GET("/:id", getBook)
		g.GET("/", getBooks)
	}
}
