package shortener

import (
	"crypto/sha1"
	"fmt"
	"io"
)

//Тут будет находиться логика сокращения ссылки
//Хэширование и запись в бд

type Url struct {
	url  string
	hash string
}

func New(url string) Url {
	var newUrl = Url{
		url: url,
	}
	return newUrl
}

func (u *Url) Hash() error {
	h := sha1.New()
	if _, err := io.WriteString(h, u.url); err != nil {
		return err
	}
	u.hash = fmt.Sprintf("%x", h.Sum(nil))
	//тут пока наибольшая проблема это дублирование значений,
	//при хэшировании может возникнуть колиззия и я хз как это хендлить пока что
	return nil
}

func (u *Url) Save() error {
	//тут будет находиться логика стука в БД и запись туда значений
	//Есть пока что у меня ощущение что надо засунуть это в postgres.go, как и isExists
	u.isExixts()
	return nil
}

func (u *Url) isExixts() error {
	//Пока закинул, мб пригодится при проверке на колиззию
	return nil
}
