module server

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/mojocn/base64Captcha v1.3.5
	golang.org/x/image v0.1.0 // indirect
)

replace server => ./server