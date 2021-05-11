module github.com/example-web

go 1.16

replace (
	github.com/example-web/conf => ./gopj/example-web/conf
	github.com/example-web/e => ./gopj/example-web/e
	github.com/example-web/middleware => ./gopj/example-web/middleware
	github.com/example-web/models => ./gopj/example-web/models
	github.com/example-web/pkg/setting/ => ./gopj/example-web/pkg/setting
	github.com/example-web/pkg/util => ./gopj/example-web/pkg/util
	github.com/example-web/routers => ./gopj/example-web/routers

)

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.7.1
	github.com/go-ini/ini v1.62.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.62.0 // indirect

)
