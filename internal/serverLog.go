package internal

import (
	"time"

	"github.com/fredallotey2000/ciq/models"
	"github.com/fredallotey2000/ciq/utils"
)

type logStruct struct {
	logs []models.ServerLog
}

type QueryLogs interface {
	NumberOfUsers() int
	NumberOfUploadsLargerThan(limit int, operation string) int
	NumberOfUploadsPerUserPerDate(username string, dateAndTime time.Time, operation string) int
}

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

func (l *logStruct) NumberOfUsers() int {

	d := getDistinctUsers(l.logs)
	return len(d)
}

func (l *logStruct) NumberOfUploadsLargerThan(limit int, operation string) int {
	logsByOperation := filterByOperation(l.logs, operation)
	logsBySize := filterBySize(logsByOperation, limit)
	return len(logsBySize)

}

func (l *logStruct) NumberOfUploadsPerUserPerDate(username string, dateAndTime time.Time, operation string) int {

	logsByUser := filterByUsername(l.logs, username)
	logsByOperation := filterByOperation(logsByUser, operation)
	logsByDateAndTime := filterByDateAndTime(logsByOperation, dateAndTime)
	return len(logsByDateAndTime)

}

func getDistinctUsers(logs []models.ServerLog) map[string]struct{} {
	users := make(map[string]struct{}, 0)
	for _, v := range logs {
		users[v.Username] = struct{}{}
	}
	return users
}

func filterByOperation(logs []models.ServerLog, operation string) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.Operation == operation {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs
}

func filterByUsername(logs []models.ServerLog, username string) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.Username == username {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs
}

func filterBySize(logs []models.ServerLog, size int) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.Size > size {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs

}

func filterByDateAndTime(logs []models.ServerLog, dateAndTime time.Time) []models.ServerLog {
	newLogs := make([]models.ServerLog, 0)
	for _, v := range logs {
		if v.DateAndTime.Truncate(24 * time.Hour).Equal(dateAndTime.Truncate(24 * time.Hour)) {
			newLogs = append(newLogs, v)
		}
	}
	return newLogs

}
