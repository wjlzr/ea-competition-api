package sponsor

import (
	"ea-competition-api/app/model/mysql/sponsorconfigure"
	orm "ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"

	"go.uber.org/zap"
)

type Sponsor struct {
	ID        int64  `gorm:"column:id" json:"id" form:"id"`
	Name      string `gorm:"column:name" json:"name" form:"name"`
	ImgPc     string `gorm:"column:img_pc" json:"img_pc" form:"img_pc"`
	ImgM      string `gorm:"column:img_m" json:"img_m" form:"img_m"`
	Type      int64  `gorm:"column:type" json:"type" form:"type"`
	Sort      int64  `gorm:"column:sort" json:"sort" form:"sort"`
	Status    int64  `gorm:"column:status" json:"status" form:"status"`
	IsRank    int64  `gorm:"column:is_rank" json:"is_rank" form:"is_rank"`
	IconS     string `gorm:"column:icon_s" json:"icon_s" form:"icon_s"`
	IconB     string `gorm:"column:icon_b" json:"icon_b" form:"icon_b"`
	TradeCode string `gorm:"column:trade_code" json:"trade_code" form:"trade_code"`
	CreateAt  int64  `gorm:"autoCreateTime" json:"create_at" form:"create_at"`
	UpdateAt  int64  `gorm:"autoUpdateTime" json:"update_at" form:"update_at"`
}

type SponsorResponse struct {
	ID               int64  `gorm:"column:id" json:"id" form:"id"`
	Name             string `gorm:"column:name" json:"name" form:"name"`
	ImgPc            string `gorm:"column:img_pc" json:"img_pc" form:"img_pc"`
	ImgM             string `gorm:"column:img_m" json:"img_m" form:"img_m"`
	Type             int64  `gorm:"column:type" json:"type" form:"type"`
	Sort             int64  `gorm:"column:sort" json:"sort" form:"sort"`
	Status           int64  `gorm:"column:status" json:"status" form:"status"`
	IsRank           int64  `gorm:"column:is_rank" json:"is_rank" form:"is_rank"`
	IconS            string `gorm:"column:icon_s" json:"icon_s" form:"icon_s"`
	IconB            string `gorm:"column:icon_b" json:"icon_b" form:"icon_b"`
	TradeCode        string `gorm:"column:trade_code" json:"trade_code" form:"trade_code"`
	SponsorConfigure []sponsorconfigure.SponsorConfigure
	CreateAt         int64 `gorm:"autoCreateTime" json:"create_at" form:"create_at"`
	UpdateAt         int64 `gorm:"autoUpdateTime" json:"update_at" form:"update_at"`
}

func (*Sponsor) TableName() string {
	return "sponsor"
}

// FindOne 查询单个
func (m *Sponsor) Find(condition map[string]interface{}) (sponsors []Sponsor, count int64, err error) {

	result := orm.Eloquent.Table(m.TableName()).Where(condition).Find(&sponsors)
	if result.Error != nil {
		log.Logger().Error("awards Find Find Err：", zap.Error(result.Error))
		return nil, 0, result.Error
	}

	return sponsors, result.RowsAffected, nil
}
