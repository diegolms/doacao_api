package doacao

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func GetAllDoacoes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
