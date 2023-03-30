package service

import (
	"YYY/internal/model"
	"github.com/go-pg/pg/v10"
	"log"
)

type Service struct {
	db *pg.DB
}

func NewService(db *pg.DB) *Service {
	return &Service{db: db}
}

func (s *Service) AddGroup(group *model.Group) error {
	_, err := s.db.Model(group).Insert()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Service) GetGroups() ([]*model.Group, error) {
	var group []*model.Group

	err := s.db.Model(&group).Column("group.id", "group.name", "group.description").Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return group, nil
}

func (s *Service) GetGroupByID(id int64) (*model.Group, error) {
	group := &model.Group{ID: id}

	err := s.db.Model(group).WherePK().Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return group, nil
}
