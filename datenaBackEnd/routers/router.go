// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	//controllers2 "datenaBackEnd/datenaBackEnd/controllers"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego"
)

func init() {
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			//&controllers2.ObjectController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			//&controllers2.UserController{},
	//		),
	//	),
	//)
	note := beego.NewNamespace("/note",
		beego.NSNamespace("/object",
			beego.NSInclude(

			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(

			),
		),
	)
	beego.AddNamespace(note)
}
