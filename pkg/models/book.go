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