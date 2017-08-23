package rpc

import (
	"context"
	"net/rpc"
	"github.com/skycoin/bbs/src/misc/boo"
	"github.com/skycoin/bbs/src/store/object"
)

type Call func() (method string, in interface{})

func Send(ctx context.Context, addresses interface{}, req Call) (goal uint64, e error) {
	for _, address := range addresses.([]string) {

		client, e := rpc.Dial("tcp", address)
		if e != nil {
			continue
		}

		methodName, in := req()
		call := client.Go(methodName, in, &goal, nil)

		select {
		case <-call.Done:
			return 0, call.Error
		case <- ctx.Done():
			return 0, ctx.Err()
		}
	}

	return 0, boo.New(boo.NotFound,
		"not found bla bla bla")
}

func NewThread(thread *object.Thread) Call {
	return func() (string, interface{}) {
		return "Gateway.NewThread", thread
	}
}

func NewPost(post *object.Post) Call {
	return func() (string, interface{}) {
		return "Gateway.NewPost", post
	}
}

func NewVote(vote *object.Vote) Call {
	return func() (string, interface{}) {
		return "Gateway.NewVote", vote
	}
}