package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ActivityController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:CurriculumController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:ScheduleItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2curriculumController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:User2scheduleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["ecnu_code/backend/controllers:UserprofileController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
