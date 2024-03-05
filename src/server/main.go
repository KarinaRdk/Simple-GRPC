package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	pb "gprs/proto/proto"
	"log"
	"math/rand"
	"net"
	"os"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	pb.UnimplementedTransmitterServer
}

func (x *server) GetEntry(ctx *emptypb.Empty, stream pb.Transmitter_GetEntryServer) error {
	// Create a log file
	logFile, err1 := os.Create("log.txt")
	if err1 != nil {
		return err1
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	//  Loging:
	log.SetOutput(logFile)
	mean := -10 + rand.Float64()*(10 - -10)
	std := 0.3 + rand.Float64()*(1.5-0.3)
	id := uuid.New().String()
	var err error

	for {
		m := &pb.Entry{
			SessionId: id,
			// NormFloat64 returns a normally distributed float64 in the range
			// [-math.MaxFloat64, +math.MaxFloat64] with standard normal distribution
			// (mean = 0, stddev = 1) from the default Source. To produce a different
			//  normal distribution, callers can adjust the output
			Frequency: rand.NormFloat64()*std + mean,
			Timestamp: timestamppb.Now(),
		}
		log.Printf("Sending message: session_id: %s, frequency: %f, timestamp: %v", m.SessionId, m.Frequency, m.Timestamp)

		//  For when connection with the client has been terminated:
		err = stream.Send(m)
		if err != nil {
			break
		}
	}

	return err
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransmitterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
