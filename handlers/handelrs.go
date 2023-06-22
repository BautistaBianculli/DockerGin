package handlers

import (
	"SatelitePacheco/routers"
	users "SatelitePacheco/routers/users"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*Handlers Pongo a escuchar al servidor para recibir los llamados HTTPs*/
func Handlers() {

	router := gin.Default() // Creo al servidor para escuchar

	router.POST("/login", routers.Login)
	userGroup := router.Group("users") // ! agregar el middlware
	{
		userGroup.GET("", 	users.GetUsers)
		userGroup.POST("", 	users.CreateUser)
	}  
	

	server := &http.Server{
		Addr:           ":8393",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
