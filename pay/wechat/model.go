package wechat

const (
	// 境外国家地区
	China         Country = 1 // 中国国内
	China2        Country = 2 // 中国国内（冗灾方案）
	SoutheastAsia Country = 3 // 东南亚
	Other         Country = 4 // 其他国家

	// URL
	baseUrlCh  = "https://api.mch.weixin.qq.com/"   // 中国国内
	baseUrlCh2 = "https://api2.mchid.weixin.qq.com/"  // 中国国内
	baseUrlHk  = "https://apihk.mch.weixin.qq.com/" // 东南亚
	baseUrlUs  = "https://apius.mch.weixin.qq.com/" // 其他

	// 正式
	microPay          = "pay/micropay"                          // 提交付款码支付
	unifiedOrder      = "pay/unifiedorder"                      // 统一下单
	orderQuery        = "pay/orderquery"                        // 查询订单
	closeOrder        = "pay/closeorder"                        // 关闭订单
	refund            = "secapi/pay/refund"                     // 申请退款
	reverse           = "secapi/pay/reverse"                    // 撤销订单
	refundQuery       = "pay/refundquery"                       // 查询退款
	downloadBill      = "pay/downloadbill"                      // 下载对账单
	downloadFundFlow  = "pay/downloadfundflow"                  // 下载资金账单
	report            = "payitil/report"                        // 交易保障
	batchQueryComment = "billcommentsp/batchquerycomment"       // 拉取订单评价数据
	transfers         = "mmpaymkttransfers/promotion/transfers" // 企业向微信用户个人付款
	authCodeToOpenid  = "tools/authcodetoopenid"                // 授权码查询openid
	entrustPublic     = "papay/entrustweb"                      // 公众号纯签约
	entrustApp        = "papay/preentrustweb"                   // APP纯签约
	entrustH5         = "papay/h5entrustweb"                    // H5纯签约
	entrustPaying     = "pay/contractorder"                     // 支付中签约
	entrustQuery      = "papay/querycontract"                   // 查询签约关系
	entrustApplyPay   = "pay/pappayapply"                       // 申请扣款
	entrustDelete     = "papay/deletecontract"                  // 申请解约
	entrustQueryOrder = "pay/paporderquery"                     // 查询扣款订单

	// SanBox
	sandboxGetSignKey   = "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"
	sandboxMicroPay     = "sandboxnew/pay/micropay"
	sandboxUnifiedOrder = "sandboxnew/pay/unifiedorder"
	sandboxOrderQuery   = "sandboxnew/pay/orderquery"
	sandboxCloseOrder   = "sandboxnew/pay/closeorder"
	sandboxRefund       = "sandboxnew/pay/refund"
	sandboxReverse      = "sandboxnew/pay/reverse"
	sandboxRefundQuery  = "sandboxnew/pay/refundquery"
	sandboxDownloadBill = "sandboxnew/pay/downloadbill"
	sandboxReport       = "sandboxnew/payitil/report"

	// 支付类型
	TradeType_Mini   = "JSAPI"  // 小程序支付
	TradeType_JsApi  = "JSAPI"  // JSAPI支付
	TradeType_App    = "APP"    // app支付
	TradeType_H5     = "MWEB"   // H5支付
	TradeType_Native = "NATIVE" // Native支付

	// 签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"
)

// Notify
type NotifyRequest struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	SignType           string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CouponFee          string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponCount        string `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1        string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0,omitempty"`
	CouponId1          string `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponFee0         string `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1         string `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
}

type UnifiedOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"returnCode,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"returnMsg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mchId,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"deviceInfo,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonceStr,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"resultCode,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"errCode,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"errCodeDes,omitempty"`
	TradeType  string `xml:"trade_type,omitempty" json:"tradeType,omitempty"`
	PrepayId   string `xml:"prepay_id,omitempty" json:"prepayId,omitempty"`
	CodeUrl    string `xml:"code_url,omitempty" json:"codeUrl,omitempty"`
	MwebUrl    string `xml:"mweb_url,omitempty" json:"mwebUrl,omitempty"`
}

type QueryOrderResponse struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	TradeState         string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CouponFee          string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CouponCount        string `xml:"coupon_count,omitempty" json:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1        string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponId0          string `xml:"coupon_id_0,omitempty" json:"coupon_id_0,omitempty"`
	CouponId1          string `xml:"coupon_id_1,omitempty" json:"coupon_id_1,omitempty"`
	CouponFee0         string `xml:"coupon_fee_0,omitempty" json:"coupon_fee_0,omitempty"`
	CouponFee1         string `xml:"coupon_fee_1,omitempty" json:"coupon_fee_1,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	TradeStateDesc     string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
}

type CloseOrderResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
}

type ReverseResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Recall     string `xml:"recall,omitempty" json:"recall,omitempty"`
}

type RefundResponse struct {
	ReturnCode          string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg           string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode          string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid               string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId               string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty" json:"sign,omitempty"`
	TransactionId       string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty" json:"settlement_refund_fee,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee             string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CashRefundFee       string `xml:"cash_refund_fee,omitempty" json:"cash_refund_fee,omitempty"`
	CouponType0         string `xml:"coupon_type_0,omitempty" json:"coupon_type_0,omitempty"`
	CouponType1         string `xml:"coupon_type_1,omitempty" json:"coupon_type_1,omitempty"`
	CouponRefundFee     string `xml:"coupon_refund_fee,omitempty" json:"coupon_refund_fee,omitempty"`
	CouponRefundFee0    string `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`
	CouponRefundFee1    string `xml:"coupon_refund_fee_1,omitempty" json:"coupon_refund_fee_1,omitempty"`
	CouponRefundCount   string `xml:"coupon_refund_count,omitempty" json:"coupon_refund_count,omitempty"`
	CouponRefundId0     string `xml:"coupon_refund_id_0,omitempty" json:"coupon_refund_id_0,omitempty"`
	CouponRefundId1     string `xml:"coupon_refund_id_1,omitempty" json:"coupon_refund_id_1,omitempty"`
}

type QueryRefundResponse struct {
	ReturnCode           string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg            string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode           string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode              string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes           string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Appid                string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId                string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr             string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                 string `xml:"sign,omitempty" json:"sign,omitempty"`
	TotalRefundCount     string `xml:"total_refund_count,omitempty" json:"total_refund_count,omitempty"`
	TransactionId        string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo           string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee             string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee   string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	FeeType              string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	CashFee              string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	RefundCount          string `xml:"refund_count,omitempty" json:"refund_count,omitempty"`
	OutRefundNo0         string `xml:"out_refund_no_0,omitempty" json:"out_refund_no_0,omitempty"`
	OutRefundNo1         string `xml:"out_refund_no_1,omitempty" json:"out_refund_no_1,omitempty"`
	RefundId0            string `xml:"refund_id_0,omitempty" json:"refund_id_0,omitempty"`
	RefundId1            string `xml:"refund_id_1,omitempty" json:"refund_id_1,omitempty"`
	RefundChannel0       string `xml:"refund_channel_0,omitempty" json:"refund_channel_0,omitempty"`
	RefundChannel1       string `xml:"refund_channel_1,omitempty" json:"refund_channel_1,omitempty"`
	RefundFee            string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	RefundFee0           string `xml:"refund_fee_0,omitempty" json:"refund_fee_0,omitempty"`
	RefundFee1           string `xml:"refund_fee_1,omitempty" json:"refund_fee_1,omitempty"`
	SettlementRefundFee0 string `xml:"settlement_refund_fee_0,omitempty" json:"settlement_refund_fee_0,omitempty"`
	SettlementRefundFee1 string `xml:"settlement_refund_fee_1,omitempty" json:"settlement_refund_fee_1,omitempty"`
	CouponType00         string `xml:"coupon_type_0_0,omitempty" json:"coupon_type_0_0,omitempty"`
	CouponType01         string `xml:"coupon_type_0_1,omitempty" json:"coupon_type_0_1,omitempty"`
	CouponType10         string `xml:"coupon_type_1_0,omitempty" json:"coupon_type_1_0,omitempty"`
	CouponType11         string `xml:"coupon_type_1_1,omitempty" json:"coupon_type_1_1,omitempty"`
	CouponRefundFee0     string `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`
	CouponRefundFee1     string `xml:"coupon_refund_fee_1,omitempty" json:"coupon_refund_fee_1,omitempty"`
	CouponRefundCount0   string `xml:"coupon_refund_count_0,omitempty" json:"coupon_refund_count_0,omitempty"`
	CouponRefundCount1   string `xml:"coupon_refund_count_1,omitempty" json:"coupon_refund_count_1,omitempty"`
	CouponRefundId00     string `xml:"coupon_refund_id_0_0,omitempty" json:"coupon_refund_id_0_0,omitempty"`
	CouponRefundId01     string `xml:"coupon_refund_id_0_1,omitempty" json:"coupon_refund_id_0_1,omitempty"`
	CouponRefundId10     string `xml:"coupon_refund_id_1_0,omitempty" json:"coupon_refund_id_1_0,omitempty"`
	CouponRefundId11     string `xml:"coupon_refund_id_1_1,omitempty" json:"coupon_refund_id_1_1,omitempty"`
	CouponRefundFee00    string `xml:"coupon_refund_fee_0_0,omitempty" json:"coupon_refund_fee_0_0,omitempty"`
	CouponRefundFee01    string `xml:"coupon_refund_fee_0_1,omitempty" json:"coupon_refund_fee_0_1,omitempty"`
	CouponRefundFee10    string `xml:"coupon_refund_fee_1_0,omitempty" json:"coupon_refund_fee_1_0,omitempty"`
	CouponRefundFee11    string `xml:"coupon_refund_fee_1_1,omitempty" json:"coupon_refund_fee_1_1,omitempty"`
	RefundStatus0        string `xml:"refund_status_0,omitempty" json:"refund_status_0,omitempty"`
	RefundStatus1        string `xml:"refund_status_1,omitempty" json:"refund_status_1,omitempty"`
	RefundAccount0       string `xml:"refund_account_0,omitempty" json:"refund_account_0,omitempty"`
	RefundAccount1       string `xml:"refund_account_1,omitempty" json:"refund_account_1,omitempty"`
	RefundRecvAccout0    string `xml:"refund_recv_accout_0,omitempty" json:"refund_recv_accout_0,omitempty"`
	RefundRecvAccout1    string `xml:"refund_recv_accout_1,omitempty" json:"refund_recv_accout_1,omitempty"`
	RefundSuccessTime0   string `xml:"refund_success_time_0,omitempty" json:"refund_success_time_0,omitempty"`
	RefundSuccessTime1   string `xml:"refund_success_time_1,omitempty" json:"refund_success_time_1,omitempty"`
}

