package main

import (
	"go-es/Funs"
	"github.com/gin-gonic/gin"
)

func main()  {


	router:=gin.Default()
	g:=router.Group("/books")
	{
		//图书列表
		g.Handle("GET","",Funs.LoadBook)
		g.Handle("GET","/press/:press",Funs.LoadBookByPress)
	}
	router.Run(":8080")


}
