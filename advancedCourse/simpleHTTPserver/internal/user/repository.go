package user

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
)

// UserRepository выполняет операции с пользователями в базе данных.
type UserRepository struct {
	Database *db.Db
}

// NewUserRepository создаёт репозиторий на основе подключения к базе данных.
func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

// Create сохраняет нового пользователя.
func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetByName ищет пользователя по имени.
func (repo *UserRepository) GetByName(name string) (*User, error) {
	var user User
	result := repo.Database.First(&user, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetByEmail ищет пользователя по адресу электронной почты.
func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// // Update обновляет переданные поля пользователя.
// func (repo *UserRepository) Update(user *User) (*User, error) {
// 	// Проверка длинны обновляемого хэша - поменять на пароль или что-то такое
// 	if len(user.Password) != 10 && user.Password != "" {
// 		return nil, errors.New("password must be 10 characters")
// 	}
// 	// Обновление
// 	result := repo.Database.Clauses(clause.Returning{}).Updates(user)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return user, nil
// }

// // Delete удаляет пользователя по идентификатору.
// func (repo *UserRepository) Delete(id uint) error {
// 	result := repo.Database.Delete(&User{}, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
