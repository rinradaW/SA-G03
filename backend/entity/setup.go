package entity

import (
	//"fmt"
	"time"
	
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("realr.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(
		&Club{}, &Activity{}, &Student{}, &ClubCommittee{}, &JoinActivityHistory{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)


	//Club Data
	computerClub := Club{
		Name:  "Computer Club",
	}
	db.Model(&Club{}).Create(&computerClub)

	automotiveClub := Club{
		Name:  "Sut Automotive Club",
	}
	db.Model(&Club{}).Create(&automotiveClub)

	/*db.Model(&Club{}).Create(&Club{
		Name:  "Computer Club",
	})
	db.Model(&Club{}).Create(&Club{
		Name:  "Sut Automotive Club",
	})*/


	//Committee data
	db.Model(&ClubCommittee{}).Create(&ClubCommittee{
		Name:  		"Rinrada",
		ID_Student: "B6210533",
		Password:	string(password),
		Club:		computerClub,
	})
	db.Model(&ClubCommittee{}).Create(&ClubCommittee{
		Name:  		"Name",
		ID_Student: "B6217092",
		Password:	string(password),
		Club:		automotiveClub,
	})

	var rinrada ClubCommittee
	var name ClubCommittee
	db.Raw("SELECT * FROM club_committees WHERE id_student = ?", "B6210533").Scan(&rinrada)
	db.Raw("SELECT * FROM club_committees WHERE id_student = ?", "B6217092").Scan(&name)


	//Activity data
	cForNewbie := Activity{
		Name:	"C#101 for Newbie",
		Time: 	time.Now(),
		Amount:	120,
		Club:	computerClub,
	}
	db.Model(&Activity{}).Create(&cForNewbie)
	studentFormula := Activity{
		Name:	"What to know about Student Formula",
		Time: 	time.Now(),
		Amount:	50,
		Club:	automotiveClub,
	}
	db.Model(&Activity{}).Create(&studentFormula)


	//Student data
	malisa := Student{
		Name:			"Malisa",
		ID_Student: 	"B6122222",
	}
	db.Model(&Student{}).Create(&malisa)
	gaga := Student{
		Name:			"Gaga",
		ID_Student: 	"B6233333",
	}
	db.Model(&Student{}).Create(&gaga)


	//JoinActivityHistory
	//History 1
	db.Model(&JoinActivityHistory{}).Create(&JoinActivityHistory{
		HourCount:	 13,
		Point:		 30,
		Timestamp:   time.Now(),
		Activity:    cForNewbie,
		Student:	 malisa,
		Editor:		 rinrada,
	})
	//History 2
	db.Model(&JoinActivityHistory{}).Create(&JoinActivityHistory{
		HourCount:	 13,
		Point:		 30,
		Timestamp:   time.Now(),
		Activity:    cForNewbie,
		Student:	 gaga,
		Editor:		 rinrada,
	})
	//History 3
	db.Model(&JoinActivityHistory{}).Create(&JoinActivityHistory{
		HourCount:	 7,
		Point:		 20,
		Timestamp:   time.Now(),
		Activity:    studentFormula,
		Student:	 gaga,
		Editor:		 name,
	})


	// ==== Query ====
	/*var target ClubCommittee
	db.Model(&ClubCommittee{}).Find(&target, db.Where("id_student = ?", "B6210533"))

	var activityHistoryList []*JoinActivityHistory
	db.Model(&JoinActivityHistory{}).
		Joins("Activity").
		Joins("Student").
		Joins("Editor").
		Find(&activityHistoryList, db.Where("editor_id = ?", target.ID))


	for _, ahl := range activityHistoryList {
		fmt.Printf("Joined Activity History: %v\n", ahl.ID)
		fmt.Printf("%v\n", ahl.Activity.Name)
		fmt.Printf("%v\n", ahl.Student.Name)
		fmt.Printf("%v\n", ahl.Editor.Name)
		fmt.Println("====")
	}*/
}
