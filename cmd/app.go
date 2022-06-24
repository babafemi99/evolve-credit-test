package cmd

import (
	"evc/dataSource/psql"
	"evc/handler"
	"evc/repository/psqlRepo"
	"evc/service/userService"
	"evc/web"
	"fmt"
	"net/http"
	"os"
)

const ()

var (
	dbConnection   = psql.GetConn()
	Repository     = psqlRepo.NewPsql(dbConnection)
	UserService    = userService.NewUserService(Repository)
	UserController = handler.NewUserController(UserService)
	WebRouter      = web.NewChiRouter()
)

func Start() {
	port := os.Getenv("port")
	if port == "" {
		port = ":9090"
	}
	WebRouter.GET("/ping", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Up and running")
		rw.Write([]byte("pong"))
	})
	WebRouter.POST("/user", UserController.SaveUser)
	WebRouter.GET("/users", UserController.GetAllUser)
	WebRouter.GET("/users/{email}", UserController.GetUserByEmail)
	WebRouter.SERVE(port)
}
