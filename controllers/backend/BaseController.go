package backend

import (
	"errors"
	"fmt"
	"goB2C/dao"
	"goB2C/model"
	"goB2C/util"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type BaseController struct {
}

// 成功提示
func (c *BaseController) Success(Ctx *gin.Context, message string, redirect string) {
	var Redirect string
	if strings.Contains(redirect, "http") {
		Redirect = redirect
	} else {
		Redirect = util.TopPath + redirect
	}
	fmt.Println("success redicet: ", Redirect)
	Ctx.HTML(200, "public_success.html", gin.H{
		"Redirect": Redirect,
		"Message":  message,
	})
}

// 错误提示
func (c *BaseController) Error(Ctx *gin.Context, message string, redirect string) {
	var Redirect string
	if strings.Contains(redirect, "http") {
		Redirect = redirect
	} else {
		Redirect = util.TopPath + redirect
	}
	fmt.Println("error redicet: ", Redirect)
	Ctx.HTML(200, "public_error.html", gin.H{
		"Redirect": Redirect,
		"Message":  message,
	})
}

// 重定向
func (c *BaseController) Goto(Ctx *gin.Context, redirect string) {
	Ctx.Redirect(302, "/"+viper.GetString("adminPath")+redirect)
}

func (c *BaseController) UploadImg(Ctx *gin.Context, picName string) (string, error) {
	ossStatus := viper.GetBool("ossStatus")
	if ossStatus == true {
		//return c.OssUploadImg(picName)
	}
	return c.LocalUploadImg(Ctx, picName)
}

func (c *BaseController) LocalUploadImg(Ctx *gin.Context, picName string) (string, error) {
	f, h, err := Ctx.Request.FormFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer f.Close()
	//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(h.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	//4、创建图片保存目录  static/upload/20200623
	day := util.FormatDay()
	dir := "static/upload/" + day

	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", err
	}
	//5、生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(util.GetUnixNano(), 10)
	//static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+extName)
	//6、保存图片
	file, _, err := Ctx.Request.FormFile(picName)
	if err != nil {
		return saveDir, err
	}
	defer file.Close()
	ff, err := os.OpenFile(saveDir, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return saveDir, err
	}
	defer ff.Close()
	io.Copy(ff, file)

	return saveDir, nil
}

/*
	func (c *BaseController) OssUploadImg(picName string) (string, error) {
		setting := model.Setting{}
		dao.DB.First(&setting)
		f, h, err := c.GetFile(picName)
		if err != nil {
			return "", err
		}
		//2、关闭文件流
		defer f.Close()
		//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
		extName := path.Ext(h.Filename)

		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			return "", errors.New("图片后缀名不合法")
		}
		//把文件流上传值OSS

		//4.1 创建OSS实例
		client, err := oss.New(setting.EndPoint, setting.Appid, setting.AppSecret)
		if err != nil {
			return "", err
		}

		// 4.2获取存储空间。
		bucket, err := client.Bucket(setting.BucketName)
		if err != nil {
			return "", err
		}
		//4.3创建图片保存目录  static/upload/20200623
		day := util.FormatDay()
		dir := "static/upload/" + day
		fileUnixName := strconv.FormatInt(util.GetUnixNano(), 10)
		//static/upload/20200623/144325235235.png
		saveDir := path.Join(dir, fileUnixName+extName)
		// 4.4上传文件流。
		err = bucket.PutObject(saveDir, f)
		if err != nil {
			return "", err
		}
		return saveDir, nil
	}
*/
func (c *BaseController) GetSetting() model.Setting {
	setting := model.Setting{Id: 1}
	dao.DB.First(&setting)
	return setting
}
