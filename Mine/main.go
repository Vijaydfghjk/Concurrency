package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {

	router := gin.Default()

	router.GET("/users", getusers)
	router.POST("/users", createuser)
	router.GET("/users/:id", Gets)

	// Print a message indicating that the server is running
	fmt.Println("Server is running on port :8080")

	router.Run(":8080")
	// http://localhost:8080/users working.
	//http://localhost:8080/users/1
	/*
	   sample json data
	   {
	       "id": 1,
	       "name": "Vijay",
	       "age": 36
	   }

	*/

}

func getusers(c *gin.Context) {

	c.JSON(http.StatusOK, users)
}

func createuser(c *gin.Context) {

	var newuser User

	if err := c.BindJSON(&newuser); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newuser)
	c.JSON(http.StatusCreated, newuser)
}

func Gets(c *gin.Context) {

	id := c.Param("id") // return string

	userid, err := strconv.Atoi(id) // id converted int

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"}) // 400 bad request
		return
	}

	var foundUser User // struct object

	for _, v := range users { // v access the slice

		if v.Id == userid {

			foundUser = v
			break
		}

	}

	if foundUser.Id != 0 {

		c.JSON(http.StatusOK, foundUser) // status ok
	} else {

		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // 404 not found
	}

}
