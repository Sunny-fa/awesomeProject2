package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
var a map[string]string
func main() {
     a=make(map[string]string)

	router := gin.Default()
	router.GET("/register", func(c *gin.Context) {
		name := c.Query("name")
		password:=c.Query("password")
		a[name]=password
		_,ok:=a[name]
		if ok==true{
			c.String(200, "该用户名已被注册！")
			return
		}
		c.String(http.StatusOK, "Hello %s", name)

	})
	router.GET("/login", func(c *gin.Context) {

		name := c.Query("name")
		password:=c.Query("password")
		if a[name]!=""&&a[name]==password{
			c.String(200,"登陆成功！")
			return
		}
		c.String(200,"登陆失败！")

		cookie := &http.Cookie{

			Name:     "login_account",
			Value:    name,
			Path:     "/",
			HttpOnly: true,

		}

		http.SetCookie(c.Writer, cookie)

		c.String(200, "hello "+name)
	})


	router.Run(":8080")

}