package produto

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	model "doacao_api/models"
	config "doacao_api/config"

)

func GetAllProdutos(c *gin.Context) {

	client := config.GetClient()
	database := client.Database("doacoes")

	produtoCollection := database.Collection("produto")

	cursor, err := produtoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var produtos []bson.M
	if err = cursor.All(context.TODO(), &produtos); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "produtos": produtos})
}

func CreateProdutos(c *gin.Context) {

	client := config.GetClient()
	database := client.Database("doacoes")

	var produtos []model.Produto
	c.Bind(&produtos)

	c.JSON(http.StatusOK, gin.H{"status": "success", "produtos": produtos})

}
