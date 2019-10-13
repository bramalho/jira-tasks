package service

import (
	"os"
	"strings"

	"github.com/bramalho/jira-tasks/model"
)

func getUsers() []model.User {
	users := []model.User{}
	if data, exists := os.LookupEnv("USERS"); exists {
		result := strings.Split(data, ",")
		for i := range result {
			if len(result[i]) > 0 {
				u := model.User{Name: result[i]}
				users = append(users, u)
			}
		}
	}

	return users
}

// New jira service
func New() []model.User {
	users := getUsers()
	return users
}
