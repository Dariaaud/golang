package main

import (
	"awesomeProject/accounts/models"
	"awesomeProject/proto"
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New() *server {
	return &server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type server struct {
	proto.UnimplementedAccountManagerServer
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	s.guard.Lock()
	defer s.guard.Unlock()
	if _, ok := s.accounts[req.Name]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "account already exists")
	}

	s.accounts[req.Name] = &models.Account{
		Name:   req.Name,
		Amount: int(req.Amount),
	}

	return &proto.CreateAccountReply{}, nil
}

func (s *server) ChangeAmountAccount(ctx context.Context, req *proto.ChangeAmountAccountRequest) (*proto.ChangeAmountAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	s.guard.Lock()
	defer s.guard.Unlock()
	if _, ok := s.accounts[req.Name]; !ok {
		return nil, status.Errorf(codes.NotFound, "account not found")
	}

	s.accounts[req.Name].Amount = int(req.NewAmount)

	return &proto.ChangeAmountAccountReply{}, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	s.guard.Lock()
	defer s.guard.Unlock()
	if _, ok := s.accounts[req.Name]; !ok {
		return nil, status.Errorf(codes.NotFound, "account not found")
	}

	delete(s.accounts, req.Name)

	return &proto.DeleteAccountReply{}, nil
}

func (s *server) ChangeNameAccount(ctx context.Context, req *proto.ChangeNameAccountRequest) (*proto.ChangeNameAccountReply, error) {
	if len(req.NewName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	if len(req.NewName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty new name")
	}

	s.guard.Lock()
	defer s.guard.Unlock()

	if _, ok := s.accounts[req.NewName]; !ok {
		return nil, status.Errorf(codes.NotFound, "account not found")
	}

	if _, ok := s.accounts[req.NewName]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "account with new name already exists")
	}

	account := s.accounts[req.NewName]
	account.Name = req.NewName

	delete(s.accounts, req.NewName)
	s.accounts[req.NewName] = account

	return &proto.ChangeNameAccountReply{}, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountReply, error) {
	if len(req.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	s.guard.Lock()
	defer s.guard.Unlock()
	if account, ok := s.accounts[req.Name]; ok {
		return &proto.GetAccountReply{
			Name:   account.Name,
			Amount: int32(account.Amount),
		}, nil
	}

	return nil, status.Errorf(codes.NotFound, "account not found")
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		panic(err)
	}

	server := New()
	s := grpc.NewServer()

	proto.RegisterBankAccountServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
