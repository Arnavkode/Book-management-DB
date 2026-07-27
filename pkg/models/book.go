package models

import (
	"github.com/Arnavkode/Book-management-DB/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) *Book {
	var getBook Book
	db.Where("Id=?", id).Find(&getBook)
	return &getBook
}

func (b *Book) UpdateBook() *Book {
	db.Save(b)
	return b
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("Id=?", id).Delete(&book)
	return book
}
