package routers

import (
	"gitlab.com/nombre_usuario_o_grupo/nombre_proyecto/controllers"
	"gitlab.com/nombre_usuario_o_grupo/nombre_proyecto/services"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(dataBase *gorm.DB) *gin.Engine {
	// Se crea el framework Gin
	router := gin.Default()

	// Se define el frontend
	router.Use(static.Serve("/", static.LocalFile("./templates", true)))

	// Se define el grupo de rutas
	api := router.Group("/api")

	// Se define la ruta con su respectivo controlador y servicio
	api.POST("/create-user", controllers.CreateUserHandler(services.CreateUser, dataBase))
	api.GET("/user/:id", controllers.GetUserHandler(services.GetUserByID, dataBase))
	api.PUT("/user/:id", controllers.UpdateUserHandler(services.UpdateUser, dataBase))
	api.DELETE("/user/:id", controllers.DeleteUserHandler(services.DeleteUser, dataBase))

	// Se pueden definir rutas en archivos separados
	SendRouter(router)

	// Retorna Gin
	return router
}
