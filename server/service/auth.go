package service

import (
	"blog-app/model/users"
	"blog-app/myerr"
	"blog-app/repo"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"os"
	"time"
)

const (
	AccessTokenTTL  = 5 * time.Minute
	RefreshTokenTTl = 24 * 7 * 3 * time.Hour
)

var (
	signingKey []byte
)

func init() {
	signingKey = []byte(os.Getenv("SIGNING_KEY"))
}

type Auth struct {
	r *repo.Repo
}

func NewAuthService(r *repo.Repo) *Auth {
	return &Auth{r}
}

type Claims struct {
	jwt.StandardClaims
	UID     int
	markant float64
}

func (a *Auth) CreateUser(user users.User) int {
	return a.r.CreateUser(user)
}

func (a *Auth) GetUser(username, pw string) users.User {
	return a.r.GetUser(username, pw)
}

func (a *Auth) GetUserById(id int) users.User {
	return a.r.GetUserById(id)
}

func (a *Auth) GenerateTokenPair(uid int) (string, string, error) {
	access := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		uid,
		markant(),
	})
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		uid,
		markant(),
	})

	at, err := access.SignedString(signingKey)
	if err != nil {
		return "", "", err
	}

	rt, err := refresh.SignedString(signingKey)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}

func (a *Auth) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(myerr.InvalidSigningMethod)
		}

		return signingKey, nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New(myerr.TokenExpired)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New(myerr.BadTokenClaimsType)
	}

	return claims.UID, nil
}

func markant() float64 {
	return rand.NormFloat64() * rand.NormFloat64()
}
