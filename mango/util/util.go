package util

import (
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/khanakia/mangobp/mango/publicid"
)

type ModelStr struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primaryKey;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *ModelStr) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.ID) == 0 {
		u.ID = publicid.Must()
	}
	return
}

type ErrorObject struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Idx     int    `json:"idx"`
}

// Convert github.com/go-ozzo/ozzo-validation errors to proper array of errors for the JSON
func ErroObjToArray(errors error) []interface{} {
	if errors == nil {
		return []interface{}{}
	}

	var errs []interface{}

	i := 0
	for field, v := range errors.(validation.Errors) {

		message := strings.Title(field) + " " + v.Error()
		errs = append(errs, &ErrorObject{
			Field:   field,
			Message: message,
			Idx:     i,
		})

		i++
	}

	return errs
}

/*
  https://gqlgen.com/reference/errors/
*/

// func OzzoErrToGraphqlErrors(errors error, ctx context.Context) {
// 	i := 0
// 	for field, v := range errors.(validation.Errors) {
// 		message := strings.Title(field) + " " + v.Error()
// 		graphql.AddError(ctx, &gqlerror.Error{
// 			Path:    graphql.GetPath(ctx),
// 			Message: message,
// 			Extensions: map[string]interface{}{
// 				"idx":   i,
// 				"field": field,
// 			},
// 		})
// 		i++
// 	}
// }

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
