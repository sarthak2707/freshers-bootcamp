package day3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitialiseRoute() {

	router := gin.New()

	//router.Use(gin.Recovery())

	var middlewareList []gin.HandlerFunc

	middlewareList = append(middlewareList, Panic)

	group := router.Group("/", middlewareList...)

	group.Handle(http.MethodGet, "/get", GetAPI)

	router.Run(":8080")

}

func GetAPI(ctx *gin.Context) {
	panic("we have a panic")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"response": "test gin",
	})
}

func Panic(ctx *gin.Context) {

	defer func(ctx *gin.Context) {

		if err := recover(); err != nil {
			fmt.Println("in recovery ", err)

		}

	}(ctx)

	fmt.Println("in middleware")

	ctx.Next()

}
