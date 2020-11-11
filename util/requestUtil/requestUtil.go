package requestUtil

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//http请求
func HttpHandle(method, urlVal,cookie,body string) []byte {
	client := &http.Client{}
	var req *http.Request
	req, _ = http.NewRequest(method, urlVal, strings.NewReader(body))
	//添加cookie，key为X-Xsrftoken，value为df41ba54db5011e89861002324e63af81
	//可以添加多个cookie
	//cookie1 := &http.Cookie{Name: "Cookie",Value: "_uab_collina=159764962840525595771827; JSESSIONID=03EEF0CABE1D947B7DF4A54860E304A2; BIGipServerotn=501743882.64545.0000; RAIL_EXPIRATION=1597909354923; RAIL_DEVICEID=sp1dl2Qfzoz3YsqjvH9f3fiFqO237MFZM9uKDEqyF_GNeFGOk0M0G6ATgyb_v4kv56GxF-ABoLsYK8TxdikBSPlCHDlYf45kwjdrYzOoRodNlpxuYHtPUMcimOaIhCY3Xt9mJvqZJ-65GQT9Onu1N8w1_5M4OZWo; BIGipServerpassport=837288202.50215.0000; route=6f50b51faa11b987e576cdb301e545c4; _jc_save_fromStation=%u6DF1%u5733%2CSZQ; _jc_save_toStation=%u8D63%u5DDE%2CGZG; _jc_save_wfdc_flag=dc; _jc_save_toDate=2020-08-20; _jc_save_fromDate=2020-09-17", HttpOnly: true}
	//req.AddCookie(cookie1)
	req.Header.Set("accept","text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01")
	req.Header.Set("user-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	req.Header.Set("Cookie",cookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return b
}

func GetURLResponseContent(url string) (body string, errs []error) {
	request := gorequest.New()
	_, body, errs = request.Get(url).
		Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36").
		Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9").
		End()
	return
}