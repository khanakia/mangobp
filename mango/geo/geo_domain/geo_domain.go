package geo_domain

import (
	"github.com/khanakia/mangobp/mango/util"
)

type Country struct {
	util.ModelStr
	// NewID       string  `json:"newId" gorm:"type:varchar(255)"`
	Name        string  `json:"name" gorm:"type:varchar(255)"`
	Iso3        string  `json:"iso3" gorm:"type:varchar(255)"`
	Iso2        string  `json:"iso2" gorm:"type:varchar(255)"`
	Region      string  `json:"region" gorm:"type:varchar(255)"`
	SubRegion   string  `json:"subRegion" gorm:"type:varchar(255)"`
	PhoneCode   string  `json:"phoneCode" gorm:"type:varchar(255)"`
	Capital     string  `json:"capital" gorm:"type:varchar(255)"`
	Currency    string  `json:"currency" gorm:"type:varchar(255)"`
	Native      string  `json:"native" gorm:"type:varchar(255)"`
	Emoji       string  `json:"emoji" gorm:"type:varchar(255)"`
	EmojiU      string  `json:"emojiU" gorm:"type:varchar(255)"`
	Symbol      string  `json:"symbol" gorm:"type:varchar(255)"`
	Lat         float64 `json:"lat" gorm:"type:varchar(255)"`
	Lng         float64 `json:"lng" gorm:"type:varchar(255)"`
	NumericCode int     `json:"numeric_code" gorm:"type:varchar(255)"`
	NumericID   int
}

func (Country) TableName() string {
	return "geo_countries"
}

type State struct {
	util.ModelStr
	// NewID            string  `json:"newId" gorm:"type:varchar(255)"`
	Name             string  `json:"name" gorm:"type:varchar(255)"`
	CountryCode      string  `json:"countryCode" gorm:"type:varchar(255)"`
	StateCode        string  `json:"stateCode" gorm:"type:varchar(255)"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	CountryID        string  `json:"countryId" gorm:"type:varchar(255)"`
	CountryNumericID int     `json:"countryNumericID"`
}

func (State) TableName() string {
	return "geo_states"
}
