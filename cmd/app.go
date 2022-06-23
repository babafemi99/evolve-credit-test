package cmd

import (
	"evc/dataSource/psql"
	"evc/handler"
	"evc/repository/psqlRepo"
	"evc/service/userService"
	"evc/web"
)

var (
	dbConnection   = psql.GetConn()
	Repository     = psqlRepo.NewPsql(dbConnection)
	UserService    = userService.NewUserService(Repository)
	UserController = handler.NewUserController(UserService)
	WebRouter      = web.NewRouterMux()
)

func Start() {

}
