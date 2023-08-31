package application

import (
	"context"
	"log"
	"sort"
	"star-golang-migrations/pkg/domain"
	"sync"
)

type gitHubService struct {
	repo domain.GitHubRepository
	gh   domain.GitHub
}

func NewGitHubService(repo domain.GitHubRepository) GitHubService {
	return &gitHubService{
		repo: repo,
		gh:   domain.GitHub{},
	}
}

var _ GitHubService = new(gitHubService)

func (s *gitHubService) ExecGitHubAPI(ctx context.Context, token string) error {
	var repos []domain.GithubRepository
	var detaiRepos []domain.ReadmeDetailsRepository

	wg := new(sync.WaitGroup)
	var lock sync.Mutex
	for _, repoNm := range domain.TargetRepository {
		wg.Add(1)
		go func(repoNm string) {
			defer wg.Done()
			repo, err := s.repo.NowGithubRepoCount(ctx, repoNm, token)
			if err != nil {
				log.Println(err)
				return
			}
			repos = append(repos, repo)
			log.Println(repoNm + " Start")
			stargazers := s.repo.GetStargazersCountByRepo(ctx, repoNm, token, repo)
			log.Println(repoNm + " DONE")
			lock.Lock()
			defer lock.Unlock()
			detaiRepos = append(detaiRepos, domain.NewDetailsRepository(repo, stargazers))
		}(repoNm)
	}

	wg.Wait()
	domain.NewGitHub(repos, detaiRepos)
	return nil
	// gh := domain.NewGitHub(repos, detaiRepos)
}

func (s *gitHubService) SortDesByStarCount() error {
	err := githubRepositorySort(s.gh.GithubRepositorys)
	if err != nil {
		return err
	}
	err = readmeDetailsRepositorySort(s.gh.ReadmeDetailsRepositorys)
	if err != nil {
		return err
	}
	return nil
}

func githubRepositorySort(grs []domain.GithubRepository) error {
	sort.Slice(grs, func(i, j int) bool {
		return grs[i].StargazersCount > grs[j].StargazersCount
	})
	return nil
}

func readmeDetailsRepositorySort(rds []domain.ReadmeDetailsRepository) error {
	sort.Slice(rds, func(i, j int) bool {
		return rds[i].StarCounts["StarCountNow"] > rds[j].StarCounts["StarCountNow"]
	})
	return nil
}
