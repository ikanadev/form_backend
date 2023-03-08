package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

// Login func to login
func Login(c *gin.Context) {
	var cred credentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Error al parsear JSON")
		return
	}
	user := db.User{
		Email: cred.Correo,
	}
	errUser := user.GetByEmail(cred.Correo)

	if errUser != nil {
		utils.MakeR(c, http.StatusBadRequest, errUser.Error())
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		utils.MakeR(c, http.StatusBadRequest, "Password incorrecto, vuelva a intentarlo.")
		return
	}
	tokenString, err := utils.GenToken(user.ID, user.Email)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Error al generar el token")
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
			"token": tokenString,
			"user":  user,
		},
	)
}
