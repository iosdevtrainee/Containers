package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    port := 3000
    engine := gin.Default()
    engine.Run(fmt.Sprintf(":%d", port))
    fmt.Printf("Server started at port %d", port)
}
