package main

import(
    "github.com/gin-gonic/gin"
    "strconv"
)

func main(){
    router := gin.Default()

    router.GET("/add/:num1/:num2", func(ctx *gin.Context) {
        num1, _ := strconv.ParseFloat(ctx.Param("num1"), 64)
        num2, _ := strconv.ParseFloat(ctx.Param("num2"), 64)
        ctx.JSON(200, gin.H{"result": num1 + num2})
    })

    router.GET("/sub/:num1/:num2", func(ctx *gin.Context) {
        num1, _ := strconv.ParseFloat(ctx.Param("num1"), 64)
        num2, _ := strconv.ParseFloat(ctx.Param("num2"), 64)
        ctx.JSON(200, gin.H{"result": num1 - num2})
    })

    router.GET("/mul/:num1/:num2", func(ctx *gin.Context) {
        num1, _ := strconv.ParseFloat(ctx.Param("num1"), 64)
        num2, _ := strconv.ParseFloat(ctx.Param("num2"), 64)
        ctx.JSON(200, gin.H{"result": num1 * num2})
    })

    router.GET("/div/:num1/:num2", func(ctx *gin.Context) {
        num1, _ := strconv.ParseFloat(ctx.Param("num1"), 64)
        num2, _ := strconv.ParseFloat(ctx.Param("num2"), 64)
        ctx.JSON(200, gin.H{"result": num1 / num2})

        if num2 := 0 {
            
        }
    })
    router.Run(":8080")
}