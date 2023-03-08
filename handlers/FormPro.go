package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/structs"
	"github.com/vmkevv/form_backend/utils"
)

// GetFormPro get profesional form
func GetFormPro(c *gin.Context) {
	nro := c.Param("nro")
	form := db.FormPro{}
	user := db.User{}
	if err := form.GetByNro(nro); err != nil {
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

// NewFormPro saves new profesional form
func NewFormPro(c *gin.Context) {
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
	var formPro db.FormPro
	if err := c.BindJSON(&formPro); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	formPro.UpdatedAt = time.Now()
	formPro.CreatedAt = time.Now()
	formPro.UserID = user.ID
	if err := formPro.Save(); err != nil {
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

// UpdateFormPro updates the profesional form
func UpdateFormPro(c *gin.Context) {
	var formPro db.FormPro
	if err := c.BindJSON(&formPro); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formPro.Update(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formPro},
	)
}

// DeleteFormPro deleted profesional form
func DeleteFormPro(c *gin.Context) {
	var formPro db.FormPro
	if err := c.BindJSON(&formPro); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formPro.Delete(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formPro},
	)
}

// GetFormProByID get form by id
func GetFormProByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := db.FormPro{}
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

// GetProQuestions get profesional questions
func GetProQuestions(c *gin.Context) {
	formPro := db.FormPro{}
	res, err := formPro.GetQuestions()
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
