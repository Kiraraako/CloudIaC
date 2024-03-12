// Copyright (c) 2015-2023 CloudJ Technology Co., Ltd.

package models

import (
	"cloudiac/portal/libs/db"
)

type User struct {
	SoftDeleteModel

	Name         string `json:"name" gorm:"size:32;not null;comment:姓名" example:"张三"`                                         // 姓名
	Email        string `json:"email" gorm:"size:64;not null;comment:邮箱" example:"mail@example.com"`                          // 邮箱
	Password     string `json:"-" gorm:"not null;comment:密码" swaggerignore:"true"`                                            // 密码
	Phone        string `json:"phone" gorm:"size:16;comment:电话" example:"18888888888"`                                        // 电话
	IsAdmin      bool   `json:"isAdmin" gorm:"default:false;comment:是否为系统管理员" example:"false"`                                // 是否为系统管理员
	Status       string `json:"status" gorm:"default:'enable';comment:用户状态" enums:"enable,disable" example:"enable"`          // 用户状态 type:enum('enable','disable');
	NewbieGuide  JSON   `json:"newbieGuide" gorm:"type:text;null;comment:新手引导状态" swaggertype:"string" example:"{\"1\"}"`      // 新手引导状态
	ActiveStatus string `json:"activeStatus" gorm:"default:'active';comment:用户激活状态" enums:"active,inactive" example:"active"` // 用户状态 type:enum('active','inactive');

	Company string `json:"company"`
}

func (User) TableName() string {
	return "iac_user"
}

func (u User) Migrate(sess *db.Session) (err error) {
	err = u.AddUniqueIndex(sess, "unique__email", "email")
	if err != nil {
		return err
	}

	return nil
}
