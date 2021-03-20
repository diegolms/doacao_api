package routes
import (
       "net/http"
        doacaoController "doacao_api/controllers/doacao"
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
					doacoes.GET("/listar", doacaoController.GetAllDoacoes)
                }
       }

       router.NoRoute(func(c *gin.Context) {
              c.AbortWithStatus(http.StatusNotFound)
       })

       router.Run()
}
