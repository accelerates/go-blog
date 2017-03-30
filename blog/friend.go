package blog

type Friend struct {
	SchemasTime
}

type FriendInfo struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	WebSite string `json:"web_site"`
}

func (f *Friend) TableName() string {
	return "friends"
}

func (f *Friend) GetFriends() []*FriendInfo {
	friends := make([]*FriendInfo, 0)
	query := "SELECT id,name,website FROM friends"

	rows, err := sqlEngine.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		friend := new(FriendInfo)
		rows.Scan(&friend.Id, &friend.Name, &friend.WebSite)
		friends = append(friends, friend)
	}

	return friends
}
