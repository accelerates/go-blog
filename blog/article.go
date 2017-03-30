package blog

import (
	"time"
)

const (
	ALL   uint8 = 0
	LOG   uint8 = 1
	WAY   uint8 = 2
	SHARE uint8 = 3
	CTF   uint8 = 4
)

//文章
type Article struct {
	ArticleInfo
	SchemasTime
}

type ArticleInfo struct {
	Id      int64   `json:"id"`
	Title   string  `json:"title"`
	Type    uint8   `json:"type"`
	Content []uint8 `json:"content"`
}

func (a *Article) TableName() string {
	return "articles"
}

func (a *Article) GetArticles(t uint8) []*Article {
	articles := make([]*Article, 0)
	query := "SELECT id,title,type,content,create_time FROM articles WHERE type = ?"
	var args string
	if t == 0 {
		args = "*"
	} else {
		args = string(t)
	}

	rows, err := sqlEngine.Query(query, args)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		article := new(Article)
		rows.Scan(&article.Id, &article.Title, &article.Type, &article.Content, &article.CreateTime)
		articles = append(articles, article)
	}

	return articles
}

func (a *Article) GetArticle(id int64) Article {
	article := new(Article)
	var createTime string
	query := "SELECT id,title,type,content,create_time FROM articles WHERE id = ?"
	err := sqlEngine.QueryRow(query, id).Scan(&article.Id, &article.Title, &article.Type, &article.Content, &createTime)
	if err != nil {
		panic(err)
	}

	loc, _ := time.LoadLocation("Local")
	article.CreateTime, err = time.ParseInLocation("2006-01-02 15:04:05", createTime, loc)
	if err != nil {
		panic(err)
	}

	return *article
}

func ArticlesStringMap(articles []*Article, tags map[int64][]*TagInfo) []map[string]interface{} {
	slice := make([]map[string]interface{}, 0)
	stringMap := make(map[string]interface{})

	for _, article := range articles {
		stringMap["id"] = article.Id
		stringMap["title"] = article.Title
		stringMap["type"] = article.Type
		stringMap["content"] = article.Content
		stringMap["tags"] = tags[article.Id]
		stringMap["create_time"] = article.CreateTime
		slice = append(slice, stringMap)
	}
	return slice
}