type MicropayResponse struct {
	ReturnCode         string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid              string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId              string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode         string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	Openid             string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType           string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType            string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee           string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	CouponFee          string `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CashFee            string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	TransactionId      string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo         string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	PromotionDetail    string `xml:"promotion_detail,omitempty" json:"promotion_detail,omitempty"`
}

type AuthCodeToOpenIdResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	Openid     string `xml:"openid,omitempty" json:"openid,omitempty"`
}

type TransfersResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	MchAppid       string `xml:"mch_appid,omitempty" json:"mch_appid,omitempty"`
	Mchid          string `xml:"mchid,omitempty" json:"mchid,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ResultCode     string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode        string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty" json:"partner_trade_no,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty" json:"payment_no,omitempty"`
	PaymentTime    string `xml:"payment_time,omitempty" json:"payment_time,omitempty"`
}

type ReportResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
}

type EntrustPublicResponse struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ResultMsg  string `xml:"result_msg,omitempty" json:"result_msg,omitempty"`
}

type EntrustAppPreResponse struct {
	ReturnCode      string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg       string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode      string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode         string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	Appid           string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId           string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	Sign            string `xml:"sign,omitempty" json:"sign,omitempty"`
	NonceStr        string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	PreEntrustwebId string `xml:"pre_entrustweb_id,omitempty" json:"pre_entrustweb_id,omitempty"`
}

type EntrustH5Response struct {
	ReturnCode  string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg   string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode  string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ResultMsg   string `xml:"result_msg,omitempty" json:"result_msg,omitempty"`
	RedirectUrl string `xml:"redirect_url,omitempty" json:"redirect_url,omitempty"`
}

type EntrustPayingResponse struct {
	ReturnCode             string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg              string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode             string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	Appid                  string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId                  string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo             string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr               string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign                   string `xml:"sign,omitempty" json:"sign,omitempty"`
	ErrCode                string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes             string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	ContractResultCode     string `xml:"contract_result_code,omitempty" json:"contract_result_code,omitempty"`
	ContractErrCode        string `xml:"contract_err_code,omitempty" json:"contract_err_code,omitempty"`
	ContractErrCodeDes     string `xml:"contract_err_code_des,omitempty" json:"contract_err_code_des,omitempty"`
	PrepayId               string `xml:"prepay_id,omitempty" json:"prepay_id,omitempty"`
	TradeType              string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	CodeUrl                string `xml:"code_url,omitempty" json:"code_url,omitempty"`
	PlanId                 string `xml:"plan_id,omitempty" json:"plan_id,omitempty"`
	RequestSerial          string `xml:"request_serial,omitempty" json:"request_serial,omitempty"`
	ContractCode           string `xml:"contract_code,omitempty" json:"contract_code,omitempty"`
	ContractDisplayAccount string `xml:"contract_display_account,omitempty" json:"contract_display_account,omitempty"`
	MwebUrl                string `xml:"mweb_url,omitempty" json:"mweb_url,omitempty"`
	OutTradeNo             string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
}

type getSignKeyResponse struct {
	ReturnCode     string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg      string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	MchId          string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	SandboxSignkey string `xml:"sandbox_signkey,omitempty" json:"sandbox_signkey,omitempty"`
}

type RefundNotifyRequest struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ReqInfo    string `xml:"req_info,omitempty" json:"req_info,omitempty"`
}

type RefundNotify struct {
	TransactionId       string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty" json:"settlement_refund_fee,omitempty"`
	RefundStatus        string `xml:"refund_status,omitempty" json:"refund_status,omitempty"`
	SuccessTime         string `xml:"success_time,omitempty" json:"success_time,omitempty"`
	RefundRecvAccout    string `xml:"refund_recv_accout,omitempty" json:"refund_recv_accout,omitempty"`
	RefundAccount       string `xml:"refund_account,omitempty" json:"refund_account,omitempty"`
	RefundRequestSource string `xml:"refund_request_source,omitempty" json:"refund_request_source,omitempty"`
}

