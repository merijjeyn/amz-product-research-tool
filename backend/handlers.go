package main

import (
	"database/sql"
	"example/hello/db"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	// credential := c.Query("credential")
	payload := struct {
		Gid   string `json:"gid"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		e := fmt.Errorf("handlers.login:: Request body could not be parsed:\n%v", err)
		fmt.Println(e)
		return e
	}

	_, err := db.GetUserWithGid(payload.Gid)
	if err == sql.ErrNoRows {
		db.InsertUserIntoDB(payload.Email, payload.Name, payload.Gid)
	} else if err != nil {
		e := fmt.Errorf("handlers.login:: getUserWithCredentialFailed:\n%v", err)
		fmt.Print(e)
		return e
	}

	return c.SendString("Success!")
}

func authenticateUser(gid string) bool {
	_, err := db.GetUserWithGid(gid)
	if err != nil {
		fmt.Printf("handlers.authenticateUser: \n%v", err)
		return false
	}
	return true
}

func analyseSearchTerms(c *fiber.Ctx) error {
	payload := struct {
		Gid        string `json:"gid"`
		SearchText string `json:"searchText"`
		DomainCode string `json:"domainCode"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		e := fmt.Errorf("handlers.analyseSearchTerms: Request body could not be parsed:\n%v", err)
		fmt.Println(e)
		return e
	}

	if !authenticateUser(payload.Gid) {
		return fmt.Errorf("handlers.analyseSearchTerms: Error authenticating user with google id")
	}

	amzApiRespMap, err := searchProductsAxesso(payload.SearchText, payload.DomainCode, 1, true)
	if err != nil {
		e := fmt.Errorf("handlers.analyseSearchTerms: Error fetching info from amazon product data api:\n%v", err)
		fmt.Println(e)
		return e
	}

	result := analyseAmazonSearch(&amzApiRespMap)

	resultMap, err := convertStructIntoMap(result)
	if err != nil {
		e := fmt.Errorf("handlers.analyseSearchTerms: Error converting result into map:\n%v", err)
		fmt.Println(e)
		return e
	}

	err = db.SaveData_MDB("searchAnalysisResults", payload.SearchText, resultMap)
	if err != nil {
		e := fmt.Errorf("handlers.analyseSearchTerms: Error saving analysis on mongodb:\n%v", err)
		fmt.Println(e)
		return e
	}

	err = db.InsertUserSearchEntryIntoDB_SQL(payload.Gid, payload.SearchText)
	if err != nil {
		e := fmt.Errorf("handlers.analyseSearchTerms: Error saving user search to sql:\n%v", err)
		fmt.Println(e)
		return e
	}

	return c.JSON(result)
}

func getUserSearches(c *fiber.Ctx) error {
	payload := struct {
		Gid string `json:"gid"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		e := fmt.Errorf("handlers.getUserSearches: Request body could not be parsed:\n%v", err)
		fmt.Println(e)
		return e
	}

	if !authenticateUser(payload.Gid) {
		e := fmt.Errorf("handlers.getUserSearches: Error authenticating user with google id")
		fmt.Println(e)
		return e
	}

	searches, err := db.GetUserSearches_SQL(payload.Gid)
	if err != nil {
		e := fmt.Errorf("handlers.getUserSearches: Error getting users from sql:\n%v", err)
		fmt.Println(e)
		return e
	}

	return c.JSON(searches)
}

func getUserSearchAnalysis(c *fiber.Ctx) error {
	payload := struct {
		Gid        string `json:"gid"`
		SearchText string `json:"searchText"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		e := fmt.Errorf("handlers.getUserSearchAnalysis: Request body could not be parsed:\n%v", err)
		fmt.Println(e)
		return e
	}

	if !authenticateUser(payload.Gid) {
		e := fmt.Errorf("handlers.getUserSearchAnalysis: Error authenticating user with google id")
		fmt.Println(e)
		return e
	}

	res, err := db.FetchSearchAnalysis_MDB(payload.SearchText)
	if err != nil {
		e := fmt.Errorf("handlers.getUserSearchAnalysis: Error fetching analysis from mongo:\n%v", err)
		fmt.Println(e)
		return e
	}

	return c.JSON(res["data"])
}
