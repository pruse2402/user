package controllers

import (
	"net/http"
	"net/url"
	"strconv"
	"user/users/models"
)

func getandConvertToInt64(query url.Values, str string) int64 {
	intData, _ := strconv.ParseInt(query.Get(str), 10, 64)
	return intData
}

// ListQueryParams which the arams struct
func ListQueryParams(r *http.Request) (models.QueryParams, error) {
	query := r.URL.Query()
	queryParams := models.QueryParams{
		CompanyName:   query.Get("companyName"),
		FirstName:     query.Get("firstName"),
		LastName:      query.Get("lastName"),
		LegalEntityID: getandConvertToInt64(query, "legalEntityID"),
	}

	return queryParams, nil
}
