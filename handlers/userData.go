package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/structs"
	"github.com/vmkevv/form_backend/utils"
)

// UserData get user data based on token
func UserData(c *gin.Context) {
	tokenString, _ := c.Get("tokenString")
	claims := &structs.Claims{}
	_, err := jwt.ParseWithClaims(tokenString.(string), claims, func(tokenString *jwt.Token) (interface{}, error) {
		return structs.JwtKey, nil
	})
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	user := db.User{}
	if err := user.GetByEmail(claims.Email); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.IsActive == false {
		utils.MakeR(c, http.StatusBadRequest, "Lo sentimos, su cuenta ha sido inhabilitada por el administrador.")
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"user": user,
		},
	)
}
