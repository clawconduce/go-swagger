package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*LoginUserOK successful operation

swagger:response loginUserOK
*/
type LoginUserOK struct {

	// In: body
	Payload string `json:"body,omitempty"`
}

// NewLoginUserOK creates LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {
	return &LoginUserOK{}
}

// WithPayload adds the payload to the login user o k response
func (o *LoginUserOK) WithPayload(payload string) *LoginUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login user o k response
func (o *LoginUserOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*LoginUserBadRequest Invalid username/password supplied

swagger:response loginUserBadRequest
*/
type LoginUserBadRequest struct {
}

// NewLoginUserBadRequest creates LoginUserBadRequest with default headers values
func NewLoginUserBadRequest() *LoginUserBadRequest {
	return &LoginUserBadRequest{}
}

// WriteResponse to the client
func (o *LoginUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}
