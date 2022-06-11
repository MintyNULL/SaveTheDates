package main

import (
	"SaveTheDates/crawler/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Date struct {
	gorm.Model
	Title   string
	Year    int8
	Month   int8
	Day     int8
	EnTitle string
}

func main() {
	db, err := gorm.Open(sqlite.Open("STD.db"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	http.Test()
	// 테이블 자동 생성
	db.AutoMigrate(&Date{})

	// 생성
	db.Create(&Date{Code: "D42", Price: 100})

	// 읽기
	var date Date
	db.First(&date, 1)                 // primary key기준으로 product 찾기
	db.First(&date, "code = ?", "D42") // code가 D42인 product 찾기

	// 수정 - product의 price를 200으로
	db.Model(&date).Update("Price", 200)
	// 수정 - 여러개의 필드를 수정하기
	db.Model(&date).Updates(Date{Price: 200, Code: "F42"})
	db.Model(&date).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 삭제 - product 삭제하기
	db.Delete(&date, 1)
}
