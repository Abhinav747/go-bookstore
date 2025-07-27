package models

import(
	"github.com/jinzhu/gorm"
	"github.com/Abhinav747/go-bookstore/pkg/config" // contain db connection logic
)

var db *gorm.DB // global variable, 

type Book struct{
	gorm.Model //ID uint -> Auto-added unique number 
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}
//init(): Automatically runs when the models package is imported.
func init(){
	config.Connect() // Initializes DB connection 
	db = config.GetDB()// Retrieves the *gorm.DB object from config
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}