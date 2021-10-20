package service

import (
	"github.com/michaelmass/hellomicro/api"
)

func MapSAToMapApiSA(input map[string][]string) map[string]*api.StringArray {
	ouput := map[string]*api.StringArray{}

	for key, values := range input {
		ouput[key] = &api.StringArray{
			Values: values,
		}
	}

	return ouput
}
