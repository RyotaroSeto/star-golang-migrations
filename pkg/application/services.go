package application

import (
	"context"
	"star-golang-migrations/pkg/domain"
)

type GitHubService interface {
	ExecGitHubAPI(ctx context.Context, token string) (*domain.GitHub, error)
}
