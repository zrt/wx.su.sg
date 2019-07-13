package main

import (
    "log"
    "net/http"
    "strings"

    "github.com/yhat/scrape"
    "golang.org/x/net/html"
    "golang.org/x/net/html/atom"
)

type Article struct {
    Title       string   `json:"title"`
    Author      string   `json:"author"`
    Url         string   `json:"url"`
    Summary     string   `json:"summary"`
    Photos      []string `json:"photos"`
    FirstImg    string   `json:"firstimg`
}

func ParseArticle(url string) *Article {
    resp, err := http.Get(url)
    if err != nil {
        log.Println("Cannot fetch content from url", url, "with error message", err)
        return nil
    }

    root, err := html.Parse(resp.Body)
    if err != nil {
        log.Println("Cannot parse the html page. Error:", err)
        return nil
    }

    title, _ := scrape.Find(root, titleMatcher)
    richMediaList, _ := scrape.Find(root, richMediaListMatcher)
    richMediaContent, _ := scrape.Find(root, richMediaContentMatcher)

    author, _ := scrape.Find(richMediaList, authorMatcher)

    article := new(Article)
    // set title
    article.Title = scrape.Text(title)
    // set url
    article.Url = url

    // set author
    article.Author = scrape.Text(author)

    // process content text
    contentNodes := scrape.FindAll(richMediaContent, contentText)
    summaryText := []string{}
    for _, node := range contentNodes {
        summaryText = append(summaryText, strings.TrimSpace(scrape.Text(node)))
    }
    article.Summary = strings.Join(summaryText, "\n")

    imageNodes := scrape.FindAll(richMediaContent, contentImage)
    images := []string{}
    for _, node := range imageNodes {
        images = append(images, scrape.Attr(node, "data-src"))
    }
    article.Photos = images
    if len(images) > 0{
        article.FirstImg = images[0]
    }else{
        article.FirstImg = "https://lepuslab.com/wp-content/uploads/2019/07/wechat.png"
    }

    return article
}

// define title matcher
func titleMatcher(n *html.Node) bool {
    if n.DataAtom == atom.H2 {
        return scrape.Attr(n, "class") == "rich_media_title"
    }
    return false
}

// define rich_media_meta_list matcher. The node should contain publish date and author info
func richMediaListMatcher(n *html.Node) bool {
    if n.DataAtom == atom.Div {
        return scrape.Attr(n, "class") == "rich_media_meta_list"
    }
    return false
}

// define rich_media_content matcher. The node contains the article content text and photos
func richMediaContentMatcher(n *html.Node) bool {
    return scrape.ById("js_content")(n)
}

// define article author matcher
func authorMatcher(n *html.Node) bool {
    return scrape.ByClass("profile_nickname")(n)
}

// define article content text nodes matcher
func contentText(n *html.Node) bool {
    return n.DataAtom == atom.Section || n.DataAtom == atom.P
}

// define article images
func contentImage(n *html.Node) bool {
    return n.DataAtom == atom.Img
}