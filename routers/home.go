package routers

import "gopkg.in/macaron.v1"

func Home(c *macaron.Context)  {
	c.HTML(200,"home")
}