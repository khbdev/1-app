package service

import (
	"1-loyiha/model"
	"1-loyiha/repostory"
)


type UserService struct {
	repo *repostory.UserRepostory
}
func NewUserService(repo *repostory.UserRepostory) *UserService {
	return &UserService{repo: repo}
}

func(s *UserService) GetAll() []model.User {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id int) *model.User {
	return s.repo.GetByID(id)
} 

func (s *UserService) Create(user model.User) model.User {
	return s.repo.Create(user)
}

func (s *UserService) Update (id int, update model.User) *model.User {
	return s.repo.Updated(id, update)
}

func (s *UserService) Delete (id int) bool {
	return  s.repo.Delete(id)
}