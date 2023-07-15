package domain

import (
	"context"
	"time"
)

var TargetRepository = []string{
	"golang-migrate/migrate",
	"pressly/goose",
	"amacneil/dbmate",
	"rubenv/sql-migrate",
	"k0kubun/sqldef",
	"schemalex/schemalex",
	"go-pg/migrations",
	"naoina/migu",
	"Boostport/migration",
}

type GitHubRepository interface {
	NowGithubRepoCount(context.Context, string, string) (*GithubRepository, error)
}

type GitHub struct {
	GithubRepositorys        []GithubRepository
	ReadmeDetailsRepositorys []ReadmeDetailsRepository
}

type GithubRepository struct {
	FullName         string    `json:"full_name"`
	URL              string    `json:"html_url"`
	Description      string    `json:"description"`
	StargazersCount  int       `json:"stargazers_count"`
	SubscribersCount int       `json:"subscribers_count"`
	ForksCount       int       `json:"forks_count"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ReadmeDetailsRepository struct {
	RepoName   string
	RepoURL    string
	StarCounts map[string]int
}
