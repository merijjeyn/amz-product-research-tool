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
		SearchText string `json:"searchText"`
		DomainCode string `json:"domainCode"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		e := fmt.Errorf("handlers.analyseSearchTerms: Request body could not be parsed:\n%v", err)
		fmt.Println(e)
		return e
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

	return c.JSON(result)
}
