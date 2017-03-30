package blog

type ArticleTag struct {
	Id        int64
	ArticleId int64 `xorm:"int notnull 'article_id'"`
	TagId     int64 `xorm:"int notnull 'tag_id'"`
}

func (at *ArticleTag) TableName() string {
	return "article_tag"
}
