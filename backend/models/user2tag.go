package models

type User2tag struct {
	Userid *User     `orm:"column(userid);rel(fk)"`
	Tagid  *Category `orm:"column(tagid);rel(fk)"`
}
