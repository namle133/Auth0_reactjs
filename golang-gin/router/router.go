package router

import (
	"CMS/config"
	"CMS/domain"
	"CMS/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

var contents []domain.News

func InitRouter() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve the frontend
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/news", NewsHandler)
		api.POST("/news/like/:newsID", LikeContent)
		api.GET("/news/:newsID", getNewsByID)
	}
	// Start the app
	router.Run(":3000")
}

// JokeHandler returns a list of jokes available (in memory)
func NewsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	us := config.UserService
	news, err := us.GetNews()
	contents = news
	if err != nil {
		log.Logrus().Error(err)
	}
	c.JSON(http.StatusOK, news)
}

func LikeContent(c *gin.Context) {
	us := config.UserService
	// Check joke ID is valid
	if newsid, err := strconv.Atoi(c.Param("newsID")); err == nil {
		// find joke and increment likes
		for i := 0; i < len(contents); i++ {
			if contents[i].ID == newsid {
				likes := contents[i].Likes + 1
				err = us.UpdateNewsLikes(newsid, likes)
				if err != nil {
					log.Logrus().Error(err)
				}

			}
		}
		c.JSON(http.StatusOK, &contents)
	} else {
		// the jokes ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// getJokesByID returns a single joke
func getNewsByID(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("newsID")); err == nil {
		for _, content := range contents {
			if content.ID == id {
				c.JSON(http.StatusOK, content)
			}
		}
		log.Logrus().Error("Joke not found")
	} else {
		// the jokes ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}
