package infrastructure

import (
	"context"
	"star-golang-migrations/pkg/domain"
)

type GitHubRepository struct {
}

var _ domain.GitHubRepository = new(GitHubRepository)

func NewGitHubRepository() *GitHubRepository {
	return &GitHubRepository{}
}

func (ghr *GitHubRepository) NowGithubRepoCount(ctx context.Context, repoNm string, token string) (domain.GithubRepository, error) {
	return domain.GithubRepository{}, nil
}

func (ghr *GitHubRepository) GetStargazersCountByRepo(ctx context.Context, name, token string, repo domain.GithubRepository) []domain.Stargazer {
	return []domain.Stargazer{}
}
