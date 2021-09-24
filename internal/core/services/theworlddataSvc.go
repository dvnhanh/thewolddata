package services

import "github.com/dvnhanh/thewolddata/internal/core/ports"

func NewTheWorldDataService(repo ports.ThewolddataMysqlRepoS) ports.ThewolddataService {
	return &theWorldDataService{
		repo: repo,
	}
}

type theWorldDataService struct {
	repo ports.ThewolddataMysqlRepoS
}

func (p *theWorldDataService) Register(email, passsword string) error {
	return p.repo.Register(email, passsword)
}
