package link

import (
	"go/links-shorter/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"unique"`
	Stats []stat.Stat `json:"stats" gorm:"foreignKey:LinkId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func NewLink(url string) *Link {
	link := &Link{Url: url}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = getRandomStr(10)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getRandomStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
