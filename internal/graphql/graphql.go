package graphql

// Code generated by graphql-codegen-golang ; DO NOT EDIT.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	*http.Client
	Url string
}

// NewClient creates a GraphQL client ready to use.
func NewClient(url string) *Client {
	return &Client{
		Client: &http.Client{},
		Url:    url,
	}
}

type GraphQLOperation struct {
	Query         string          `json:"query"`
	OperationName string          `json:"operationName,omitempty"`
	Variables     json.RawMessage `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   json.RawMessage `json:"data,omitempty"`
	Errors []GraphQLError  `json:"errors,omitempty"`
}

type GraphQLError map[string]interface{}

func (err GraphQLError) Error() string {
	return fmt.Sprintf("graphql: %v", map[string]interface{}(err))
}

func (resp *GraphQLResponse) Error() string {
	if len(resp.Errors) == 0 {
		return ""
	}
	errs := strings.Builder{}
	for _, err := range resp.Errors {
		errs.WriteString(err.Error())
		errs.WriteString("\n")
	}
	return errs.String()
}

func execute(client *http.Client, req *http.Request) (*GraphQLResponse, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return unmarshalGraphQLReponse(body)
}

func unmarshalGraphQLReponse(b []byte) (*GraphQLResponse, error) {
	resp := GraphQLResponse{}
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	if len(resp.Errors) > 0 {
		return &resp, &resp
	}
	return &resp, nil
}

//
// mutation CreateProject($slug: String!, $teamId: String!)
//

type CreateProjectVariables struct {
	Slug   String `json:"slug"`
	TeamId String `json:"teamId"`
}

type CreateProjectResponse struct {
	CreateProject struct {
		ID        string `json:"id"`
		Slug      string `json:"slug"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"createProject"`
}

type CreateProjectRequest struct {
	*http.Request
}

func NewCreateProjectRequest(url string, vars *CreateProjectVariables) (*CreateProjectRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation CreateProject($slug: String!, $teamId: String!) {
  createProject(slug: $slug, teamId: $teamId) {
    id
    slug
    createdAt
    updatedAt
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &CreateProjectRequest{req}, nil
}

func (req *CreateProjectRequest) Execute(client *http.Client) (*CreateProjectResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result CreateProjectResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateProject(url string, client *http.Client, vars *CreateProjectVariables) (*CreateProjectResponse, error) {
	req, err := NewCreateProjectRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) CreateProject(vars *CreateProjectVariables) (*CreateProjectResponse, error) {
	return CreateProject(client.Url, client.Client, vars)
}

//
// mutation CreateStripeCheckoutBillingPortalUrl($teamId: String!)
//

type CreateStripeCheckoutBillingPortalUrlVariables struct {
	TeamId String `json:"teamId"`
}

type CreateStripeCheckoutBillingPortalUrlResponse struct {
	CreateStripeCheckoutBillingPortalUrl string `json:"createStripeCheckoutBillingPortalUrl"`
}

type CreateStripeCheckoutBillingPortalUrlRequest struct {
	*http.Request
}

func NewCreateStripeCheckoutBillingPortalUrlRequest(url string, vars *CreateStripeCheckoutBillingPortalUrlVariables) (*CreateStripeCheckoutBillingPortalUrlRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation CreateStripeCheckoutBillingPortalUrl($teamId: String!) {
  createStripeCheckoutBillingPortalUrl(teamId: $teamId)
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &CreateStripeCheckoutBillingPortalUrlRequest{req}, nil
}

func (req *CreateStripeCheckoutBillingPortalUrlRequest) Execute(client *http.Client) (*CreateStripeCheckoutBillingPortalUrlResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result CreateStripeCheckoutBillingPortalUrlResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateStripeCheckoutBillingPortalUrl(url string, client *http.Client, vars *CreateStripeCheckoutBillingPortalUrlVariables) (*CreateStripeCheckoutBillingPortalUrlResponse, error) {
	req, err := NewCreateStripeCheckoutBillingPortalUrlRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) CreateStripeCheckoutBillingPortalUrl(vars *CreateStripeCheckoutBillingPortalUrlVariables) (*CreateStripeCheckoutBillingPortalUrlResponse, error) {
	return CreateStripeCheckoutBillingPortalUrl(client.Url, client.Client, vars)
}

//
// mutation CreateStripeCheckoutSession($plan: PaidPlan!, $teamId: String!)
//

type CreateStripeCheckoutSessionVariables struct {
	Plan   PaidPlan `json:"plan"`
	TeamId String   `json:"teamId"`
}

type CreateStripeCheckoutSessionResponse struct {
	CreateStripeCheckoutSession string `json:"createStripeCheckoutSession"`
}

type CreateStripeCheckoutSessionRequest struct {
	*http.Request
}

func NewCreateStripeCheckoutSessionRequest(url string, vars *CreateStripeCheckoutSessionVariables) (*CreateStripeCheckoutSessionRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation CreateStripeCheckoutSession($plan: PaidPlan!, $teamId: String!) {
  createStripeCheckoutSession(plan: $plan, teamId: $teamId)
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &CreateStripeCheckoutSessionRequest{req}, nil
}

func (req *CreateStripeCheckoutSessionRequest) Execute(client *http.Client) (*CreateStripeCheckoutSessionResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result CreateStripeCheckoutSessionResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateStripeCheckoutSession(url string, client *http.Client, vars *CreateStripeCheckoutSessionVariables) (*CreateStripeCheckoutSessionResponse, error) {
	req, err := NewCreateStripeCheckoutSessionRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) CreateStripeCheckoutSession(vars *CreateStripeCheckoutSessionVariables) (*CreateStripeCheckoutSessionResponse, error) {
	return CreateStripeCheckoutSession(client.Url, client.Client, vars)
}

//
// mutation CreateTeam($name: String!)
//

type CreateTeamVariables struct {
	Name String `json:"name"`
}

type CreateTeamResponse struct {
	CreateTeam struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		PaidPlan string `json:"paidPlan"`
	} `json:"createTeam"`
}

type CreateTeamRequest struct {
	*http.Request
}

func NewCreateTeamRequest(url string, vars *CreateTeamVariables) (*CreateTeamRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation CreateTeam($name: String!) {
  createTeam(name: $name) {
    id
    name
    slug
    paidPlan
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &CreateTeamRequest{req}, nil
}

func (req *CreateTeamRequest) Execute(client *http.Client) (*CreateTeamResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result CreateTeamResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateTeam(url string, client *http.Client, vars *CreateTeamVariables) (*CreateTeamResponse, error) {
	req, err := NewCreateTeamRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) CreateTeam(vars *CreateTeamVariables) (*CreateTeamResponse, error) {
	return CreateTeam(client.Url, client.Client, vars)
}

//
// query GetProject($id: String, $teamId: String, $slug: String)
//

type GetProjectVariables struct {
	ID     *String `json:"id,omitempty"`
	TeamId *String `json:"teamId,omitempty"`
	Slug   *String `json:"slug,omitempty"`
}

type GetProjectResponse struct {
	Project struct {
		ID        string `json:"id"`
		Slug      string `json:"slug"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"project"`
}

type GetProjectRequest struct {
	*http.Request
}

func NewGetProjectRequest(url string, vars *GetProjectVariables) (*GetProjectRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `query GetProject($id: String, $teamId: String, $slug: String) {
  project(id: $id, teamId: $teamId, slug: $slug) {
    id
    slug
    createdAt
    updatedAt
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &GetProjectRequest{req}, nil
}

func (req *GetProjectRequest) Execute(client *http.Client) (*GetProjectResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result GetProjectResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetProject(url string, client *http.Client, vars *GetProjectVariables) (*GetProjectResponse, error) {
	req, err := NewGetProjectRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) GetProject(vars *GetProjectVariables) (*GetProjectResponse, error) {
	return GetProject(client.Url, client.Client, vars)
}

//
// query GetTeam($id: String, $slug: String)
//

type GetTeamVariables struct {
	ID   *String `json:"id,omitempty"`
	Slug *String `json:"slug,omitempty"`
}

type GetTeamResponse struct {
	Team struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		PaidPlan string `json:"paidPlan"`
	} `json:"team"`
}

type GetTeamRequest struct {
	*http.Request
}

func NewGetTeamRequest(url string, vars *GetTeamVariables) (*GetTeamRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `query GetTeam($id: String, $slug: String) {
  team(id: $id, slug: $slug) {
    id
    name
    slug
    paidPlan
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &GetTeamRequest{req}, nil
}

func (req *GetTeamRequest) Execute(client *http.Client) (*GetTeamResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result GetTeamResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetTeam(url string, client *http.Client, vars *GetTeamVariables) (*GetTeamResponse, error) {
	req, err := NewGetTeamRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) GetTeam(vars *GetTeamVariables) (*GetTeamResponse, error) {
	return GetTeam(client.Url, client.Client, vars)
}

//
// query GetTeamWithUsersAndProjects($id: String, $slug: String)
//

type GetTeamWithUsersAndProjectsVariables struct {
	ID   *String `json:"id,omitempty"`
	Slug *String `json:"slug,omitempty"`
}

type GetTeamWithUsersAndProjectsResponse struct {
	Team struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		PaidPlan string `json:"paidPlan"`
		Users    struct {
			Edges *[]struct {
				Role string `json:"role"`
				Node struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Email string `json:"email"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"users"`
		Projects struct {
			Nodes *[]struct {
				ID        string `json:"id"`
				Slug      string `json:"slug"`
				CreatedAt string `json:"createdAt"`
				UpdatedAt string `json:"updatedAt"`
			} `json:"nodes"`
		} `json:"projects"`
	} `json:"team"`
}

type GetTeamWithUsersAndProjectsRequest struct {
	*http.Request
}

func NewGetTeamWithUsersAndProjectsRequest(url string, vars *GetTeamWithUsersAndProjectsVariables) (*GetTeamWithUsersAndProjectsRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `query GetTeamWithUsersAndProjects($id: String, $slug: String) {
  team(id: $id, slug: $slug) {
    id
    name
    slug
    paidPlan
    users(first: 100) {
      edges {
        role
        node {
          id
          name
          email
        }
      }
    }
    projects(first: 20) {
      nodes {
        id
        slug
        createdAt
        updatedAt
      }
    }
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &GetTeamWithUsersAndProjectsRequest{req}, nil
}

func (req *GetTeamWithUsersAndProjectsRequest) Execute(client *http.Client) (*GetTeamWithUsersAndProjectsResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result GetTeamWithUsersAndProjectsResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetTeamWithUsersAndProjects(url string, client *http.Client, vars *GetTeamWithUsersAndProjectsVariables) (*GetTeamWithUsersAndProjectsResponse, error) {
	req, err := NewGetTeamWithUsersAndProjectsRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) GetTeamWithUsersAndProjects(vars *GetTeamWithUsersAndProjectsVariables) (*GetTeamWithUsersAndProjectsResponse, error) {
	return GetTeamWithUsersAndProjects(client.Url, client.Client, vars)
}

//
// query GetViewer
//

type GetViewerResponse struct {
	Viewer struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Teams struct {
			Nodes *[]struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
				PaidPlan string `json:"paidPlan"`
			} `json:"nodes"`
		} `json:"teams"`
	} `json:"viewer"`
}

type GetViewerRequest struct {
	*http.Request
}

func NewGetViewerRequest(url string) (*GetViewerRequest, error) {
	b, err := json.Marshal(&GraphQLOperation{
		Query: `query GetViewer {
  viewer {
    id
    name
    email
    teams(first: 100) {
      nodes {
        id
        name
        slug
        paidPlan
      }
    }
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &GetViewerRequest{req}, nil
}

func (req *GetViewerRequest) Execute(client *http.Client) (*GetViewerResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result GetViewerResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetViewer(url string, client *http.Client) (*GetViewerResponse, error) {
	req, err := NewGetViewerRequest(url)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) GetViewer() (*GetViewerResponse, error) {
	return GetViewer(client.Url, client.Client)
}

//
// query GetViewerTokens
//

type GetViewerTokensResponse struct {
	Viewer struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Teams struct {
			Nodes *[]struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
				PaidPlan string `json:"paidPlan"`
			} `json:"nodes"`
		} `json:"teams"`
		Tokens struct {
			Nodes *[]struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				CreatedAt string `json:"createdAt"`
			} `json:"nodes"`
		} `json:"tokens"`
	} `json:"viewer"`
}

type GetViewerTokensRequest struct {
	*http.Request
}

func NewGetViewerTokensRequest(url string) (*GetViewerTokensRequest, error) {
	b, err := json.Marshal(&GraphQLOperation{
		Query: `query GetViewerTokens {
  viewer {
    id
    name
    email
    teams(first: 100) {
      nodes {
        id
        name
        slug
        paidPlan
      }
    }
    tokens(first: 100) {
      nodes {
        id
        name
        createdAt
      }
    }
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &GetViewerTokensRequest{req}, nil
}

func (req *GetViewerTokensRequest) Execute(client *http.Client) (*GetViewerTokensResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result GetViewerTokensResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetViewerTokens(url string, client *http.Client) (*GetViewerTokensResponse, error) {
	req, err := NewGetViewerTokensRequest(url)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) GetViewerTokens() (*GetViewerTokensResponse, error) {
	return GetViewerTokens(client.Url, client.Client)
}

//
// mutation InviteToTeam($email: String!, $teamId: String!)
//

type InviteToTeamVariables struct {
	Email  String `json:"email"`
	TeamId String `json:"teamId"`
}

type InviteToTeamResponse struct {
	InviteToTeam string `json:"inviteToTeam"`
}

type InviteToTeamRequest struct {
	*http.Request
}

func NewInviteToTeamRequest(url string, vars *InviteToTeamVariables) (*InviteToTeamRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation InviteToTeam($email: String!, $teamId: String!) {
  inviteToTeam(email: $email, teamId: $teamId)
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &InviteToTeamRequest{req}, nil
}

func (req *InviteToTeamRequest) Execute(client *http.Client) (*InviteToTeamResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result InviteToTeamResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func InviteToTeam(url string, client *http.Client, vars *InviteToTeamVariables) (*InviteToTeamResponse, error) {
	req, err := NewInviteToTeamRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) InviteToTeam(vars *InviteToTeamVariables) (*InviteToTeamResponse, error) {
	return InviteToTeam(client.Url, client.Client, vars)
}

//
// mutation RemoveUserFromTeam($userId: String!, $teamId: String!)
//

type RemoveUserFromTeamVariables struct {
	UserId String `json:"userId"`
	TeamId String `json:"teamId"`
}

type RemoveUserFromTeamResponse struct {
	RemoveUserFromTeam struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		PaidPlan string `json:"paidPlan"`
		Users    struct {
			Nodes *[]struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"nodes"`
		} `json:"users"`
	} `json:"removeUserFromTeam"`
}

type RemoveUserFromTeamRequest struct {
	*http.Request
}

func NewRemoveUserFromTeamRequest(url string, vars *RemoveUserFromTeamVariables) (*RemoveUserFromTeamRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation RemoveUserFromTeam($userId: String!, $teamId: String!) {
  removeUserFromTeam(userId: $userId, teamId: $teamId) {
    id
    name
    slug
    paidPlan
    users(first: 100) {
      nodes {
        id
        name
        email
      }
    }
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &RemoveUserFromTeamRequest{req}, nil
}

func (req *RemoveUserFromTeamRequest) Execute(client *http.Client) (*RemoveUserFromTeamResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result RemoveUserFromTeamResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func RemoveUserFromTeam(url string, client *http.Client, vars *RemoveUserFromTeamVariables) (*RemoveUserFromTeamResponse, error) {
	req, err := NewRemoveUserFromTeamRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) RemoveUserFromTeam(vars *RemoveUserFromTeamVariables) (*RemoveUserFromTeamResponse, error) {
	return RemoveUserFromTeam(client.Url, client.Client, vars)
}

//
// mutation UpdateTeam($name: String, $slug: String, $teamId: String!)
//

type UpdateTeamVariables struct {
	Name   *String `json:"name,omitempty"`
	Slug   *String `json:"slug,omitempty"`
	TeamId String  `json:"teamId"`
}

type UpdateTeamResponse struct {
	UpdateTeam struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		PaidPlan string `json:"paidPlan"`
	} `json:"updateTeam"`
}

type UpdateTeamRequest struct {
	*http.Request
}

func NewUpdateTeamRequest(url string, vars *UpdateTeamVariables) (*UpdateTeamRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation UpdateTeam($name: String, $slug: String, $teamId: String!) {
  updateTeam(name: $name, slug: $slug, teamId: $teamId) {
    id
    name
    slug
    paidPlan
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &UpdateTeamRequest{req}, nil
}

func (req *UpdateTeamRequest) Execute(client *http.Client) (*UpdateTeamResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result UpdateTeamResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateTeam(url string, client *http.Client, vars *UpdateTeamVariables) (*UpdateTeamResponse, error) {
	req, err := NewUpdateTeamRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) UpdateTeam(vars *UpdateTeamVariables) (*UpdateTeamResponse, error) {
	return UpdateTeam(client.Url, client.Client, vars)
}

//
// mutation UpdateUser($name: String, $userId: String!)
//

type UpdateUserVariables struct {
	Name   *String `json:"name,omitempty"`
	UserId String  `json:"userId"`
}

type UpdateUserResponse struct {
	UpdateUser struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Teams struct {
			Nodes *[]struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
				PaidPlan string `json:"paidPlan"`
			} `json:"nodes"`
		} `json:"teams"`
	} `json:"updateUser"`
}

type UpdateUserRequest struct {
	*http.Request
}

func NewUpdateUserRequest(url string, vars *UpdateUserVariables) (*UpdateUserRequest, error) {
	variables, err := json.Marshal(vars)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(&GraphQLOperation{
		Variables: variables,
		Query: `mutation UpdateUser($name: String, $userId: String!) {
  updateUser(name: $name, userId: $userId) {
    id
    name
    email
    teams(first: 100) {
      nodes {
        id
        name
        slug
        paidPlan
      }
    }
  }
}`,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return &UpdateUserRequest{req}, nil
}

func (req *UpdateUserRequest) Execute(client *http.Client) (*UpdateUserResponse, error) {
	resp, err := execute(client, req.Request)
	if err != nil {
		return nil, err
	}
	var result UpdateUserResponse
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateUser(url string, client *http.Client, vars *UpdateUserVariables) (*UpdateUserResponse, error) {
	req, err := NewUpdateUserRequest(url, vars)
	if err != nil {
		return nil, err
	}
	return req.Execute(client)
}

func (client *Client) UpdateUser(vars *UpdateUserVariables) (*UpdateUserResponse, error) {
	return UpdateUser(client.Url, client.Client, vars)
}

//
// Scalars
//

type Int int32
type Float float64
type Boolean bool
type String string
type ID string
type DateTime string

//
// Enums
//

type PaidPlan string

const (
	PaidPlanPro PaidPlan = "pro"
)

//
// Inputs
//

//
// Objects
//

type Mutation struct {
	CreateProject                        *Project `json:"createProject,omitempty"`
	CreateStripeCheckoutBillingPortalUrl *String  `json:"createStripeCheckoutBillingPortalUrl,omitempty"`
	CreateStripeCheckoutSession          *String  `json:"createStripeCheckoutSession,omitempty"`
	CreateTeam                           *Team    `json:"createTeam,omitempty"`
	DeleteToken                          *Token   `json:"deleteToken,omitempty"`
	InviteToTeam                         *Boolean `json:"inviteToTeam,omitempty"`
	RemoveUserFromTeam                   *Team    `json:"removeUserFromTeam,omitempty"`
	UpdateTeam                           *Team    `json:"updateTeam,omitempty"`
	UpdateUser                           *User    `json:"updateUser,omitempty"`
}

type PageInfo struct {
	EndCursor       *String `json:"endCursor,omitempty"`
	HasNextPage     Boolean `json:"hasNextPage"`
	HasPreviousPage Boolean `json:"hasPreviousPage"`
	StartCursor     *String `json:"startCursor,omitempty"`
}

type Project struct {
	CreatedAt DateTime `json:"createdAt"`
	ID        String   `json:"id"`
	Slug      String   `json:"slug"`
	Team      Team     `json:"team"`
	TeamId    String   `json:"teamId"`
	UpdatedAt DateTime `json:"updatedAt"`
}

type ProjectConnection struct {
	Edges    *[]ProjectEdge `json:"edges,omitempty"`
	Nodes    *[]Project     `json:"nodes,omitempty"`
	PageInfo PageInfo       `json:"pageInfo"`
}

type ProjectEdge struct {
	Cursor String   `json:"cursor"`
	Node   *Project `json:"node,omitempty"`
}

type Query struct {
	Project *Project `json:"project,omitempty"`
	Team    *Team    `json:"team,omitempty"`
	Token   *Token   `json:"token,omitempty"`
	Viewer  *User    `json:"viewer,omitempty"`
}

type Team struct {
	CreatedAt DateTime             `json:"createdAt"`
	ID        String               `json:"id"`
	Name      String               `json:"name"`
	PaidPlan  *PaidPlan            `json:"paidPlan,omitempty"`
	Projects  *ProjectConnection   `json:"projects,omitempty"`
	Slug      String               `json:"slug"`
	UpdatedAt DateTime             `json:"updatedAt"`
	Users     *TeamUsersConnection `json:"users,omitempty"`
}

type TeamUsersConnection struct {
	Edges    *[]TeamUsersEdge `json:"edges,omitempty"`
	Nodes    *[]User          `json:"nodes,omitempty"`
	PageInfo PageInfo         `json:"pageInfo"`
}

type TeamUsersEdge struct {
	Cursor String  `json:"cursor"`
	Node   *User   `json:"node,omitempty"`
	Role   *String `json:"role,omitempty"`
}

type Token struct {
	CreatedAt DateTime `json:"createdAt"`
	ID        String   `json:"id"`
	Name      String   `json:"name"`
	User      *User    `json:"user,omitempty"`
	UserId    *String  `json:"userId,omitempty"`
}

type TokenConnection struct {
	Edges    *[]TokenEdge `json:"edges,omitempty"`
	Nodes    *[]Token     `json:"nodes,omitempty"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type TokenEdge struct {
	Cursor String `json:"cursor"`
	Node   *Token `json:"node,omitempty"`
}

type User struct {
	CreatedAt DateTime             `json:"createdAt"`
	Email     String               `json:"email"`
	ID        String               `json:"id"`
	Name      *String              `json:"name,omitempty"`
	Team      *Team                `json:"team,omitempty"`
	Teams     *UserTeamsConnection `json:"teams,omitempty"`
	Tokens    *TokenConnection     `json:"tokens,omitempty"`
	UpdatedAt DateTime             `json:"updatedAt"`
}

type UserTeamsConnection struct {
	Edges      *[]UserTeamsEdge `json:"edges,omitempty"`
	Nodes      *[]Team          `json:"nodes,omitempty"`
	PageInfo   PageInfo         `json:"pageInfo"`
	TotalCount *Int             `json:"totalCount,omitempty"`
}

type UserTeamsEdge struct {
	Cursor String  `json:"cursor"`
	Node   *Team   `json:"node,omitempty"`
	Role   *String `json:"role,omitempty"`
}
