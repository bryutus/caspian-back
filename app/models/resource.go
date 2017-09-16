package models

import "github.com/jinzhu/gorm"

// アルバム・ソング情報
type Resource struct {
	gorm.Model
	HistoryId  uint   `gorm:"type:int(10);not null;index:idx_history_id"` // History ID
	Name       string `gorm:"type:text;not null;"`                        // アルバム・ソング名
	Url        string `gorm:"type:text;not null;"`                        // アルバム・ソング URL
	ArtworkUrl string `gorm:"type:text;not null;"`                        // ジャケ写画像 URL
	ArtistName string `gorm:"type:text;not null;"`                        // アーティスト名
	ArtistUrl  string `gorm:"type:text;not null;"`                        // アーティストページ URL
	Copyright  string `gorm:"type:text;not null;"`                        // コピーライト
}
