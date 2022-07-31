package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/admin"
)

type Users struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Username  string `json:"username"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Referrer  string `json:"referrer"`
}

// Define another GORM-backend model
type News struct {
	gorm.Model
	Title          string `json:"title"`
	TitleContent   string `json:"title_content"`
	Content        string `json:"content"`
	CreatorContent string `gorm:"foreignKey:UserName"`
}

func main() {

	dsn := "host=127.0.0.1 user=postgres password=Namle311 dbname=root port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Users{}, &News{})
	// Serve the frontend
	Admin := admin.New(&admin.AdminConfig{DB: db})

	// 	// Create resources from GORM-backend model
	Admin.AddResource(&Users{})
	Admin.AddResource(&News{})
	router := gin.Default()
	mux := http.NewServeMux()

	// 	// Mount admin to the mux
	Admin.MountTo("/admin", mux)
	// Serve the frontend
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.Any("/admin/*resources", gin.WrapH(mux))
	// Start the app
	router.Run(":3000")
}
