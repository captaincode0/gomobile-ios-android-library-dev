package transformer

import m "googlemapsearch/model"
import "fmt"

type LocationTransformer interface {
	Transformer
}

func (location *m.Location) ToMap() (outputMap map[string]string) {
	outputMap["Latitude"] = fmt.Sprintf("%f", location.Latitude)
	outputMap["Longitude"] = fmt.Sprintf("%f", location.Longitude)

	return outputMap
}
