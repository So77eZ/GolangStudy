package link

import (
	"GolangCourse/advancedCourse/simpleHTTPserver/internal/stat"
	"math/rand/v2"

	"gorm.io/gorm"
)

// Link хранит исходный URL, короткий хэш и связанные дневные счётчики.
type Link struct {
	gorm.Model
	URL   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL"`
}

func NewLink(url string) *Link {
	link := &Link{
		URL: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(8)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
