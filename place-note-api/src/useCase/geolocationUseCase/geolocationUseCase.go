package geolocationUseCase

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	placeNote "placeNote/src/gen/proto"

	"github.com/bufbuild/connect-go"
)

func GetLatLonFromAddressByGeocodingApi(address string) (*connect.Response[placeNote.GetLatLonFromAddressResponse], *connect.Error) {
	// Geocoding.jp APIから結果を取得
	resp, err := http.Get("https://www.geocoding.jp/api/?q=" + url.QueryEscape(address))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	defer resp.Body.Close()

	respByteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// xmlをパース
	type GeocodingXML struct {
		XMLName xml.Name `xml:"result"`
		Lat     *float64 `xml:"coordinate>lat"`
		Lon     *float64 `xml:"coordinate>lng"`
	}
	parseResult := GeocodingXML{}
	err = xml.Unmarshal(respByteArray, &parseResult)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if parseResult.Lat != nil && parseResult.Lon != nil {
		latLon := placeNote.LatLon{
			Lat: *parseResult.Lat,
			Lon: *parseResult.Lon,
		}
		return connect.NewResponse(&placeNote.GetLatLonFromAddressResponse{
			LatLon: &latLon,
		}), nil
	} else {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
}
