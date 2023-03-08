package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/utils"
)

// UserList returns the list of users of type user
func UserList(c *gin.Context) {
	users, err := db.GetAll()
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"users": users,
			"token": c.GetString("tokenString"),
		},
	)
}
