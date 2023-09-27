package main

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"

	seller_score "github.com/Longoria-Yu/test-scripts/cmd/seller-score"
	contentfeed_proto "github.com/Longoria-Yu/test-scripts/pb/content-feed"
)

var (
	// prod
	host = flag.String("host", "10.240.100.126", "The service hostname or IP address")
	port = flag.Int64("port", 8601, "The service gRPC port")

	// stage
	//host = flag.String("host", "10.140.16.161", "The service hostname or IP address")
	//port = flag.Int64("port", 8601, "The service gRPC port")

	// local
	// host = flag.String("host", "127.0.0.1", "The service hostname or IP address")
	// port = flag.Int64("port", 9281, "The service gRPC port")
)

func main() {
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *host, *port)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := contentfeed_proto.NewContentFeedServiceClient(conn)
	_ = c

	seller_score.CallSearchCGProductListingsV10(c)
}
