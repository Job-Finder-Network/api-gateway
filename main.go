package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Job-Finder-Network/api-gateway/endpoint"
	"github.com/Job-Finder-Network/api-gateway/entity"
	"github.com/Job-Finder-Network/api-gateway/repository"
	"github.com/Job-Finder-Network/api-gateway/service"
	"github.com/Job-Finder-Network/api-gateway/transport"
	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"

	"net/http"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	entity.InitDb()
	ctx := context.Background()
	var userService service.UserService
	{
		repository := repository.UserRepository{}
		userService = service.NewUserService(repository, logger)
	}
	var loginService service.LoginServiceInterface
	{
		loginService = service.LoginServiceCreate(logger)
	}
	endpoints := endpoint.MakeEndpoints(userService,loginService)
	var httpAddr = flag.String("http", ":8080", "http listen address")
	fmt.Println("listening on port", *httpAddr)
	handler := transport.NewHTTPServer(ctx, endpoints)
	http.ListenAndServe(*httpAddr, handler)

}
