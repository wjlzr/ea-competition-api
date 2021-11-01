package geo

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

//根据ip获取国家代码
func GetCountryCode(ipAddress string) string {
	if geoDB, err := geoip2.Open("./public/GeoLite2-Country.mmdb"); err == nil {
		defer geoDB.Close()
		ip := net.ParseIP(ipAddress)
		if record, err := geoDB.Country(ip); err == nil {
			return record.Country.IsoCode
		}
	}
	return ""
}
