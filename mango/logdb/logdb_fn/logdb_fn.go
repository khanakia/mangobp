package logdb_fn

import (
	"errors"
	"reflect"

	"github.com/goccy/go-json"
	"github.com/khanakia/mangobp/mango/logdb/logdb_domain"
	"gorm.io/gorm"
)

func interfaceToString(val interface{}) string {
	if val == nil {
		return ""
	}

	xt := reflect.TypeOf(val).Kind().String()

	// fmt.Println("XT", xt)
	var p []byte
	var err error
	if xt == "map" {
		p, err = json.Marshal(val)
	} else {
		p = []byte(val.(string))
	}

	if err != nil {
		return ""
	}

	return string(p)
}

func Create(args logdb_domain.CreateArgs, db *gorm.DB) (*logdb_domain.Log, error) {
	if len(args.Flag) == 0 {
		return nil, errors.New("flag is required")
	}

	// Prevent sending the "null" value and if data is empty then do not marshal otherwise it will add double quotes
	// in database "{}" instead of {}
	data := []byte("{}")
	if args.Data != nil {
		data, _ = json.Marshal(args.Data)
	}

	record := &logdb_domain.Log{
		CreatedAt: args.CreatedAt,
		Flag:      args.Flag,
		Name:      args.Name,
		ClientIP:  args.ClientIP,
		Subject:   args.Subject,
		Ref:       args.Ref,
		RefID:     args.RefID,
		Ref2:      args.Ref2,
		Ref2ID:    args.Ref2ID,
		Ref3:      args.Ref3,
		Ref3ID:    args.Ref3ID,
		Request:   interfaceToString(args.Request),
		Response:  interfaceToString(args.Response),
		Data:      data,
	}
	err := db.Save(record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}
