package controllers

import "github.com/gin-gonic/gin"

func ErrorResponse(code uint, msg string) gin.H {
	return gin.H{"code": code, "error": msg}
}
