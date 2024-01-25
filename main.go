// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"time"

	"github.com/Narongrit089/go-gorm-db/db"
	"github.com/Narongrit089/go-gorm-db/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	database, err := db.ConnectDatabase(dbType, dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = database.AutoMigrate(&models.Item{})
	err = database.AutoMigrate(&models.Student{})
	err = database.AutoMigrate(&models.Subject{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	itemRepo := models.NewItemRepository(database)
	stydentRepo := models.NewStudentRepository(database)
	subRepo := models.NewSubjectRepository(database)

	r := gin.Default()

	// กำหนด cors (Cross-Origin Resource Sharing)
	r.Use(cors.New(cors.Config{
		// 3000 คือ port ที่ใช้งานใน frontend react
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/items", itemRepo.GetItems)
	r.POST("/items", itemRepo.PostItem)
	r.GET("/items/:id", itemRepo.GetItem)
	r.PUT("/items/:id", itemRepo.UpdateItem)
	r.DELETE("/items/:id", itemRepo.DeleteItem)

	r.GET("/students", stydentRepo.GetStudents)
	r.POST("/students", stydentRepo.PostStudent)
	r.GET("/students/:id", stydentRepo.GetStudent)
	r.PUT("/students/:id", stydentRepo.UpdateStudent)
	r.DELETE("/students/:id", stydentRepo.DeleteStudent)

	r.GET("/subjects", subRepo.GetSubjects)
	r.POST("/subjects", subRepo.PostSubject)
	r.GET("/subjects/:id", subRepo.GetSubject)
	r.PUT("/subjects/:id", subRepo.UpdateSubject)
	r.DELETE("/subjects/:id", subRepo.DeleteSubject)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Run the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Server is not running: %v", err)
	}
}
