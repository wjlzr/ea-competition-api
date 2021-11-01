package img

import "ea-competition-api/config"

const (
	Images      = "images"
	ImagesRoute = "public/images/"
)

//获取图片地址
func GetImage(name, img string) string {
	return config.Conf().StaticResources.Url + img
}
