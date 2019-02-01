package routers

import (
	"Macaron-Demo/models"
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
)

func ReadProfile(c *macaron.Context)  {
	c.Write([]byte(fmt.Sprint(models.ReadProfile())))
	return
}
func ShowRegister(c *macaron.Context)  {
	c.HTML(200,"register")
	c.Data["Success"] = false
}

func CreateProfile(c *macaron.Context) {
	username := c.Query("username")
	if isEmpty(username, "Username", c) {
		return
	}

	name := c.Query("name")
	if isEmpty(name, "Name", c){
		return
	}
	company := c.Query("company")
	if isEmpty(company, "Company", c){
		return
	}
	position := c.Query("position")
	if isEmpty(position, "Position", c){
		return
	}
	password := c.Query("password")
	if isEmpty(password, "Password", c){
		return
	}
	log.Println(username, name, company, position, password)
	err := models.CreateProfile(username,name,company,position,password)
	if err!= nil{
		c.Write([]byte(err.Error()))
		return
	}
	c.Data["Success"] = true
	c.HTML(200,"register")
	return
}

func ShowLogin(c *macaron.Context)  {
	c.Data["Success"] = false
	c.HTML(200,"login")

}

func Login(c *macaron.Context)  {
	username := c.Query("username")
	if isEmpty(username, "Username", c) {
		return
	}

	password := c.Query("password")
	if isEmpty(password, "Password", c){
		return
	}

	ok, err := models.MatchUsernamePass(username, password)
	if ok{
		c.Data["Success"] = true
		c.Data["Username"] = username
		c.HTML(200, "login")
	}else{
		c.Write([]byte("Username & Password doesn't match!"))
	}

	if err!= nil {
		c.Write([]byte(err.Error()))
	}
}

func isEmpty(s string, t string, c *macaron.Context) bool  {
	if s == "" {
		c.Write([]byte(t+" is empty!\n"))
		return true
	}
	return false
}