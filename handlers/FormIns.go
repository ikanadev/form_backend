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

// NewFormIns saves new form
func NewFormIns(c *gin.Context) {
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
	var formIns db.FormIns
	err = c.BindJSON(&formIns)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	formIns.UpdatedAt = time.Now()
	formIns.CreatedAt = time.Now()
	formIns.UserID = user.ID
	if err := formIns.Save(); err != nil {
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

// UpdateFormIns updates student form
func UpdateFormIns(c *gin.Context) {
	var formIns db.FormIns
	if err := c.BindJSON(&formIns); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formIns.Update(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formIns},
	)
}

// GetFormIns gets form student
func GetFormIns(c *gin.Context) {
	nro := c.Param("nro")
	form := db.FormIns{}
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

// DeleteFormIns deletes the doc form
func DeleteFormIns(c *gin.Context) {
	var formIns db.FormIns
	if err := c.BindJSON(&formIns); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formIns.Delete(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formIns},
	)
}

// GetFormInsByID get form by id
func GetFormInsByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := db.FormIns{}
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

// GetInsQuestions get institutions questions
func GetInsQuestions(c *gin.Context) {
	formIns := db.FormIns{}
	res, err := formIns.GetQuestions()
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(c, http.StatusOK, gin.H{"res": res})
}
