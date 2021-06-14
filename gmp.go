package hcpairing

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type PlaceAPIResult struct {
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

type PlaceAPIResp struct {
	Candidates []PlaceAPIResult `json:"candidates"`
	Status     string           `json:"status"`
}

type GoogleMapAPIGateway interface {
	GetRatingByGoogleMapsAPI(name string) (PlaceAPIResult, error)
}

type googleMapAPIGateway struct {
	apiKey string
}

var (
	GMPGateway = NewGoogleMapAPIGateway()
)

func NewGoogleMapAPIGateway() GoogleMapAPIGateway {
	instance := googleMapAPIGateway{
		apiKey: Config.GetGoogleMapAPIKey(),
	}
	return &instance
}

func (g *googleMapAPIGateway) GetRatingByGoogleMapsAPI(name string) (PlaceAPIResult, error) {
	queryString := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=%s&inputtype=textquery&fields=name,rating&key=%s",
		url.QueryEscape(name),
		g.apiKey,
	)

	req, _ := http.NewRequest("GET", queryString, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		return PlaceAPIResult{}, errors.New("api unavailable")
	}

	result := PlaceAPIResp{}
	json.NewDecoder(resp.Body).Decode(&result)
	if len(result.Candidates) == 0 {
		return PlaceAPIResult{}, errors.New("no match results")
	}
	return result.Candidates[0], nil
}
