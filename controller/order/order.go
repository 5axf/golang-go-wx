package order

import (
	"easy123/util/requestUtil"
	"encoding/json"
	"fmt"
	"log"
	_ "net/url"
	url2 "net/url"
)

func SubmitOrderRequest()  {

	value,err := url2.ParseQuery("NEyLWQ%2FRvGVPhce%2F1j3QJuhpyMrc%2Fx01mAgNwEIlLddXrmoBlP0yhzSDMAlWxTXN%2FN4uOvxoCAiN%0AAL0nq6T14XotFFlM1u%2BplTLE3F3%2FlwAQ3momG1SsNwcf8g3nWvZARVvN2tC4RVo4Us0NNgfQ%2F8sZ%0AANPwkrUHXy%2BLofzocAdGU1xNKn%2BQgaM2%2F5cYu30WMo7XYmHLY6SwewJJk72HGmUCndSU5GqDpJGQ%0ABQOENvovuom78fEfL1loQ3UCOhSIRm%2F7zq40Nac0K3dFPdN7LKRclwOHo16CSR4MH1a%2BipT2%2BeNn%0A7Z8mvQ%3D%3D")
	if err != nil {
		log.Println(err)
	}
	var secretStr = ""
	for k,_ := range value{
		secretStr = k
	}
	url := fmt.Sprintf("https://kyfw.12306.cn/otn/leftTicket/submitOrderRequest")
	var cookie = "JSESSIONID=B04265B9AEAC7ED9809CDBEF98DAE0F6; tk=wBIJbcETTE02O7h1Wcmv3OmRuyn5plaQRS1AJAwea1a0; BIGipServerotn=501743882.64545.0000; BIGipServerpassport=837288202.50215.0000; route=6f50b51faa11b987e576cdb301e545c4; ten_key=1DK2nRKQfXWkCiLKAma5HODTecNTXIR9; ten_js_key=1DK2nRKQfXWkCiLKAma5HODTecNTXIR9; _jc_save_wfdc_flag=dc; _jc_save_toStation=%u8D63%u5DDE%2CGZG; RAIL_EXPIRATION=1598849580895; RAIL_DEVICEID=i_K94c9mJApT_XggbW6gfqPsNdxRnVDuPRADAZHVCAvUbrFCX6My0EJBPkVUvtOWjlOUgAN9PJj-xH2CeZCKR06vFAkb-a1mDHv1FQz2m7xpGFSWzvVfBLYLVKCioCmwV2iwKsbkUOmWU5aIp86muQxrrISE_8Uh; uKey=4e4b9b2c80a892f0f54318477579501ff94dd99d6676cab5a7c47f051619ff4a; current_captcha_type=Z; BIGipServerportal=3151233290.17183.0000; _jc_save_fromStation=%u6DF1%u5733%2CSZQ; _jc_save_toDate=2020-08-28; _jc_save_fromDate=2020-09-25"

	var requestParam = make(map[string]interface{})
	// 查询车票时返回结果索引为0的数据
	requestParam["secretStr"] = secretStr
	// 出发日期
	requestParam["train_date"] = "2020-09-24"
	// 当前查询日期（如果是往返票，则为返程日期，也即是出发日期）
	requestParam["back_train_date"] = "2020-09-24"
	// 车票类型（dc：单程   wc：往返）
	requestParam["tour_flag"] = "dc"
	// 成人票
	requestParam["purpose_codes"] = "ADULT"
	// 出发站
	requestParam["query_from_station_name"] = "深圳"
	// 到达站
	requestParam["query_to_station_name"] = "赣州"

	bodyBety,err := json.Marshal(requestParam)
	if err != nil {
		log.Fatal(err)
	}
	responseData := requestUtil.HttpHandle("POST",url,cookie, string(bodyBety))
	var resultMap = make(map[string]interface{})
	json.Unmarshal(responseData,&resultMap)
	//if resultMap["result_code"].(string) == "0" {
	//}
}