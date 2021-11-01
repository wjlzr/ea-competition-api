package configure

import (
	orm "ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"

	"go.uber.org/zap"
)

type Configure struct {
	ID                 int64 `gorm:"column:id" json:"id" form:"id"`
	Type               int64 `gorm:"column:type" json:"type" form:"type"`
	GlobalPrizePool    int64 `gorm:"column:global_prize_pool" json:"global_prize_pool" form:"global_prize_pool"`
	Trophy             int64 `gorm:"column:trophy" json:"trophy" form:"trophy"`
	Certificate        int64 `gorm:"column:certificate" json:"certificate" form:"certificate"`
	WarmUpTime         int64 `gorm:"column:warm_up_time" json:"warm_up_time" form:"warm_up_time"`
	SignUpStartTime    int64 `gorm:"column:sign_up_start_time" json:"sign_up_start_time" form:"sign_up_start_time"`
	SignUpEndTime      int64 `gorm:"column:sign_up_end_time" json:"sign_up_end_time" form:"sign_up_end_time"`
	EaStartTime        int64 `gorm:"column:ea_start_time" json:"ea_start_time" form:"ea_start_time"`
	EaEndTime          int64 `gorm:"column:ea_end_time" json:"ea_end_time" form:"ea_end_time"`
	RankAuditStartTime int64 `gorm:"column:rank_audit_start_time" json:"rank_audit_start_time" form:"rank_audit_start_time"`
	RankAuditEndTime   int64 `gorm:"column:rank_audit_end_time" json:"rank_audit_end_time" form:"rank_audit_end_time"`
	Announce           int64 `gorm:"column:announce" json:"announce" form:"announce"`
	NowTime            int64 `gorm:"-" json:"now_time" form:"now_time"`
	CreateAt           int64 `gorm:"autoCreateTime" json:"create_at" form:"create_at"`
	UpdateAt           int64 `gorm:"autoUpdateTime" json:"update_at" form:"update_at"`
}

func (*Configure) TableName() string {
	return "configure"
}

// FindOne 查询单个
func (m *Configure) FindOne(condition map[string]interface{}) (configure Configure, count int64, err error) {

	result := orm.Eloquent.Table(m.TableName()).Where(condition).FirstOrInit(&configure)
	if result.Error != nil {
		log.Logger().Error("configure FindOne FirstOrInit Err：", zap.Error(result.Error))
		return Configure{}, 0, result.Error
	}

	return configure, result.RowsAffected, nil
}
