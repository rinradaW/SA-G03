package entity

import (
	"time"

	"gorm.io/gorm"
)

//B6210564 นายภัทรพงษ์ พิมหอม ระบบบันทึกข้อมูลชมรม

type StudentCouncil struct {
	gorm.Model
	Name       string
	ID_Student string `gorm:"uniqueIndex"`
	Password   string
	Clubs      []Club `gorm:"foreignKey:AdderID"`
}

type Teacher struct {
	gorm.Model
	Name  string
	Clubs []Club `gorm:"foreignKey:AdviserID"`
}

type TypeClub struct {
	gorm.Model
	Name  string
	Clubs []Club `gorm:"foreignKey:TypeClubID"`
}

type Club struct {
	gorm.Model
	Name string

	AdderID *uint
	Adder   StudentCouncil

	AdviserID *uint
	Adviser   Teacher

	TypeClubID *uint
	TypeClub   TypeClub

	// 1 club can create many activities
	Activities []Activity `gorm:"foreignKey:ClubID"`

	// 1 club can have many committees
	Editors []ClubCommittee `gorm:"foreignKey:ClubID"`

	// 1 Club สมารถมี ClubMembers ได้หลายคน
	ClubMembership []ClubMembership `gorm:"foreignKey: ClubID"`
}

/*อันนี้ของเลดี้นะครับ*/

type Activity struct {
	gorm.Model
	Name   string
	Time   time.Time
	Amount uint

	// 1 activities can be in many JoinActivity
	JoinActivityHistories []JoinActivityHistory `gorm:"foreignKey:ActivityID"`

	BudgetProposals []BudgetProposal `gorm:"foreignkey:ActivityID"`

	Joinings []Joining `gorm:"foreignKey:ActivityID"`

	ReserveLocations []ReserveLocation `gorm:"foreignKey:ActivityID"`

	ClubID *uint
	Club   Club
}

type Student struct {
	gorm.Model
	Name       string
	ID_Student string `gorm:"uniqueIndex"`
	Password   string
	// 1 user can be in many JoinActivity
	JoinActivityHistories []JoinActivityHistory `gorm:"foreignKey:StudentID"`

	Joinings []Joining `gorm:"foreignKey:StudentID"`

	// 1 User สมารถเป็น Clubmembers ได้หลายชมรม
	ClubMembership []ClubMembership `gorm:"foreignKey:StudentID"`
}

type ClubCommittee struct {
	gorm.Model
	Name       string
	ID_Student string `gorm:"uniqueIndex"`
	Password   string

	ClubID *uint
	Club   Club

	// 1 ClubCommittee can create many JoinActivity
	JoinActivityHistories []JoinActivityHistory `gorm:"foreignKey:EditorID"`

	ReserveLocations []ReserveLocation `gorm:"foreignKey:RequestID"`
}

type JoinActivityHistory struct {
	gorm.Model
	HourCount uint
	Point     uint
	Timestamp time.Time

	ActivityID *uint
	Activity   Activity `gorm:"references:ID"`

	StudentID *uint
	Student   Student `gorm:"references:ID"`

	EditorID *uint
	Editor   ClubCommittee `gorm:"references:ID"`
}

/*จบของเลดี้เพียงเท่านี้*/
// แบงค์เอง ,ใช่เหย๋อ

type BudgetCategory struct {
	gorm.Model
	Name            string
	BudgetProposals []BudgetProposal `gorm:"foreignKey:BudgetCategoryID"`
}
type BudgetType struct {
	gorm.Model
	Name            string
	BudgetProposals []BudgetProposal `gorm:"foreignKey:BudgetTypeID"`
}

type BudgetProposal struct {
	gorm.Model
	BudgetPrice uint

	// ActivityID ทำหน้าที่เป็น FK
	ActivityID *uint
	Activity   Activity

	// CategoryID ทำหน้าที่เป็น FK
	BudgetCategoryID *uint
	BudgetCategory   BudgetCategory

	// TypeBudgetId ทำหน้าที่เป็น Fk
	BudgetTypeID *uint
	BudgetType   BudgetType
}

// จบแบงค์

//อ๊อฟ
type Joinstatus struct {
	gorm.Model
	Name string
	// 1 Joinstatus เป็นเจ้าของได้หลาย joining
	Joinings []Joining `gorm:"foreignKey:JoinstatusID"`
}

type Joining struct {
	gorm.Model
	Joining_time time.Time

	//StudentID	ทำหน้าที่เป็น FK
	StudentID *uint
	Student   Student

	//ActivityID ทำหน้าที่เป็น FK
	ActivityID *uint
	Activity   Activity

	//JoinstatusID ทำหน้าที่เป็น FK
	JoinstatusID *uint
	Joinstatus   Joinstatus
}

//ของนุ

type Location struct {
	gorm.Model
	Name             string
	ReserveLocations []ReserveLocation `gorm:"foreignKey:LocationID"`
}
type ReserveStatus struct {
	gorm.Model
	Label            string
	ReserveLocations []ReserveLocation `gorm:"foreignKey:ReserveStatusID"`
}

type ReserveLocation struct {
	gorm.Model

	DateStart time.Time
	DateEnd   time.Time

	LocationID *uint
	Location   Location `gorm:"references:id"`

	RequestID *uint
	Request   ClubCommittee `gorm:"references:id"`

	ActivityID *uint
	Activity   Activity `gorm:"references:id"`

	ReserveStatusID *uint
	ReserveStatus   ReserveStatus `gorm:"references:id"`
}

//จบของนุ

// ของโอม

type Authority struct {
	gorm.Model
	Name string
	// 1 Authority สมารถมีอยู่ใน ClubMembership ได้หลายความสัมพันธ์
	ClubMembership []ClubMembership `gorm:"foreignKey: AuthorityID"`
}

type MembershipStatus struct {
	gorm.Model
	Name string
	// 1 MembershipStatus สามารถมีอยู่ใน ClubMembership ได้หลายความสัมพันธ์
	ClubMembership []ClubMembership `gorm:"foreignKey: MembershipStatusID"`
}

type ClubMembership struct {
	gorm.Model
	RegisterDate time.Time

	// UserID ทำหน้าที่เป็น FK
	StudentID *uint
	Student   Student

	// AuthorityID ทำหน้าที่เป็น FK
	AuthorityID *uint
	Authority   Authority

	// ClubID ทำหน้าที่เป็น FK
	ClubID *uint
	Club   Club

	// RequestStatusID ทำหน้าที่เป็น FK
	MembershipStatusID *uint
	MembershipStatus   MembershipStatus
}

// จบของโอม
