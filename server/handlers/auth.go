package handlers

import (
	"blog-app/model/users"
	"blog-app/myerr"
	"crypto/sha512"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type signInInput struct {
	Username string `json:"nickname"`
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
	var sii signInInput
	err := c.BindJSON(&sii)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrInvalidCredentials)
		return
	}

	u := h.AuthService.GetUser(sii.Username, hashcode(sii.Password))
	if u.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusBadRequest)
		return
	}

	at, rt, err := h.AuthService.GenerateTokenPair(u.Id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrSignInError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "refresh-token",
		Value:   rt,
		Path:    "/",
		Expires: time.Now().Add(24 * 7 * 3 * time.Hour),
		//Secure:   true,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, map[string]string{
		"token": at,
	})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrPasswordDoesntMatch)
		return
	}

	h.AuthService.CreateUser(users.User{
		Name:         sui.Name,
		Surname:      sui.Surname,
		FatherName:   sui.FatherName,
		Nickname:     sui.Nickname,
		PasswordHash: hashcode(sui.Password),
	})
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}

func (h *handlers) refresh(c *gin.Context) {
	at := c.GetHeader("Authorization")
	rt, err := c.Cookie("refresh-token")
	if err != nil || at == "" {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrUnauthorized)
		return
	}

	_, err = h.AuthService.ParseToken(at)
	if err != nil {
		log.Println(err)
		if err.Error() == myerr.InvalidSigningMethod || err.Error() == myerr.BadTokenClaimsType {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		return
	}

	uid, err := h.AuthService.ParseToken(rt)
	if err != nil {
		log.Println(err)
		if err.Error() == myerr.InvalidSigningMethod || err.Error() == myerr.BadTokenClaimsType {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		if err.Error() == myerr.TokenExpired {
			c.AbortWithStatusJSON(http.StatusUnauthorized, myerr.RefreshTokenExpired)
		}
		return
	}

	newAT, newRT, err := h.AuthService.GenerateTokenPair(uid)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh-token",
		Value:    newRT,
		Path:     "/",
		Expires:  time.Now().Add(24 * 7 * 3 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, map[string]string{
		"token": newAT,
	})
}

func (h *handlers) signOut(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh-token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		Secure:   true,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, map[string]string{
		"message": "logged out",
	})
}

func hashcode(pw string) string {
	hash := sha512.New()
	hash.Write([]byte(pw))
	return hex.EncodeToString(hash.Sum(nil))
}
