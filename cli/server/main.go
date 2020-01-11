package main

import (
	"flag"
	grpcSetup "github.com/etokatlian/m-highscore/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "connection address for m-highscore")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc server for m-highscore")
	}
}
