// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	server.Serve(lis)

}
