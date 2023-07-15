package test

import (
	"fmt"
	"golang-gin-note/models"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (c TestController) SetSession(r *gin.Context) {
	session := sessions.Default(r)
	session.Options(sessions.Options{
		MaxAge: 10, // //6hrs MaxAge单位是秒
		// Path:   "/",
		// HttpOnly: true,
		// Secure: false,
	})
	session.Set("age", "18")
	session.Save()
	r.String(200, "设置session成功")
}

func (c TestController) GetSession(r *gin.Context) {

	session := sessions.Default(r)
	age := session.Get("age")
	r.String(200, "获取session成功", age)
}
func (c TestController) SetCookie(r *gin.Context) {
	r.SetCookie("name", "kingcwt", 10, "/", "localhost", false, true)
	r.String(http.StatusOK, "ok")
}

func (c TestController) GetCookie(r *gin.Context) {
	s, err := r.Cookie("name")
	if err != nil {
		fmt.Println(err)
		r.String(http.StatusOK, "获取cookie失败")
		return
	}
	r.String(http.StatusOK, "cookie="+s)
}
func (c TestController) DeleteCookie(r *gin.Context) {
	r.SetCookie("name", "kingcwt", -1, "/", "localhost", false, true)
	r.String(http.StatusOK, "删除cookie成功")
}

func (con TestController) Index(r *gin.Context) {
	fmt.Println("【我是第二个执行的】我是返回的数据 get test")
	time.Sleep(time.Second)
	r.String(200, "get test")
}

// 单文件上传
func fileUpload(r *gin.Context) {
	username := r.PostForm("username")
	file, err := r.FormFile("files") //单文件

	if err != nil {
		fmt.Println(err)
	} else {
		dst := path.Join("./static/upload/", file.Filename)
		r.SaveUploadedFile(file, dst)
	}
	r.JSON(200, gin.H{
		"username": username,
		"file":     file.Filename,
		"size":     file.Size,
		"type":     file.Header["Content-Type"],
		"key":      file.Filename,
		"createAt": time.Now().Format("2006-01-02 15:04:05"),
		"updateAt": time.Now().Format("2006-01-02 15:04:05"),
		"deleteAt": time.Now().Format("2006-01-02 15:04:05"),
		"status":   "ok",
	})

}

// 多文件上传
func MultipartFileUpload(r *gin.Context) {
	username := r.PostForm("username")
	form, err := r.MultipartForm()
	if err != nil {
		fmt.Println(err)
	} else {
		files := form.File["files"]
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			dst := path.Join("./static/upload/", filename)
			if err := r.SaveUploadedFile(file, dst); err != nil {
				r.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}
		r.String(http.StatusOK, "Uploaded successfully %d files with fields name=%s", len(files), username)

	}
}

func (con TestController) Upload(r *gin.Context) {
	fileUpload(r)
	// MultipartFileUpload(r)

}

func (con TestController) GetUpload(r *gin.Context) {
	r.HTML(200, "test/index1.html", gin.H{
		"num":  8 << 20,
		"num2": 8 << 10,
	})
}

// 图片上传和根据当前时间分组保存
func (con TestController) GetFile(r *gin.Context) {
	r.HTML(200, "test/index2.html", gin.H{})
}
func (con TestController) UploadFile(r *gin.Context) {
	username := r.PostForm("username")
	file, err := r.FormFile("files")
	if err != nil {
		fmt.Println(err)
		return
	}
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
		".svg":  true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		r.String(200, "文件格式不正确")
		return
	}
	day := models.GetDay()
	dir := "./static/upload/" + day

	fileDir, err := os.Open(dir)
	if err != nil {
		error := os.MkdirAll(dir, os.ModePerm)
		if error != nil {
			r.String(200, "创建文件夹失败")
			return
		}
	}
	defer fileDir.Close()
	filName := models.GetFormattedTime() + extName
	dst := path.Join(dir, filName)
	if err := r.SaveUploadedFile(file, dst); err != nil {
		r.String(200, "上传文件失败")
		return
	}
	fmt.Println("----------------------------------------返回数据执行")
	// r.Redirect(http.StatusMovedPermanently, "/test/get/file")
	r.JSON(200, gin.H{
		"username": username,
		"file":     filName,
		"size":     file.Size,
		"createAt": time.Now().Format("2006-01-02 15:04:05"),
		"msg":      "上传成功 3秒钟后 跳转原页面",
	})
}
