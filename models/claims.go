package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	AccountName  string `json:"accountName"`
	AccountEmail string `json:"accountEmail"`
	jwt.StandardClaims
}