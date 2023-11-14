package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "golang_express_grpc/siswa"

	"google.golang.org/grpc"
)

func main() {
	serverAddr := "localhost:50051"

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSiswaServiceClient(conn)

	siswaID := "1"
	req := &pb.SiswaRequest{Id: siswaID}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	siswa, err := client.GetSiswa(ctx, req)
	if err != nil {
		log.Fatalf("Error getting siswa: %v", err)
	}

	fmt.Printf("Siswa: %+v\n", siswa)
}

