package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func searchProductsAxesso(searchText string, domainCode string, page int, excludeSponsored bool) (map[string]interface{}, error) {
	if searchText == "" || domainCode == "" || page == 0 {
		return nil, fmt.Errorf("api.searchProductsAxesso: Invalid parameters passed")
	}

	pageStr := strconv.Itoa(page)
	excludeSponsoredStr := strconv.FormatBool(excludeSponsored)

	url := fmt.Sprintf(
		"https://axesso-axesso-amazon-data-service-v1.p.rapidapi.com/amz/amazon-search-by-keyword-asin?domainCode=%s&keyword=%s&page=%s&excludeSponsored=%s&sortBy=relevanceblender&withCache=true",
		domainCode, searchText, pageStr, excludeSponsoredStr,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("api.SearchProductsAxesso: Error creating new request:\n%v", err)
	}

	req.Header.Add("X-RapidAPI-Key", "1a286522e7mshb6e48f0e32c3f44p1c4440jsnff3a804db90e")
	req.Header.Add("X-RapidAPI-Host", "axesso-axesso-amazon-data-service-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("api.SearchProductsAxesso: Error sending request:\n%v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("api.SearchProductsAxesso: Error reading response body:\n%v", err)
	}

	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, fmt.Errorf("api.SearchProductsAxesso: Error converting resp body to json\n%v", err)
	}

	return m, nil
}
