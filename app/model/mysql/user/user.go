package user

import (
	orm "ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"

	"go.uber.org/zap"
)

type User struct {
	ID          int64  `json:"id"`
	UserId      string `json:"user_id"`
	OpenId      string `json:"open_id"`
	AccountType string `json:"account_type"`
	AreaCode    string `json:"area_code"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CountryCode string `json:"country_code"`
	CreateAt    int64  `json:"create_at" xorm:"created"`
}

func (*User) TableName() string {
	return "user"
}

// Create 创建
func (m *User) Create() (*User, error) {

	result := orm.Eloquent.Table(m.TableName()).Create(&m)

	if result.Error != nil {
		log.Logger().Error("hotdata Create Create Err：", zap.Error(result.Error))
		return nil, result.Error
	}

	return m, nil
}
