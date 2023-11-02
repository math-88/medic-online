package controller

import (
	"context"

	"time"

	"net"
	"net/http"
	"strings"

	"github.com/math-88/medic-online/pkg/logger"
	pb "github.com/math-88/medic-online/proto/v1/gen/go"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

func Run(
	log *logger.Logger,
	// authServer pb.AuthServiceServer,
	// medicalServer pb.MedicalServer,
	// jwt *JWTManager,
	listener net.Listener,
	grpcServerEndpoint string,

) error {

	<-time.After(3 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(
		// convert header in response(going from gateway) from metadata received.
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")

			// send all the headers received from the client
			md := metadata.Pairs("authorization", header)
			return md
		}),
		// runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
		// 	//creating a new HTTTPStatusError with a custom status, and passing error

		// 	newError := runtime.HTTPStatusError{
		// 		HTTPStatus: 400,
		// 		Err:        err,
		// 	}
		// 	// using default handler to do the rest of heavy lifting of marshaling error and adding headers
		// 	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		// }),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithMarshalerOption("application/json+pretty", &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				Indent:    "  ",
				Multiline: true, // Optional, implied by presence of "Indent".
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// register auth service
	if err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, dialOptions); err != nil {
		return err
	}

	// register medical service
	if err := pb.RegisterMedicalServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, dialOptions); err != nil {
		return err
	}
	// pretty json
	prettier := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// checking Values as map[string][]string also catches ?pretty and ?pretty=
			// r.URL.Query().Get("pretty") would not.
			if _, ok := r.URL.Query()["pretty"]; ok {
				r.Header.Set("Accept", "application/json+pretty")
			}
			h.ServeHTTP(w, r)
		})
	}

	// websocket
	// go http.ListenAndServe("localhost:8843", wsproxy.WebsocketProxy(mux))

	// Start HTTP server (and proxy calls to gRPC server endpoint)

	log.Info("Start REST server at %s", listener.Addr().String())

	// return http.ListenAndServe("localhost:8082", prettier(mux))

	return http.Serve(listener, prettier(mux))
}

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in allowed header, don't send the header
	return s, false
}
