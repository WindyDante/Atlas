package service

import (
	"context"

	docv1 "github.com/ToAtlas/AtlasBackend/api/gen/go/doc/service/v1"
)

type DocService struct {
	docv1.UnimplementedDocServer
}

func NewDocService() *DocService {
	return &DocService{}
}

func (s *DocService) GetDoc(ctx context.Context, req *docv1.GetDocRequest) (*docv1.GetDocResponse, error) {
	return &docv1.GetDocResponse{
		Id:      req.Id,
		Title:   "demo",
		Content: "hello doc",
	}, nil
}
