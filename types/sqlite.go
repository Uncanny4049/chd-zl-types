package types

import "time"

type ActionRecord struct {
	Idx    int       `gorm:"column:idx;primaryKey"`
	Copy   string    `gorm:"column:copy"`
	Action string    `gorm:"column:action"`
	Item   string    `gorm:"column:item"`
	Date   time.Time `gorm:"column:date"`
	Role   string    `gorm:"column:role"`
}

type CopyRecord struct {
	Idx   int       `gorm:"column:idx;primaryKey"`
	Role  string    `gorm:"column:role"`
	Types string    `gorm:"column:types"`
	Date  time.Time `gorm:"column:date"`
}

type Income struct {
	Idx   int       `gorm:"column:idx;primaryKey"`
	Types int       `gorm:"column:type" json:"type"`
	Item  string    `gorm:"column:item"`
	Num   int       `gorm:"column:num"`
	Date  time.Time `gorm:"column:date" json:"date"`
	Role  string    `gorm:"column:role"`
}

type Player struct {
	Idx   int    `gorm:"column:idx;primaryKey"`
	Name  string `gorm:"column:name"`
	Level int    `gorm:"column:level"`
}

type RecordType struct {
	Nid  int    `gorm:"column:nid"`
	Name string `gorm:"column:name"`
}

type TaskInfo struct {
	Idx   int       `gorm:"column:idx;primaryKey"`
	Role  string    `gorm:"column:role"`
	Times int       `gorm:"column:times"`
	Date  time.Time `gorm:"column:date"`
}
