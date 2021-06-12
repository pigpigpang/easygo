package entity

import "gorm.io/gorm"

// User users表模型, 模型采用 gorm 自动迁移生成, 不推荐直接通过 sql 修改表结构
// 各字段含义如下:
//   Username : 手机注册的账号，默认与手机号绑定
//   Password: 用户密码
//   Status: 用户状态 (是否被管理员封禁) 1活跃 9封禁

//以下是个人信息

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(255);not null"`
	Personal string `gorm:"type:varchar(255);not null"`
}

func (u *User) TableName() string {
	return "easy_go_user"
}
