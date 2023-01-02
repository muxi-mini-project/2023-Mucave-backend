package service

//服务器本地储存
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"os"
//)
//
//func UploadFile(c *gin.Context, id interface{}) string {
//	// Multipart form
//	task := c.Query("task")
//	form, _ := c.MultipartForm()
//	files := form.File["file"]
//	var dst string
//	var path string
//	if task == "post" {
//		path = fmt.Sprintf("C:/temp/Post/%v", id)
//		os.MkdirAll(path, 0766)
//	} else if task == "chat" {
//		path = fmt.Sprintf("C:/temp/Chat/%v", id)
//		os.MkdirAll(path, 0766)
//		dst = fmt.Sprintf(path+"/%s", files[0].Filename)
//		c.SaveUploadedFile(files[0], dst)
//		return dst
//	} else if task == "avatar" {
//		path = fmt.Sprintf("C:/temp/Avatar/%v", id)
//		os.MkdirAll(path, 0766)
//		dst = path + "/avatar." + c.Query("format")
//		c.SaveUploadedFile(files[0], dst)
//		return dst
//	}
//	for index, file := range files {
//		dst = fmt.Sprintf(path+"/%d_%s", index, file.Filename)
//		c.SaveUploadedFile(file, dst)
//	}
//	return path
//}
//func LoadFile(c *gin.Context, path string) error {
//	filePath := path
//	//打开文件
//	fileTmp, errByOpenFile := os.Open(filePath)
//	if errByOpenFile != nil {
//		c.Redirect(http.StatusFound, "/404")
//		return errByOpenFile
//	}
//	defer fileTmp.Close()
//	fileName := fileTmp.Name()
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", "attachment; filename="+fileName)
//	c.Header("Content-Transfer-Encoding", "binary")
//	c.Header("Cache-Control", "no-cache")
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", "attachment; filename="+fileName)
//	c.Header("Content-Transfer-Encoding", "binary")
//	fmt.Println("?????")
//	c.File(filePath)
//	return nil
//}
