package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"restapi/cmdx"
	"restapi/service"
	"restapi/spec"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	advertiseURLs = cmdx.FlagStrArr(cmdx.EnvOrStrArr("ADVERTISE_URL", []string{"http://127.0.0.1:9898", "http://localhost:9898"}))

	httpPort = 9898

	// S3
	s3Endpoint  string
	s3AccessKey string
	s3SecretKey string
	s3UseSSL    bool
)

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	flag.IntVar(&httpPort, "http-port", cmdx.EnvOrInt("HTTP_PORT", httpPort), "[HTTP_PORT] listen port for HTTP server")

	flag.Var(&advertiseURLs, "advertise-url", "[ADVERTISE_URL] the url will advertise as the service address")
}

func main() {
	ctx, quit := signal.NotifyContext(context.Background(), os.Interrupt)
	defer quit()

	logWriter := zerolog.SyncWriter(os.Stdout)
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		logWriter = zerolog.NewConsoleWriter()
	}

	logger := zerolog.New(logWriter).With().Timestamp().Logger()

	ctx = logger.WithContext(ctx)

	flag.Parse()

	if err := run(ctx); err != nil {
		logger.Fatal().Stack().Err(err).Msgf("program exited with an error: %+v", err)
	}
}

func run(ctx context.Context) error {
	log := zerolog.Ctx(ctx)

	// echo
	e := echo.New()
	e.HTTPErrorHandler = service.EchoErrorHandler
	e.Use(middleware.BodyLimit("2M"))
	e.Pre(middleware.MethodOverride())
	e.Pre(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowHeaders:     []string{},
		AllowCredentials: true,
	}))

	// upload
	e.Static("/", "public")
	e.Use(middleware.Recover())

	log.Debug().Interface("advertise-urls", advertiseURLs).Msg("")

	apiRouterV1 := e.Group("/v1")

	sp, err := service.NewServiceProvider(service.ServiceOptions{
		APIAdvertiseURLs: advertiseURLs,
	})

	if err != nil {
		return errors.WithStack(err)
	}

	apiRouterV1.Use(oapimiddleware.OapiRequestValidatorWithOptions(sp.APIService().Spec(), &oapimiddleware.Options{
		Options: openapi3filter.Options{
			ExcludeRequestBody:    false,
			ExcludeResponseBody:   false, // oapi-codegen middleware does not support yet validation response
			IncludeResponseStatus: true,  // oapi-codegen middleware does not support yet validation status code
			MultiError:            true,
		},
	}))

	spec.RegisterHandlers(apiRouterV1, sp.APIService())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: e,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	srvErrCh := make(chan error)
	go func() {
		defer close(srvErrCh)

		if err := srv.ListenAndServe(); err != nil {
			srvErrCh <- errors.WithStack(err)
		}
	}()

	log.Info().Str("listen_addr", srv.Addr).Msg("the app has been started.")

	select {
	case err := <-srvErrCh:
		return errors.WithStack(err)

	case <-ctx.Done():
		log.Debug().Msg("graceful shutdown has been started.")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return errors.WithStack(err)
		}

		log.Debug().Msg("graceful shutdown has been completed.")

		return nil
	}
}
