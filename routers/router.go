// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"yuwenlanzi/controllers"
)

func init() {
	//ns := beego.NewNamespace("/",
	//	//beego.NSNamespace("/v1/object",
	//	//	beego.NSInclude(
	//	//		&controllers.ObjectController{},
	//	//	),
	//	//),
	//	//beego.NSNamespace("/v1/user",
	//	//	beego.NSInclude(
	//	//		&controllers.UserController{},
	//	//	),
	//	//),
	//
	//	beego.NSInclude(
	//		&controllers.WechatController{},
	//		),
	//)
	//beego.AddNamespace(ns)

	beego.Router("/",&controllers.WechatController{},"*:Index")
}
