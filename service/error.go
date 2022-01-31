package service

import (
	"restapi/model"
	"restapi/spec"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func EchoErrorHandler(err error, c echo.Context) {
	log := zerolog.Ctx(c.Request().Context())

	httpErr := &echo.HTTPError{}
	if errors.As(err, &httpErr) && httpErr.Internal != nil {
		intErr := httpErr.Internal

		if mulErr, ok := intErr.(openapi3.MultiError); ok {
			intErr = mulErr[0]
		}

		switch e := intErr.(type) {
		case *openapi3filter.RequestError:
			err = &model.APIError{
				Source:  model.ErrorSourceClient,
				Code:    spec.ErrorCodeINVALIDREQUEST,
				Message: e.Error(),
				Cause:   e,
			}

		case *openapi3filter.SecurityRequirementsError:
			err = &model.APIError{
				Source:  model.ErrorSourceClientSecurity,
				Code:    spec.ErrorCodeSECURITYFAILED,
				Message: e.Error(),
				Cause:   e,
			}

		default:
			if httpErr.Code == 400 {
				err = &model.APIError{
					Source:  model.ErrorSourceClient,
					Code:    spec.ErrorCodeINVALIDREQUEST,
					Message: e.Error(),
					Cause:   e,
				}
			}
		}
	}

	apiErr := &model.APIError{}
	if !errors.As(err, &apiErr) {
		apiErr.Source = model.ErrorSourceServer
		apiErr.Code = spec.ErrorCodeGENERIC
		apiErr.Message = "An internal error occurred."
		apiErr.Cause = err
	}

	logErr := apiErr.Cause
	if logErr == nil {
		logErr = apiErr
	}

	log.Warn().Stack().Err(logErr).Msg("an error occurred while processing the API request")

	if err := c.JSON(int(apiErr.Source), &spec.ErrorResponse{
		Error: spec.Error{
			Code:    apiErr.Code,
			Message: apiErr.Message,
		},
	}); err != nil {
		log.Error().Err(err).Msg("unable to send error response")
	}
}
