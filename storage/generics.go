package storage

import "generics/model"

type Models interface {
	model.User | model.Transaction
}

func Get[entity Models](s *Storage, e entity, id int) (*entity, error) {

	result := s.Db.
		Where("id = ?", id).
		Find(&e)

	if er := result.Error; er != nil {
		return nil, er
	}

	return &e, nil
}

func Create[entity Models](s *Storage, e entity) (*entity, error) {
	result := s.Db.
		Create(&e)
	if er := result.Error; er != nil {
		return nil, er
	}

	return &e, nil
}

func Update[entity Models](s *Storage, e entity, id int) (*entity, error) {
	result := s.Db.
		Where("id = ?", id).
		Updates(&e)
	if er := result.Error; er != nil {
		return nil, er
	}

	return &e, nil
}

func Delete[entity Models](s *Storage, e entity, id int) error {
	result := s.Db.
		Delete(&e, "id = ?", id)

	if er := result.Error; er != nil {
		return er
	}

	return nil
}
