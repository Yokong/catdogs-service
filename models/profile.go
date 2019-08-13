package models

var (
	profile = Mdb.DB("account").C("profile")
)

type Profile struct {
	Name     string
	Gender   string
	Age      int32
	PhoneNum string
	Email    string
	Birthday string
	City     string
	Address  string
}

func (p *Profile) Set() {
	profile.Insert(p)
}
