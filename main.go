package main

import (
    r "Macaron-Demo/routers"
    "gopkg.in/macaron.v1"
    "log"
    "net/http"
)

func main() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())

    m.Get("/", r.Home)
    m.Get("/register",r.ShowRegister)
    m.Post("/register",r.CreateProfile)
    m.Get("/profile",r.ReadProfile)
    m.Get("/login",r.ShowLogin)
    m.Post("/login",r.Login)

    log.Fatal(http.ListenAndServe(":4000",m))
}