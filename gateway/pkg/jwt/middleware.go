package jwt

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/maxim12233/crypto-app-server/gateway/internal/config"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/common"
)

// Validating jwt token in cookies and validating audience
func AuthMiddleware(validRoles []uint) func(c *gin.Context) {
	return func(c *gin.Context) {
		config := config.GetConfig()

		fmt.Println(c.GetString("Authorization"))
		tokenString := c.GetString("Authorization")
		if tokenString == "0" || tokenString == "" {
			common.ReturnAnauthorized(c)
			return
		}

		// Decode/validate it
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.Server.SecretKey), nil
		})
		if err != nil {
			common.ReturnAnauthorized(c)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				common.ReturnAnauthorized(c)
				return
			}

			rolesVal := claims["roles"].([]interface{})
			if rolesVal == nil {
				common.ReturnAnauthorized(c)
				return
			}

			roles := make([]uint, len(rolesVal))
			for i, val := range rolesVal {
				roles[i] = uint(val.(float64))
			}
			if len(roles) == 0 {
				common.ReturnAnauthorized(c)
				return
			}

			validation := make(map[uint]uint)
			for _, val := range roles {
				validation[val] = 0
			}

			for _, val := range validRoles {
				if _, ok := validation[val]; !ok {
					common.ReturnAnauthorized(c)
					return
				}
			}

			c.Set("claims", claims)
			c.Set("ID", claims["sub"])

			c.Next()
		} else {
			common.ReturnAnauthorized(c)
		}
	}
}

func AuthorizationToken(c *gin.Context) {
	var token string

	token = c.GetHeader("Authorization")
	if token != "" {
		c.Set("Authorization", token)
		return
	}

	token, ok := c.GetQuery("authorization")
	if ok {
		c.Set("Authorization", token)
		return
	}

	token, ok = c.GetPostForm("authorization")
	if ok {
		c.Set("Authorization", token)
		return
	}
}
