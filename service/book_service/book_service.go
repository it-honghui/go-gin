package book_service

import (
	"go-gin/domain/entity"
)

func FindBooks() ([]entity.Book, error) {
	var books = []entity.Book{
		{Id: 1, Name: "西游记", Price: 1},
		{Id: 2, Name: "红楼梦", Price: 2},
	}
	return books, nil
}
