package logdb_domain

import (
	"time"

	"github.com/khanakia/mangobp/mango/publicid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	NATS_LOGDB_LOG = "logdb.log"
)

type Log struct {
	ID        string         `json:"id" gorm:"type:varchar(50);primaryKey;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Flag      string         `json:"flag" gorm:"type:varchar(255)"` // info|error|warning
	Name      string         `json:"name" gorm:"type:varchar(255)"` // any name lets say mail_api, register_api etc
	ClientIP  string         `json:"clientIp" gorm:"type:varchar(255)"`
	Subject   string         `json:"subject" gorm:"type:varchar(255)"`
	Ref       string         `json:"ref" gorm:"type:varchar(255)"`
	RefID     string         `json:"refId" gorm:"type:varchar(255)"`
	Ref2      string         `json:"ref2" gorm:"type:varchar(255)"`
	Ref2ID    string         `json:"ref2Id" gorm:"type:varchar(255)"`
	Ref3      string         `json:"ref3" gorm:"type:varchar(255)"`
	Ref3ID    string         `json:"ref3Id" gorm:"type:varchar(255)"`
	Request   string         `json:"request"`
	Response  string         `json:"response"`
	Data      datatypes.JSON `json:"data"` // any other extra data
}

func (Log) TableName() string {
	return "logdb_logs"
}

func (u *Log) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.ID) == 0 {
		u.ID = "log_" + publicid.Must36()
	}
	return
}

type CreateArgs struct {
	CreatedAt time.Time   `json:"createdAt"`
	Flag      string      `json:"flag"`
	Name      string      `json:"name"`
	ClientIP  string      `json:"clientIp"`
	Subject   string      `json:"subject"`
	Ref       string      `json:"ref"`
	RefID     string      `json:"refId"`
	Ref2      string      `json:"ref2"`
	Ref2ID    string      `json:"ref2Id"`
	Ref3      string      `json:"ref3"`
	Ref3ID    string      `json:"ref3Id"`
	Request   interface{} `json:"request"`
	Response  interface{} `json:"response"`
	Data      interface{} `json:"data"`
}
