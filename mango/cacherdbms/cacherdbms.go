package cacherdbms

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/khanakia/mangobp/mango/publicid"
	"github.com/ubgo/goutil"
	"gorm.io/gorm"
)

// Cache ...
type Cache struct {
	ID        string `json:"id" gorm:"type:varchar(36);primaryKey;"`
	CreatedAt time.Time
	Key       string `gorm:"type:varchar(100);unique_index"`
	Value     string
	Expires   int
}

func (u *Cache) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.ID) == 0 {
		u.ID = publicid.Must()
	}
	return
}

func (Cache) TableName() string {
	return getTableName()
}

func getTableName() string {
	return goutil.Env("PKG_CACHERDBMS_TABLENAME", "caches")
}

type Rdbms struct {
	Config
}

func (pkg Rdbms) Version() string {
	return "0.01"
}

type Config struct {
	DB *gorm.DB
}

func (pkg Rdbms) MigrateDb() {
	pkg.DB.AutoMigrate(&Cache{})
}

// New initialize
func New(config Config) *Rdbms {
	rdbms := &Rdbms{Config: config}
	return rdbms
}

// ttl - in seconds
func (a *Rdbms) Put(key string, val interface{}, ttl int) (bool, error) {
	p, err := json.Marshal(val)
	if err != nil {
		return false, err
	}

	expire := int(time.Now().UnixNano()/int64(time.Second) + int64(ttl))

	var entity Cache
	res := a.Config.DB.First(&entity, &Cache{Key: key})
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		entity := &Cache{
			Key:     key,
			Value:   string(p),
			Expires: expire,
		}

		err1 := a.Config.DB.Create(entity).Error
		if err1 != nil {
			return false, err
		}
	} else {
		entity.Value = string(p)
		entity.Expires = expire
		a.Config.DB.Save(&entity)
	}

	return true, nil
}

func (a *Rdbms) Get(key string) interface{} {
	var entity Cache
	res := a.Config.DB.First(&entity, &Cache{Key: key})

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	now := int(time.Now().UnixNano() / int64(time.Second))
	if now > entity.Expires {
		a.Del(key)
		return nil
	}

	var v string
	err := json.Unmarshal([]byte(entity.Value), &v)
	if err != nil {
		return entity.Value
	}
	// val, _ := strconv.Unquote(entity.Value)

	return v
}

func (a *Rdbms) Del(key string) {
	a.Config.DB.Where("key = ?", key).Delete(&Cache{})
}

func (a *Rdbms) Flush() {
	a.Config.DB.Exec("Truncate TABLE " + getTableName() + " RESTART IDENTITY;")
}
