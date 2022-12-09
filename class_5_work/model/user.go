package model

import "github.com/dgrijalva/jwt-go"

type Register struct {
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
	CheckQuestion string `form:"check question" json:"check question" binding:"required"`
	CheckAnswer   string `form:"check answer" json:"check answer" binding:"required"`
}
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type Change struct {
	NewPassword string `form:"new password" json:"new password" binding:"required"`
}
type Forget struct {
	Username string `form:"username" json:"username" binding:"required"`
}
type Question struct {
	Answer   string `form:"answer" json:"answer" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
}
type AddComment struct {
	Comment string `form:"comment" json:"comment" binding:"required"`
}
type Num struct {
	Num string `form:"num" json:"num" binding:"required"`
}
type Like struct {
	Like string `form:"like" json:"like" binding:"required"`
}
type CancelLike struct {
	CancelLike string `form:"cancel like" json:"cancel like" binding:"required"`
}
