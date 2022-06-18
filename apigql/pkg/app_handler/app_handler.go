package app_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ubgo/gofm/ginserver"
	"github.com/ubgo/gofm/gormdb"
	"github.com/ubgo/gofm/logger"
)

type Config struct {
	GormDB gormdb.GormDB
	Logger logger.Logger
	Server ginserver.Server
}

type AppHandler struct {
	Config
}

func New(config Config) AppHandler {
	pkg := AppHandler{
		Config: config,
	}

	router := config.Server.Router
	// router.Any("/ping", pkg.PingHandler)
	router.Any("/ma/client_ip", pkg.ClientIpHandler)

	// routeSpace := router.Group("/spaces/:spaceid") // yeh readability key liye use kiya naki security ke liye
	// routeSpace.Use(middleware.SpaceApiKeyMiddleware(config.GormDB.DB))
	// routeSpace.GET("/request/run/:id", pkg.RunRequest)
	// routeSpace.GET("/token/refresh/:id", pkg.RefreshOat)

	// routesOauth := router.Group("/oauth")
	// routesOauth.GET("/login", pkg.OauthLogin)
	// routesOauth.GET("/verify", pkg.OauthVerify)

	return pkg
}

func (pkg AppHandler) ClientIpHandler(c *gin.Context) {

	userIP := getUserIP(c.Writer, c.Request)
	c.JSON(200, gin.H{
		"message": userIP,
	})
}

func (pkg AppHandler) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Get the IP address of the server's connected user.
func getUserIP(httpWriter http.ResponseWriter, httpServer *http.Request) string {
	var userIP string
	if len(httpServer.Header.Get("CF-Connecting-IP")) > 1 {
		userIP = httpServer.Header.Get("CF-Connecting-IP")
	} else if len(httpServer.Header.Get("X-Forwarded-For")) > 1 {
		userIP = httpServer.Header.Get("X-Forwarded-For")
	} else if len(httpServer.Header.Get("X-Real-IP")) > 1 {
		userIP = httpServer.Header.Get("X-Real-IP")
	} else {
		userIP = httpServer.RemoteAddr
		// if strings.Contains(userIP, ":") {
		// 	fmt.Println(net.ParseIP(strings.Split(userIP, ":")[0]))
		// } else {
		// 	fmt.Println(net.ParseIP(userIP))
		// }
	}

	return userIP
}
