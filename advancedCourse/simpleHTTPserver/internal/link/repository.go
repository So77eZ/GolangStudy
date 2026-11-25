package link

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/pkg/db"
	"errors"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	// Проверка длинны обновляемого хэша
	if len(link.Hash) != 8 && link.Hash != "" {
		return nil, errors.New("hash must be 8 characters")
	}
	// Обновление
	result := repo.Database.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link
	result := repo.Database.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

// Count подсчитывает активные ссылки с тем же фильтром, что и GetAll, без пагинации.
func (repo *LinkRepository) Count() (int64, error) {
	var count int64
	err := repo.Database.
		Table("links").
		Where("deleted_at is null").
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetAll возвращает страницу активных ссылок в порядке возрастания ID.
func (repo *LinkRepository) GetAll(limit, offset uint) ([]Link, error) {
	// Инициализируем пустой срез, чтобы пустая страница кодировалась как [].
	links := make([]Link, 0)

	// Scan выполняет запрос и записывает строки в links.
	err := repo.Database.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Scan(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}
