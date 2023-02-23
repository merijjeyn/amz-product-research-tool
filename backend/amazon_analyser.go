package main

import "fmt"

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
		res.Products = append(res.Products, *processedProd)
	}

	return nil
}

func processProduct(prodData map[string]interface{}) *AmazonSearchAnalysisProductEntry {
	// var res AmazonSearchAnalysisProductEntry

	// price, priceOk := prodData["price"].(float64)
	// reviewCount, reviewOk := prodData["countReview"].(int)
	// ratingString, ratingOk := prodData["productRating"].(string)

	// if !priceOk || !reviewOk || !ratingOk {
	// 	// !!LEFT HERE
	// }
	return new(AmazonSearchAnalysisProductEntry)
}
