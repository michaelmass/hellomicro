package service

import (
	api "github.com/michaelmass/hellomicro/api/proto"
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
