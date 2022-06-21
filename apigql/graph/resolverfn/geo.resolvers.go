package resolverfn

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/khanakia/mangobp/apigql/graph/model"
	"github.com/khanakia/mangobp/mango/geo/geo_domain"
)

func (r *queryResolver) GeoCountries(ctx context.Context) ([]*model.Country, error) {
	db := r.GormDB.DB

	var records []geo_domain.Country
	err := db.Find(&records).Error
	if err != nil {
		return nil, errors.New("server error")
	}

	var result []*model.Country
	for _, record := range records {
		result = append(result, &model.Country{
			ID:   record.ID,
			Name: record.Name,
		})
	}

	return result, nil
}

func (r *queryResolver) GeoStates(ctx context.Context, countryID *string, filter *model.FilterInput, limit int, offset int, orderBy []*model.SortOrderInput) ([]*model.State, error) {
	query := r.GormDB.DB

	if countryID != nil && len(*countryID) != 0 {
		query = query.Where(geo_domain.State{CountryID: *countryID})
	}

	var records []geo_domain.State
	err := query.Find(&records).Error
	if err != nil {
		return nil, errors.New("server error")
	}

	var result []*model.State
	for _, record := range records {
		result = append(result, &model.State{
			ID:   record.ID,
			Name: record.Name,
		})
	}

	return result, nil
}
