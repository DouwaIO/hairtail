package server

import (
	"github.com/gin-gonic/gin"
	"github.com/DouwaIO/hairtail/src/store"

	"github.com/DouwaIO/hairtail/src/model"
)


func Test(c *gin.Context) {
	//c.JSON(200, "test teste")
	user, err := store.FromContext(c).GetUser(1)
	if err != nil {
		new_user := &model.User{
			Login:  "test",
			Token: "test test",
		}
		err = store.FromContext(c).CreateUser(new_user)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.JSON(200, new_user)

	} else {
		c.JSON(200, user)
	}
}

