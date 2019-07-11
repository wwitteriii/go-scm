package fake

import (
	"net/url"

	"github.com/jenkins-x/go-scm/scm"
)

// NewDefault returns a new fake client.
// The Data object lets you pre-load resources into the fake driver or check for results after the
// scm operations have been performed
func NewDefault() (*scm.Client, *Data) {
	data := NewData()
	client := &wrapper{new(scm.Client)}
	client.BaseURL = &url.URL{
		Host: "fake.com",
		Path: "/",
	}
	// initialize services
	client.Driver = scm.DriverFake

	client.Git = &gitService{client: client, data: data}
	client.Issues = &issueService{client: client, data: data}
	client.Organizations = &organizationService{client: client, data: data}
	client.PullRequests = &pullService{client: client, data: data}
	client.Repositories = &repositoryService{client: client, data: data}
	client.Reviews = &reviewService{client: client, data: data}

	// TODO
	/*
		client.Contents = &contentService{client}
		client.Users = &userService{client}
		client.Webhooks = &webhookService{client}
	*/
	return client.Client, data
}

type wrapper struct {
	*scm.Client
}
