package main

import (
	"log"

	"github.com/codelikesuraj/gdsc-challenge-day-nineteen-twenty/controllers"
	"github.com/codelikesuraj/gdsc-challenge-day-nineteen-twenty/middlewares"
	"github.com/codelikesuraj/gdsc-challenge-day-nineteen-twenty/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	PORT = "3000"
)

func main() {
	log.Fatalln(SetupRouter(SetupDB()).Run(":" + PORT))
}

func SetupDB() *gorm.DB {
	// initialize database
	db, err := gorm.Open(sqlite.Open("bookstoreapi.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to database:", err.Error())
	}

	// run migrations
	if err := db.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		log.Fatalln("Error running migrations:", err.Error())
	}

	return db
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	BookController := controllers.BookController{DB: db}
	UserController := controllers.UserController{DB: db}
	AuthMiddleware := middlewares.Authenticated{DB: db}

	r.POST("/register", UserController.Register)
	r.POST("/login", UserController.Login)
	r.POST("/refresh-token", UserController.RefreshToken)

	r.Group("", AuthMiddleware.Authenticate).
		GET("/books", BookController.GetAllBooks).
		GET("/books/:id", BookController.GetABook).
		POST("/books", BookController.CreateBook)

	return r
}

func RefreshDatabase(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.Book{}, &models.User{})
	if err != nil {
		log.Fatalln("Error clearing migrations:", err.Error())
	}

	// run migrations
	if err := db.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		log.Fatalln("Error running migrations:", err.Error())
	}
}
