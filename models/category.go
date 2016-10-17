package models

type Category struct {
	Id         int    `orm:"pk;auto;"`
	Title      string `orm:"default(0);size(32);"`
	CreateTime uint   `orm:size(10)`
}
