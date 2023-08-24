package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Al final del archivo esta la documentación completa de la estructura de una funcion controlador

func CreateUserHandler(serviceFunc func(c *gin.Context, db *gorm.DB), db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceFunc(c, db)
	}
}

func GetUserHandler(serviceFunc func(c *gin.Context, db *gorm.DB, userID uint64), db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parseUserID(c)
		serviceFunc(c, db, userID)
	}
}

func UpdateUserHandler(serviceFunc func(c *gin.Context, db *gorm.DB, userID uint64), db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parseUserID(c)
		serviceFunc(c, db, userID)
	}
}

func DeleteUserHandler(serviceFunc func(c *gin.Context, db *gorm.DB, userID uint64), db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := parseUserID(c)
		serviceFunc(c, db, userID)
	}
}

func parseUserID(c *gin.Context) uint64 {
	userIDStr := c.Param("id")
	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	return userID
}

/*
func GetUserHandler( 											  // Esta es la funcion controlador
	serviceFunc func(c *gin.Context, db *gorm.DB, userID uint64), // Esta es la funcion de servicio con el nombre de 'serviceFunc'
	db *gorm.DB) gin.HandlerFunc { 								  // Esta es la conexion a la base de datos
		             ^- esto es lo que retorna la funcion
	return func(c *gin.Context) { 								  // El cuerpo de la funcion controlador
		userID := parseUserID(c)							      // Se parsea el id del usuario (aqui van las validaciones)
		serviceFunc(c, db, userID) 								  // Se llama a la funcion de servicio 'serviceFunc'
	}
}
*/
