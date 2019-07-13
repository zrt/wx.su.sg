package main

import (
	"net/http"
	"github.com/labstack/echo"
	"strings"
	"bytes"
	"html/template"
)

func Handler(c echo.Context) error {
	r := c.Request()
	path := r.URL.Path[1:]
	if path == "favicon.ico"{
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if strings.HasPrefix(path, "https://mp.weixin.qq.com/s/"){
		path = path[len("https://mp.weixin.qq.com/s/"):]
	}else if strings.HasPrefix(path, "mp.weixin.qq.com/s/"){
		path = path[len("mp.weixin.qq.com/s/"):]
	}else if strings.HasPrefix(path, "s/"){
		path = path[len("s/"):]
	}

	fullURL := "https://mp.weixin.qq.com/s/"+path

	// referer 301
	if ! strings.Contains(r.Referer(), "TwitterBot"){
		return c.Redirect(http.StatusFound, fullURL)
	}

	t, err := template.New("og").Parse(`{{define "T"}}
<!doctype html>
<html lang="zh-CN" prefix="og: http://ogp.me/ns#">
<head>
<meta charset="utf-8">
<meta name="generator" content="wx.su.sg">
<meta property="og:title" content="{{.Title}}">
<meta property="og:locale" content="zh_CN">
<meta property="og:type" content="website">
<meta property="og:description" content="{{.Summary}}">
<meta property="og:url" itemprop="url" content="{{.Url}}">
<meta property="og:site_name" content="{{.Author}}">
<meta property="og:image" content="{{.FirstImg}}">
</head>
<body>
<a href="{{.Url}}">{{.Url}}</a>
<script>
window.location.href="{{.Url}}";
</script>
</body>
</html>
{{end}}
`)
	if err != nil{
		return err
	}
	article := ParseArticle(fullURL)
	if article == nil{
		return echo.NewHTTPError(http.StatusNotFound)
	}
	article.Summary = strings.ReplaceAll(article.Summary, "\r","")
	for strings.Contains(article.Summary, "\n\n") {
		article.Summary = strings.ReplaceAll(article.Summary, "\n\n","\n")
	}
	var buff bytes.Buffer
	err = t.ExecuteTemplate(&buff, "T", article)
	if err != nil{
		return err
	}
	return c.HTML(http.StatusOK, buff.String())

}