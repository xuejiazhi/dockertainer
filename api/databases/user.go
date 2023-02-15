package databases

type TUser struct {
	UserName string `gorm:"column(user_name);size(32);primary_key"`
	PassWord string `gorm:"column(pass_word);size(32)"`
}

// QueryUser 查询用户信息
func QueryUser(userName string) (userInfo TUser) {
	//插入默认的数据
	DBConn.First(&userInfo, "user_name=?", userName)
	return
}

// AddUser 新增用户信息
func AddUser(userInfo *TUser) error {
	return DBConn.Create(userInfo).Error
}

// ModifyPassword 更新密码
func ModifyPassword(user, password string) error {
	tUser := TUser{}
	DBConn.Model(tUser).Where("user_name=?", user).Take(&tUser)
	tUser.PassWord = password
	return DBConn.Save(&tUser).Error
}
