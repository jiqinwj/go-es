package main

import (
	"go-es/gg"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main()  {

	schema:=gg.NewLogQuerySchema()
	h:=handler.New(&handler.Config{
		Schema:&schema,
	})
   router:=gin.Default()
   g:=router.Group("/logs")
   {
		g.Handle("POST","/", func(context *gin.Context) {
			 h.ServeHTTP(context.Writer,context.Request)
		})
   }
   router.Run(":8086")



}