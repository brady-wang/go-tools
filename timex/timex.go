package timex

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type TimeFormat struct {
}

// NowFormat 当前时间的字符串
func (t *TimeFormat) NowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// NowSqlNullTime 当前时间的 sql.NULLTime
func (t *TimeFormat) NowSqlNullTime() sql.NullTime {
	return sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

// StringToTime 字符串格式的 2020-01-01 00:00:00 转换为time类型
func (t *TimeFormat) StringToTime(timeString string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return time.Time{}, errors.New("时区获取失败")
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", timeString, loc)
	if err != nil {
		fmt.Printf("时间转换失败 %s \n",timeString)
		return time.Time{}, errors.New("转换失败")
	}
	return timeObj, nil
}

// StringToSqlNullTime 字符串格式的 2020-01-01 00:00:00 转换为sql.NullTime格式 存入数据库
func (t *TimeFormat) StringToSqlNullTime(timeString string) sql.NullTime {

	if timeString == "0000-00-00 00:00:00" {
		return sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	}
	var res sql.NullTime
	if len(timeString) <= 0 {
		res = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	} else {
		timeRes, err := t.StringToTime(timeString)
		if err != nil {
			res = sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			}
		} else {
			res = sql.NullTime{
				Time:  timeRes,
				Valid: true,
			}
		}
	}

	return res
}

// SqlNullTimeToString sql.NullTime查询出来转为字符串
func (t *TimeFormat) SqlNullTimeToString(sqlTime sql.NullTime) string {
	var res string
	if sqlTime.Valid {
		res = sqlTime.Time.Format("2006-01-02 15:04:05")
	} else {
		res = ""
	}
	return res
}
