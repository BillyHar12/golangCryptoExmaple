package main

import "time"

type GameCrypto struct {
	ID          uint   `gorm:"primary_key"`
	Coin        string `gorm:"type:varchar(255);not null"`
	Price       uint   `gorm:"type:integer;not null"`
	MarketCap   uint   `gorm:"type:integer;not null"`
	Volume24h   uint   `gorm:"type:integer;not null"`
	Change1h    uint   `gorm:"type:integer;not null"`
	Change24h   uint   `gorm:"type:integer;not null"`
	Change7d    uint   `gorm:"type:integer;not null"`
	Change30d   uint   `gorm:"type:integer;not null"`
	Change60d   uint   `gorm:"type:integer;not null"`
	Change90d   uint   `gorm:"type:integer;not null"`
	AllTimeHigh uint   `gorm:"type:integer;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
