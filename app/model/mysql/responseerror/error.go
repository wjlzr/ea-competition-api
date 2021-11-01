package responseerror

import (
	orm "ea-competition-api/boot/db/mysql"
)

type Error struct {
	Code string `json:"code"`
	ZhCN string `json:"zh-CN"`
	ZhHK string `json:"zh-HK"`
	ZhTW string `json:"zh-TW"`
	En   string `json:"en"`
	Vi   string `json:"vi"`
	Th   string `json:"th"`
	Fr   string `json:"fr"`
	Id   string `json:"id"`
	Es   string `json:"es"`
	Ru   string `json:"ru"`
	De   string `json:"de"`
	Tl   string `json:"tl"`
	It   string `json:"it"`
	Hi   string `json:"hi"`
	Ja   string `json:"ja"`
	Ko   string `json:"ko"`
	Pt   string `json:"pt"`
}

func GetError(code, lang string) string {

	var results []map[string]interface{}

	orm.Eloquent.Table("error").Where("code = ?", code).Find(&results)
	for _, val := range results {
		for k, v := range val {
			if k == lang {
				return v.(string)
			}
		}
	}
	return ""
}
