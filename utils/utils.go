package utils

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/fredallotey2000/ciq/models"
)

const (
	serverLogInputPath string = "./input/server_log.csv"
	dateFormat         string = "Mon Jan 02 15:04:05 UTC 2006"
	InputdateFormat    string = "Jan-02-2006"
)

// import logs not struct
func ImportServerLogs() ([]models.ServerLog, error) {
	//initialize a logs slice
	logs := make([]models.ServerLog, 0)
	//read csv log file
	input, err := readCsv(serverLogInputPath)
	if err != nil {
		return nil, err
	}
	// each line in the file is a log
	for _, line := range input {
		dateAndTime, err := time.Parse(dateFormat, line[0])
		// bad csv line continue import
		if err != nil {
			continue
		}
		user := line[1]
		operation := line[2]
		downloadSize, err := strconv.Atoi(line[3])
		// bad csv line continue import
		if err != nil {
			continue
		}
		//create a serverlog to be added to struct
		s := models.ServerLog{
			DateAndTime: dateAndTime,
			Username:    user,
			Operation:   operation,
			Size:        downloadSize,
		}
		//add each log entry to slice
		logs = append(logs, s)

	}

	return logs, nil
}

// readCsv reads file and converts it to an array of strings
func readCsv(filename string) ([][]string, error) {

	//open file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	//read all into a 2D array
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
