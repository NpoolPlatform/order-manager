package payment

import (
	"github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	payment.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	payment.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
