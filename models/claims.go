package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	AccountID    uint   `json:"accountID"`
	AccountName  string `json:"accountName"`
	AccountEmail string `json:"accountEmail"`
	Access       uint8  `json:"access"`
	jwt.StandardClaims
}
