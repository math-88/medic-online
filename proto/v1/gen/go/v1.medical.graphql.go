// Code generated by proroc-gen-graphql, DO NOT EDIT.
package v1

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	gql_ptypes_timestamppb "github.com/ysugimoto/grpc-graphql-gateway/ptypes/timestamppb"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__enum_role                *graphql.Enum        // enum role in proto/v1/medical.proto
	gql__type_UserGetRequest      *graphql.Object      // message UserGetRequest in proto/v1/medical.proto
	gql__type_UserGetReply        *graphql.Object      // message UserGetReply in proto/v1/medical.proto
	gql__type_User                *graphql.Object      // message User in proto/v1/medical.proto
	gql__type_ProtocolGetRequest  *graphql.Object      // message ProtocolGetRequest in proto/v1/medical.proto
	gql__type_ProtocolGetReply    *graphql.Object      // message ProtocolGetReply in proto/v1/medical.proto
	gql__type_Protocol            *graphql.Object      // message Protocol in proto/v1/medical.proto
	gql__type_Anamnesis           *graphql.Object      // message Anamnesis in proto/v1/medical.proto
	gql__type_Actions             *graphql.Object      // message Actions in proto/v1/medical.proto
	gql__input_UserGetRequest     *graphql.InputObject // message UserGetRequest in proto/v1/medical.proto
	gql__input_UserGetReply       *graphql.InputObject // message UserGetReply in proto/v1/medical.proto
	gql__input_User               *graphql.InputObject // message User in proto/v1/medical.proto
	gql__input_ProtocolGetRequest *graphql.InputObject // message ProtocolGetRequest in proto/v1/medical.proto
	gql__input_ProtocolGetReply   *graphql.InputObject // message ProtocolGetReply in proto/v1/medical.proto
	gql__input_Protocol           *graphql.InputObject // message Protocol in proto/v1/medical.proto
	gql__input_Anamnesis          *graphql.InputObject // message Anamnesis in proto/v1/medical.proto
	gql__input_Actions            *graphql.InputObject // message Actions in proto/v1/medical.proto
)

func Gql__enum_role() *graphql.Enum {
	if gql__enum_role == nil {
		gql__enum_role = graphql.NewEnum(graphql.EnumConfig{
			Name: "V1_Enum_role",
			Values: graphql.EnumValueConfigMap{
				"user": &graphql.EnumValueConfig{
					Value: 0,
				},
				"admin": &graphql.EnumValueConfig{
					Value: 1,
				},
			},
		})
	}
	return gql__enum_role
}

func Gql__type_UserGetRequest() *graphql.Object {
	if gql__type_UserGetRequest == nil {
		gql__type_UserGetRequest = graphql.NewObject(graphql.ObjectConfig{
			Name:        "V1_Type_UserGetRequest",
			Description: `The request message containing the login and password`,
			Fields: graphql.Fields{
				"user_id": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_UserGetRequest
}

func Gql__type_UserGetReply() *graphql.Object {
	if gql__type_UserGetReply == nil {
		gql__type_UserGetReply = graphql.NewObject(graphql.ObjectConfig{
			Name:        "V1_Type_UserGetReply",
			Description: `The response message containing the answer`,
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: Gql__type_User(),
				},
			},
		})
	}
	return gql__type_UserGetReply
}

func Gql__type_User() *graphql.Object {
	if gql__type_User == nil {
		gql__type_User = graphql.NewObject(graphql.ObjectConfig{
			Name: "V1_Type_User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"surname": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"role": &graphql.Field{
					Type: Gql__enum_role(),
				},
				"createdAt": &graphql.Field{
					Type: gql_ptypes_timestamppb.Gql__type_Timestamp(),
				},
				"updated_at": &graphql.Field{
					Type: gql_ptypes_timestamppb.Gql__type_Timestamp(),
				},
			},
		})
	}
	return gql__type_User
}

func Gql__type_ProtocolGetRequest() *graphql.Object {
	if gql__type_ProtocolGetRequest == nil {
		gql__type_ProtocolGetRequest = graphql.NewObject(graphql.ObjectConfig{
			Name:        "V1_Type_ProtocolGetRequest",
			Description: `The request message containing the login and password`,
			Fields: graphql.Fields{
				"user_id": &graphql.Field{
					Type: graphql.String,
				},
				"anamnesis": &graphql.Field{
					Type: Gql__type_Anamnesis(),
				},
			},
		})
	}
	return gql__type_ProtocolGetRequest
}

func Gql__type_ProtocolGetReply() *graphql.Object {
	if gql__type_ProtocolGetReply == nil {
		gql__type_ProtocolGetReply = graphql.NewObject(graphql.ObjectConfig{
			Name:        "V1_Type_ProtocolGetReply",
			Description: `The response message containing the protocol`,
			Fields: graphql.Fields{
				"protocol": &graphql.Field{
					Type: Gql__type_Protocol(),
				},
			},
		})
	}
	return gql__type_ProtocolGetReply
}

