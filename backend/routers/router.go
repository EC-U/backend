// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ecnu_code/backend/controllers"

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

		beego.NSNamespace("/curriculum",
			beego.NSInclude(
				&controllers.CurriculumController{},
			),
		),

		beego.NSNamespace("/migrations",
			beego.NSInclude(
				&controllers.MigrationsController{},
			),
		),

		beego.NSNamespace("/schedule_item",
			beego.NSInclude(
				&controllers.ScheduleItemController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/user2curriculum",
			beego.NSInclude(
				&controllers.User2curriculumController{},
			),
		),

		beego.NSNamespace("/user2schedule",
			beego.NSInclude(
				&controllers.User2scheduleController{},
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
