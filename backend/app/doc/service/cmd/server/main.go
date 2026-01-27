package main

import (
	"context"
	"log"

	docv1 "github.com/WindyDante/Atlas/backend/api/gen/go/atlas/doc/v1"
	"github.com/go-kratos/kratos/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

type DocumentService struct {
	docv1.UnimplementedDocumentServiceServer
}

// 示例实现：先保证能启动，逻辑后面再补
func (s *DocumentService) ListDocuments(ctx context.Context, req *docv1.ListDocumentsRequest) (*docv1.ListDocumentsResponse, error) {
	return &docv1.ListDocumentsResponse{}, nil
}

func main() {
	svc := &DocumentService{}

	grpcSrv := kgrpc.NewServer(
		kgrpc.Address(":9000"),
	)
	docv1.RegisterDocumentServiceServer(grpcSrv, svc)

	httpSrv := khttp.NewServer(
		khttp.Address(":8000"),
	)
	// 前提：你已经生成了 go-http 的绑定代码（会有 RegisterXXXHTTPServer）
	docv1.RegisterDocumentServiceHTTPServer(httpSrv, svc)

	app := kratos.New(
		kratos.Name("doc-service"),
		kratos.Server(grpcSrv, httpSrv),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
