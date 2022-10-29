module server

go 1.16

require (
	github.com/armon/consul-api v0.0.0-20180202201655-eb2c6b5be1b6 // indirect
	github.com/coreos/etcd v3.3.10+incompatible // indirect
	github.com/coreos/go-etcd v2.0.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-gonic/gin v1.8.1
	github.com/mojocn/base64Captcha v1.3.5
	github.com/spf13/viper v1.13.0 // indirect
	github.com/xordataexchange/crypt v0.0.3-0.20170626215501-b2862e3d0a77 // indirect
	golang.org/x/image v0.1.0 // indirect
)

replace server => ./server