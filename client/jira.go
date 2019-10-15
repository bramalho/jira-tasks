package client

import (
	"os"
	"strings"

	"github.com/andygrunwald/go-jira"
)

// InitClient for Jira
func InitClient() *jira.Client {
	base, _ := os.LookupEnv("ATLASSIAN_URL")
	email, _ := os.LookupEnv("ATLASSIAN_EMAIL")
	token, _ := os.LookupEnv("ATLASSIAN_TOKEN")

	tp := jira.BasicAuthTransport{
		Username: email,
		Password: token,
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}

	return jiraClient
}

// Query JQL
func Query(c *jira.Client, u string, q string) int {
	query, _ := os.LookupEnv(q)
	query = strings.Replace(query, "%user%", u, 1)
	issues, _, _ := c.Issue.Search(query, nil)

	return len(issues)
}
