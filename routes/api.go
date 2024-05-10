package routes

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/http/controllers"
)

func Api(router contracts.HttpRouter) {

	router.Get("/wrk/{name}", func(request contracts.HttpRequest) any {
		return request.Param("name")
	})

	router.Post("/queue", controllers.DemoJob)

	router.Get("/", controllers.HelloWorld)
	//router.Get("/", controllers.HelloWorld, ratelimiter.Middleware(100))
	router.Post("/login", controllers.LoginExample)

	authRouter := router.Group("", auth.Guard("jwt"))
	authRouter.Get("/myself", controllers.GetCurrentUser, auth.Guard("jwt"))
	router.Get("/users/{id}", controllers.GetUser)

	router.Post("/mail", controllers.SendEmail)
}
