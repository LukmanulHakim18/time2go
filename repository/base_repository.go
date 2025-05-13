package repository

import (
	"fmt"

	"github.com/LukmanulHakim18/time2go/repository/repoiface"
)

type RepoConf interface {
	Init(*Repository) error
	GetRepoName() string
}

func (r *Repository) PrintHealthy(repoName string) {
	fmt.Println(fmt.Sprintf("+ %s Repository is healthy!", repoName))
}

func (r *Repository) PrintNotHealthy(repoName string) {
	fmt.Println(fmt.Sprintf("- %s Repository is not healthy!", repoName))
}

var repositoriesPointer *Repository

func NewRepository(rf []RepoConf) (*Repository, error) {
	if repositoriesPointer != nil {
		return repositoriesPointer, nil
	}

	repositoriesPointer = &Repository{}
	for _, rc := range rf {
		err := rc.Init(repositoriesPointer)
		if err != nil {
			repositoriesPointer.PrintNotHealthy(rc.GetRepoName())
			return nil, err
		}
		repositoriesPointer.PrintHealthy(rc.GetRepoName())
	}

	return repositoriesPointer, nil
}

type Repository struct {
	Redis repoiface.Redis
}
