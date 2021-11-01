package apply

import (
	"ea-competition-api/app/model/mysql/apply"
	"ea-competition-api/boot/log"
	"ea-competition-api/services/skyeyecloud"
	"fmt"

	"github.com/k0kubun/pp"

	"go.uber.org/zap"
)

// CheckStatus 检查更新所有报名资格
func CheckStatus() {

	pp.Println("-----开始检查更新所有报名资格-----")

	var apply1 apply.Apply
	result, count, err := apply1.QueryList()
	pp.Println(result)
	if err != nil || count <= 0 {
		return
	}

	for _, v := range result {

		var request skyeyecloud.StatusRequest
		request.UserId = v.UserId
		request.PageSize = 1
		request.PageNum = 1

		// 先判断此人是否有参赛资格
		res, err := skyeyecloud.Status(request)
		if err != nil {
			continue
		}

		// 判断是否有参赛资格
		if res.Vps == 0 || res.Available1 <= 0 || res.Available3 == 1 {
			// 无参赛资格
			if v.Qualifications == 1 {
				apply1.Qualifications = 2
				if _, err = apply1.Update(map[string]interface{}{"user_id": v.UserId}); err != nil {
					fmt.Printf("用户ID: %s \n", v.UserId)
					log.Logger().Error("apply Update 更新状态失败 Err：", zap.Error(err))
					continue
				}
			}
		} else {
			if v.Qualifications == 2 {
				apply1.Qualifications = 1
				if _, err = apply1.Update(map[string]interface{}{"user_id": v.UserId}); err != nil {
					fmt.Printf("用户ID: %s \n", v.UserId)
					log.Logger().Error("apply Update 更新状态失败 Err：", zap.Error(err))
					continue
				}
			}
		}
	}

}
