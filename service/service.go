package service

import (
	"xingtu/dao"
	"xingtu/model"
)

type Service struct {
    dao *dao.DAO
}

func NewService(dao *dao.DAO) *Service {
    return &Service{dao: dao}
}

func (s *Service) CreateUser(user *model.User) error {
    return s.dao.CreateUser(user)
}

func (s *Service) Login(openID, nickname, avatarURL string) (*model.User, error) {
    user, err := s.dao.GetUserByOpenID(openID)
    if err == nil {
        return user, nil
    }
    user = &model.User{
        OpenID:    openID,
        Nickname:  nickname,
        AvatarURL: avatarURL,
    }
    if err := s.dao.CreateUser(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *Service) GetQuestions(category, questionType string, limit int) ([]model.Question, error) {
    return s.dao.GetQuestions(category, questionType, limit)
}

func (s *Service) SubmitTest(userID, category string, score int) error {
    record := &model.TestRecord{
        UserID:   userID,
        Category: category,
        Score:    score,
    }
    return s.dao.CreateTestRecord(record)
}