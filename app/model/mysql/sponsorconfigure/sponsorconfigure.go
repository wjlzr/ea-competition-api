package sponsorconfigure

import (
	orm "ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"

	"go.uber.org/zap"
)

type SponsorConfigure struct {
	ID        int64  `gorm:"column:id" json:"id" form:"id"`
	SponsorId int64  `gorm:"column:sponsor_id" json:"sponsor_id" form:"sponsor_id"`
	Name      string `gorm:"column:name" json:"name" form:"name"`
	Type      int64  `gorm:"column:type" json:"type" form:"type"`
	Award     string `gorm:"column:award" json:"award" form:"award"`
	Marks     string `gorm:"column:marks" json:"marks" form:"marks"`
	CreateAt  int64  `gorm:"autoCreateTime" json:"create_at" form:"create_at"`
	UpdateAt  int64  `gorm:"autoUpdateTime" json:"update_at" form:"update_at"`
}

func (*SponsorConfigure) TableName() string {
	return "sponsor_configure"
}

// FindOne 查询单个
func (m *SponsorConfigure) FindOne(condition map[string]interface{}) (sponsorConfigure []SponsorConfigure, count int64, err error) {

	result := orm.Eloquent.Table(m.TableName()).Where(condition).Find(&sponsorConfigure)
	if result.Error != nil {
		log.Logger().Error("sponsorconfigure FindOne FirstOrInit Err：", zap.Error(result.Error))
		return nil, 0, result.Error
	}

	return sponsorConfigure, result.RowsAffected, nil
}
