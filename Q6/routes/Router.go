package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tinder-server/middlewares"
	"tinder-server/providders"
	"tinder-server/routes/v1Router"
)

type Command struct {
	Name          string   // Command name(case-sensitive).
	Usage         string   // A brief line description about its usage, eg: gf build main.go [OPTION]
	Brief         string   // A brief info that describes what this command will do.
	Description   string   // A detailed description.
	Run           function // Custom function.
	Additional    string   // Additional info about this command, which will be appended to the end of help info.
	Strict        bool     // Strict parsing options, which means it returns error if invalid option given.
	CaseSensitive bool     // CaseSensitive parsing options, which means it parses input options in case-sensitive way.
	Config        string   // Config node name, which also retrieves the values from config component along with command line.

}

type function func() (err error)

var (
	Main = Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Run: func() (err error) {
			Routers.Run()
			return nil
		},
	}
)

type routerStruct struct {
}

var (
	Routers routerStruct
)

func (r *routerStruct) Run() {

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		fmt.Errorf("%s", err)
	} else {
		router.Use(
			middlewares.Middleware.CORSMiddleware,
			middlewares.Middleware.DebugMiddleware,
			middlewares.Middleware.RequestMiddleware,
			middlewares.Middleware.ResponseMiddleware,
		)
		rootV1 := router.Group("/")
		{
			v1Router.V1Router.InitRouter(rootV1)
		}
		router.Run(fmt.Sprintf(":%s", providders.Provider.GetEnv("port", "8080")))
	}

}
