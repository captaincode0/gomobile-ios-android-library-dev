package transformer

import (
	"fmt"
	m "googlemapsearch/model"
)

type SearchTransformer interface {
	Transformer
}

func (search *m.Search) ToMap() (outputMap map[string]string) {
	outputMap["Latitude"] = fmt.Sprintf("%f", search.Latitude)
	outputMap["Longitude"] = fmt.Sprintf("%f", search.Longitude)
	outputMap["Name"] = search.Name
	outputMap["CityName"] = search.CityName
	outputMap["CityCode"] = search.CityCode
	outputMap["PostCode"] = search.PostCode
	outputMap["CountryCode"] = search.CountryCode

	return outputMap
}