func Gql__type_Protocol() *graphql.Object {
	if gql__type_Protocol == nil {
		gql__type_Protocol = graphql.NewObject(graphql.ObjectConfig{
			Name: "V1_Type_Protocol",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"actions": &graphql.Field{
					Type: graphql.NewList(Gql__type_Actions()),
				},
			},
		})
	}
	return gql__type_Protocol
}

func Gql__type_Anamnesis() *graphql.Object {
	if gql__type_Anamnesis == nil {
		gql__type_Anamnesis = graphql.NewObject(graphql.ObjectConfig{
			Name: "V1_Type_Anamnesis",
			Fields: graphql.Fields{
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_Anamnesis
}

func Gql__type_Actions() *graphql.Object {
	if gql__type_Actions == nil {
		gql__type_Actions = graphql.NewObject(graphql.ObjectConfig{
			Name: "V1_Type_Actions",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"drug": &graphql.Field{
					Type:        graphql.String,
					Description: `string name = 2;`,
				},
				"dosage": &graphql.Field{
					Type: graphql.String,
				},
				"time": &graphql.Field{
					Type: graphql.String,
				},
				"duration": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_Actions
}

func Gql__input_UserGetRequest() *graphql.InputObject {
	if gql__input_UserGetRequest == nil {
		gql__input_UserGetRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_UserGetRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_UserGetRequest
}

func Gql__input_UserGetReply() *graphql.InputObject {
	if gql__input_UserGetReply == nil {
		gql__input_UserGetReply = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_UserGetReply",
			Fields: graphql.InputObjectConfigFieldMap{
				"user": &graphql.InputObjectFieldConfig{
					Type: Gql__input_User(),
				},
			},
		})
	}
	return gql__input_UserGetReply
}

func Gql__input_User() *graphql.InputObject {
	if gql__input_User == nil {
		gql__input_User = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_User",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"surname": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"email": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"role": &graphql.InputObjectFieldConfig{
					Type: Gql__enum_role(),
				},
				"createdAt": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamppb.Gql__input_Timestamp(),
				},
				"updated_at": &graphql.InputObjectFieldConfig{
					Type: gql_ptypes_timestamppb.Gql__input_Timestamp(),
				},
			},
		})
	}
	return gql__input_User
}

func Gql__input_ProtocolGetRequest() *graphql.InputObject {
	if gql__input_ProtocolGetRequest == nil {
		gql__input_ProtocolGetRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_ProtocolGetRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"user_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"anamnesis": &graphql.InputObjectFieldConfig{
					Type: Gql__input_Anamnesis(),
				},
			},
		})
	}
	return gql__input_ProtocolGetRequest
}

func Gql__input_ProtocolGetReply() *graphql.InputObject {
	if gql__input_ProtocolGetReply == nil {
		gql__input_ProtocolGetReply = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_ProtocolGetReply",
			Fields: graphql.InputObjectConfigFieldMap{
				"protocol": &graphql.InputObjectFieldConfig{
					Type: Gql__input_Protocol(),
				},
			},
		})
	}
	return gql__input_ProtocolGetReply
}

func Gql__input_Protocol() *graphql.InputObject {
	if gql__input_Protocol == nil {
		gql__input_Protocol = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_Protocol",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"description": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"actions": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_Actions()),
				},
			},
		})
	}
	return gql__input_Protocol
}

func Gql__input_Anamnesis() *graphql.InputObject {
	if gql__input_Anamnesis == nil {
		gql__input_Anamnesis = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_Anamnesis",
			Fields: graphql.InputObjectConfigFieldMap{
				"description": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_Anamnesis
}

func Gql__input_Actions() *graphql.InputObject {
	if gql__input_Actions == nil {
		gql__input_Actions = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "V1_Input_Actions",
			Fields: graphql.InputObjectConfigFieldMap{
				"id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"drug": &graphql.InputObjectFieldConfig{
					Description: `string name = 2;`,
					Type:        graphql.String,
				},
				"dosage": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"time": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"duration": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"description": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_Actions
}

// graphql__resolver_MedicalService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_MedicalService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_MedicalService creates pointer of service struct
func new_graphql_resolver_MedicalService(conn *grpc.ClientConn) *graphql__resolver_MedicalService {
	return &graphql__resolver_MedicalService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_MedicalService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_MedicalService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"GetUser": &graphql.Field{
			Type: Gql__type_UserGetReply(),
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req UserGetRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for GetUser")
				}
				client := NewMedicalServiceClient(conn)
				resp, err := client.UserGet(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC UserGet")
				}
				return resp, nil
			},
		},
		"getProtocol": &graphql.Field{
			Type: Gql__type_ProtocolGetReply(),
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"anamnesis": &graphql.ArgumentConfig{
					Type: Gql__input_Anamnesis(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ProtocolGetRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for getProtocol")
				}
				client := NewMedicalServiceClient(conn)
				resp, err := client.ProtocolGet(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ProtocolGet")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_MedicalService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterMedicalServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterMedicalServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterMedicalServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service MedicalService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterMedicalServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_MedicalService(conn))
}
