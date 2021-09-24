package services

import "github.com/dvnhanh/thewolddata/internal/core/port"

func NewTheWorldDataService(repo port.TheworlddataMysqlRepoS) port.TheworlddataService {
	return &theWorldDataService{
		repo: repo,
	}
}

type theWorldDataService struct {
	repo port.TheworlddataMysqlRepoS
}

func (p *theWorldDataService) Register(email, passsword string) error {
	return p.repo.Register(email, passsword)
}
