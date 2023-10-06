package postgres

import (
	"fmt"

	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB{
	fmt.Println(">>>>>>>>>> connect with DB <<<<<<<<<<")
	return nil
}