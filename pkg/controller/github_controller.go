package controller

import (
	"context"
	"log"
	"star-golang-migrations/pkg/di"
)

func Start(ctx context.Context, token string) error {
	ghs := di.InitializeGitHubController()
	gh, err := ghs.ExecGitHubAPI(ctx, token)
	if err != nil {
		return err
	}

	log.Println(gh)
	// err = gh.SortDesByStarCount()
	// if err != nil {
	// 	return err
	// }

	// err = gh.MakeChart()
	// if err != nil {
	// 	return err
	// }

	// err = gh.Edit()
	// if err != nil {
	// 	return err
	// }

	return nil
}
