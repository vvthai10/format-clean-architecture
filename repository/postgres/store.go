// TODO: File này xử lý các tác vụ mà có sự tương tác qua lại giữa các table
package postgres

import (
	"gorm.io/gorm"
)

type Store struct {
	*ProductRepo
	*UserRepo
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store{
	return &Store{
		db: db,
	}
}

func (s *Store) HandleProductUser() error{
	return nil
}