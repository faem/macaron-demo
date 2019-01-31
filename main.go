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
    m.Post("/profile",r.CreateProfile)
    m.Get("/profile",r.ReadProfile)

    log.Fatal(http.ListenAndServe(":4000",m))
}