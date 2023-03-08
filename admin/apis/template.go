package apis

import (
	"github.com/LINQQ1212/common2/global"
	"github.com/LINQQ1212/common2/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"path"
	"strings"
)

type option struct {
	Value any    `json:"value"`
	Label string `json:"label"`
}

func TemplateList(c *gin.Context) {
	response.OkWithData(map[string]any{
		"list":    dirFile("list"),
		"article": dirFile("article"),
		"category_link": []option{
			{0, "自动"},
			{1, "数字"},
			{2, "字母"},
			{3, "长数字"},
			{4, "名称+数字"},
			{5, "数字+名称"},
			{6, "名称+长数字"},
			{7, "长数字+名称"},
			{8, "字母+名称"},
			{9, "名称+字母"},
		},
		"product_link": []option{
			{0, "自动"},
			{1, "数字"},
			{2, "字母"},
			{3, "长数字"},
			{4, "名称+数字"},
			{5, "数字+名称"},
			{6, "名称+字母"},
			{7, "字母+名称"},
			{8, "无分类+名称+数字"},
			{9, "无分类+数字+名称"},
			{10, "无分类+名称+字母"},
			{11, "无分类+字母+名称"},
			{12, "长数字+名称"},
			{13, "名称+长数字"},
			{14, "无分类+长数字+名称"},
			{15, "无分类+名称+长数字"},
		},
	}, c)
}

func dirFile(dir string) []string {
	var data []string
	fs, err := os.ReadDir(path.Join(global.CONFIG.System.MainDir, "views", dir))
	if err != nil {
		global.LOG.Error("dirFile", zap.Error(err))
		return data
	}
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		if !strings.HasSuffix(f.Name(), ".html") && !strings.HasSuffix(f.Name(), ".jet") {
			continue
		}
		data = append(data, path.Join(dir, f.Name()))
	}
	return data
}
