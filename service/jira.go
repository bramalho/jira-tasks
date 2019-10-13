package service

import (
	"os"
	"strings"

	"github.com/bramalho/jira-tasks/client"
	"github.com/bramalho/jira-tasks/model"
)

func getUserNames() []string {
	users := []string{}
	if data, exists := os.LookupEnv("USERS"); exists {
		result := strings.Split(data, ",")
		for i := range result {
			if len(result[i]) > 0 {
				users = append(users, result[i])
			}
		}
	}

	return users
}

// New jira service
func New() []model.User {
	users := []model.User{}
	userNames := getUserNames()
	jiraClient := client.InitClient()

	for _, u := range userNames {
		user, _, err := jiraClient.User.Get(u)

		if err == nil {
			users = append(users, model.User{
				Name:       user.DisplayName,
				Avatar:     user.AvatarUrls.One6X16,
				ToDo:       client.Query(jiraClient, u, "QUERY_TODO"),
				InProgress: client.Query(jiraClient, u, "QUERY_IN_PROGRESS"),
				ToReview:   client.Query(jiraClient, u, "QUERY_TO_REVIEW"),
				Done:       client.Query(jiraClient, u, "QUERY_DONE"),
			})
		}
	}

	return users
}
