package main

import (
	//"easy123/controller/qzone"
	_ "easy123/controller/stationNameCodeInfo"
	_ "easy123/util/dbConn"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//qzone.GetImages()
	//router := gin.Default()

	//查询车站名及对应的代码，并保存进数据库
	//stationNameCodeInfo.SaveInfo2db()

	//获取验证码
	//login.GetCaptchaImage()

	//查询火车信息
	//queryTicket.QueryTicket("2020-09-25","深圳","赣州",true)

	//order.SubmitOrderRequest()

	//定时任务
	//c := cron.New()
	//c.AddFunc("*/3 * * * * *", func() {
	//	queryTicket.QueryTicket("2020-10-01","深圳","赣州",true)
	//})
	//c.Start()

	//端口
	//router.Run(viper.GetString("port"))
}
