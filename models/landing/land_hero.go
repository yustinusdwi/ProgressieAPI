package models

import "time"

type Land_Hero struct {
	HeroComponentID       int       `gorm:"column:hero_component_id" json:"hero_component_id"` // Primary key
	HeroComponentTitle    string    `gorm:"column:hero_component_title" json:"hero_component_title"`
	HeroComponentSubtitle string    `gorm:"column:hero_component_subtitle" json:"hero_component_subtitle"`
	HeroComponentImage    string    `gorm:"column:hero_component_image" json:"hero_component_image"`
	HeroComponentCoverImg string    `gorm:"column:hero_component_cover_img" json:"hero_component_cover_img"`
	Tooltip               string    `gorm:"column:tooltip" json:"tooltip"`
	CreatedBy             string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt             time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy             string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt             time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (l *Land_Hero) TableName() string {
	return "land_hero"
}
