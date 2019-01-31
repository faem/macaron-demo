package models

type Profile struct {
        UserName string `xorm:"pk not null"`
        Name     string
        Company  string
        Position string
        Password string
}

func ReadProfile(username string) {
        var profile = Profile{
                UserName: username,
        }
        _, err := engine.Get(&profile)
        printError(err)
}

func CreateProfile(username, name, company, position, password string) {
        var profile = Profile{
                UserName: username,
                Name:     name,
                Company:  company,
                Position: position,
                Password: password,
        }
        _, err := engine.Insert(&profile)
        printError(err)
}

func printError(err error) {
        if err != nil {
                panic(err.Error())
        }
}
