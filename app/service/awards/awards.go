package awards

import (
	"ea-competition-api/app/model/mysql/awards"
	"errors"
)

// FindAwards
func FindAwards() (*[]awards.Awards, error) {

	var s awards.Awards

	// 检查账号是否已存在
	r, count, err := s.Find()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("500")
	}

	return &r, nil
}
