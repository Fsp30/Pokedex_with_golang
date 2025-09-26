package handlers
import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
	
	"github.com/Fsp30/Pokedex_with_golang/models"
	"github.com/Fsp30/Pokedex_with_golang/services"
)

type UserHandler struct{
	Service *services.UserService
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	user := models.User{
		ID:        primitive.NewObjectID(),
		Username:      body.Name,
	}

	_, err := h.Service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create user successfully",
		"user":    user,
	})
}
