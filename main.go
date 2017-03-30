package main

import (
	"encoding/json"
	"github.com/go-blog/blog"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

func main() {

	m := martini.Classic()

	m.Use(core())

	m.Group("/api", func(r martini.Router) {

		r.Get("", func() string {
			return "hello world"
		})

		r.Get("/article/:id", func(params martini.Params) string {
			MyBlog := new(blog.Article)
			BlogId, _ := strconv.ParseInt(params["id"], 10, 64)
			article := MyBlog.GetArticle(BlogId)
			data, _ := json.Marshal(article)
			return string(data)
		})

		r.Get("/article", func() string {
			MyBlog := new(blog.Article)
			MyTag := new(blog.Tag)
			articles := MyBlog.GetArticles(blog.ALL)

			articleIds := []int64{}
			for _, article := range articles {
				articleIds = append(articleIds, article.Id)
			}

			tags := MyTag.GetArticleTags(articleIds)

			result := blog.ArticlesStringMap(articles, tags)

			json, _ := json.Marshal(result)
			return string(json)
		})
		r.Get("/friend", func() string {
			MyFriends := new(blog.Friend)
			friends := MyFriends.GetFriends()
			data, _ := json.Marshal(friends)
			return string(data)
		})

		r.Get("/log", func() string {
			MyBlog := new(blog.Article)
			articles := MyBlog.GetArticles(blog.LOG)
			data := make(map[string]interface{})
			data["data"] = articles
			json, _ := json.Marshal(data)
			return string(json[:])
		})
		r.Get("/way", func() string {
			MyBlog := new(blog.Article)
			articles := MyBlog.GetArticles(blog.WAY)
			data := make(map[string]interface{})
			data["data"] = articles
			json, _ := json.Marshal(data)
			return string(json[:])
		})
		r.Get("/ctf", func() string {
			MyBlog := new(blog.Article)
			articles := MyBlog.GetArticles(blog.CTF)
			data := make(map[string]interface{})
			data["data"] = articles
			json, _ := json.Marshal(data)
			return string(json[:])
		})

	})

	m.RunOnAddr(":9000")
}

func core() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		return
	}
}
