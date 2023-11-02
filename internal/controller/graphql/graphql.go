package graphql

import (
	"net"
	"net/http"

	"github.com/math-88/medic-online/pkg/logger"
	pb "github.com/math-88/medic-online/proto/v1/gen/go"

	runtime_graphql "github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(
	log *logger.Logger,
	listener net.Listener,
	grpcServerEndpoint string,
) error {
	mux := runtime_graphql.NewServeMux(

		runtime_graphql.Cors(),

		// @TODO - library grpc-graphql-gateway hasn`t metadata context
		// runtime_graphql.WithMetadata(
		// 	func(ctx context.Context, request *http.Request) metadata.MD {
		// 		header := request.Header.Get("Authorization")

		// 		// send all the headers received from the client
		// 		md := metadata.Pairs("authorization", header)
		// 		return md
		// 	},
		// ),
	)

	grpcCon, err := grpc.Dial(grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	if err := pb.RegisterAuthServiceGraphqlHandler(mux, grpcCon); err != nil {
		return err
	}

	if err := pb.RegisterMedicalServiceGraphqlHandler(mux, grpcCon); err != nil {
		return err
	}

	// return http.Serve(listener, mux)

	http.Handle("/graphql", mux)

	log.Info("Start graphQL server at %s", listener.Addr().String())

	return http.Serve(listener, mux)
}

// var allowedHeaders = map[string]struct{}{
// 	"x-request-id": {},
// }

// func isHeaderAllowed(s string) (string, bool) {
// 	// check if allowedHeaders contain the header
// 	if _, isAllowed := allowedHeaders[s]; isAllowed {
// 		// send uppercase header
// 		return strings.ToUpper(s), true
// 	}
// 	// if not in allowed header, don't send the header
// 	return s, false
// }
