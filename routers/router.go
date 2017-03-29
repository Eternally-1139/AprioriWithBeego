package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Include(&controllers.AprioriController{},
		)

}
