package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email""`
	Age   int    `json:"age"`
}

var users = []user{
	{ID: "1", Name: "beanlam", Email: "beanlan@126.com", Age: 30},
	{ID: "2", Name: "yanru", Email: "yanru@126.com", Age: 29},
	{ID: "3", Name: "uti", Email: "uti@126.com", Age: 2},
	{ID: "4", Name: "uyun", Email: "uyun@126.com", Age: 1},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/user/:id", getUserByID)
	router.GET("/health", health)
	router.POST("/user", postUser)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	router.Run("0.0.0.0:" + httpPort)
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func postUser(context *gin.Context) {
	var newUser user
	if err := context.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	context.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(context *gin.Context) {
	id := context.Param("id")
	for _, u := range users {
		if u.ID == id {
			context.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found by id " + id})
}

func health(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"Status": "Ok!"})
}
