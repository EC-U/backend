package models

type Text2tag struct {
	TextID *Article  `orm:"column(TextID);rel(fk)"`
	TagID  *Category `orm:"column(TagID);rel(fk)"`
}
