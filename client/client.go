package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"splinter/grpcapi"

	"google.golang.org/grpc"
)

func banner() {
	fmt.Println("               .__  .__        __                 ")
	fmt.Println("  ____________ |  | |__| _____/  |_  ___________  ")
	fmt.Println(" /  ___/\\____ \\|  | |  |/    \\   __\\/ __ \\_  __ \\ ")
	fmt.Println(" \\___ \\ |  |_> >  |_|  |   |  \\  | \\  ___/|  | \\/ ")
	fmt.Println("/____  >|   __/|____/__|___|  /__|  \\___  >__|    ")
	fmt.Println("     \\/ |__|                \\/          \\/        ")
}

func main() {
	banner()
	var (
		opts   []grpc.DialOption
		conn   *grpc.ClientConn
		err    error
		client grpcapi.AdminClient
	)

	opts = append(opts, grpc.WithInsecure())
	if conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", 9090), opts...); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = grpcapi.NewAdminClient(conn)
	var cmd = new(grpcapi.Command)
	cmd.In = os.Args[1]
	ctx := context.Background()
	cmd, err = client.RunCommand(ctx, cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Out)
}
