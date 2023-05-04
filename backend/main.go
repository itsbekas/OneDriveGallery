package main

import (
	"fmt"
	"log"

	"github.com/itsbekas/onedrivegallery/graphhelper"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Graph Tutorial")

	fmt.Println()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	graphHelper := graphhelper.NewGraphHelper()

	initializeGraph(graphHelper)

	greetUser(graphHelper)

	var choice int64 = -1

	for {
		fmt.Println("Please choose one of the following options:")
		fmt.Println("0. Exit")
		fmt.Println("1. Display access token")
		fmt.Println("2. List my inbox")
		fmt.Println("3. Display drive details")

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			// Exit the program
			fmt.Println("Goodbye...")
		case 1:
			// Display access token
			displayAccessToken(graphHelper)
		case 2:
			// List emails from user's inbox
			listInbox(graphHelper)
		case 3:
			// Send an email message
			displayDrive(graphHelper)
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		if choice == 0 {
			break
		}
	}
}

func initializeGraph(graphHelper *graphhelper.GraphHelper) {
	err := graphHelper.InitializeGraphForUserAuth()
	if err != nil {
		log.Panicf("Error initializing Graph for user auth: %v\n", err)
	}
}

func greetUser(graphHelper *graphhelper.GraphHelper) {
	user, err := graphHelper.GetUser()
	if err != nil {
		log.Panicf("Error getting user: %v\n", err)
	}

	fmt.Printf("Hello, %s!\n", *user.GetDisplayName())

	// For Work/school accounts, email is in Mail property
	// Personal accounts, email is in UserPrincipalName
	email := user.GetMail()
	if email == nil {
		email = user.GetUserPrincipalName()
	}

	fmt.Printf("Email: %s\n", *email)
	fmt.Println()
}

func displayAccessToken(graphHelper *graphhelper.GraphHelper) {
	token, err := graphHelper.GetUserToken()
	if err != nil {
		log.Panicf("Error getting user token: %v\n", err)
	}

	fmt.Printf("User token: %s", *token)
	fmt.Println()
}

func displayDrive(graphHelper *graphhelper.GraphHelper) {
	drive, err := graphHelper.GetDrive()
	if err != nil {
		log.Panicf("Error getting user's drive: %v", err)
	}

	// Output drive details
	fmt.Printf("Drive ID: %s\n", *drive.GetId())
	fmt.Printf("Drive owner: %s\n", *drive.GetOwner().GetUser().GetDisplayName())
	fmt.Printf("Drive quota: %d bytes\n", *drive.GetQuota().GetTotal())
	fmt.Printf("Drive used: %d bytes\n", *drive.GetQuota().GetUsed())
	fmt.Println()
}

func displayDrives(graphHelper *graphhelper.GraphHelper) {
	drives, err := graphHelper.GetDrives()
	if err != nil {
		log.Panicf("Error getting user's drives: %v", err)
	}

	// Output drive details
	for _, drive := range *drives.value {
		fmt.Printf("Drive ID: %s\n", *drive.GetId())
		fmt.Printf("Drive owner: %s\n", *drive.GetOwner().GetUser().GetDisplayName())
		fmt.Printf("Drive quota: %d bytes\n", *drive.GetQuota().GetTotal())
		fmt.Printf("Drive used: %d bytes\n", *drive.GetQuota().GetUsed())
		fmt.Println()
	}
}
