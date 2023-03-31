// Code generated by goctl. DO NOT EDIT.
// Source: pet.proto

package petclient

import (
	"context"

	"boat/rpc/pet/pb/pet"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest   = pet.IdRequest
	PetResponse = pet.PetResponse

	Pet interface {
		GetPet(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*PetResponse, error)
	}

	defaultPet struct {
		cli zrpc.Client
	}
)

func NewPet(cli zrpc.Client) Pet {
	return &defaultPet{
		cli: cli,
	}
}

func (m *defaultPet) GetPet(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	client := pet.NewPetClient(m.cli.Conn())
	return client.GetPet(ctx, in, opts...)
}