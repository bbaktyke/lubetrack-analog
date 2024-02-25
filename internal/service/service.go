package service

import (
	repo "bbaktyke/lubetrack-analog.git/internal/repository"
	"context"
)

type Service interface {
	GetAnalysis(ctx context.Context, id string) error
}

type OilAnalysis struct {
	repository repo.Database
}

func New(repo repo.Database) (Service, error) {
	return OilAnalysis{
		repository: repo,
	}, nil
}

func (s OilAnalysis) GetAnalysis(ctx context.Context, id string) error {
	return nil
}
