package models

type List struct {
	Name string `orm:"column(name);size(20);null"`
	Age  int    `orm:"column(age);null"`
}
