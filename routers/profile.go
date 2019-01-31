package routers

import (
	"Macaron-Demo/models"
	"gopkg.in/macaron.v1"
)

func ReadProfile(c *macaron.Context)  {

}

func CreateProfile(c *macaron.Context) {
	username := c.Query("username")
	isEmpty(username, "Username", c)
	name := c.Query("name")
	isEmpty(name, "Name", c)
	company := c.Query("company")
	isEmpty(company, "Company", c)
	position := c.Query("position")
	isEmpty(position, "Position", c)
	password := c.Query("password")
	isEmpty(password, "Password", c)

	models.CreateProfile(username,name,company,position,password)
	c.Write([]byte("Profile Created Successfully!"))
	return
}

func isEmpty(s string, t string, c *macaron.Context)  {
	if s == "" {
		c.Write([]byte(t+" is empty!"))
		return
	}
}