type Code2SessionRsp struct {
	SessionKey string `json:"session_key,omitempty"` // 会话密钥
	ExpiresIn  string `json:"expires_in,omitempty"`  // SessionKey超时时间（秒）
	Openid     string `json:"openid,omitempty"`      // 用户唯一标识
	Unionid    string `json:"unionid,omitempty"`     // 用户在开放平台的唯一标识符
	Errcode    string `json:"errcode,omitempty"`     // 错误码
	Errmsg     string `json:"errmsg,omitempty"`      // 错误信息
}

type PaidUnionId struct {
	Unionid string `json:"unionid,omitempty"` // 用户在开放平台的唯一标识符
	Errcode string `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"` // 获取到的凭证
	ExpiresIn   string `json:"expires_in,omitempty"`   // SessionKey超时时间（秒）
	Errcode     string `json:"errcode,omitempty"`      // 错误码
	Errmsg      string `json:"errmsg,omitempty"`       // 错误信息
}

type UserInfo struct {
	Subscribe      int    `json:"subscribe,omitempty"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid         string `json:"openid,omitempty"`          // 用户唯一标识
	Nickname       string `json:"nickname,omitempty"`        // 用户的昵称
	Sex            int    `json:"sex,omitempty"`             // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language       string `json:"language,omitempty"`        // 用户的语言，简体中文为zh_CN
	City           string `json:"city,omitempty"`            // 用户所在城市
	Province       string `json:"province,omitempty"`        // 用户所在省份
	Country        string `json:"country,omitempty"`         // 用户所在国家
	Headimgurl     string `json:"headimgurl,omitempty"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime  int    `json:"subscribe_time,omitempty"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid        string `json:"unionid,omitempty"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string `json:"remark,omitempty"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid        int    `json:"groupid,omitempty"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int  `json:"tagid_list,omitempty"`      // 用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene,omitempty"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他
	QrScene        int    `json:"qr_scene,omitempty"`        // 二维码扫码场景（开发者自定义）
	Errcode        int    `json:"errcode,omitempty"`         // 错误码
	Errmsg         string `json:"errmsg,omitempty"`          // 错误信息
}

// 微信小程序解密后 用户手机号结构体
type UserPhone struct {
	PhoneNumber     string         `json:"phoneNumber,omitempty"`
	PurePhoneNumber string         `json:"purePhoneNumber,omitempty"`
	CountryCode     string         `json:"countryCode,omitempty"`
	Watermark       *watermarkInfo `json:"watermark,omitempty"`
}

// 微信小程序解密后 用户信息结构体
type AppletUserInfo struct {
	OpenId    string         `json:"openId,omitempty"`
	NickName  string         `json:"nickName,omitempty"`
	Gender    int            `json:"gender,omitempty"`
	City      string         `json:"city,omitempty"`
	Province  string         `json:"province,omitempty"`
	Country   string         `json:"country,omitempty"`
	AvatarUrl string         `json:"avatarUrl,omitempty"`
	UnionId   string         `json:"unionId,omitempty"`
	Watermark *watermarkInfo `json:"watermark,omitempty"`
}

type watermarkInfo struct {
	Appid     string `json:"appid,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

// 授权码查询openid 返回
type OpenIdByAuthCodeRsp struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	Appid      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	Openid     string `xml:"openid,omitempty" json:"openid,omitempty"` // 用户唯一标识
}

// App应用微信第三方登录，code换取access_token
type AppLoginAccessToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	Openid       string `json:"openid,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Unionid      string `json:"unionid,omitempty"`
	Errcode      string `json:"errcode,omitempty"` // 错误码
	Errmsg       string `json:"errmsg,omitempty"`  // 错误信息
}

// 刷新App应用微信第三方登录后，获取的 access_token
type RefreshAppLoginAccessTokenRsp struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	Openid       string `json:"openid,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Errcode      string `json:"errcode,omitempty"` // 错误码
	Errmsg       string `json:"errmsg,omitempty"`  // 错误信息
}
