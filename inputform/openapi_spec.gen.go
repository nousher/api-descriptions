// Package newForm provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package newForm

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9SUTU/bQBCG/8pq2qOJA60q1TdAiYhUQtSCKrXisNiTeMH70dkJIUL+79WuE/LlBEFP",
	"Pdnyzr7zzjMzfobcamcNGvaQPYPPS9Qyvp4WBaH3AzO2pCUra8JXR9YhscIYc654Hp48dwgZeCZlJlAn",
	"8IMl454TQuSF+OsRJ60hv5Q7t0VM0LiDDJThL58hWQYrwzhBgrpOgPDPVBEWkP3eyp80JSwNr5Rv6wTO",
	"rWGZ82ECWFWtNrR8UnqqIfu6aymBnpaqai3twmocldbgO1TjveFU3yG94/ZPSw/vy7yFeN3GstTA8wJl",
	"xWVeYv5wdXePOe/i7BFZukTv5QT3Dta0bW522hzjkk3F4GLx3ierz7ApdtPDofT/QnjL4bqnIc76lnTP",
	"MM13DbVv4kfCMWTwIV1tcLpY37Tlxt55PqTTcqOd4CGR7fA6gRuP9AYb2+EB5YKubQapXXMTY1+R56HU",
	"7Z39JvcebjVupbN267YOYcqMbRRQXAWFwXB0c33Uv/p+eXQ6GkACj0g+moPjTrfTDYmtQyOdggw+xU8J",
	"OMlldJxKp9LH45TRc1qutiecTTCuT4E+J+WaimEtRowtCS5RhMvCIz2qHL3wcS86EPNSJDUoIIPeEyMZ",
	"Wa1nCWV7Z41v+J10u+GRW8NoYnLpXKXyKJLe+4Z507PXOrr7K4j8Nqu5LtXSsFBeyGom514Q8pQMFmJW",
	"olkWFs6vHjqxV36qtQx7tPjhiBYe0qkXJh1x47EQA9MQqOZRZYO9wVkYLHzZT+tb6PdMMbLKsGArlPFI",
	"LArJMmY2OAtPvUt+Y/ebSUPPZ7aYv4m2YtT+Newbueq6bkb7v+5xzxSilboybGPDA3qWdxV2Ysl/AwAA",
	"///L/l1/8wgAAA==",
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

