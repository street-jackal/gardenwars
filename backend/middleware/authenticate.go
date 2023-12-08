package middleware

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/street-jackal/gardenwars/env"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get the raw token from the request header
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// parse the token
		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				// Validate the alg is what we require
				method, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok || method != jwt.SigningMethodHS256 {
					return nil, errors.New("Unexpected signing method. Method: " + method.Name)
				}

				return env.GetUserJWTSecret(), nil
			},
		)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// TODO: use fields from the Claims that we need (for expiration etc)
		mapClaims := token.Claims.(jwt.MapClaims)
		fmt.Printf("%+v\n", mapClaims)

		c.Next()
	}
}
