package main

import (
	"github.com/layeh/murmur-cli/MurmurRPC"

	"google.golang.org/grpc"
)

func initTree(conn *grpc.ClientConn) {
	tree := MurmurRPC.NewTreeServiceClient(conn)

	cmd := root.Add("tree")

	cmd.Add("get", func(args []string) {
		server := MustServer(args)
		Output(tree.Get(ctx, server))
	})
}