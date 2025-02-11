package models

import (
	"time"

	"github.com/iAmImran007/draw-app-js-go/pkg/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Drawing struct {
	gorm.Model
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	Data      string    `json:"data"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

//insilaze daabase 
func init() {
	config.ConectDb()
	DB = config.GetDb()
  DB.AutoMigrate(&Drawing{})
}

//dd create draw query
func (d *Drawing) CreateDraw() *Drawing {
  result := DB.Create(&d)
	if result.Error != nil {
		return nil
	}
	return d
}
//db getdraw query
func GetDraw() []Drawing {
	var drawings []Drawing
	result := DB.Find(&drawings)
	if result.Error != nil {
		return nil
	}
	return drawings
}

//db get book query
func GetDrawById(Id int64) (*Drawing, error) {
	var getDraw Drawing
	result := DB.Where("id = ?", Id).First(&getDraw)
	if result.Error != nil {
			return nil, result.Error
	}
	return &getDraw, nil
}

//db delete query
func DeleteDraw(ID int64) Drawing {
  var draw Drawing 
	 DB.Where("id=?", ID).Delete(&draw)
	return draw
}
