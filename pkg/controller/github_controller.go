package controller

import (
	"context"
	"log"
	"star-golang-migrations/pkg/application"
)

type GitHubController struct {
	service application.GitHubService
}

func NewGitHubController(service application.GitHubService) *GitHubController {
	return &GitHubController{service: service}
}

func Run(ctx context.Context, token string, ghs *GitHubController) error {
	err := ghs.service.ExecGitHubAPI(ctx, token)
	if err != nil {
		return err
	}
	log.Println(ghs)

	// if err := ghs.service.SortDesByStarCount(); err != nil {
	// 	return err
	// }

	// err = ghs.service.MakeChart()
	// if err != nil {
	// 	return err
	// }

	// err = ghs.service.Edit()
	// if err != nil {
	// 	return err
	// }

	return nil
}
