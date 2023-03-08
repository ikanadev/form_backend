package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/utils"
	"golang.org/x/crypto/bcrypt"
)

// NewUser registers a new user in the app
func NewUser(c *gin.Context) {
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.MakeRes(false, "Error recibiendo los datos"),
		)
		return
	}
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	user.IsActive = true
	user.Type = "user"
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Ci), bcrypt.DefaultCost)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "No se pudo encriptar el password")
		return
	}
	user.Password = string(hashedPass)
	existEmail := user.CheckEmail(db.DBCon)
	if existEmail {
		utils.MakeR(c, http.StatusBadRequest, "El correo ingresado ya esta siendo usado.")
		return
	}
	if err := user.Save(db.DBCon); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"user":  user,
			"token": c.GetString("tokenString"),
		},
	)
}
