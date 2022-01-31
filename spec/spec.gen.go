// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Defines values for ErrorCode.
const (
	ErrorCodeGENERIC ErrorCode = "GENERIC"

	ErrorCodeINVALIDREQUEST ErrorCode = "INVALID_REQUEST"

	ErrorCodeRELOGINREQUIRED ErrorCode = "RELOGIN_REQUIRED"

	ErrorCodeSECURITYFAILED ErrorCode = "SECURITY_FAILED"
)

// Error defines model for Error.
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// ErrorCode defines model for ErrorCode.
type ErrorCode string

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Language defines model for Language.
type Language struct {
	CultureCode  string `json:"culture_code"`
	DisplayOrder int    `json:"display_order"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Published    bool   `json:"published"`
	Rtl          bool   `json:"rtl"`
}

// LanguageResponse defines model for LanguageResponse.
type LanguageResponse []Language

// InternalError defines model for InternalError.
type InternalError ErrorResponse

// SecurityError defines model for SecurityError.
type SecurityError ErrorResponse

// GetAPISpecParamsFormat defines parameters for GetAPISpec.
type GetAPISpecParamsFormat string

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Send feedback
	// (POST /feedbacks)
	SendFeedback(ctx echo.Context) error
	// Uploads a file.
	// (POST /fileupload1)
	FileUpload1(ctx echo.Context) error
	// Uploads a file.
	// (POST /fileupload2)
	FileUpload2(ctx echo.Context) error
	// Get languages.
	// (GET /languages)
	Languages(ctx echo.Context) error
	// Get OpenAPI 3 specification.
	// (GET /spec{format})
	GetAPISpec(ctx echo.Context, format GetAPISpecParamsFormat) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// SendFeedback converts echo context to params.
func (w *ServerInterfaceWrapper) SendFeedback(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SendFeedback(ctx)
	return err
}

// FileUpload1 converts echo context to params.
func (w *ServerInterfaceWrapper) FileUpload1(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FileUpload1(ctx)
	return err
}

// FileUpload2 converts echo context to params.
func (w *ServerInterfaceWrapper) FileUpload2(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FileUpload2(ctx)
	return err
}

// Languages converts echo context to params.
func (w *ServerInterfaceWrapper) Languages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Languages(ctx)
	return err
}

// GetAPISpec converts echo context to params.
func (w *ServerInterfaceWrapper) GetAPISpec(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "format" -------------
	var format GetAPISpecParamsFormat

	err = runtime.BindStyledParameterWithLocation("simple", false, "format", runtime.ParamLocationPath, ctx.Param("format"), &format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAPISpec(ctx, format)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/feedbacks", wrapper.SendFeedback)
	router.POST(baseURL+"/fileupload1", wrapper.FileUpload1)
	router.POST(baseURL+"/fileupload2", wrapper.FileUpload2)
	router.GET(baseURL+"/languages", wrapper.Languages)
	router.GET(baseURL+"/spec:format", wrapper.GetAPISpec)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xX3W7jNhN9FYLfd6la3mQL7OqqbmIHQt1sajsFisAIaGpkMUuRKjnK1gj07gWpH0ux",
	"vUnbbVG0vbIscuYMz5wZDZ8o13mhFSi0NHqiBmyhlQX/J1YIRjE5NUYb94JrhaDQPbKikIIzFFqFD1Yr",
	"987yDHLmnv5vIKUR/V+49x7Wqzb03hYNDq2qKqAJWG5E4ZzRqIMlfueIVgFdAi+NwN1fHEkL20VSBY1r",
	"z08XTWF0AQZFTRvXCbwK+sJtrAKag7Vs621wVwCNqEUj1NbjGfi5FAYSGt3Vnvf710G7X28egKPztXcc",
	"PVFQZe7srqbX00V8QQO6nF7cLuLVT/ezSTyfXtKAxtc/Tubx5f1i+sPtdLmiAV1M5x+u4mv/Jl5ML3s4",
	"bVwNTsfdAQfQUvMiCQeHrE2PnW3O1LZsiHpGeSmxNHDfUn8QbiJsIdnuXpsEfFypNjlDGlGhkHZYQiFs",
	"wTgLkbxqm2L5ccSi3EhhM0h6qxutJTDllg3KYwvPuBAJDYaHaxBrB32U52f8HIH9vAmE3L6UqY75qvPK",
	"jGG7uiKaKlm6zXU+NsAMmEmJ2f7frKXyorSoc9qUkj++37DnN0MsHBTX+qOA1o1wJVm/ammIKOb3Fqx1",
	"9boPrRDfwa4uaKFS3TYMxn3DgJwJ6bOY6m9SBJ5JtrEj7mNq3M5WwDMyZxtLA1oa2QQVheHQ4KBlTIgF",
	"JDolk5vYEtSESak/kV6TssTJxzCObvmTwIxgBmRlWALf642QQOzOIuQjdyKB0sXTX11Ml6u0lB6BBvQR",
	"jK2x34zGo7ELSRegWCFoRM9H49G50wnDzGcmTAGSDeMf/b9CW0+JqyUfXJz4tqeSWbON1oIEi9/qZPes",
	"9ealRFEwg6Grk68ShqxuO1wnrgxcd0ZkPMsbg8Z21ZRRzrYQFmobkPrxoYDec/e4FWmv9R5W/xCjK9mN",
	"UMzs6JHmlcCj4MfL9nQvDqgt60p6sU+fbtBVvbX3kT0bv3U/QxWtMiBtokjGLNkAKGJBIbEl52BtWkq5",
	"89/Gt+PzU+XbAYXDD2gV0K/H45ethgNAv9ZpdLd2hOS5o7iWTBexo5xtrWPiQue5VnTtbMNUSCgLqVny",
	"5rT4ZkLCbbPpP+19Ce052v+5yqu1Ygkj7pyjl7V39hrtnX1We3tpDGbNk9nvqj+ge7H9XttC/WbTf3fq",
	"ZTO8+JNv4Uje592OA57GX+y2cTB8HblwLABLoyzpQv77kH0F2A/rBNe2AP5UC7Lq0f3sgofENAd1k4+z",
	"EWlDqhudMBPWTTcOZZioK8DJTbwsgPuZxrAcEIwL4wgGZ4psgJQWEjdpCZU4DCAJWNcsScsCqeN1aH7E",
	"dLPSfhJsqqvfZNGU0P8itNeskRdEQEe7XPoflssjl6dq/QdVdtgZEH7B0MMd29gBH+jtQwFqchOT89GY",
	"LPtp+JME1OENk96X06xUCXPfUyadpoYOhzeLO2pgW0pm6LoKnga3hf6SiwLMY6uTwTQvNWcy0xaj9+/e",
	"vwsf39BqXf0aAAD//xccowEVEQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}