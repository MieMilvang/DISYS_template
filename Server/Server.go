package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/MieMilvang/DISYS_gRPC_Eample/Proto"
	"google.golang.org/grpc"
)
 
const (
   port = ":8080"
)
 
type Server struct{
   pb.UnimplementedServiceServer
}

func main(){
	fmt.Printf("try")
}