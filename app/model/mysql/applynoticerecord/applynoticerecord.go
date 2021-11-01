package applynoticerecord

import (
	orm "ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"

	"go.uber.org/zap"
)

type ApplyNoticeRecord struct {
	ID            int64  `gorm:"column:id" json:"id" form:"id"`
	UserId        string `gorm:"column:user_id" json:"user_id" form:"user_id"`
	AreaCode      string `gorm:"column:area_code" json:"area_code" form:"area_code"`
	Communication string `gorm:"column:communication" json:"communication" form:"communication"`
	Type          int64  `gorm:"column:type" json:"type" form:"type"`
	CountryCode   string `gorm:"column:country_code" json:"country_code" form:"country_code"`
	Result        int64  `gorm:"column:result" json:"result" form:"result"`
	RetCode       int64  `gorm:"column:ret_code" json:"ret_code" form:"ret_code"`
	Message       string `gorm:"column:message" json:"message" form:"message"`
	SessionNo     string `gorm:"column:session_no" json:"session_no" form:"session_no"`
	CreateAt      int64  `gorm:"autoCreateTime" json:"create_at" form:"create_at"`
}

func (*ApplyNoticeRecord) TableName() string {
	return "apply_notice_record"
}

// Create 创建
func (m *ApplyNoticeRecord) Create() (*ApplyNoticeRecord, error) {

	result := orm.Eloquent.Table(m.TableName()).Create(&m)
	if result.Error != nil {
		log.Logger().Error("applynoticerecord Create Create Err：", zap.Error(result.Error))
		return nil, result.Error
	}

	return m, nil
}
