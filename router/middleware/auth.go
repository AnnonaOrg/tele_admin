package middleware

// import (
// 	"fmt"

// 	"github.com/AnnonaOrg/osenv"
// 	"github.com/AnnonaOrg/pkg/errno"
// 	"github.com/AnnonaOrg/pkg/token"
// 	"github.com/gin-gonic/gin"
// 	"github.com/umfaka/umfaka_core/handler"
// 	"github.com/umfaka/umfaka_core/internal/log"
// 	"github.com/umfaka/umfaka_core/internal/service"
// 	"github.com/umfaka/umfaka_core/pkg/jwt_token"
// )

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if _, err := token.ParseRequest(c); err != nil {
// 			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func AdminAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		header := c.Request.Header.Get("Authorization")

// 		secret := osenv.GetJWTSecret()
// 		if len(header) == 0 {
// 			c.Abort()
// 			return
// 		}

// 		var token string
// 		fmt.Sscanf(header, "Bearer %s", &token)

// 		// //校验cookie是否存在
// 		// cookie, err := c.Request.Cookie("admin_token")
// 		// if err != nil {
// 		// 	restful.UnLoginErr(c)
// 		// 	c.Abort()
// 		// 	return
// 		// }
// 		// token := cookie.Value

// 		//校验token是否成功解码
// 		claims, err := jwt_token.DecodeToken(token, secret)
// 		if err != nil {
// 			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
// 			c.Abort()
// 			return
// 		}

// 		//校验一致性
// 		userID, isOk := service.CheckUserLoginToken(token)
// 		if !isOk {
// 			handler.SendResponse(c, errno.ErrBadRequest, nil)
// 			c.Abort()
// 			return
// 		}

// 		action, _ := claims["action"].(string)
// 		log.Debugf("action: %s %s", userID, action)

// 		c.Next()
// 	}
// }
