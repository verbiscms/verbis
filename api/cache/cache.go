package cache

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	Store *cache.Cache
)

type Cacher interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
	Flush()
}

const (
	// For use with functions that take an expiration time.
	RememberForever time.Duration = -1
	postIdKey                     = "post-id-"
)

// Init set-ups go-cache with defaults
func Init() {
	Store = cache.New(5*time.Minute, 10*time.Minute)
}

func ClearPostCache(id int) {
	Store.Delete(GetPostKey(id))
}

func ClearUserCache(userId int, posts []domain.Post) {
	for _, v := range posts {
		if v.UserId == userId {
			ClearPostCache(v.Id)
		}
	}
}

//func ClearCategoryCache(categoryId int, posts []domain.Post) {
//
//	for _, v := range posts {
//		if v == userId {
//			ClearPostCache(v.Id)
//		}
//	}
//}

func GetPostKey(id int) string {
	return fmt.Sprintf("%s%d", postIdKey, id)
}
