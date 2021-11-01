package apk

import (
	"ea-competition-api/boot/log"

	"github.com/lunny/axmlParser"
	"go.uber.org/zap"
)

func GetApkVersion(filePath string) (v string, err error) {
	listener := new(axmlParser.AppNameListener)
	_, err = axmlParser.ParseApk(filePath, listener)
	if err != nil {
		log.Logger().Error("apk GetApkVersion Err：", zap.Error(err))
		return "", err
	}
	return listener.VersionName, nil
}
