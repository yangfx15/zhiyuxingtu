package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId    string `json:"user_id" gorm:"index:uid;type:varchar(32)"`
	OpenID    string `json:"open_id" gorm:"index:open_id;type:varchar(64);comment:OpenID"`
	Nickname  string `json:"nick_name" gorm:"type:varchar(256)"`
	AvatarURL string `json:"avatar_url" gorm:"type:varchar(1024)"`
}

type Question struct {
	gorm.Model
	QuestionId string `json:"question_id" gorm:"index:qid;type:varchar(32)"`
	Category   string `json:"category" gorm:"index:category-type;type:varchar(128)"`
	Type       string `json:"type" gorm:"index:category-type;type:varchar(128)"`
	Difficulty int    `json:"difficulty" type:tinyint(4)"`
	Content    string `json:"content" gorm:"type:varchar(1024)"`
	Options    string `json:"options" gorm:"type:varchar(1024)"`
	Answer     string `json:"answer" gorm:"type:varchar(1024)"`
}

type TestRecord struct {
	gorm.Model
	RecordId  string `json:"record_id" gorm:"index:rid;type:varchar(32)"`
	UserID    string `json:"user_id" gorm:"index:user_id;type:varchar(64)"`
	Category  string `json:"category" gorm:"index:category-type;type:varchar(128)"`
	Type      string `json:"type" gorm:"index:category-type;type:varchar(128)"`
	Score     int    `json:"score" type:tinyint(4)"`
}
