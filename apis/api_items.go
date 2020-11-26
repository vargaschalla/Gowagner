package apis

import (
	"net/http"

	"github.com/vargaschalla/Gowagner/models"

	"github.com/gin-gonic/gin"
)

func ItemsIndex(c *gin.Context) {
	s := models.Item{Title: "Sean", Notes: "nnn"}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hola we " + s.Title,
	})
}
