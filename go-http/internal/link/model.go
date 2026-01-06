package link

import (
	"app/adv-http/internal/stat/model"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	URL   string      `json:"url"`
	Hash  string      `gorm:"uniqueIndex" json:"hash"`
	Stats []stat.Stat `json:"stats" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {

	link := &Link{
		URL: url,
	}
	link.GenerateHash()
	return link
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRunes(10)
}
