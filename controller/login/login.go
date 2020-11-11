package login

import (
	"easy123/util/requestUtil"
	"encoding/json"
	"fmt"
	"time"
)

//获取图片验证码
func GetCaptchaImage()  {
	url := fmt.Sprintf("https://kyfw.12306.cn/passport/captcha/captcha-image64?login_site=E&module=login&rand=sjrand&1598236991557&_=%s",time.UnixDate)
	var cookie = "_passport_ct=10c3fce35df14b7899899f31fa42276et7056; _passport_session=b0c0585f9c014b17a8cfa9aa6f79b6f13574; BIGipServerotn=501743882.64545.0000; BIGipServerpassport=837288202.50215.0000; route=6f50b51faa11b987e576cdb301e545c4; _jc_save_fromStation=%u6DF1%u5733%2CSZQ; _jc_save_toStation=%u8D63%u5DDE%2CGZG; RAIL_EXPIRATION=1598528953802; RAIL_DEVICEID=kJA82zSLEOanJizU6jSL2iOr9sPkum-1zHDpi5LpHQBLzX0hk0ZOia_p6isTTfGrhQr-3kOnU4jtksv8joMekMU-deOlEKse1oo5JQ2jjcJopdzbeFVrx5embiHoOpxpqk1E7FtZ8XnZzGD6Iqi4O6I4RUKkoNrn; _jc_save_fromDate=2020-09-22; _jc_save_toDate=2020-09-22; _jc_save_wfdc_flag=wf; ten_key=1DK2nRKQfXWkCiLKAma5HODTecNTXIR9; ten_js_key=1DK2nRKQfXWkCiLKAma5HODTecNTXIR9"
	responseData := requestUtil.HttpHandle("GET",url,cookie,"")
	var resultMap = make(map[string]interface{})
	json.Unmarshal(responseData,&resultMap)
	if resultMap["result_code"].(string) == "0" {
		//验证码图片
		//fmt.Println("data:image/png;base64,"+resultMap["image"].(string))
	}

	//ciphertxt,err := sm4.Sm4Encrypt([]byte("tiekeyuankp12306"),[]byte("123456"))
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Printf("加密结果: %x\n", ciphertxt)

}
