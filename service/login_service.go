package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/go-kit/kit/log"
)

type LoginServiceInterface interface {
	GenerateToken(email string, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type loginService struct {
	secretKey string
	issure    string
	logger    log.Logger
}

func LoginServiceCreate(logger log.Logger) LoginServiceInterface {
	return &loginService{
		secretKey: "secret",
		issure:    "test",
		logger:logger,
	}
}

func (service loginService) GenerateToken(email string, role string) (string, error) {
	claims := &authCustomClaims{
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (service loginService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("Invalid token")
		}
		return []byte(service.secretKey), nil
	})
}
