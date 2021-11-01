package sponsor

import (
	"ea-competition-api/app/model/mysql/sponsor"
	"ea-competition-api/app/model/mysql/sponsorconfigure"
	"encoding/json"
	"errors"
)

// FindSponsor
func FindSponsor() (resp []sponsor.SponsorResponse, err error) {

	var s sponsor.Sponsor
	var sc sponsorconfigure.SponsorConfigure

	r, count, err := s.Find(map[string]interface{}{"status": 1})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("500")
	}
	jsonStu, err := json.Marshal(r)
	_ = json.Unmarshal(jsonStu, &resp)
	// 查询赞助商所提供的奖品
	for k, v := range resp {
		r2, count1, err := sc.FindOne(map[string]interface{}{"sponsor_id": v.ID})
		if err != nil || count1 == 0 {
			continue
		}
		resp[k].SponsorConfigure = r2
	}

	return resp, nil
}
