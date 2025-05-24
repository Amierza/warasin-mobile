package dto

import (
	"errors"

	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
)

const (
	// failed
	MESSAGE_FAILED_GET_LIST_CITY     = "failed get list city"
	MESSAGE_FAILED_GET_LIST_PROVINCE = "failed get list province"

	// success
	MESSAGE_SUCCESS_GET_LIST_CITY     = "success get list city"
	MESSAGE_SUCCESS_GET_LIST_PROVINCE = "success get list province"
)

var (
	ErrGetCityByID    = errors.New("failed get city by id")
	ErrGetAllProvince = errors.New("failed get list province")
)

type (
	ProvinceResponse struct {
		ID   *uuid.UUID `json:"province_id"`
		Name string     `json:"province_name"`
	}

	CityResponse struct {
		ID       *uuid.UUID       `json:"city_id"`
		Name     string           `json:"city_name"`
		Type     string           `json:"city_type"`
		Province ProvinceResponse `json:"province"`
	}

	ProvincesResponse struct {
		Data []ProvinceResponse `json:"data"`
	}

	CityResponseCustom struct {
		ID   *uuid.UUID `json:"city_id"`
		Name string     `json:"city_name"`
		Type string     `json:"city_type"`
	}

	CityQueryRequest struct {
		ProvinceID string `json:"province_id" form:"province_id"`
	}

	CitiesResponse struct {
		Data []CityResponseCustom `json:"data"`
	}

	AllProvinceRepositoryResponse struct {
		Provinces []entity.Province
	}

	AllCityRepositoryResponse struct {
		Cities []entity.City
	}
)
