package stationNameCodeInfo

import (
	"easy123/util/dbConn"
	"easy123/util/requestUtil"
	"fmt"
	"log"
	"strings"
)

type StationInfo struct{
	Id        	uint `gorm:"primary_key"`
	AtName 		string
	NameZhCn 	string
	NameCode	string
	NameSpell	string
	NameSpellAd	string
}

func init()  {
	table := dbConn.MysqlConn.HasTable(&StationInfo{})
	if !table {
		dbConn.MysqlConn.CreateTable(&StationInfo{})
	}
}

//获取车站名与代码保存到数据库
func SaveInfo2db()  {
	body, errs := requestUtil.GetURLResponseContent("https://kyfw.12306.cn/otn/resources/js/framework/station_name.js?station_version=1.8968")
	if len(errs) != 0 {
	} else {
		if !strings.Contains(body,"|"){
			log.Println("数据查询出错")
			return
		}
		ss := strings.Split(body, "'")
		stationInfo := strings.Split(ss[1], "|")
		var sinfo []StationInfo
		var station []string
		for k, v := range stationInfo {
			station = append(station, v)
			k2 := (k + 1) % 5
			if k2 == 0 {
				//一个车站信息
				var s StationInfo
				s.AtName = station[0]
				s.NameZhCn = station[1]
				s.NameCode = station[2]
				s.NameSpell = station[3]
				s.NameSpellAd = station[4]

				//检测数据是否存在-存在则跳过
				var e StationInfo
				dbConn.MysqlConn.Where("name_code = ?",s.NameCode).First(&e)
				if e.Id != 0 {
					log.Println(e.Id)
					continue
				}

				err := dbConn.MysqlConn.Create(&s).Error
				if err != nil {
					fmt.Println("添加车站信息数据操作失败：",err)
				}
				sinfo = append(sinfo, s)

				//重置数组
				station = station[0:0]
			}
		}
		//批量保存信息到数据库
		//sql := "insert into station_name_code(id,at_name,name_zh_cn,name_code,name_spell,name_spell_ad) values"
		//for k, v := range sinfo {
		//	if len(sinfo)-1 == k {
		//		sql += fmt.Sprintf("(null,%s,%s,%s,%s,%s);", v.AtName, v.NameZhCn, v.NameCode, v.NameSpell, v.NameSpellAd)
		//	} else {
		//		sql += fmt.Sprintf("(null,%s,%s,%s,%s,%s),", v.AtName, v.NameZhCn, v.NameCode, v.NameSpell, v.NameSpellAd)
		//	}
		//}
		//err := dbConn.MysqlConn.Exec(sql).Error
		//if err != nil {
		//	fmt.Println("添加车站信息数据操作失败：",err)
		//}
	}
}