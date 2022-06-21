package xmail_dm

import "github.com/khanakia/mangobp/mango/util"

const (
	NATS_XMAIL_SEND = "xmail.send"
)

// For now we will add only SMTP as channel i gave the struct name channel because in future
// we can have thrid party api as channel too e.g. sendgrid, twillo, mandrill so these can be used
// contain apikey, apipass as their settings so naming channel make sense
type Channel struct {
	util.ModelStr
	From     string `json:"from" gorm:"type:varchar(255)"`     // default From Name
	Host     string `json:"host" gorm:"type:varchar(100)"`     // smtp host
	Port     uint   `json:"port" gorm:"type:varchar(3)"`       // smtp port
	User     string `json:"user" gorm:"type:varchar(255)"`     // smtp username
	Password string `json:"password" gorm:"type:varchar(255)"` // smtp pass
}

func (Channel) TableName() string {
	return "xmail_channels"
}
