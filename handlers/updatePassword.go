package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type passCredentials struct {
	Old string `json:"old"`
	New string `json:"new"`
}

// UpdatePassword updates the password of an user
func UpdatePassword(c *gin.Context) {
	userEmail, _ := c.Get("userEmail")
	var cred passCredentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Error al parsear JSON")
		return
	}
	user := db.User{
		Email: userEmail.(string),
	}
	if err := user.GetByEmail(userEmail.(string)); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Correo no encontrado")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Old)); err != nil {
		utils.MakeR(c, http.StatusBadRequest, "Password incorrecto, vuelva a intentar.")
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(cred.New), bcrypt.DefaultCost)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "No se pudo encriptar el password")
		return
	}
	if err := user.UpdatePassword(string(hashedPass)); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Error actualizando password")
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{},
	)
}
