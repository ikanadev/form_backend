package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/structs"
	"github.com/vmkevv/form_backend/utils"
)

// Auth handles token validation
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Abort()
			c.JSON(http.StatusInternalServerError, utils.MakeRes(false, "No se reconoce el token en los headers."))
			return
		}
		claims := &structs.Claims{}
		tokenStruct, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return structs.JwtKey, nil
		})
		if err != nil {
			c.Abort()
			if strings.Contains(err.Error(), "token is expired by") {
				c.JSON(http.StatusBadRequest, utils.MakeRes(false, "Ha estado inactivo por más de 25 minutos, por favor ingrese de nuevo."))
				return
			}
			c.JSON(http.StatusInternalServerError, utils.MakeRes(false, err.Error()))
			return
		}
		if !tokenStruct.Valid {
			c.Abort()
			c.JSON(http.StatusInternalServerError, utils.MakeRes(false, "Token no valido."))
			return
		}
		tokenString, err := utils.GenToken(claims.ID, claims.Email)
		if err != nil {
			c.Abort()
			utils.MakeR(c, http.StatusInternalServerError, "Error generando token")
			return
		}
		c.Set("tokenString", tokenString)
		c.Set("userEmail", claims.Email)
	}
}
