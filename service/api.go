package service

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"restapi/spec"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type APIServiceDeps interface {
	CommonService() *CommonService
}

type APIService struct {
	deps APIServiceDeps

	apiSpec *openapi3.T
}

func newAPIService(deps APIServiceDeps, advertiseURLs []string) (*APIService, error) {
	apiSpec, err := spec.GetSwagger()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	apiSpec.Servers = make([]*openapi3.Server, len(advertiseURLs))
	for i, str := range advertiseURLs {
		u, err := url.Parse(str)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		u.Path = path.Join(u.Path, "v1")

		apiSpec.Servers[i] = &openapi3.Server{
			URL: u.String(),
		}
	}

	return &APIService{
		deps:    deps,
		apiSpec: apiSpec,
	}, nil
}

func (s *APIService) Spec() *openapi3.T {
	return s.apiSpec
}

func (s *APIService) GetAPISpec(ctx echo.Context, format spec.GetAPISpecParamsFormat) error {
	switch format {
	case ".json":
		return errors.WithStack(ctx.JSONPretty(http.StatusOK, s.apiSpec, ""))

	case ".yml", ".yaml":
		b, err := s.apiSpec.MarshalJSON()
		if err != nil {
			return errors.WithStack(err)
		}

		b, err = yaml.JSONToYAML(b)
		if err != nil {
			return errors.WithStack(err)
		}

		return errors.WithStack(ctx.Blob(http.StatusOK, "application/x-yaml; charset=UTF-8", b))

	default:
		panic("here is a bug")
	}
}

// common

func (s *APIService) Languages(ctx echo.Context) error {
	fmt.Println("Languages")
	return s.deps.CommonService().HandleLanguages(ctx)
}

func (s *APIService) SendFeedback(ctx echo.Context) error {
	fmt.Println("SendFeedback")
	return s.deps.CommonService().HandleSendFeedback(ctx)
}

func (s *APIService) FileUpload1(ctx echo.Context) error {
	fmt.Println("FileUpload1")
	return s.deps.CommonService().HandleFileUpload1(ctx)
}

func (s *APIService) FileUpload2(ctx echo.Context) error {
	fmt.Println("FileUpload2")
	return s.deps.CommonService().HandleFileUpload2(ctx)
}
