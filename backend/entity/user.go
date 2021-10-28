package entity

import (
	"time"

	"gorm.io/gorm"
)

type Club struct {
	gorm.Model
	Name string
	// 1 club can create many activities
	Activities []Activity `gorm:"foreignKey:ClubID"`
	// 1 club can have many committees
	Editors []ClubCommittee `gorm:"foreignKey:ClubID"`
}

type Activity struct {
	gorm.Model
	Name string
	Time time.Time
	Amount uint 
	// 1 activities can be in many JoinActivity
	JoinActivityHistories []JoinActivityHistory `gorm:"foreignKey:ActivityID"`

	ClubID 		*uint
	Club 		Club
}

type Student struct {
	gorm.Model
	Name      string
	ID_Student string `gorm:"uniqueIndex"`
	// 1 user can be in many JoinActivity
	JoinActivityHistories []JoinActivityHistory `gorm:"foreignKey:StudentID"`
}

type ClubCommittee struct {
	gorm.Model
	Name      string
	ID_Student string `gorm:"uniqueIndex"`
	Password string
	// 1 ClubCommittee can create many JoinActivity
	JoinActivityHistories []JoinActivityHistory `gorm:"foreignKey:EditorID"`

	ClubID *uint
	Club   Club
}

type JoinActivityHistory struct {
	gorm.Model
	HourCount uint
	Point     uint
	Timestamp time.Time

	ActivityID      *uint
	Activity        Activity 		`gorm:"references:ID"`

	StudentID       *uint
	Student         Student 		`gorm:"references:ID"`
	
	EditorID 		*uint
	Editor   ClubCommittee 			`gorm:"references:ID"`
}
