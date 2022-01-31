package spec

import "github.com/getkin/kin-openapi/openapi3"

//go:generate oapi-codegen --package=spec --generate=types,server,spec -o spec.gen.go spec.v1.yaml

func init() {
	openapi3.DefineStringFormat("uuid", "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[4][0-9a-fA-F]{3}-[89ABab][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
	openapi3.DefineStringFormat("base64", "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")
}
