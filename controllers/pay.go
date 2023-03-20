package controllers

import(
"sanskrit_api/pay/wechat"
"sanskrit_api/util"
"github.com/astaxie/beego"
"time"
"sanskrit_api/models"
"fmt"
)

type PayController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [post]
func (this *PayController) Post() {
	source:=this.getInt("source",1)
	uid := this.getInt64("uid",-1)
	body := this.getString("body",DEFAULT_ISNULL, DEFAULT_MIN_SIZE)
	fee := this.getInt("fee",1)
	receipt:=this.getString("receipt",0,0)
	appId := "wx1bd7f51f4f97f248"
	mchId := "1590226511"
	apiKey := "GFDS8j98rewnmgl45wHTt980jg543abc"
	client := wechat.NewClient(appId, mchId, apiKey, false)
	//设置国家	
	client.SetCountry(wechat.China)

	number := util.GetRandomString(32)

	//初始化参数Map
	bm := make(util.BodyMap)
	bm.Set("nonce_str", util.GetRandomString(32))
	bm.Set("body", body)
	bm.Set("out_trade_no", number)
	bm.Set("total_fee", fee)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "http://zd112.com:8091/")
	bm.Set("trade_type", wechat.TradeType_App)
	bm.Set("device_info", "WEB")
	if len(receipt) >0{
		bm.Set("receipt",receipt)
	}
	bm.Set("sign_type", wechat.SignType_MD5)

	// sceneInfo := make(map[string]map[string]string)
	// h5Info := make(map[string]string)
	// h5Info["type"] = "Wap"
	// h5Info["wap_url"] = "http://www.gopay.ink"
	// h5Info["wap_name"] = "H5测试支付"
	// sceneInfo["h5_info"] = h5Info
	// bm.Set("scene_info", sceneInfo)

	//body.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")

	// 正式	
	//sign := gopay.GetWeChatParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	// 沙箱
	//sign, _ := gopay.GetWeChatSanBoxParamSign("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", body)
	//body.Set("sign", sign)

	//请求支付下单，成功后得到结果
	pay := new(models.Pay)
	pay.Source = source
	pay.Uid = uid
	pay.Appid=appId
	pay.Mchid=mchId
	pay.PackageValue="Sign=WXPay"
	pay.TradeType = wechat.TradeType_App
	wxRsp, err := client.UnifiedOrder(bm)
	if err != nil {
		pay.Status=-1
		pay.Reason = fmt.Sprintf("%s",err)
		pay.Apikey = apiKey
		id,err := pay.Add()
		beego.Info("id:",id," | err:",err)
		this.false(-1,err)
	}else{
		pay.Status = 1
		pay.NonceStr = wxRsp.NonceStr
		pay.Sign = wxRsp.Sign
		pay.PrepayId = wxRsp.PrepayId
		pay.Date = time.Now().Unix()
		pay.Reason = wxRsp.ErrCode

		this.response(pay)
		pay.Apikey = apiKey
		if id,err := pay.Add();err != nil{
			beego.Info("pay~id:",id," | err:",err)	
		}
	}
}