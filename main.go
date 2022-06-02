package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fredallotey2000/ciq/internal"
	"github.com/fredallotey2000/ciq/utils"
)

func main() {

	// Define the flags for this program
	dateAndTime := flag.String("d", "", "Operation date")
	operation := flag.String("o", "", "Operation")
	username := flag.String("u", "", "Username of user")
	size := flag.Int("s", 0, "Download size")
	help := flag.Bool("help", false, "Display help")
	flag.Parse()

	//help to show user how to use the program
	if *help {
		fmt.Println("This program allows a server administrator to gather metrics from server logs")
		fmt.Println("Usage")
		fmt.Println("Usage: go run main.go with <OPTIONS> or ./main with <OPTIONS>")
		fmt.Println("Options")
		fmt.Println("   -d Operation date")
		fmt.Println("   -o Operation name")
		fmt.Println("   -u Username of user")
		fmt.Println("   -s Download or Upload size")
		fmt.Println("   -u all --This flag specifies all users.eg. -a all")
		os.Exit(0)
	}
	//default operation is upload, this is used when user doesnt specify the opertion
	var operatn string = "upload"
	if *operation != "" {
		operatn = *operation
	}
	//get a new instance of the querylogs
	queryLogs, err := internal.NewQueryLogs()
	if err != nil {
		panic(err)
	}
	//check input provided by user for date, username or size and call appropriate method
	if *dateAndTime != "" && *username != "" {
		parsedDate, err := time.Parse(utils.InputdateFormat, *dateAndTime)
		if err != nil {
			panic(err)
		}

		results := queryLogs.NumberOfUploadsPerUserPerDate(*username, parsedDate, operatn)
		fmt.Printf("Number of %v for %v on %v is %v", operatn, *username, strings.Title(*dateAndTime), results)
	} else if !(*size == 0) {
		results := queryLogs.NumberOfUploadsLargerThan(*size, operatn)
		fmt.Printf("Number of %v with size greater than %v is %v", operatn, *size, results)
	} else if !(*username == "") {
		results := queryLogs.NumberOfUsers()
		fmt.Printf("Number of users is %v", results)
	}

}
