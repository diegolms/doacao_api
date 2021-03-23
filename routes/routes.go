package routes
import (
       "net/http"
        doacaoController "doacao_api/controllers/doacao"
        produtoController "doacao_api/controllers/produto"
       "github.com/gin-gonic/gin"
       cors "github.com/rs/cors/wrapper/gin"       
)

func StartGin() {
       router := gin.Default()

       router.Use(cors.Default())

       v1 := router.Group("/api/v1")
       {
                doacoes := v1.Group("/doacoes")
                {
					doacoes.GET("/", doacaoController.GetAllDoacoes)
                                   //doacoes.POST("/", doacaoController.CreateDoacao)
                                   //doacoes.DELETE("/", doacaoController.DeleteDoacao)
                }

                produtos := v1.Group("/produtos")
                {
                                   produtos.GET("/", produtoController.GetAllProdutos)
                                   produtos.POST("/", produtoController.CreateProdutos)
                                   //produtos.DELETE("/", produtoController.DeleteProduto)
                }
       }

       router.NoRoute(func(c *gin.Context) {
              c.AbortWithStatus(http.StatusNotFound)
       })

       router.Run()
}
