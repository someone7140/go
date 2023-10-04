package server

import (
	"context"
	placeNote "placeNote/src/gen/proto"
	"placeNote/src/useCase/geolocationUseCase"

	"github.com/bufbuild/connect-go"
)

type GeoLocationServer struct{}

func (s *GeoLocationServer) GetLatLonFromAddress(
	ctx context.Context,
	req *connect.Request[placeNote.GetLatLonFromAddressRequest],
) (*connect.Response[placeNote.GetLatLonFromAddressResponse], error) {
	latLonResponse, err := geolocationUseCase.GetLatLonFromAddressByGeocodingApi(req.Msg.Address)
	if err != nil {
		return nil, err
	}
	return latLonResponse, nil
}
