package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"time"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("schema.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&Club{},
		&Teacher{},
		&StudentCouncil{},
		&TypeClub{},
		&Activity{},
		&Student{},
		&ClubCommittee{},
		&JoinActivityHistory{},
		&BudgetCategory{},
		&BudgetType{},
		&BudgetProposal{},
		&Joining{},
		&Location{},
		&ReserveStatus{},
		&ReserveLocation{},
		&Authority{},
		&MembershipStatus{},
		&ClubMembership{},
		&Joinstatus{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//student council
	db.Model(&StudentCouncil{}).Create(&StudentCouncil{
		Name:       "นารา สิงห์ใจ",
		ID_Student: "B6223412",
		Password:   string(password),
	})

	db.Model(&StudentCouncil{}).Create(&StudentCouncil{
		Name:       "มีนา น่าอยู่",
		ID_Student: "B6178531",
		Password:   string(password),
	})

	db.Model(&StudentCouncil{}).Create(&StudentCouncil{
		Name:       "ปศิมา สาครศิลป์",
		ID_Student: "B6289123",
		Password:   string(password),
	})

	//teacher
	Teacher1 := Teacher{
		Name: "อ.ศิลกรรม ประยูร",
	}
	db.Model(&Teacher{}).Create(&Teacher1)
	Teacher2 := Teacher{
		Name: "อ.ประจักร วาติการ",
	}
	db.Model(&Teacher{}).Create(&Teacher2)
	Teacher3 := Teacher{
		Name: "อ.ดร.คาคง นิสันต์",
	}
	db.Model(&Teacher{}).Create(&Teacher3)
	Teacher4 := Teacher{
		Name: "อ.ปรนัย จงจำมั่น",
	}
	db.Model(&Teacher{}).Create(&Teacher4)
	Teacher5 := Teacher{
		Name: "อ.ประสาสี กลิ่นแก้ว",
	}
	db.Model(&Teacher{}).Create(&Teacher5)
	Teacher6 := Teacher{
		Name: "อ.ดร.ทรงประศรี ก้อนทอง",
	}
	db.Model(&Teacher{}).Create(&Teacher6)

	//type club
	sportsClubs := TypeClub{
		Name: "ด้านกีฬา",
	}
	db.Model(&TypeClub{}).Create(&sportsClubs)

	relationClubs := TypeClub{
		Name: "ด้านนักศึกษาสัมพันธ์",
	}
	db.Model(&TypeClub{}).Create(&relationClubs)

	socialClubs := TypeClub{
		Name: "ด้านพัฒนาสังคมและบำเพ็ญประโยชน์",
	}
	db.Model(&TypeClub{}).Create(&socialClubs)

	culturalClubs := TypeClub{
		Name: "ด้านศิลปวัฒนธรรม",
	}
	db.Model(&TypeClub{}).Create(&culturalClubs)

	academicClubs := TypeClub{
		Name: "ด้านวิชาการ",
	}
	db.Model(&TypeClub{}).Create(&academicClubs)

	// lady gnai krab
	//Lady ผู้โดนลบชื่อ ;-; - -*
	//--------------------------------------------<สร้าง club>
	//Club Data

	var adder1 StudentCouncil
	var adder2 StudentCouncil
	var adder3 StudentCouncil
	db.Raw("SELECT * FROM student_councils WHERE ID_Student = ?", "B6223412").Scan(&adder1)
	db.Raw("SELECT * FROM student_councils WHERE ID_Student = ?", "B6178531").Scan(&adder2)
	db.Raw("SELECT * FROM student_councils WHERE ID_Student = ?", "B6289123").Scan(&adder3)

	c1 := Club{
		Name:     "Astronomy Club",
		Adder:    adder1,
		Adviser:  Teacher1,
		TypeClub: academicClubs,
	}
	db.Model(&Club{}).Create(&c1)
	c2 := Club{
		Name:     "Football Club",
		Adder:    adder1,
		Adviser:  Teacher2,
		TypeClub: sportsClubs,
	}
	db.Model(&Club{}).Create(&c2)
	c3 := Club{
		Name:     "Computer Club",
		Adder:    adder2,
		Adviser:  Teacher3,
		TypeClub: academicClubs,
	}
	db.Model(&Club{}).Create(&c3)

	c4 := Club{
		Name:     "Sut Automotive Club",
		Adder:    adder3,
		Adviser:  Teacher4,
		TypeClub: academicClubs,
	}
	db.Model(&Club{}).Create(&c4)

	// แบงค์เอง
	var club1 Club
	var club2 Club
	var club3 Club
	var club4 Club
	db.Raw("SELECT * FROM clubs WHERE name = ?", "Astronomy Club").Scan(&club1)
	db.Raw("SELECT * FROM clubs WHERE name = ?", "Football Club").Scan(&club2)
	db.Raw("SELECT * FROM clubs WHERE name = ?", "Computer Club").Scan(&club3)
	db.Raw("SELECT * FROM clubs WHERE name = ?", "Sut Automotive Club").Scan(&club4)

	// แบงค์เอง

	//Committee data

	db.Model(&ClubCommittee{}).Create(&ClubCommittee{
		Name:       "Rinrada",
		ID_Student: "B6210533",
		Password:   string(password),
		Club:       c3,
	})
	db.Model(&ClubCommittee{}).Create(&ClubCommittee{
		Name:       "Bunyarit",
		ID_Student: "B6217082",
		Password:   string(password),
		Club:       c4,
	})
	astroCommittee1 := ClubCommittee{
		Name:       "Phanuwat",
		ID_Student: "B622222",
		Password:   string(password),
		Club:       c1,
	}
	db.Model(&ClubCommittee{}).Create(&astroCommittee1)

	var phanuwat ClubCommittee
	var rinrada ClubCommittee
	var bangk ClubCommittee
	db.Raw("SELECT * FROM club_committees WHERE id_student = ?", "B6210533").Scan(&rinrada)
	db.Raw("SELECT * FROM club_committees WHERE id_student = ?", "B6217082").Scan(&bangk)
	db.Raw("SELECT * FROM club_committees WHERE id_student = ?", "B622222").Scan(&phanuwat)

	//Activity data
	cForNewbie := Activity{
		Name:   "C#101 for Newbie",
		Time:   time.Now(),
		Amount: 120,
		Club:   c3,
	}
	db.Model(&Activity{}).Create(&cForNewbie)
	ps := Activity{
		Name:   "Pre-Coding",
		Time:   time.Now(),
		Amount: 100,
		Club:   c3,
	}
	db.Model(&Activity{}).Create(&ps)
	studentFormula := Activity{
		Name:   "What to know about Student Formula",
		Time:   time.Now(),
		Amount: 70,
		Club:   c4,
	}
	db.Model(&Activity{}).Create(&studentFormula)
	Astronomycamp := Activity{
		Name:   "Astronomy camp#1",
		Amount: 150,
		Club:   c1,
		Time:   time.Date(2019, 11, 19, 17, 30, 00, 000, time.UTC),
	}
	db.Model(&Activity{}).Create(&Astronomycamp)
	AstronomyDay := Activity{
		Name:   "Astronomy Day#1",
		Amount: 150,
		Club:   c1,
		Time:   time.Date(2020, 10, 19, 17, 30, 00, 000, time.UTC),
	}
	db.Model(&Activity{}).Create(&AstronomyDay)
	formulaRacing := Activity{
		Name:   "Formula Racing 2021",
		Time:   time.Now(),
		Amount: 70,
		Club:   c4,
	}
	db.Model(&Activity{}).Create(&formulaRacing)

	Astronomynight1 := Activity{
		Name:   "Astronomy Night#1",
		Amount: 150,
		Club:   c1,
		Time:   time.Date(2022, 11, 19, 17, 30, 00, 000, time.UTC),
	}
	db.Model(&Activity{}).Create(&Astronomynight1)
	SeeTheStar := Activity{
		Name:   "SeeTheStar",
		Amount: 30,
		Club:   c1,
		Time:   time.Now(),
	}
	db.Model(&Activity{}).Create(&SeeTheStar)

	Footballtourwatch1 := Activity{
		Name:   "Football Tourwatch#1",
		Amount: 1500,
		Club:   c1,
		Time:   time.Date(2022, 12, 12, 13, 00, 00, 000, time.UTC),
	}
	db.Model(&Activity{}).Create(&Footballtourwatch1)

	//Student data
	malisa := Student{
		Name:       "นางสาวมาลิสา โคนนท์",
		ID_Student: "B6122222",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&malisa)
	gaga := Student{
		Name:       "นางสาวกาก้า ปิรันย่า",
		ID_Student: "B6233333",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&gaga)
	chinatip := Student{
		Name:       "นายชินาธิป ชนะราวี",
		ID_Student: "B6210540",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&chinatip)
	punpun := Student{
		Name:       "นายปันปัน ปุนปุน",
		ID_Student: "B1234567",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&punpun)
	thana := Student{
		Name:       "นายธนะ การนา",
		ID_Student: "B6210541",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&thana)

	phimpa := Student{
		Name:       "นางพิมพา พาเเก้ว",
		ID_Student: "B6210542",
		Password:   string(password),
	}
	db.Model(&Student{}).Create(&phimpa)

	//JoinActivityHistory
	//History 1
	db.Model(&JoinActivityHistory{}).Create(&JoinActivityHistory{
		HourCount: 13,
		Point:     30,
		Timestamp: time.Now(),
		Activity:  cForNewbie,
		Student:   malisa,
		Editor:    rinrada,
	})
	//History 2
	db.Model(&JoinActivityHistory{}).Create(&JoinActivityHistory{
		HourCount: 13,
		Point:     30,
		Timestamp: time.Now(),
		Activity:  cForNewbie,
		Student:   gaga,
		Editor:    rinrada,
	})
	//History 3
	db.Model(&JoinActivityHistory{}).Create(&JoinActivityHistory{
		HourCount: 7,
		Point:     20,
		Timestamp: time.Now(),
		Activity:  studentFormula,
		Student:   gaga,
		Editor:    bangk,
	})

	// ของโอม

	// Authority Data
	// สิทธิ์ในชมรมที่ไม่ได้เป็นกรรมการบริหารชมรม
	clubMember := Authority{
		Name: "club member",
	}
	db.Model(&Authority{}).Create(&clubMember)

	clubSecretary := Authority{
		Name: "club secretary",
	}
	db.Model(&Authority{}).Create(&clubSecretary)

	clubVicePresident := Authority{
		Name: "club vice president",
	}
	db.Model(&Authority{}).Create(&clubVicePresident)

	clubPresident := Authority{
		Name: "club president",
	}
	db.Model(&Authority{}).Create(&clubPresident)

	// Membership Status Data
	// รอการอนุมัติเข้าชมรม
	pendingApproval := MembershipStatus{
		Name: "pending approval",
	}
	db.Model(&MembershipStatus{}).Create(&pendingApproval)

	// สามารถดำเนินกิจกรรมทางชมรมได้
	active := MembershipStatus{
		Name: "active",
	}
	db.Model(&MembershipStatus{}).Create(&active)

	// ถูกปฏิเสธการเข้าร่วมชมรม
	reject := MembershipStatus{
		Name: "reject",
	}
	db.Model(&MembershipStatus{}).Create(&reject)

	// ไม่สามารถดำเนินกิจกรรมทางชมรมได้
	inactive := MembershipStatus{
		Name: "inactive",
	}
	db.Model(&MembershipStatus{}).Create(&inactive)

	// Club Membership Data
	db.Model(&ClubMembership{}).Create(&ClubMembership{
		Student:          malisa,
		Authority:        clubMember,
		MembershipStatus: active,
		Club:             c3,
		RegisterDate:     time.Now(),
	})

	db.Model(&ClubMembership{}).Create(&ClubMembership{
		Student:          gaga,
		Authority:        clubPresident,
		MembershipStatus: active,
		Club:             c4,
		RegisterDate:     time.Now(),
	})

	db.Model(&ClubMembership{}).Create(&ClubMembership{
		Student:          phimpa,
		Authority:        clubMember,
		MembershipStatus: pendingApproval,
		Club:             c1,
		RegisterDate:     time.Now(),
	})

	// จบของโอม

	//อ๊อฟ
	j1 := Joinstatus{
		Name: "รอผลการเข้าร่วม",
	}
	db.Model(&Joinstatus{}).Create(&j1)

	j2 := Joinstatus{
		Name: "ได้เข้าร่วมเเล้ว",
	}
	db.Model(&Joinstatus{}).Create(&j2)

	j3 := Joinstatus{
		Name: "ยกเลิกการเข้าร่วมเเล้ว",
	}
	db.Model(&Joinstatus{}).Create(&j3)
	//

	// ของนุ
	lanmolum := Location{
		Name: "ลานหมอลำ",
	}
	db.Model(&Location{}).Create(&lanmolum)
	B5101 := Location{
		Name: "B5101",
	}
	db.Model(&Location{}).Create(&B5101)
	Samsan := Location{
		Name: "สระสามแสน",
	}
	db.Model(&Location{}).Create(&Samsan)
	Status1 := ReserveStatus{
		Label: "คำร้องรอการอนุมัติ",
	}

	db.Model(&ReserveStatus{}).Create(&Status1)
	Status2 := ReserveStatus{
		Label: "คำร้องได้รับการอนุมัติ",
	}
	db.Model(&ReserveStatus{}).Create(&Status2)
	Status3 := ReserveStatus{
		Label: "คำร้องไม่ได้รับการอนุมัติ",
	}
	db.Model(&ReserveStatus{}).Create(&Status3)
	db.Model(&ReserveLocation{}).Create(&ReserveLocation{
		DateStart:     time.Date(2021, 11, 18, 17, 30, 00, 000, time.UTC),
		DateEnd:       time.Date(2021, 11, 19, 17, 30, 00, 000, time.UTC),
		Location:      lanmolum,
		Activity:      Astronomynight1,
		Request:       phanuwat,
		ReserveStatus: Status2,
	})
	db.Model(&ReserveLocation{}).Create(&ReserveLocation{
		DateStart:     time.Date(2020, 10, 18, 17, 30, 00, 000, time.UTC),
		DateEnd:       time.Date(2020, 10, 19, 17, 30, 00, 000, time.UTC),
		Location:      lanmolum,
		Activity:      AstronomyDay,
		Request:       phanuwat,
		ReserveStatus: Status3,
	})
	//

	//แบงค์ ep2

	// budgetcategory Data
	cate1 := BudgetCategory{
		Name: "โภชนาการ",
	}
	db.Model(&BudgetCategory{}).Create(&cate1)

	cate2 := BudgetCategory{
		Name: "เครื่องมืออุปกรณ์",
	}
	db.Model(&BudgetCategory{}).Create(&cate2)
	cate3 := BudgetCategory{
		Name: "เบ็ดเตล็ด",
	}
	db.Model(&BudgetCategory{}).Create(&cate3)

	cate4 := BudgetCategory{
		Name: "การเดินทาง",
	}
	db.Model(&BudgetCategory{}).Create(&cate4)

	// budgettype Data
	money1 := BudgetType{
		Name: "เบิกเงินเต็มจำนวน",
	}
	db.Model(&BudgetType{}).Create(&money1)

	money2 := BudgetType{
		Name: "ทยอยเบิกเงิน",
	}
	db.Model(&BudgetType{}).Create(&money2)

	//Budgetps1
	db.Model(&BudgetProposal{}).Create(&BudgetProposal{
		Activity:       ps,
		BudgetType:     money1,
		BudgetCategory: cate4,
		BudgetPrice:    50000,
	})
	//Budgetps2
	db.Model(&BudgetProposal{}).Create(&BudgetProposal{
		Activity:       studentFormula,
		BudgetType:     money2,
		BudgetCategory: cate2,
		BudgetPrice:    1000,
	})

	//จบแบงค์ep2

}
