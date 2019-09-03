// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ecnu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/activity",
			beego.NSInclude(
				&controllers.ActivityController{},
			),
		),

		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),

		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),

		beego.NSNamespace("/migrations",
			beego.NSInclude(
				&controllers.MigrationsController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/userprofile",
			beego.NSInclude(
				&controllers.UserprofileController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
