package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	//intinnya middleware interceptor, ditaro di tengah logic flow
	//kalau kondisi gak memenuhi, middleware gak akan jalan
	return func(ctx *gin.Context) {

		if ctx.Request.URL.Path == "/api/auth/login" {
			ctx.Next()
		} else {
			h := AuthHeader{}
			if err := ctx.ShouldBindHeader(&h); err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				ctx.Abort()
			}
			//header bisa catet lu pake broeswer apa, metadata pendukung
			//token itu informsi tambahan
			//header data pendukung

			//body --> itu nyimpen data kayak username sama apssword, tapi data pendukung ada di header

			if h.AuthorizationHeader == "ini_token" {
				ctx.Next()
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "token invalid",
				})
				ctx.Abort()
			}
		}
	}
}
