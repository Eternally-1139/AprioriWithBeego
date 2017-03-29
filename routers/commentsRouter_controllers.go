package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["test/controllers:AprioriController"] = append(beego.GlobalControllerRouter["test/controllers:AprioriController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
