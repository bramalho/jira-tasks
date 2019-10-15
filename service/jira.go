package service

import (
	"os"
	"sort"
	"strings"
	"sync"

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

	sort.Strings(users)

	return users
}

// New jira service
func New() []*model.User {
	userNames := getUserNames()
	jiraClient := client.InitClient()

	users := make([]*model.User, len(userNames))

	var wg sync.WaitGroup
	wg.Add(len(userNames))

	for i, u := range userNames {
		go func(i int, u string) {
			defer wg.Done()

			user, _, err := jiraClient.User.Get(u)

			if err == nil {
				newUser := model.User{
					Name:       user.DisplayName,
					Avatar:     user.AvatarUrls.One6X16,
					ToDo:       client.Query(jiraClient, u, "QUERY_TODO"),
					InProgress: client.Query(jiraClient, u, "QUERY_IN_PROGRESS"),
					ToReview:   client.Query(jiraClient, u, "QUERY_TO_REVIEW"),
					Done:       client.Query(jiraClient, u, "QUERY_DONE"),
				}

				users[i] = &newUser
			}
		}(i, u)
	}

	wg.Wait()

	return users
}
