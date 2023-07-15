package di

import (
	"star-golang-migrations/pkg/application"
	"star-golang-migrations/pkg/infrastructure"
)

func InitializeGitHubController() application.GitHubService {
	gitHubRepository := infrastructure.NewGitHubRepository()
	gitHubService := application.NewGitHubService(gitHubRepository)
	return gitHubService
}
