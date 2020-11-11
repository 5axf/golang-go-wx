package queryTicket

import (
	"easy123/controller/stationNameCodeInfo"
	"easy123/util/dbConn"
	"easy123/util/goMail"
	"easy123/util/requestUtil"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

//
type QueryParam struct {
	//查询日期
	TrainDate	time.Time
	//出发地
	FromStation	string
	//目的地
	ToStation	string
	//乘客类型（成人:ADULT，学生:0X00）
	PurposeCodes	string
}


/*

返回数据分割后数据对应的下标

secretStr：0
备注：1
车票号：2
车次：3
start_station_code:起始站：4
end_station_code终点站：5
from_station_code:出发站：6
to_station_code:到达站：7
start_time:出发时间：8
arrive_time:达到时间：9
历时：10
能否购买：11 ， Y 可以
start_train_date:车票出发日期：13
高级软卧：21
其他：22
软卧：23
软座：24
无座：26
硬卧：28
硬座：29
二等座：30
一等座：31
商务特等座：32
动卧：33

*/
func QueryTicket(dateTime,form,to string,adult bool)  {
	var formCode stationNameCodeInfo.StationInfo
	err := dbConn.MysqlConn.Where("name_zh_cn = ?",form).First(&formCode).Error
	if err != nil {
		log.Fatal(err)
	}
	var toCode stationNameCodeInfo.StationInfo
	err = dbConn.MysqlConn.Where("name_zh_cn = ?",to).First(&toCode).Error
	if err != nil {
		log.Fatal(err)
	}
	var purposeCodes = "ADULT"
	if !adult {
		purposeCodes = "0X00"
	}
	param := QueryParam{
		FromStation: formCode.NameCode,
		ToStation:	toCode.NameCode,
		PurposeCodes: purposeCodes,
	}
	tf, _ := time.Parse("2006-01-02",dateTime)
	url := fmt.Sprintf("https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=%s&leftTicketDTO.from_station=%s&leftTicketDTO.to_station=%s&purpose_codes=%s",tf.Format("2006-01-02"),param.FromStation,param.ToStation,param.PurposeCodes)
	var cookie = "_uab_collina=159764962840525595771827; JSESSIONID=03EEF0CABE1D947B7DF4A54860E304A2; BIGipServerotn=501743882.64545.0000; RAIL_EXPIRATION=1597909354923; RAIL_DEVICEID=sp1dl2Qfzoz3YsqjvH9f3fiFqO237MFZM9uKDEqyF_GNeFGOk0M0G6ATgyb_v4kv56GxF-ABoLsYK8TxdikBSPlCHDlYf45kwjdrYzOoRodNlpxuYHtPUMcimOaIhCY3Xt9mJvqZJ-65GQT9Onu1N8w1_5M4OZWo; BIGipServerpassport=837288202.50215.0000; route=6f50b51faa11b987e576cdb301e545c4"
	responseData := requestUtil.HttpHandle("GET",url,cookie,"")
	var resultMap = make(map[string]interface{})
	json.Unmarshal(responseData,&resultMap)
	if resultMap["data"] != nil {
		data := resultMap["data"].(map[string]interface{})
		result := data["result"].([]interface{})
		mapData := data["map"].(map[string]interface{})
		fmt.Println("secretStr      车次    出发站   出发时间    到达站    到达时间    历时    硬座    是否能购买")
		for _, v := range result {
			s2 := v.(string)
			s3 := strings.Split(s2, "|")

			if s3[11] == "Y" && s3[29] != "" && (s3[3] == "Z182" || s3[3] == "Z332" || s3[3] == "D728") {
				//发送通知邮件
				//定义收件人
				mailTo := []string{
					"1318591661@qq.com",
				}
				//邮件主题为"Hello"
				subject := "火车票查询通知"
				// 邮件正文
				body := "你查询的："+dateTime+"，由【"+ mapData[s3[6]].(string)+"】【"+s3[8]+"】出发-->【"+s3[9]+"】到达【"+mapData[s3[7]].(string)+"】"+s3[3]+"次列车有余票可购买，请及时购买"
				//发送邮件
				er := goMail.SendMail(mailTo, subject, body)
				if er != nil {
					log.Println("邮件发送失败：",er)
					fmt.Println("send fail")
					return
				}
				fmt.Println("send successfully")
			}
			fmt.Println(s3[0],s3[3],s3[6], mapData[s3[6]],s3[8],s3[7], mapData[s3[7]],s3[9], s3[10], s3[29],s3[11])
		}
	}

}