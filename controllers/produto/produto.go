package produto

import (
	"net/http"
	"github.com/gin-gonic/gin"
	model "doacao_api/models"

)

func GetAllProdutos(c *gin.Context) {w

	client := config.GetClient()
	database := client.Database("doacoes")

	produtoCollection := database.Collection("produto")

	var produtos []bson.M
	err := produtoCollection.Find(context.TODO(), bson.D{}).All(&produtos)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "produtos": produtos})

}
