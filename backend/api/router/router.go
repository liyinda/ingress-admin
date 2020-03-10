package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/liyinda/ingress-admin/backend/api/apis"
	"net/http"
	//"fmt"
	//"github.com/liyinda/viewdns/backend/middleware/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	/*
	   router.Use(cors.New(cors.Config{
	       //AllowOrigins:     []string{"*"},
	       //AllowOrigins:     []string{"http://172.20.52.36", "http://localhost"},
	       AllowOrigins:     []string{"http://172.20.52.36:9528"},
	       AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE"},
	       //AllowMethods:     []string{"*"},
	       AllowHeaders:     []string{"content-type","Content-Type","Authorization","X-Token","Access-Control-Allow-Origin"},
	       //AllowHeaders:     []string{"*"},
	       ExposeHeaders:    []string{"Content-Length"},
	       AllowCredentials: true,
	       AllowOriginFunc: func(origin string) bool {
	           return origin == "*"
	       },
	   }))
	*/

	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		//AllowHeaders:     []string{"Authorization", "Content-Length", "Content-Type","X-Token", "Access-Control-Allow-Origin"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "X-Token", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		AllowOrigins:     []string{"http://localhost:9528"},
	}))

	//设置sessions
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//登录入口
	passport := router.Group("/passport")
	{
		passport.POST("/login", Login)
		passport.POST("/logout", Logout)
	}

	//用户管理入口
	home := router.Group("/home")
	//home.Use(jwt.JWT())
	{
		home.GET("/userinfo", Userinfo)
		home.GET("/table", Ingresslist)
		//        home.GET("/table", Table)
		//        home.POST("/adddomain", AddDomain)
		//        home.POST("/deldomain", DelDomain)
	}
	//    //home.Use(AuthRequired())
	//

	//定义默认路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})

	return router
}
