package session

import (
	// "os"
	"github.com/gin-gonic/gin"

	"gitee.com/douwatech/account/src/model"
)

func User(c *gin.Context) *model.User {
	v, ok := c.Get("user")
	if !ok {
		return nil
	}
	u, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return u
}

// func MustAdmin(grpc_client pb.GreeterClient) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		role, err := GetRole(c, grpc_client)
// 		if err != nil {
// 			c.AbortWithError(404, err)
// 		}
//
// 		if role == "admin" {
// 			c.Next()
// 		} else {
// 			c.String(403, "User not authorized")
// 			c.Abort()
// 		}
// 	}
// }
//
// func MustRepoAdmin() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		user := User(c)
// 		//perm := Perm(c)
// 		switch {
// 		case user == nil:
// 			c.String(403, "User not authorized")
// 			c.Abort()
// 		//case perm.Admin == false:
// 		//	c.String(403, "User not authorized")
// 		//	c.Abort()
// 		default:
// 			c.Next()
// 		}
// 	}
// }
//
// func MustUser(grpc_client pb.GreeterClient) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		role, err := GetRole(c, grpc_client)
// 		if err != nil {
// 			c.AbortWithError(404, err)
// 		}
//
// 		if role == "admin" || role == "developer" {
// 			c.Next()
// 		} else {
// 			c.String(403, "User not authorized")
// 			c.Abort()
// 		}
// 	}
// }
