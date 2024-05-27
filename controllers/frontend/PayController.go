package frontend

import (
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/objcoding/wxpay"
	"github.com/skip2/go-qrcode"
	"github.com/smartwalle/alipay/v3"
)

type PayController struct {
	BaseController
}

func (P *PayController) Alipay(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	AliId, err1 := strconv.Atoi(idTemp)
	if err1 != nil {
		Ctx.Redirect(302, Ctx.Request.Referer())
	}
	orderitem := []model.OrderItem{}
	dao.DB.Where("order_id=?", AliId).Find(&orderitem)
	var privateKey = "xxxxxxx" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("2021001186696588", privateKey, true)
	client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021001186696588.certfile") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.certfile")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.certfile")       // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return
	}

	//计算总价格
	var TotalAmount float64
	for i := 0; i < len(orderitem); i++ {
		TotalAmount = TotalAmount + orderitem[i].ProductPrice
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = "xxxxxxx"
	p.ReturnURL = "xxxxxxx"
	p.TotalAmount = "0.01"
	p.Subject = "订单order——" + time.Now().Format("200601021504")
	p.OutTradeNo = "WF" + time.Now().Format("200601021504") + "_" + strconv.Itoa(AliId)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err4 = client.TradePagePay(p)
	if err4 != nil {
		fmt.Println(err4)
	}
	var payURL = url.String()
	Ctx.Redirect(302, payURL)
}

func (c *PayController) AlipayNotify(Ctx *gin.Context) {
	var privateKey = "xxxxxxxxxxxxxxx" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New("2021001186696588", privateKey, true)

	client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021001186696588.certfile") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.certfile")                   // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.certfile")       // 加载支付宝公钥证书

	if err != nil {
		fmt.Println(err)
		return
	}

	req := Ctx.Request
	req.ParseForm()
	err = client.VerifySign(req.Form)
	if err != nil {
		Ctx.Redirect(302, Ctx.Request.Referer())
	}
	rep := Ctx.Writer
	var noti, _ = client.GetTradeNotification(req)
	if noti != nil {
		fmt.Println("交易状态为:", noti.TradeStatus)
		if string(noti.TradeStatus) == "TRADE_SUCCESS" {
			order := model.Order{}
			temp := strings.Split(noti.OutTradeNo, "_")[1]
			id, _ := strconv.Atoi(temp)
			dao.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.OrderStatus = 1
			dao.DB.Save(&order)
		}
	}
	alipay.AckNotification(rep) // 确认收到通知消息
}
func (c *PayController) AlipayReturn(Ctx *gin.Context) {
	Ctx.Redirect(302, "/user/order")
}

func (c *PayController) WxPay(Ctx *gin.Context) {
	idTemp := Ctx.Query("id")
	WxId, err := strconv.Atoi(idTemp)
	if err != nil {
		Ctx.Redirect(302, Ctx.Request.Referer())
	}
	orderitem := []model.OrderItem{}
	dao.DB.Where("order_id=?", WxId).Find(&orderitem)
	//1、配置基本信息
	account := wxpay.NewAccount(
		"xxxxxxxx", //appid
		"xxxxxxxx", //商户号
		"xxxxxxxx", //appkey
		false,
	)
	client := wxpay.NewClient(account)
	var price int64
	for i := 0; i < len(orderitem); i++ {
		price = 1
	}
	//2、获取ip地址,订单号等信息
	ip := strings.Split(Ctx.Request.RemoteAddr, ":")[0]
	template := "202001021504"
	tradeNo := time.Now().Format(template)
	//3、调用统一下单
	params := make(wxpay.Params)
	params.SetString("body", "order——"+time.Now().Format(template)).
		SetString("out_trade_no", tradeNo+"_"+strconv.Itoa(WxId)).
		SetInt64("total_fee", price).
		SetString("spbill_create_ip", ip).
		SetString("notify_url", "http://xxxxxx/wxpay/notify"). //配置的回调地址
		// SetString("trade_type", "APP")//APP端支付
		SetString("trade_type", "NATIVE") //网站支付需要改为NATIVE

	p, err1 := client.UnifiedOrder(params)
	log.Println(p)
	if err1 != nil {
		Ctx.Redirect(302, Ctx.Request.Referer())
	}
	//4、获取code_url生成支付二维码
	var pngObj []byte
	log.Println(p)
	pngObj, _ = qrcode.Encode(p["code_url"], qrcode.Medium, 256)
	Ctx.Writer.WriteString(string(pngObj))
}

func (c *PayController) WxPayNotify(Ctx *gin.Context) {
	//1、获取表单传过来的xml数据，在配置文件里设置 copyrequestbody = true
	body, _ := io.ReadAll(Ctx.Request.Body)
	xmlStr := string(body)
	postParams := wxpay.XmlToMap(xmlStr)
	log.Println(postParams)

	//2、校验签名
	account := wxpay.NewAccount(
		"xxxxxxxx",
		"xxxxxxxx",
		"xxxxxxxx",
		false,
	)
	client := wxpay.NewClient(account)
	isValidate := client.ValidSign(postParams)
	// xml解析
	params := wxpay.XmlToMap(xmlStr)
	log.Println(params)
	if isValidate == true {
		if params["return_code"] == "SUCCESS" {
			idStr := strings.Split(params["out_trade_no"], "_")[1]
			id, _ := strconv.Atoi(idStr)
			order := model.Order{}
			dao.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.PayType = 1
			order.OrderStatus = 1
			dao.DB.Save(&order)
		}
	} else {
		Ctx.Redirect(302, Ctx.Request.Referer())
	}
}
