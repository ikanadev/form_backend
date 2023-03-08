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

// NewFormEst saves new form
func NewFormEst(c *gin.Context) {
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
	var formEst db.FormEst
	err = c.BindJSON(&formEst)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	formEst.UpdatedAt = time.Now()
	formEst.CreatedAt = time.Now()
	formEst.UserID = user.ID
	if err := formEst.Save(db.DBCon); err != nil {
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

// UpdateFormEst updates student form
func UpdateFormEst(c *gin.Context) {
	var formEst db.FormEst
	if err := c.BindJSON(&formEst); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formEst.Update(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formEst},
	)
}

// DeleteFormEst deletes the form
func DeleteFormEst(c *gin.Context) {
	var formEst db.FormEst
	if err := c.BindJSON(&formEst); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formEst.Delete(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formEst},
	)
}

// GetFormEst gets form student
func GetFormEst(c *gin.Context) {
	nro := c.Param("nro")
	form := db.FormEst{}
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

// GetFormEstByID get form by id
func GetFormEstByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := db.FormEst{}
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

// GetEstQuestions gets student questions
func GetEstQuestions(c *gin.Context) {
	formEst := db.FormEst{}
	res, err := formEst.GetQuestions()
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
