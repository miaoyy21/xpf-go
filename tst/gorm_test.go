package tst

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"testing"
)

type Score struct {
	StudentID string
	Score     int
}

type Student struct {
	ID   string
	Name string

	Score *Score `gorm:"embedded"`
	Nil   *sql.NullInt32
}

func Test(t *testing.T) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&loc=UTC&parseTime=true", "root", "root", "127.0.0.1", "3306", "tst")
	db, err := gorm.Open("mysql", source)
	if err != nil {
		logrus.Fatalf("gorm.Open() Failure :: %s", err.Error())
		return
	}
	db.LogMode(true)

	if db.HasTable(&Student{}) {
		db.DropTable(&Student{})
	}
	db.CreateTable(&Student{})

	db.Create(&Student{ID: "1001", Name: "test01",Score: &Score{StudentID: "1111",Score: 65}})
	db.Create(&Student{ID: "1002", Name: "test02"})
	db.Create(&Student{ID: "1003", Name: "test03"})

	result := Student{}
	//db.Table("students").First(&result)
	db.First(&result)
	db.First(&result)

	t.Logf("Student is %#v\n", result)

}
