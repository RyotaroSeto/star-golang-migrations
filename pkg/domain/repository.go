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
	NowGithubRepoCount(context.Context, string, string) (GithubRepository, error)
	GetStargazersCountByRepo(context.Context, string, string, GithubRepository) []Stargazer
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

type Stargazer struct {
	StarredAt time.Time `json:"starred_at"`
}

func NewGitHub(gr []GithubRepository, dr []ReadmeDetailsRepository) GitHub {
	return GitHub{
		GithubRepositorys:        gr,
		ReadmeDetailsRepositorys: dr,
	}
}

func NewDetailsRepository(repo GithubRepository, stargazers []Stargazer) ReadmeDetailsRepository {
	r := &ReadmeDetailsRepository{
		RepoName: repo.FullName,
		RepoURL:  repo.URL,
		StarCounts: map[string]int{
			"StarCount72MouthAgo": 0,
			"StarCount60MouthAgo": 0,
			"StarCount48MouthAgo": 0,
			"StarCount36MouthAgo": 0,
			"StarCount24MouthAgo": 0,
			"StarCount12MouthAgo": 0,
			"StarCountNow":        0,
		},
	}
	r.calculateStarCount(stargazers)
	r.RepoName = repo.FullName
	r.RepoURL = repo.URL
	return *r
}

func (r *ReadmeDetailsRepository) calculateStarCount(stargazers []Stargazer) {
	for _, star := range stargazers {
		r.updateStarCount("StarCountNow", star.StarredAt, 0)
		r.updateStarCount("StarCount12MouthAgo", star.StarredAt, -12)
		r.updateStarCount("StarCount24MouthAgo", star.StarredAt, -24)
		r.updateStarCount("StarCount36MouthAgo", star.StarredAt, -36)
		r.updateStarCount("StarCount48MouthAgo", star.StarredAt, -48)
		r.updateStarCount("StarCount60MouthAgo", star.StarredAt, -60)
		r.updateStarCount("StarCount72MouthAgo", star.StarredAt, -72)
	}
}

func (r *ReadmeDetailsRepository) updateStarCount(period string, starredAt time.Time, monthsAgo int) {
	var targetTime time.Time
	if monthsAgo == 0 {
		targetTime = time.Now().UTC()
	} else {
		targetTime = time.Now().UTC().AddDate(0, monthsAgo, 0)
	}

	if starredAt.Before(targetTime) {
		r.StarCounts[period]++
	}
}
