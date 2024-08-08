package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Maged-Zaki/gin-rest-api/models"
	"github.com/Maged-Zaki/gin-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var User models.User

	// Bind the JSON data to the user struct
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	User.Password, err = utils.HashPassword(User.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error hashing password: " + err.Error(),
		})
	}

	err = User.Save()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := utils.FormatResponse("Created Successfully", User)

	c.JSON(http.StatusCreated, response)
}

func Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error binding JSON: " + err.Error(),
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error validating user credentials: " + err.Error(),
		})
		return
	}

	secretKey := os.Getenv("JWT_SECRET")
	claims := map[string]any{
		"userId": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := utils.GenerateToken(secretKey, claims)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error generating token: " + err.Error(),
		})
		return
	}

	response := utils.FormatResponse("Logged in successfully", map[string]string{
		"token": token,
	})

	c.JSON(200, response)
}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("%s is not a valid id", idStr),
		})
		return
	}

	var user models.User
	user.ID = id
	err = user.Delete()

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error Deleting: " + err.Error(),
		})
		return
	}

	response := utils.FormatResponse("Deleted Successfully", nil)

	c.JSON(200, response)
}
