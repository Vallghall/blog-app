package handlers

import (
	"blog-app/model/users"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signUpInput struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	FatherName      string `json:"father_name"`
	Nickname        string `json:"nickname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (h *handlers) signIn(c *gin.Context) {

}

func (h *handlers) signUp(c *gin.Context) {
	var sui signUpInput
	err := c.BindJSON(&sui)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidCredentials)
		return
	}

	if sui.Password != sui.ConfirmPassword {
		log.Println(ErrPasswordDoesntMatch)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrPasswordDoesntMatch)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(sui.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrBcryptError)
		return
	}

	h.AuthService.CreateUser(users.User{
		Name:         sui.Name,
		Surname:      sui.Surname,
		FatherName:   sui.FatherName,
		Nickname:     sui.Nickname,
		PasswordHash: hex.EncodeToString(hash),
	})
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}

func (h *handlers) refresh(c *gin.Context) {

}

func (h *handlers) signOut(c *gin.Context) {

}
