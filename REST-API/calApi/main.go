package main

import(
    "github.com/gin-gonic/gin"
    "strconv"
)

func main(){
    router := gin.Default()

    performOperation := func(ctx *gin.Context, operation func(float64, float64) float64) {
        num1, _ := strconv.ParseFloat(ctx.Param("num1"), 64)
        num2, _ := strconv.ParseFloat(ctx.Param("num2"), 64)

        if num2 == 0 && operation == func(float64, float64) float64 { return float64(0) } {
            ctx.JSON(200, gin.H{"result": "Can not divide"})
            return
        }

        ctx.JSON(200, gin.H{"result": operation(num1, num2)})
    }

    router.GET("/add/:num1/:num2", func(ctx *gin.Context) {
        performOperation(ctx, func(a, b float64) float64 { return a + b })
    })

    router.GET("/sub/:num1/:num2", func(ctx *gin.Context) {
        performOperation(ctx, func(a, b float64) float64 { return a - b })
    })

    router.GET("/mul/:num1/:num2", func(ctx *gin.Context) {
        performOperation(ctx, func(a, b float64) float64 { return a * b })
    })

    router.GET("/div/:num1/:num2", func(ctx *gin.Context) {
        performOperation(ctx, func(a, b float64) float64 { return a / b })
    })

    router.Run(":8080")
}
