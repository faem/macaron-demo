package models

import "log"

type Profile struct {
        UserName string `xorm:"pk not null"`
        Name     string
        Company  string
        Position string
        Password string
}

func ReadProfile(username string) []Profile {
        var profile = make([]Profile,1)
        err := engine.Find(&profile)
        printError(err)

        return profile
}

func CreateProfile(username, name, company, position, password string) {
        var profile = Profile{
                UserName: username,
                Name:     name,
                Company:  company,
                Position: position,
                Password: password,
        }
        log.Println(profile)
        log.Println("engine:",engine)
        _, err := engine.Insert(profile)
        printError(err)
}

func printError(err error) {
        if err != nil {
                panic(err.Error())
        }
}
