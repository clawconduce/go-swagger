package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/todo-list/models"
)

// NewUpdateOneParams creates a new UpdateOneParams object
// with the default values initialized.
func NewUpdateOneParams() *UpdateOneParams {
	var ()
	return &UpdateOneParams{}
}

/*UpdateOneParams contains all the parameters to send to the API endpoint
for the update one operation typically these are written to a http.Request
*/
type UpdateOneParams struct {

	/*Body*/
	Body *models.Item
	/*ID*/
	ID string
}

// WithBody adds the body to the update one params
func (o *UpdateOneParams) WithBody(body *models.Item) *UpdateOneParams {
	o.Body = body
	return o
}

// WithID adds the id to the update one params
func (o *UpdateOneParams) WithID(id string) *UpdateOneParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.Item)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
