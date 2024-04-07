package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

var JwtSecretKey = []byte("jwt-secret-key")

type Login struct {
    User     string `json:"user"`
    Password string `json:"password"`
}

type UserLoginClaims struct {
    User string `json:"user"`
    jwt.StandardClaims
}

func JwtAuthTokenGenerate(c *gin.Context) {
    var login Login
    if err := c.ShouldBindJSON(&login); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    claims := UserLoginClaims{
        User: login.User,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
            Issuer:    "Your_app",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString(JwtSecretKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func JwtAuthMiddleware(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
        c.Abort()
        return
    }

    claims := &UserLoginClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return JwtSecretKey, nil
    })

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        c.Abort()
        return
    }

    if !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    c.Next()
}

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })

    loginMiddleware := func(c *gin.Context) {
        fmt.Println("API is running, and Request Type:", c.Request.Method)
        c.Next()
    }

    apiRoutes := router.Group("/api", loginMiddleware)
    {
        apiRoutes.GET("/UP", func(ctx *gin.Context) {
            ctx.JSON(200, gin.H{
                "message": "Hello World",
                "Time":    time.Now().UTC().Format(time.RFC3339),
            })
        })

        apiRoutes.GET("/updates", func(ctx *gin.Context) {
            name := ctx.DefaultQuery("name", "Guest")
            ctx.String(200, "Hello %s", name)
        })

        apiRoutes.GET("/", func(ctx *gin.Context) {
            ctx.JSON(200, gin.H{
                "Date": time.Now(),
            })
        })

        apiRoutes.GET("/user/:name", func(ctx *gin.Context) {
            name := ctx.Param("name")
            ctx.JSON(200, gin.H{
                "Name": name,
                "Time": time.Now().UTC().Format(time.RFC3339),
            })
        })

        // JWT Authentication routes
        apiRoutes.POST("/login", JwtAuthTokenGenerate)
        apiRoutes.GET("/protected", JwtAuthMiddleware, func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "This is a protected endpoint"})
        })
    }

    router.Run(":8080")
}