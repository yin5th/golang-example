package routers

import (
	"dothis.top/example/beeblog_example/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
