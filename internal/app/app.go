package app

import (
	"log"
	"net"
	"net/http"

	"github.com/math-88/medic-online/internal/config"
	graphql "github.com/math-88/medic-online/internal/controller/graphql"
	controller "github.com/math-88/medic-online/internal/controller/grpc"
	rest "github.com/math-88/medic-online/internal/controller/rest"
	"github.com/math-88/medic-online/internal/openai"
	"github.com/math-88/medic-online/internal/usecase"
	"github.com/math-88/medic-online/pkg/logger"
	"github.com/math-88/medic-online/storage"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

func Run() error {

	cfg, err := config.New()
	if err != nil {
		return err
	}

	// logger init
	lg := logger.New("info")
	lg.Info("App started")

	// storage init
	store := storage.New(cfg)
	// test data
	err = seedUsers(store)
	if err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	// openAI init
	oa := openai.New(cfg)

	// Use case protocol
	protocolUseCase := usecase.New(store, oa)

	// auth
	jwtManager := controller.NewJWTManager(secretKey, tokenDuration)
	authServer := controller.NewAuthServer(store, jwtManager)

	medicalServer := controller.NewMedicalServer(store, protocolUseCase)

	listenerGRPC, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	defer listenerGRPC.Close()

	listenerREST, err := net.Listen("tcp", *restServerEndpoint)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	defer listenerREST.Close()

	listenerGraphQL, err := net.Listen("tcp", *graphQLServerEndpoint)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	defer listenerREST.Close()

	// app := model.NewApp()

	var group errgroup.Group

	group.Go(func() error {
		return controller.RunGRPC(lg, authServer, medicalServer, jwtManager, listenerGRPC)
	})

	group.Go(func() error {
		return rest.Run(lg, listenerREST, *grpcServerEndpoint)
	})

	group.Go(func() error {
		return graphql.Run(lg, listenerGraphQL, *grpcServerEndpoint)
	})

	group.Go(func() error {
		return http.ListenAndServe(":2662", promhttp.Handler())
	})

	// group.Go(func() error {
	// 	return http.ListenAndServe(":8843", wsproxy.WebsocketProxy()
	// })

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
