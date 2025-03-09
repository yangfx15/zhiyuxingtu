package model

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    OpenID    string   `json:"open_id" gorm:"index:open_id;type:varchar(64);comment:OpenID"`
    Nickname  string
    AvatarURL string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Question struct {
    gorm.Model
    Category   string
    Type       string
    Difficulty int
    Content    string
    Options    string
    Answer     string
    CreatedAt  time.Time
    UpdatedAt  time.Time
}

type TestRecord struct {
    gorm.Model
    UserID    uint
    Category  string
    Score     int
    CreatedAt time.Time
}