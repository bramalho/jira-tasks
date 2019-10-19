package main

import (
	"os"
	"sort"
	"strings"
	"sync"
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

// GetUserData from jira
func GetUserData() []*User {
	userNames := getUserNames()
	jiraClient := InitClient()

	users := make([]*User, len(userNames))

	var wg sync.WaitGroup
	wg.Add(len(userNames))

	for i, u := range userNames {
		go func(i int, u string) {
			defer wg.Done()

			user, _, err := jiraClient.User.Get(u)

			if err == nil {
				newUser := User{
					Name:       user.DisplayName,
					Avatar:     user.AvatarUrls.One6X16,
					ToDo:       Query(jiraClient, u, "QUERY_TODO"),
					InProgress: Query(jiraClient, u, "QUERY_IN_PROGRESS"),
					ToReview:   Query(jiraClient, u, "QUERY_TO_REVIEW"),
					Done:       Query(jiraClient, u, "QUERY_DONE"),
				}

				users[i] = &newUser
			}
		}(i, u)
	}

	wg.Wait()

	return users
}
