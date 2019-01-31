package main

import (
    r "Macaron-Demo/routers"
    "gopkg.in/macaron.v1"
)

func main() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())
    m.Get("/", r.Home)
    //m.Post("profile",r.CreateProfile)
    m.Run()
}