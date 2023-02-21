//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// GraphqlPostReader is a Reader for the GraphqlPost structure.
type GraphqlPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GraphqlPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGraphqlPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGraphqlPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGraphqlPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGraphqlPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGraphqlPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGraphqlPostOK creates a GraphqlPostOK with default headers values
func NewGraphqlPostOK() *GraphqlPostOK {
	return &GraphqlPostOK{}
}

/*
GraphqlPostOK describes a response with status code 200, with default header values.

Successful query (with select).
*/
type GraphqlPostOK struct {
	Payload *models.GraphQLResponse
}

// IsSuccess returns true when this graphql post o k response has a 2xx status code
func (o *GraphqlPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this graphql post o k response has a 3xx status code
func (o *GraphqlPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql post o k response has a 4xx status code
func (o *GraphqlPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this graphql post o k response has a 5xx status code
func (o *GraphqlPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql post o k response a status code equal to that given
func (o *GraphqlPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the graphql post o k response
func (o *GraphqlPostOK) Code() int {
	return 200
}

func (o *GraphqlPostOK) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostOK  %+v", 200, o.Payload)
}

func (o *GraphqlPostOK) String() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostOK  %+v", 200, o.Payload)
}

func (o *GraphqlPostOK) GetPayload() *models.GraphQLResponse {
	return o.Payload
}

func (o *GraphqlPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.GraphQLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphqlPostUnauthorized creates a GraphqlPostUnauthorized with default headers values
func NewGraphqlPostUnauthorized() *GraphqlPostUnauthorized {
	return &GraphqlPostUnauthorized{}
}

/*
GraphqlPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type GraphqlPostUnauthorized struct{}

// IsSuccess returns true when this graphql post unauthorized response has a 2xx status code
func (o *GraphqlPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql post unauthorized response has a 3xx status code
func (o *GraphqlPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql post unauthorized response has a 4xx status code
func (o *GraphqlPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this graphql post unauthorized response has a 5xx status code
func (o *GraphqlPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql post unauthorized response a status code equal to that given
func (o *GraphqlPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the graphql post unauthorized response
func (o *GraphqlPostUnauthorized) Code() int {
	return 401
}

func (o *GraphqlPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostUnauthorized ", 401)
}

func (o *GraphqlPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostUnauthorized ", 401)
}

func (o *GraphqlPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewGraphqlPostForbidden creates a GraphqlPostForbidden with default headers values
func NewGraphqlPostForbidden() *GraphqlPostForbidden {
	return &GraphqlPostForbidden{}
}

/*
GraphqlPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GraphqlPostForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this graphql post forbidden response has a 2xx status code
func (o *GraphqlPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql post forbidden response has a 3xx status code
func (o *GraphqlPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql post forbidden response has a 4xx status code
func (o *GraphqlPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this graphql post forbidden response has a 5xx status code
func (o *GraphqlPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql post forbidden response a status code equal to that given
func (o *GraphqlPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the graphql post forbidden response
func (o *GraphqlPostForbidden) Code() int {
	return 403
}

func (o *GraphqlPostForbidden) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostForbidden  %+v", 403, o.Payload)
}

func (o *GraphqlPostForbidden) String() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostForbidden  %+v", 403, o.Payload)
}

func (o *GraphqlPostForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GraphqlPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphqlPostUnprocessableEntity creates a GraphqlPostUnprocessableEntity with default headers values
func NewGraphqlPostUnprocessableEntity() *GraphqlPostUnprocessableEntity {
	return &GraphqlPostUnprocessableEntity{}
}

/*
GraphqlPostUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type GraphqlPostUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this graphql post unprocessable entity response has a 2xx status code
func (o *GraphqlPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql post unprocessable entity response has a 3xx status code
func (o *GraphqlPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql post unprocessable entity response has a 4xx status code
func (o *GraphqlPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this graphql post unprocessable entity response has a 5xx status code
func (o *GraphqlPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql post unprocessable entity response a status code equal to that given
func (o *GraphqlPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the graphql post unprocessable entity response
func (o *GraphqlPostUnprocessableEntity) Code() int {
	return 422
}

func (o *GraphqlPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GraphqlPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GraphqlPostUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GraphqlPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphqlPostInternalServerError creates a GraphqlPostInternalServerError with default headers values
func NewGraphqlPostInternalServerError() *GraphqlPostInternalServerError {
	return &GraphqlPostInternalServerError{}
}

/*
GraphqlPostInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type GraphqlPostInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this graphql post internal server error response has a 2xx status code
func (o *GraphqlPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql post internal server error response has a 3xx status code
func (o *GraphqlPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql post internal server error response has a 4xx status code
func (o *GraphqlPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this graphql post internal server error response has a 5xx status code
func (o *GraphqlPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this graphql post internal server error response a status code equal to that given
func (o *GraphqlPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the graphql post internal server error response
func (o *GraphqlPostInternalServerError) Code() int {
	return 500
}

func (o *GraphqlPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostInternalServerError  %+v", 500, o.Payload)
}

func (o *GraphqlPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /graphql][%d] graphqlPostInternalServerError  %+v", 500, o.Payload)
}

func (o *GraphqlPostInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GraphqlPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
