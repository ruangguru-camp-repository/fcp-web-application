package middleware

import (
	cfg "a21hc3NpZ25tZW50/config"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		token, err := ctx.Cookie("session_token")
		s := ctx.ContentType()
		jsonType := "application/json"

		if err != nil && s == jsonType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			return
		}
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusSeeOther, model.NewErrorResponse("redirect to login page"))
			return
		}
		tokenStr, err2 := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return cfg.Config.JWTKey, nil
		})

		if err2 != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse("unauthorized"))
			return
		}

		claims, ok := tokenStr.Claims.(jwt.MapClaims)
		if !ok || ! tokenStr.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
		}

		b, _ := json.Marshal(claims)
		var customClaims model.Claims
		json.Unmarshal(b, &customClaims)

		ctx.Set("email", customClaims.Email)

		// TODO: answer here 
	})
}
