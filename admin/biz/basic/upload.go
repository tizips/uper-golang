package basic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "github.com/tizips/uper-go/admin/http/request/basic"
	"github.com/tizips/uper-go/admin/http/response/basic"
	"path/filepath"
)

func DoUploadByFile(ctx context.Context, c *app.RequestContext) {

	var request req.DoUploadByFile

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		http.Fail(c, "上传失败：%v", err)
		return
	}

	fp, err := file.Open()

	if err != nil {
		http.Fail(c, "文件读取失败：%v", err)
		return
	}

	filename := facades.Snowflake.Generate().String() + filepath.Ext(file.Filename)
	uri := request.Dir + "/" + filename

	if err = facades.Storage.Put(uri, fp, file.Size); err != nil {
		http.Fail(c, "上传失败：%v", err)
		return
	}

	responses := basic.DoUploadByFile{
		Name: filename,
		Uri:  uri,
		Url:  facades.Storage.Url(uri),
	}

	http.Success(c, responses)
}
