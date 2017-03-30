package blog

import (
	"time"
)

type SchemasTime struct {
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
