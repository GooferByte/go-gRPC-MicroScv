package account

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type fgrpcServer struct{
	service Service
}

func ListenGRPC(s Service, port int) error{
	lis, err := net.Listen("tcp", fmt.Sprint(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.(serv,)
	reflection.Register(serv)
	return serv.Server(lis)
}
func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest)(*pb.PostAccountResponse, error){
	a, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id: a.ID,
		Name: a.Name,
	}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountRespose, error){
	a, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountRespose{
		Account: &pb.Account{
			Id: a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r.*pb.GetAccountsRequest) (*pb.GetAccountsRespose, error){
	a, err := s.service.GetAccounts(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(accounts, 
			&pb.Account{
				Id: p.ID,
				Name: p.name,
			},
		)
	}
	return &pb.GetAccountsRespose{Accounts: accounts}, nil
}

