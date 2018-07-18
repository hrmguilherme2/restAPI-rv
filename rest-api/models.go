package main

import jwt "github.com/dgrijalva/jwt-go"

//Estruturas para o JSON
type _User struct {
	Id       int
	Name     string
	Gravatar string
}

type _Widgets struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float32 `json:"price"`
	Inventory int     `json:"inventory"`
	Melts     bool    `json:"melts"`
}

type User struct {
	Username string
	Password string
	Profile  UserProfile
}

// UserProfile
type UserProfile struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// UserClaims (Conjunto de claims que contem userprofile)
type UserClaims struct {
	Profile UserProfile `json:"profile"`
	jwt.StandardClaims
}
