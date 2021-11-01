package apply

import (
	orm "ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"

	"go.uber.org/zap"
)

type Apply struct {
	ID             int64  `gorm:"column:id" json:"id" form:"id"`
	UserId         string `gorm:"column:user_id" json:"user_id" form:"user_id"`
	AreaCode       string `gorm:"-" json:"area_code" form:"area_code"`
	Phone          string `gorm:"-" json:"phone" form:"phone"`
	Qualifications int64  `gorm:"column:qualifications;default:2" json:"qualifications" form:"qualifications"`
	CreateAt       int64  `gorm:"autoCreateTime" json:"create_at" form:"create_at"`
	UpdateAt       int64  `gorm:"autoUpdateTime" json:"update_at" form:"update_at"`
}

func (*Apply) TableName() string {
	return "apply"
}

// QueryList 查询全部
func (m *Apply) QueryList() (apply []Apply, count int64, err error) {

	result := orm.Eloquent.Table(m.TableName()).Find(&apply)

	if result.Error != nil {
		log.Logger().Error("apply QueryList Find Err：", zap.Error(result.Error))
		return nil, 0, result.Error
	}

	return apply, result.RowsAffected, nil
}

// Create 创建
func (m *Apply) Create() (*Apply, error) {

	result := orm.Eloquent.Table(m.TableName()).Create(&m)
	if result.Error != nil {
		log.Logger().Error("apply Create Create Err：", zap.Error(result.Error))
		return nil, result.Error
	}

	return m, nil
}

// FindOne 查询单个
func (m *Apply) FindOne() (apply *Apply, count int64, err error) {

	result := orm.Eloquent.Table(m.TableName()).Where(m).FirstOrInit(&apply)
	if result.Error != nil {
		log.Logger().Error("apply FindOne FirstOrInit Err：", zap.Error(result.Error))
		return nil, 0, result.Error
	}

	return apply, result.RowsAffected, nil
}

// Update 编辑
func (m *Apply) Update(condition map[string]interface{}) (apply *Apply, err error) {

	result := orm.Eloquent.Table(m.TableName()).Where(condition).Updates(&m)
	if result.Error != nil {
		log.Logger().Error("apply Update Updates Err：", zap.Error(result.Error))
		return nil, result.Error
	}

	return m, nil
}

// QueryList 查询全部
func (m Apply) QueryStatistics(startTime, endTime, qualifications int64) (Apply []Apply, err error) {

	result := orm.Eloquent.Table(m.TableName()).Where("create_at >= ? and create_at < ? and qualifications = ?", startTime, endTime, qualifications).Find(&Apply)

	if result.Error != nil {
		log.Logger().Error("apply QueryStatistics Pluck Err：", zap.Error(result.Error))
		return nil, result.Error
	}

	return Apply, nil
}
