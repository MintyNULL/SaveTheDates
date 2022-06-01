package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}
	// 테이블 자동 생성
	db.AutoMigrate(&Product{})

	// 생성
	db.Create(&Product{Code: "D42", Price: 100})

	// 읽기
	var product Product
	db.First(&product, 1)                 // primary key기준으로 product 찾기
	db.First(&product, "code = ?", "D42") // code가 D42인 product 찾기

	// 수정 - product의 price를 200으로
	db.Model(&product).Update("Price", 200)
	// 수정 - 여러개의 필드를 수정하기
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// 삭제 - product 삭제하기
	db.Delete(&product, 1)
}
