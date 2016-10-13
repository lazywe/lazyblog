package models

type User struct {
	Id        int    `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
	Avatar    string
	Email     string  `orm:"null"`
	Url       string  `orm:"null"`
	Signature string  `orm:"null;size(1000)"`
	InTime    int     `orm:"0;size(10)"`
	Roles     []*Role `orm:"rel(m2m)"`
	Remark    string  `orm:"null"`
}
