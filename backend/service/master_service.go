package service

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/repository"
)

type (
	IMasterService interface {
		// Get Province & City
		GetAllProvince(ctx context.Context) (dto.ProvincesResponse, error)
		GetAllCity(ctx context.Context, req dto.CityQueryRequest) (dto.CitiesResponse, error)
	}

	MasterService struct {
		masterRepo repository.IMasterRepository
		jwtService IJWTService
	}
)

func NewMasterService(masterRepo repository.IMasterRepository, jwtService IJWTService) *MasterService {
	return &MasterService{
		masterRepo: masterRepo,
		jwtService: jwtService,
	}
}

// Get Province & City
func (ms *MasterService) GetAllProvince(ctx context.Context) (dto.ProvincesResponse, error) {
	data, err := ms.masterRepo.GetAllProvince(ctx, nil)
	if err != nil {
		return dto.ProvincesResponse{}, dto.ErrGetAllProvince
	}

	var datas []dto.ProvinceResponse
	for _, province := range data.Provinces {
		data := dto.ProvinceResponse{
			ID:   &province.ID,
			Name: province.Name,
		}

		datas = append(datas, data)
	}

	return dto.ProvincesResponse{
		Data: datas,
	}, nil
}
func (ms *MasterService) GetAllCity(ctx context.Context, req dto.CityQueryRequest) (dto.CitiesResponse, error) {
	data, err := ms.masterRepo.GetAllCity(ctx, nil, req)
	if err != nil {
		return dto.CitiesResponse{}, dto.ErrGetAllProvince
	}

	var datas []dto.CityResponseCustom
	for _, city := range data.Cities {
		data := dto.CityResponseCustom{
			ID:   &city.ID,
			Name: city.Name,
			Type: city.Type,
		}

		datas = append(datas, data)
	}

	return dto.CitiesResponse{
		Data: datas,
	}, nil
}
