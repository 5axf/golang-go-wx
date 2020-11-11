package test

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 公众号绑定url接口
func BindUrl(c *gin.Context)  {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	println(signature)
	println(timestamp)
	println(nonce)
	println(echostr)

	// 忽略校验过程

	c.String(200,echostr)
}

// 微信推送返回信息实体类
type Info struct {
	XMLName  		xml.Name 	`xml:"xml"`
	ToUserName 		string 		`xml:"ToUserName"`		// 	开发者微信号
	FromUserName 	string 		`xml:"FromUserName"`	// 	发送方帐号（一个OpenID）
	CreateTime 		string 		`xml:"CreateTime"`		// 	消息创建时间 （整型）
	MsgType 		string 		`xml:"MsgType"`			//	消息类型，event
	Event 			string 		`xml:"Event"`			// 	事件类型，subscribe(订阅)、unsubscribe(取消订阅)
	EventKey 		string 		`xml:"EventKey"`		// 	事件KEY值，qrscene_为前缀，后面为二维码的参数值
}

// 公众号用户信息实体类
type UserInfo struct {
	Subscribe 			string 		`json:"subscribe"`			//	用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid 				string 		`json:"openid"`				//	用户的标识，对当前公众号唯一
	Nickname 			string 		`json:"nickname"`			//	用户的昵称
	Sex 				int 		`json:"sex"`				//	用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language 			string 		`json:"language"`			//	用户所在城市
	City 				string 		`json:"city"`				//	用户所在国家
	Province 			string 		`json:"province"`			//	用户所在省份
	Country 			string 		`json:"country"`			//	用户的语言，简体中文为zh_CN
	Headimgurl 			string 		`json:"headimgurl"`			//	用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime 		string 		`json:"subscribe_time"`		//	用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid 			string 		`json:"unionid"`			//	只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark 				string 		`json:"remark"`				//	公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid 			int 		`json:"groupid"`			//	用户所在的分组ID（兼容旧的用户分组接口）
	TagidList 			[]int 		`json:"tagid_list"`			//	用户被打上的标签ID列表
	SubscribeScene 		string 		`json:"subscribe_scene"`	//	返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_OTHERS 其他
	QrScene 			int 		`json:"qr_scene"`			//	二维码扫码场景（开发者自定义）
	QrSceneStr 			string 		`json:"qr_scene_str"`		//	二维码扫码场景描述（开发者自定义）
}

// 用户关注、取消关注接口微信推送接口
func IsSubscribe(c *gin.Context)  {
	body,_ := ioutil.ReadAll(c.Request.Body)
	var info Info
	xml.Unmarshal(body,&info)
	println(info.FromUserName)
	println(info.ToUserName)
	println(info.CreateTime)
	println(info.MsgType)
	println(info.Event)
	println(info.EventKey)

	token := getToken()
	if token != "" {
		openid := info.FromUserName
		// 根据用户openid获取用户信息
		userInfoData := httpHandle("GET","https://api.weixin.qq.com/cgi-bin/user/info?access_token="+token+"&openid="+openid+"&lang=zh_CN","","")
		var userInfo UserInfo
		json.Unmarshal(userInfoData,&userInfo)
		println(userInfo.Openid)
		println(userInfo.Nickname)
		println(userInfo.Subscribe)
	}
}

// 获取token
func getToken() string {
	// 公众号appID
	appid := "wx024bead69a59b953"
	// 公众号appsecret
	secret := "b8eb837e79e42889e582b55bfe8433c8"
	// 获取token
	responseData := httpHandle("GET","https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+appid+"&secret="+secret,"","")
	var tokenMap map[string]interface{}
	json.Unmarshal(responseData,&tokenMap)
	if tokenMap["errcode"] != nil {
		log.Print("获取token失败，错误代码：",tokenMap["errcode"])
		return ""
	}else {
		token := tokenMap["access_token"].(string)
		return token
	}
}

// http请求
func httpHandle(method, urlVal,cookie,body string) []byte {
	client := &http.Client{}
	var req *http.Request
	req, _ = http.NewRequest(method, urlVal, strings.NewReader(body))
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}