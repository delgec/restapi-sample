package service

import (
	"net/http"

	"restapi/spec"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type CommonServiceDeps interface {
	CommonService() *CommonService
}

type CommonService struct {
	deps CommonServiceDeps
}

func newCommonService(deps CommonServiceDeps) (*CommonService, error) {
	return &CommonService{
		deps: deps,
	}, nil
}

func (s *CommonService) HandleLanguages(c echo.Context) error {
	var items []spec.Language = make([]spec.Language, 1)

	items[0] = spec.Language{
		Id:           1,
		CultureCode:  "tr-TR",
		Name:         "Türkçe",
		Rtl:          false,
		Published:    true,
		DisplayOrder: 1,
	}

	return errors.WithStack(c.JSON(http.StatusOK, items))
}
