package awards

import (
	orm "ea-competition-api/boot/db/mysql"
	"ea-competition-api/boot/log"

	"go.uber.org/zap"
)

type Awards struct {
	ID          int64   `gorm:"column:id" json:"id" form:"id"`
	Rank        int64   `gorm:"column:rank" json:"rank" form:"rank"`
	Cash        float64 `gorm:"column:cash" json:"cash" form:"cash"`
	FirmCapital float64 `gorm:"column:firm_capital" json:"firm_capital" form:"firm_capital"`
	Vps         int64   `gorm:"column:vps" json:"vps" form:"vps"`
	Unit        int64   `gorm:"column:unit" json:"unit" form:"unit"`
	CreateAt    int64   `gorm:"autoCreateTime" json:"create_at" form:"create_at"`
	UpdateAt    int64   `gorm:"autoUpdateTime" json:"update_at" form:"update_at"`
}

func (*Awards) TableName() string {
	return "awards"
}

// FindOne 查询单个
func (m *Awards) Find() (awards []Awards, count int64, err error) {

	result := orm.Eloquent.Table(m.TableName()).Find(&awards)
	if result.Error != nil {
		log.Logger().Error("awards Find Find Err：", zap.Error(result.Error))
		return nil, 0, result.Error
	}

	return awards, result.RowsAffected, nil
}
