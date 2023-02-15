package databases

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("./api/data/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed Connected databases")
	}
	DBConn.AutoMigrate(&TUser{}, &TNodePoint{})

	//插入默认的数据
	userInfo := QueryUser("admin")
	if "" == userInfo.UserName {
		AddUser(&TUser{
			UserName: "admin",
			PassWord: "admin",
		})
	}
}
