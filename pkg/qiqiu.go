package qiniu

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"mime/multipart"
	"time"
)

type Qiniu struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

var Q Qiniu

func Load() {
	Q = Qiniu{
		AccessKey: viper.GetString("oss.access_key"),
		SecretKey: viper.GetString("oss.secret_key"),
		Bucket:    viper.GetString("oss.bucket_name"),
		Domain:    viper.GetString("oss.domain_name"),
	}
}
func UploadQiniu(file *multipart.FileHeader) (int, string) {
	src, err := file.Open()
	if err != nil {
		return 10011, err.Error()
	}

	defer src.Close()

	putPolicy := storage.PutPolicy{
		Scope: Q.Bucket,
	}

	mac := qbox.NewMac(Q.AccessKey, Q.SecretKey)

	// 获取上传凭证
	upToken := putPolicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传返回后的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 自定义文件名及后缀
	key := "(" + time.Now().String() + ")" + file.Filename

	if err := formUploader.Put(context.Background(), &ret,
		upToken, key, src, file.Size, &putExtra); err != nil {
		return 501, err.Error()
	}

	return 0, Q.Domain + "/" + ret.Key
}

func UploadFile(c *gin.Context) ([]string, error) {

	form, err := c.MultipartForm()
	files := form.File["file"]
	if err != nil {
		return nil, err
	}
	urls := make([]string, 0)
	for _, file := range files {
		_, url := UploadQiniu(file)
		urls = append(urls, url)
	}
	return urls, nil
}
