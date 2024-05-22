package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/hunterhug/go_image"
	"github.com/spf13/viper"
)

//在这里定义一些基础库

// 封装一个产生随机数的函数  由四位数字组成
func GetRandomNum() string {
	var str string
	for i := 0; i < 4; i++ {
		current := rand.Intn(10)
		str += strconv.Itoa(current)
	}
	return str
}

// 将时间戳转化为日期格式
func TimestampToData(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func GetUnix() int64 {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()
}

// 获取时间戳的nano时间
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// 获取当前日期
func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Md5 加密
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

// 验证邮箱
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 获取日期
func FormatDade() string {
	return time.Now().Format("20060102")
}

// 随机生成订单号
func GenerateOrderId() string {
	return time.Now().Format("200601021504") + GetRandomNum()
}

// 发送验证码，这里先写一个固定值吧
func SendMsg(str string) {
	ioutil.WriteFile("test_send.txt", []byte(str), 06666)
}

// 重新裁剪图片
func ResizeImage(filename string) {
	extName := path.Ext(filename) //后缀名
	resize := strings.Split(viper.GetString("resizeImageSize"), ",")

	for i := 0; i < len(resize); i++ {
		w := resize[i]
		width, _ := strconv.Atoi(w)
		savepath := filename + "_" + w + "x" + w + extName
		err := go_image.ThumbnailF2F(filename, savepath, width, width)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// 格式化图片
func FormatImage(picName string) string {
	ossStatus := viper.GetBool("ossStatus")
	if ossStatus {
		return viper.GetString("ossDomain") + "/" + picName
	} else {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName

	}
}

// 格式化标题
func FormatAttribute(str string) string {
	md := []byte(str)
	htmlByte := markdown.ToHTML(md, nil, nil)
	return string(htmlByte)
}

//计算乘法

func Mul(price float64, num int) float64 {
	return price * float64(num)
}
