package di

import (
	"star-golang-migrations/pkg/application"
	"star-golang-migrations/pkg/controller"
	"star-golang-migrations/pkg/infrastructure"
)

func InitializeGitHubController() *controller.GitHubController {
	gitHubRepository := infrastructure.NewGitHubRepository()
	gitHubService := application.NewGitHubService(gitHubRepository)
	gitHubController := controller.NewGitHubController(gitHubService)
	return gitHubController
}
