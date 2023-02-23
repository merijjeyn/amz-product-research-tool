package main

import (
	"encoding/json"
	"fmt"
)

func convertStructIntoMap(v any) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	jsonData, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("db/mongo.convertStructIntoMap: Error converting struct into json:\n%v", err)
	}
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		return nil, fmt.Errorf("db/mongo.convertStructIntoMap: Error converting json into map:\n%v", err)
	}
	return m, nil
}
