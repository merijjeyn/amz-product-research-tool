package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {
	credential := c.Query("credential")
	user, err := getUserWithCredential(credential)
	if err != nil {
		return fmt.Errorf("handlers.login:: getUserWithCredentialFailed:\n%v", err)
	}

	// User found
	email := c.Query("email")
	if user.id == 0 {
		insertUserIntoDB(email, email, credential)
	}

	return c.SendString("Success!")
}

func authenticateUser(credential string) bool {
	user, err := getUserWithCredential(credential)
	if err != nil {
		fmt.Printf("handlers.authenticateUser: \n%v", err)
		return false
	}

	return user.id != 0
}
