package main

import (
	"context"
	"github.com/gofrs/uuid"
	pb "github.com/rasha-hantash/grpc-gateway/rewardsrefund"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddUser to the in-memory store.
func (b *Backend) AddUser(ctx context.Context, _ *pb.AddUserRequest) (*pb.User, error) {
	user := &pb.User{
		Id: uuid.Must(uuid.NewV4()).String(),
	}
	b.users = append(b.users, user)

	return user, nil
}

// ListUser lists all users in the store.
func (b *Backend) ListUser(_ *pb.ListUsersRequest, srv pb.RefundService_ListUserServer) error {

	for _, user := range b.users {
		err := srv.Send(user)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateUserRefund lists all users in the store.
func (b *Backend) CreateUserRefund(ctx context.Context, req *pb.UserRefundRequest) (*pb.User, error) {

	var u *pb.User
	for _, user := range b.users {
		if user.Id == req.GetUser().GetId() {
			u = user
			break
		}
	}

	u.AccountBalance = u.GetAccountBalance() + req.RefundItemCost

	return u, nil
}

// UpdateUserBalance lists all users in the store.
func (b *Backend) UpdateUserBalance(ctx context.Context, req *pb.UserBalanceRequest) (*pb.User, error) {

	var u *pb.User
	for _, user := range b.users {
		if user.Id == req.GetUser().GetId() {
			u = user
			break
		}
	}

	u.AccountBalance = req.GetBalance()

	return u, nil
}

// RedeemRewards lists all users in the store.
func (b *Backend) RedeemRewards(ctx context.Context, req *pb.UserRewardsRequest) (*pb.User, error) {

	var u *pb.User
	for _, user := range b.users {
		if user.Id == req.GetUser().GetId() {
			u = user
			break
		}
	}

	if req.GetRedemption() > u.AccountBalance {
		return u, status.Errorf(codes.InvalidArgument, "Not enough balance to redeem")
	}

	u.AccountBalance = u.GetAccountBalance() - req.GetRedemption()

	return u, nil
}
