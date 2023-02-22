package main

type AmazonSearchAnalysis struct {
	Products         []AmazonSearchAnalysisProductEntry
	OpportunityScore float64
	AvgPrice         float64
	AvgReviews       float64
	AvgSellerRating  float64
	AvgMonthlySales  float64
}

type AmazonSearchAnalysisProductEntry struct {
	Name     string
	Brand    string
	Price    float64
	Category string
	Sales    int
	Reviews  int
	Rating   float64
	Sellers  int
	Lqs      float64
}
