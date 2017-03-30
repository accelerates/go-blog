package blog

import (
	"strconv"
)

type Tag struct {
	TagInfo
	SchemasTime
}

type TagInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (t *Tag) TableName() string {
	return "tags"
}

func SetTagInfo(id int64, name string) *TagInfo {
	return &TagInfo{id, name}
}

func (t *Tag) GetTags() []*Tag {
	tags := make([]*Tag, 0)
	return tags
}

//获取article的tags
func (t *Tag) GetArticleTags(ids []int64) (result map[int64][]*TagInfo) {

	result = make(map[int64][]*TagInfo)

	var in string
	for i, id := range ids {
		if i != 0 {
			in += ","
		}
		in += strconv.FormatInt(id, 10)
	}

	query := "SELECT at.article_id,t.id,t.name" +
		" FROM tags t RIGHT JOIN article_tag as at ON at.tag_id = t.id WHERE at.article_id IN (?) ORDER BY create_time DESC"
	rows, err := sqlEngine.Query(query, in)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var articleId, id int64
		var name string
		rows.Scan(&articleId, &id, &name)
		result[articleId] = append(result[articleId], SetTagInfo(id, name))
	}

	return result
}

func (t *Tag) GetTag(id int64) Tag {
	tag := new(Tag)
	return *tag
}
