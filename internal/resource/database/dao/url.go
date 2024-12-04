package dao

import (
	"time"

	"github.com/hcd233/Aris-url-gen/internal/resource/database/model"
	"gorm.io/gorm"
)

// UserDAO 用户DAO
//
//	@author centonhuang
//	@update 2024-10-17 02:30:24
type URLDAO struct {
	baseDAO[model.URL]
}

// GetByEmail 通过邮箱获取用户
//
//	@receiver dao *UserDAO
//	@param db *gorm.DB
//	@param email string
//	@param fields []string``
//	@return user *model.User
//	@return err error
//	@author centonhuang
//	@update 2024-10-17 05:08:00
func (dao *URLDAO) GetByOriginalUrl(db *gorm.DB, originalUrl string, fields, preloads []string) (url *model.URL, err error) {
	sql := db.Select(fields)
	for _, preload := range preloads {
		sql = sql.Preload(preload)
	}
	err = sql.Where(model.URL{OriginalURL: originalUrl}).First(&url).Error
	return
}

func (dao *URLDAO) GetByShortUrl(db *gorm.DB, shortUrl string, fields, preloads []string) (url *model.URL, err error) {
	sql := db.Select(fields)
	for _, preload := range preloads {
		sql = sql.Preload(preload)
	}
	err = sql.Where(model.URL{ShortURL: shortUrl}).First(&url).Error
	return
}

func (dao *URLDAO) BatchGetExpiredURLs(db *gorm.DB, fields, preloads []string) (urls []*model.URL, err error) {
	sql := db.Select(fields)
	for _, preload := range preloads {
		sql = sql.Preload(preload)
	}
	err = sql.Where("expire_at < ?", time.Now()).Find(&urls).Error
	return
}
