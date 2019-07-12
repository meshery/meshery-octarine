// Copyright 2019 The Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main implements a client for the Octarine adapter.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/layer5io/meshery-octarine/meshes"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10004"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("test_client <init <kubeconfig filename><context name>>|<install|delete>|<install-bookinfo|delete-bookinfo <namespace>>")
}

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMeshServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if os.Args[1] == "init" {
		config, err := ioutil.ReadFile(os.Args[2])
		contextName := os.Args[3]			
		_, err = c.CreateMeshInstance(ctx, &pb.CreateMeshInstanceRequest{K8SConfig: config, ContextName: contextName})
		if err != nil {
			log.Fatalf("could not initialize client: %v", err)
		}
	} else if os.Args[1] == "install" || os.Args[1] == "delete" {
		_, err = c.ApplyOperation(ctx, &pb.ApplyRuleRequest{OpName: "octarine_install",
															DeleteOp: os.Args[1] == "delete",
															Namespace: "octarine-dataplane"})
		if err != nil {
			log.Fatalf("could not install octarine: %v", err)
		}
	} else if os.Args[1] == "install-bookinfo" || os.Args[1] == "delete-bookinfo" {
		_, err = c.ApplyOperation(ctx, &pb.ApplyRuleRequest{OpName: "install_book_info",
															DeleteOp: os.Args[1] == "delete-bookinfo",
															Namespace: os.Args[2]})
		if err != nil {
			log.Fatalf("could not install octarine: %v", err)
		}
	} else {
		usage()
	}
}
