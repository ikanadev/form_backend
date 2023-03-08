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

// GetFormDoc get a doc form
func GetFormDoc(c *gin.Context) {
	nro := c.Param("nro")
	form := db.FormDoc{}
	user := db.User{}
	if err := form.GetByNro(nro); err != nil {
		fmt.Println(err.Error())
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

// NewFormDoc saves new doc form
func NewFormDoc(c *gin.Context) {
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
	var formDoc db.FormDoc
	if err := c.BindJSON(&formDoc); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	formDoc.UpdatedAt = time.Now()
	formDoc.CreatedAt = time.Now()
	formDoc.UserID = user.ID
	if err := formDoc.Save(); err != nil {
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

// UpdateFormDoc update docente form
func UpdateFormDoc(c *gin.Context) {
	var formDoc db.FormDoc
	if err := c.BindJSON(&formDoc); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formDoc.Update(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formDoc},
	)
}

// DeleteFormDoc deletes the doc form
func DeleteFormDoc(c *gin.Context) {
	var formDoc db.FormDoc
	if err := c.BindJSON(&formDoc); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := formDoc.Delete(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{"data": formDoc},
	)
}

// GetFormDocByID get form by id
func GetFormDocByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	form := db.FormDoc{}
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

// GetDocQuestions get docentes questions
func GetDocQuestions(c *gin.Context) {
	formDoc := db.FormDoc{}
	res, err := formDoc.GetQuestions()
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
