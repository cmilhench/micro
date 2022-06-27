package consignments

import (
	"context"
	"io"
	"math/rand"
	"time"

	"consignments/model"
)

func init() {

	rand.Seed(time.Now().UnixNano())
}

type Service struct {
	Repo model.Repository
}

func (s *Service) CreateConsignment(ctx context.Context, req *model.Consignment) (*model.Consignment, error) {
	consignment, err := s.Repo.Create(req)
	if err != nil {
		return nil, err
	}
	return consignment, nil
}

func (s *Service) GetConsignments(ctx context.Context) ([]*model.Consignment, error) {
	consignments := s.Repo.GetAll()
	if rand.Intn(100) > 75 {
		return nil, io.EOF
	}
	return consignments, nil
}
