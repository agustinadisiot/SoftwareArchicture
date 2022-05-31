package controllers

import (
	jwt "auth"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	domain "voter_api/domain/vote"
	"voter_api/logic"
	proto "voter_api/proto/authService"
	pb "voter_api/proto/voteService"
)

type VoterServer struct {
	server pb.VoteServiceServer
}

type AuthServer struct {
	server     proto.AuthServiceServer
	jwtManager *jwt.Manager
}

var voteReply pb.VoteReply

func RegisterServicesServer(grpcServer *grpc.Server, jwtManager *jwt.Manager) {
	voteServer := VoterServer{}
	pb.RegisterVoteServiceServer(grpcServer, &voteServer)

	server := AuthServer{jwtManager: jwtManager}
	proto.RegisterAuthServiceServer(grpcServer, &server)
}

func (server *AuthServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := logic.FindVoter(request.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !logic.IsCorrectPassword(user, request.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user.Username, user.Id, user.Role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &proto.LoginResponse{AccessToken: token}
	return res, nil
}

func (newVote *VoterServer) Vote(ctx context.Context, req *pb.VoteRequest) (*pb.VoteReply, error) {
	user, err := logic.FindVoter(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	voteModel := &domain.VoteModel{
		Id:              req.GetId(),
		CivicCredential: req.GetCivicCredential(),
		Department:      req.GetDepartment(),
		Circuit:         req.GetCircuit(),
		Candidate:       req.GetCandidate(),
		PoliticalParty:  req.GetPoliticalParty(),
	}
	err = logic.StoreVote(voteModel)
	if err != nil {
		return &pb.VoteReply{Message: "Error"}, status.Errorf(codes.Internal, "cannot store vote: %v", err)
	}

	message := user.Username + " voted correctly"
	return &pb.VoteReply{Message: message}, status.Errorf(codes.OK, "vote stored")

	// rabbitmq test
	//api_voter.sendCertificate(idVoter)
	//api_voter.PrintCertificate()
}