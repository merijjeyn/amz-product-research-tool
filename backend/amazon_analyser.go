package main

import "fmt"

func analyseAmazonSearch(amzApiRespMap map[string]interface{}) AmazonSearchAnalysis {
	var result AmazonSearchAnalysis

	return result
}

func processProducts(amzApiRespMap map[string]interface{}, res *AmazonSearchAnalysis) error {
	searchProductDetails, ok := amzApiRespMap["searchProductDetails"].([]map[string]interface{})
	if !ok {
		return fmt.Errorf("amazon_analyser.processProducts: failed getting searchProductDetails\n")
	}

	for i, prod := range searchProductDetails {

	}
}

func processProduct(prodData map[string]interface{}) AmazonSearchAnalysisProductEntry {
	var res AmazonSearchAnalysisProductEntry

	price, priceOk := prodData["price"].(float64)
	reviewCount, reviewOk := prodData["countReview"].(int)
	ratingString, ratingOk := prodData["productRating"].(string)

	if !priceOk || !reviewOk || !ratingOk {

	}
}
