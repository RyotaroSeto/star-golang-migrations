package application

import (
	"context"
)

type GitHubService interface {
	ExecGitHubAPI(ctx context.Context, token string) error
	// SortDesByStarCount() error
	// MakeChart() error
	// Edit() error
}
