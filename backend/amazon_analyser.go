package main

import (
	"fmt"
	"strconv"
	"strings"
)

func analyseAmazonSearch(amzApiRespMap *map[string]interface{}) AmazonSearchAnalysis {
	var result AmazonSearchAnalysis

	processProducts(amzApiRespMap, &result)

	return result
}

func processProducts(amzApiRespMap *map[string]interface{}, res *AmazonSearchAnalysis) error {
	searchProductDetails, ok := (*amzApiRespMap)["searchProductDetails"].([]map[string]interface{})
	if !ok {
		return fmt.Errorf("amazon_analyser.processProducts: failed getting searchProductDetails\n")
	}

	for _, prod := range searchProductDetails {
		processedProd := processProduct(prod)
		res.Products = append(res.Products, processedProd)
	}

	return nil
}

func processProduct(prodData map[string]interface{}) AmazonSearchAnalysisProductEntry {
	var res AmazonSearchAnalysisProductEntry

	price, ok := prodData["price"].(float64)
	if ok {
		res.Price = price
	}

	reviewCount, ok := prodData["countReview"].(int)
	if ok {
		res.Reviews = reviewCount
	}

	ratingString, ok := prodData["productRating"].(string)
	if ok {
		splitted := strings.Split(ratingString, " ")
		if len(splitted) > 0 {
			rating, error := strconv.ParseFloat(splitted[0], 64)
			if error == nil {
				res.Rating = rating
			}
		}
	}

	return res
}
