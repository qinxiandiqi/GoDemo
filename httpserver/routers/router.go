package routers

import (
	"github.com/qinxiandiqi/GoDemo/httpserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
