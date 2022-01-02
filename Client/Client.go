package main
 
import (
   "context"
   "fmt"
   "log"
   "os"
   "strconv"
   "strings"
 
   pb "github.com/MieMilvang/DISYS_template/Proto"
   "google.golang.org/grpc"
)

const(
	address = "localhost: 8080"
 )
  
 var (
	clientid                    int64
 )
 func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err!= nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.ServiceServer(conn)
	

 }
