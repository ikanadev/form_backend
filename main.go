package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vmkevv/form_backend/db"
	"github.com/vmkevv/form_backend/handlers"
	"github.com/vmkevv/form_backend/middleware"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello World",
	})
}

func main() {
	db.DBCon = db.Connect()
	r := gin.Default()
	r.Use(middleware.WithCors())

	public := r.Group("/api")
	public.POST("/login", handlers.Login)
	public.POST("/student", handlers.SearchStudent)
	public.GET("/user/forms/:id", handlers.GetFormList)

	admin := r.Group("/api")
	admin.Use(middleware.Auth())
	admin.POST("/user", handlers.NewUser)
	admin.GET("/user", handlers.UserData)
	admin.PUT("/user", handlers.UpdatePassword)
	admin.GET("/users", handlers.UserList)
	admin.POST("/user/active", handlers.Active)
	admin.POST("/user/reset", handlers.Reset)
	admin.POST("/user/admin", handlers.ChangePermissions)

	admin.GET("/form/form-est/:nro", handlers.GetFormEst)
	admin.GET("/form-est/:id", handlers.GetFormEstByID)
	admin.POST("/form/form-est", handlers.NewFormEst)
	admin.PUT("/form/form-est", handlers.UpdateFormEst)
	admin.DELETE("/form/form-est", handlers.DeleteFormEst)
	admin.GET("/form-est", handlers.GetEstQuestions)

	admin.GET("/form/form-pro/:nro", handlers.GetFormPro)
	admin.GET("/form-pro/:id", handlers.GetFormProByID)
	admin.POST("/form/form-pro", handlers.NewFormPro)
	admin.PUT("/form/form-pro", handlers.UpdateFormPro)
	admin.DELETE("/form/form-pro", handlers.DeleteFormPro)
	admin.GET("/form-pro", handlers.GetProQuestions)

	admin.GET("/form/form-pre/:nro", handlers.GetFormPre)
	admin.GET("/form-pre/:id", handlers.GetFormPreByID)
	admin.POST("/form/form-pre", handlers.NewFormPre)
	admin.PUT("/form/form-pre", handlers.UpdateFormPre)
	admin.DELETE("/form/form-pre", handlers.DeleteFormPre)
	admin.GET("/form-pre", handlers.GetPreQuestions)

	admin.GET("/form/form-doc/:nro", handlers.GetFormDoc)
	admin.GET("/form-doc/:id", handlers.GetFormDocByID)
	admin.POST("/form/form-doc", handlers.NewFormDoc)
	admin.PUT("/form/form-doc", handlers.UpdateFormDoc)
	admin.DELETE("/form/form-doc", handlers.DeleteFormDoc)
	admin.GET("/form-doc", handlers.GetDocQuestions)

	admin.GET("/form/form-ins/:nro", handlers.GetFormIns)
	admin.GET("/form-ins/:id", handlers.GetFormInsByID)
	admin.POST("/form/form-ins", handlers.NewFormIns)
	admin.PUT("/form/form-ins", handlers.UpdateFormIns)
	admin.DELETE("/form/form-ins", handlers.DeleteFormIns)
	admin.GET("/form-ins", handlers.GetInsQuestions)

	r.GET("/", hello)
	r.Run(":4000")
}
