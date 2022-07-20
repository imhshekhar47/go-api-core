package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/imhshekhar47/go-api-core/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

var (
	restServerLogger = logger.GetLogger("core/server/rest")
)

type RestServer struct {
	port        uint16
	basePath    string
	ctx         context.Context
	ctxCancelFn context.CancelFunc
	httpServer  *http.ServeMux
	gwMux       *runtime.ServeMux
	dialOptions []grpc.DialOption
	grpcAddr    string
	swaggerDoc  string
}

func NewRestServer(port uint16, basePath string, grpcAddr string) *RestServer {
	ctx, ctxCancelFn := context.WithCancel(context.Background())
	return &RestServer{
		port:        port,
		basePath:    basePath,
		ctx:         ctx,
		ctxCancelFn: ctxCancelFn,
		httpServer:  http.NewServeMux(),
		gwMux:       runtime.NewServeMux(),
		dialOptions: []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		grpcAddr:    grpcAddr,
	}
}

func (s *RestServer) GetHttpServer() *http.ServeMux {
	return s.httpServer
}

func (s *RestServer) GetBasePath() string {
	return s.basePath
}

func (s *RestServer) GetContext() context.Context {
	return s.ctx
}

func (s *RestServer) GetMultiplex() *runtime.ServeMux {
	return s.gwMux
}

func (s *RestServer) GetDialoptions() []grpc.DialOption {
	return s.dialOptions
}

func (s *RestServer) getAddress() string {
	return fmt.Sprintf("0.0.0.0:%d", s.port)
}

func (s *RestServer) GetGrpcAddress() string {
	return s.grpcAddr
}

func (s *RestServer) Run() error {
	defer s.ctxCancelFn()
	restServerLogger.Debugf("entry: Run()")
	address := fmt.Sprintf("0.0.0.0:%d", s.port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	apiDocUrl := fmt.Sprintf("%s/swagger/actuator", s.basePath)
	s.httpServer.HandleFunc(
		apiDocUrl,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.Write([]byte(s.swaggerDoc))
		})

	s.httpServer.Handle(
		fmt.Sprintf("%s/swagger/", s.basePath),
		httpSwagger.Handler(
			httpSwagger.URL(apiDocUrl), //The url pointing to API definition
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		))

	// Alow cors
	muxCors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		}),
		handlers.AllowedHeaders([]string{
			"Authorization",
			"Content-Type",
			"Accept-Encoding",
			"Accept",
		}),
	)(s.gwMux)

	//s.httpServer.Handle("/", s.gwMux)
	s.httpServer.Handle("/", muxCors)

	restServerLogger.Debugf("exit: Run(), HTTP Server listening on %s", listener.Addr())
	return http.Serve(listener, s.httpServer)
}
