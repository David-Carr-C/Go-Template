package services

import (
	"net/http"

	"gitlab.com/nombre_usuario_o_grupo/nombre_proyecto/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserByID obtiene un usuario por su ID
func GetUserByID(c *gin.Context, db *gorm.DB, userID uint64) { // obtiene por parametro a gin, gorm y el id del usuario
	var user models.User              // se crea una variable de tipo User
	result := db.First(&user, userID) // se busca en la base de datos el usuario con el id especificado
	if result.Error != nil {          // si hay un error, se retorna el error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user}) // si no hay error, se retorna el usuario
}

// CreateUser crea un nuevo usuario
func CreateUser(c *gin.Context, db *gorm.DB) { // obtiene por parametro a gin y gorm
	var newUser models.User                            // se crea una variable de tipo User
	if err := c.ShouldBindJSON(&newUser); err != nil { // se obtiene el json enviado por el cliente y se lo asigna a la variable newUser
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // si hay un error, se retorna el error
		return
	}

	if err := newUser.SetPassword(newUser.Password); err != nil { // se cifra la contraseña del usuario
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // si hay un error, se retorna el error
		return
	}

	result := db.Create(&newUser) // se crea el usuario en la base de datos
	if result.Error != nil {      // si hay un error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	newUser.Password = "" // se elimina la contraseña del usuario para no mostrarla al cliente

	c.JSON(http.StatusOK, gin.H{"data": newUser}) // se retorna el usuario creado
}

// UpdateUser actualiza un usuario existente
func UpdateUser(c *gin.Context, db *gorm.DB, userID uint64) { // obtiene por parametro a gin, gorm y el id del usuario
	var user models.User              // se crea una variable de tipo User
	result := db.First(&user, userID) // se busca en la base de datos el usuario con el id especificado
	if result.Error != nil {          // si hay un error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	err := c.ShouldBindJSON(&user) // se obtiene el json enviado por el cliente y se lo asigna a la variable user
	if err != nil {                // si hay un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.SetPassword(user.Password); err != nil { // se cifra la contraseña del usuario
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // si hay un error, se retorna el error
		return
	}

	result = db.Save(&user)  // se actualiza el usuario en la base de datos
	if result.Error != nil { // si hay un error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	user.Password = "" // se elimina la contraseña del usuario para no mostrarla al cliente

	c.JSON(http.StatusOK, gin.H{"data": user}) // se retorna el usuario actualizado
}

// DeleteUser elimina un usuario por su ID
func DeleteUser(c *gin.Context, db *gorm.DB, userID uint64) { // obtiene por parametro a gin, gorm y el id del usuario
	result := db.Delete(&models.User{}, userID) // se elimina el usuario de la base de datos
	if result.Error != nil {                    // si hay un error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"}) // se retorna un mensaje de exito
}

// Para evitar tantas comparaciones de errores, usar try catch o una funcion
