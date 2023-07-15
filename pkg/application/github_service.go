package application

import (
	"context"
	"log"
	"star-golang-migrations/pkg/domain"
)

type gitHubService struct {
	repo domain.GitHubRepository
}

func NewGitHubService(repo domain.GitHubRepository) GitHubService {
	return &gitHubService{
		repo: repo,
	}
}

var _ GitHubService = new(gitHubService)

func (s *gitHubService) ExecGitHubAPI(ctx context.Context, token string) (*domain.GitHub, error) {
	// var repos []domain.GithubRepository
	// var detaiRepos []domain.ReadmeDetailsRepository

	// wg := new(sync.WaitGroup)
	// var lock sync.Mutex
	for _, repoNm := range domain.TargetRepository {
		// 	wg.Add(1)
		// 	go func(repoNm string) {
		// 		defer wg.Done()
		repo, err := s.repo.NowGithubRepoCount(ctx, repoNm, token)
		if err != nil {
			log.Println(err)
			// return
		}
		log.Println(repo)
		// repos = append(repos, repo)
		// log.Println(repoNm + " Start")
		// stargazers := getStargazersCountByRepo(ctx, repoNm, token, repo)
		// log.Println(repoNm + " DONE")
		// lock.Lock()
		// defer lock.Unlock()
		// detaiRepos = append(detaiRepos, NewDetailsRepository(repo, stargazers))
		// }(repoNm)
	}

	// wg.Wait()
	// gh := NewGitHub(repos, detaiRepos)
	return nil, nil
	// return gh, nil
}
