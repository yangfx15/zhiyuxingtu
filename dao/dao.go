package dao

import (
    "gorm.io/gorm"
    "xingtu/model"
)

type DAO struct {
    db *gorm.DB
}

func NewDAO(db *gorm.DB) *DAO {
    return &DAO{db: db}
}

func (d *DAO) GetUserByOpenID(openID string) (*model.User, error) {
    var user model.User
    if err := d.db.Where("open_id = ?", openID).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (d *DAO) CreateUser(user *model.User) error {
    return d.db.Create(user).Error
}

func (d *DAO) GetQuestions(category, questionType string, limit int) ([]model.Question, error) {
    var questions []model.Question
    if err := d.db.Where("category = ? AND type = ?", category, questionType).Limit(limit).Find(&questions).Error; err != nil {
        return nil, err
    }
    return questions, nil
}

func (d *DAO) CreateTestRecord(record *model.TestRecord) error {
    return d.db.Create(record).Error
}