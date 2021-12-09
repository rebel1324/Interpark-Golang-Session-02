package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func initLogger() *zap.Logger {
	// 새로운 Zap 로거 설정
	logger, err := zap.NewProduction()
	//예외처리
	if err != nil {
		// Zap 로거 설정 실패
		panic("Unable to initialize zap logger")
	}
	return logger
}

func CustomLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 시작 전
		logger.Info("request-start", zap.String("url", c.Request.URL.Path))
		// 다음 요청으로 넘어가기
		// 이게 호출이 안되면 여기서 요청이 끝납니다
		c.Next()
		// 다음 요청 완료 후
		logger.Info("request-done", zap.String("url", c.Request.URL.Path))
	}
}

func main() {
	// 안?뇽?
	fmt.Println("hello world")

	// gin 핸들러 초기화
	r := gin.Default()

	logger := initLogger()
	// flush (close)같은 느낌,
	// 다 못쓴 로그를 다 쓸때까지 기다린다
	defer logger.Sync()

	r.Use(CustomLogger(logger))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // hardcoded 8080 listen
}
