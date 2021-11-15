package server

import (
	"context"
	pb "github.com/anchamber/lab/proto/line"
)

type Server struct {
	pb.UnimplementedLineServer
}

func (s *Server) CreateLine(ctx context.Context, req *pb.CreateLineRequest) (*pb.CreateLineResponse, error) {
	return &pb.CreateLineResponse{}, nil
}

func (s *Server) ActivateLine(ctx context.Context, req *pb.ActivateLineRequest) (*pb.ActivateLineResponse, error) {
	return &pb.ActivateLineResponse{}, nil
}

func (s *Server) DeactivateLine(ctx context.Context, req *pb.DeactivateLineRequest) (*pb.DeactivateLineResponse, error) {
	return &pb.DeactivateLineResponse{}, nil
}

func (s *Server) DeleteLine(ctx context.Context, req *pb.DeleteLineRequest) (*pb.DeleteLineResponse, error) {
	return &pb.DeleteLineResponse{}, nil
}
