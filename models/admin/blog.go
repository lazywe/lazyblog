package admin

type Blog struct {
	Id         int    `orm:"pk;auto"`
	Title      string `orm:"default(0);size(32)"`
	Blog       string `orm:"type(text)"`
	CreateTime uint   `orm:"default(0);size(10)"`
}
