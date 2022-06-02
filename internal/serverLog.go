package internal

import (
	"time"

	"github.com/fredallotey2000/ciq/models"
	"github.com/fredallotey2000/ciq/utils"
)

//log struct to be used as receiver function for operations
type logStruct struct {
	logs []models.ServerLog
}

//Query log interface with all required methods
type QueryLogs interface {
	NumberOfUsers() int
	NumberOfUploadsLargerThan(limit int, operation string) int
	NumberOfUploadsPerUserPerDate(username string, dateAndTime time.Time, operation string) int
}

//creates a new instance of queryLogs
func NewQueryLogs() (QueryLogs, error) {
	l, err := utils.ImportServerLogs()
	if err != nil {
		return nil, err
	}
	ls := &logStruct{
		logs: l,
	}
	return ls, nil
}

//Function to return number of distinct users of the solution
func (l *logStruct) NumberOfUsers() int {
	//retrieve distinct users
	d := getDistinctUsers(l.logs)
	return len(d)
}

//function to return number of uplplads greater than a given file size
func (l *logStruct) NumberOfUploadsLargerThan(limit int, operation string) int {
	//first filter by the operation provided by the user
	logsByOperation := filterByOperation(l.logs, operation)
	//second filter on the size specified
	logsBySize := filterBySize(logsByOperation, limit)
	return len(logsBySize)

}

//Function to return the number of uploads or downloads per user per date
func (l *logStruct) NumberOfUploadsPerUserPerDate(username string, dateAndTime time.Time, operation string) int {
	//first filter by username
	logsByUser := filterByUsername(l.logs, username)
	//second filter by operation type, upload or download
	logsByOperation := filterByOperation(logsByUser, operation)
	//last filter by date
	logsByDateAndTime := filterByDateAndTime(logsByOperation, dateAndTime)
	return len(logsByDateAndTime)

}

//Function to filter logs for distinct users
func getDistinctUsers(logs []models.ServerLog) map[string]struct{} {
	users := make(map[string]struct{}, 0)
	for _, v := range logs {
		users[v.Username] = struct{}{}
	}
	return users
}

//Function to filter logs by operation type.i.e upload or download
func filterByOperation(logs []models.ServerLog, operation string) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.Operation == operation {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs
}

//Function to fiter logs by username
func filterByUsername(logs []models.ServerLog, username string) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.Username == username {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs
}

//Function to filter logs by upload or download size
func filterBySize(logs []models.ServerLog, size int) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.Size > size {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs

}

//Function to filter logs by date uploaded or downloaded
func filterByDateAndTime(logs []models.ServerLog, dateAndTime time.Time) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.DateAndTime.Truncate(24 * time.Hour).Equal(dateAndTime.Truncate(24 * time.Hour)) {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs

}
