package dapp_type

// type Space struct {
// 	core.ModelStr
// 	ByUserID string `json:"byUserId" gorm:"type:varchar(36);"` // create by userid
// 	Name     string `json:"name" gorm:"type:varchar(255);"`
// 	Status   bool   `json:"status"`
// 	APIKey   string `json:"-" gorm:"type:varchar(36)"`
// 	Descr    string `json:"descr" gorm:"type:varchar(255);"`
// }

// func (tbl *Space) FillDefaults() {
// 	if len(tbl.APIKey) <= 0 {
// 		tbl.APIKey = publicid.Must36()
// 	}
// }

// type Entity struct {
// 	core.ModelStr
// 	SpaceID string `json:"spaceId" gorm:"type:varchar(36);"`
// }

// type MyApp struct {
// 	core.ModelStr
// 	Name   string `json:"name" gorm:"type:varchar(255);"`
// 	Slug   string `json:"slug" gorm:"type:varchar(255);"`
// 	Status bool   `json:"status"`
// 	Descr  string `json:"descr" gorm:"type:varchar(255);"`
// 	Icon   string `json:"icon" gorm:"type:varchar(2000);"`
// }
