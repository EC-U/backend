package models

type User2text struct {
	Userid *User     `orm:"column(userid);rel(fk)"`
	Textid *Activity `orm:"column(textid);rel(fk)"`
}
