package models

import "github.com/jinzhu/gorm"

// APIからのデータ取得履歴
type History struct {
	gorm.Model
	ApiUpdatedAt string `gorm:"type:datetime;not null"`             // API 更新日時
	ResourceType string `gorm:"type:enum('album','song');not null"` // リソースタイプ
	ApiUrl       string `gorm:"type:text;not null"`                 // API URL
	Resources    []Resource
}
