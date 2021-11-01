package configure

import (
	"ea-competition-api/app/model/mysql/configure"
	"errors"
	"time"
)

// FindConfigure
func FindConfigure() (*configure.Configure, error) {

	var s configure.Configure

	// 检查账号是否已存在
	r, count, err := s.FindOne(map[string]interface{}{"id": 1})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("500")
	}

	r.NowTime = time.Now().Unix()

	return &r, nil
}
