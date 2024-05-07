package utils

import (
	"time"

	"github.com/sirupsen/logrus"
)

const defaultTimeFormat = `2006-01-02 15:04:34`

var timeLocation, _ = time.LoadLocation("Asia/Shanghai")

func Transfer(v interface{}) int64 {
	var userId int64
	if v, ok := v.(float64); ok {
		userId = int64(v)
	} else {
		logrus.Info("类型断言失败,无法将变量转换为string类型")
	}
	return userId
}

func ConvertTimestampToString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(defaultTimeFormat)
}

func ConvertStringToTimestampDefault(date string) int64 {
	t, _ := time.ParseInLocation(defaultTimeFormat, date, timeLocation)
	return t.Unix()
}
func StringToUnixTime(timeStr string) (int64, error) {
    // 将字符串类型的时间转换为 time.Time 类型
    t, err := time.Parse(time.RFC3339, timeStr)
    if err != nil {
        return 0, err
    }
    // 将 time.Time 类型的时间转换为 UNIX 时间戳
    unixTime := t.Unix()
    return unixTime, nil
}