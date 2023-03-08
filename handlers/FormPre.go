package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/structs"
	"github.com/vmkevv/form_backend/utils"
)

// GetFormPre gets prefas form
func GetFormPre(c *gin.Context) {
	nro := c.Param("nro")
	form := db.FormPre{}
	user := db.User{}
	if err := form.GetByNro(nro); err != nil {
		fmt.Print(err.Error())
		utils.MakeR(c, http.StatusBadRequest, "No existe el form nro "+nro)
		return
	}
	user.ID = form.UserID
	if err := user.GetByID(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"form": form,
			"user": user,
		},
	)
}

// NewFormPre saves new profesional form
func NewFormPre(c *gin.Context) {
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
	var formPre db.FormPre
	if err := c.BindJSON(&formPre); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	formPre.UpdatedAt = time.Now()
	formPre.CreatedAt = time.Now()
	formPre.UserID = user.ID
	if err := formPre.Save(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"token": c.GetString("tokenString"),
		},
	)
}

// UpdateFormPre updates the prefas form
func UpdateFormPre(c *gin.Context) {
	var formPre db.FormPre
	if err := c.BindJSON(&formPre); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formPre.Update(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formPre},
	)
}

// DeleteFormPre deleted prefas form
func DeleteFormPre(c *gin.Context) {
	var formPre db.FormPre
	if err := c.BindJSON(&formPre); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formPre.Delete(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formPre},
	)
}

// GetFormPreByID get form by id
func GetFormPreByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := db.FormPre{}
	user := db.User{}
	form.ID = id
	if err := form.GetByID(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = form.UserID
	if err := user.GetByID(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"form": form,
			"user": user,
		},
	)
}

// GetPreQuestions gets all pre questions
func GetPreQuestions(c *gin.Context) {
	formPre := db.FormPre{}
	res, err := formPre.GetQuestions()
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"res": res,
		},
	)
}
