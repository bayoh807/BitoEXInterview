package v1Router

const (
	authPrefix = "/auth"
)

//func (r *v1Router) unAuthGroup(rg *gin.RouterGroup) {
//	pg := rg.Group(authPrefix)
//	{
//		pg.POST("login", v1Controller.V2AuthController.Login)
//	}
//
//}

//func (r *v2Router) authGroup(rg *gin.RouterGroup) {
//	pg := rg.Group(authPrefix)
//	{
//		pg.GET("profile", v1Controller.V2AuthController.GetProfile)
//		pg.GET("register/hasUnPaymentIntent", v1Controller.V2AuthController.HasUnPaymentIntent)
//	}
//
//}
