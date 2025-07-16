// this file return a variable db that helps other file to use db and interact with database
package config

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  //The _ means this package is imported for its side effects only (it registers the MySQL driver with GORM, so GORM can use MySQL). You don’t call anything from this package directly
//This import registers the MySQL driver with GORM, enabling GORM to connect to MySQL databases.

)
// pacakage level variable (global variable)
var (
	db *gorm.DB // it is a variable that holds a pointer to your database connection, created and managed by GORM
)

func Connect() { // helps us to open a connection to database
	//d is the database connection object.
	d, err := gorm.Open("mysql", "abhinav:Abhixan@120@/simplerest?charset=utf8&parseTime=True&loc=local")
	//The connection string
	// abhinav: MySQL username
	// Abhixan@120: MySQL password
	// simplerest: Database name
	// charset=utf8&parseTime=True&loc=local: Additional connection options
	if err != nil {
		panic(err)
	}
	db = d 
	
}
func GetDB() *gorm.DB {
	return db
}
