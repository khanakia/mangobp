package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khanakia/mangobp/mango/util"
	"github.com/spf13/viper"
	"github.com/ubgo/goutil"
	"github.com/ubgo/gqlgenfn"
)

// Block the host by ip address
// add below code to your default.yaml file
/*
	blocked_hosts:
  	- localhost:2143
		- 10.0.0.1
*/
func BlockHostsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		hosts := viper.GetStringSlice("blocked_hosts")
		_, ok := goutil.StringIndex(hosts, c.Request.Host)
		if !ok {
			ip := util.GetIP(c.Request)
			// fmt.Println(ip)
			_, ok = goutil.StringIndex(hosts, ip)
		}
		if ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"message": "Request blocked.",
				},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-captcha-res, X-CAPTCHA-BYPASS, x-auth-token, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func GetClientIP(ctx context.Context) string {
	gc, err := gqlgenfn.GinContextFromContext(ctx)
	if err != nil {
		return ""
	}

	userIP := GetUserIP(gc.Writer, gc.Request)
	return userIP
}

// Get the IP address of the server's connected user.
func GetUserIP(httpWriter http.ResponseWriter, httpServer *http.Request) string {
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
