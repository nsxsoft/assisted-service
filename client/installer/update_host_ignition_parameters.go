// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// NewUpdateHostIgnitionParams creates a new UpdateHostIgnitionParams object
// with the default values initialized.
func NewUpdateHostIgnitionParams() *UpdateHostIgnitionParams {
	var ()
	return &UpdateHostIgnitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHostIgnitionParamsWithTimeout creates a new UpdateHostIgnitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateHostIgnitionParamsWithTimeout(timeout time.Duration) *UpdateHostIgnitionParams {
	var ()
	return &UpdateHostIgnitionParams{

		timeout: timeout,
	}
}

// NewUpdateHostIgnitionParamsWithContext creates a new UpdateHostIgnitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateHostIgnitionParamsWithContext(ctx context.Context) *UpdateHostIgnitionParams {
	var ()
	return &UpdateHostIgnitionParams{

		Context: ctx,
	}
}

// NewUpdateHostIgnitionParamsWithHTTPClient creates a new UpdateHostIgnitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateHostIgnitionParamsWithHTTPClient(client *http.Client) *UpdateHostIgnitionParams {
	var ()
	return &UpdateHostIgnitionParams{
		HTTPClient: client,
	}
}

/*UpdateHostIgnitionParams contains all the parameters to send to the API endpoint
for the update host ignition operation typically these are written to a http.Request
*/
type UpdateHostIgnitionParams struct {

	/*ClusterID
	  The cluster of the host whose ignition file should be updated.

	*/
	ClusterID strfmt.UUID
	/*HostIgnitionParams
	  Ignition config overrides.

	*/
	HostIgnitionParams *models.HostIgnitionParams
	/*HostID
	  The host whose ignition file should be updated.

	*/
	HostID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update host ignition params
func (o *UpdateHostIgnitionParams) WithTimeout(timeout time.Duration) *UpdateHostIgnitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update host ignition params
func (o *UpdateHostIgnitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update host ignition params
func (o *UpdateHostIgnitionParams) WithContext(ctx context.Context) *UpdateHostIgnitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update host ignition params
func (o *UpdateHostIgnitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update host ignition params
func (o *UpdateHostIgnitionParams) WithHTTPClient(client *http.Client) *UpdateHostIgnitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update host ignition params
func (o *UpdateHostIgnitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the update host ignition params
func (o *UpdateHostIgnitionParams) WithClusterID(clusterID strfmt.UUID) *UpdateHostIgnitionParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update host ignition params
func (o *UpdateHostIgnitionParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostIgnitionParams adds the hostIgnitionParams to the update host ignition params
func (o *UpdateHostIgnitionParams) WithHostIgnitionParams(hostIgnitionParams *models.HostIgnitionParams) *UpdateHostIgnitionParams {
	o.SetHostIgnitionParams(hostIgnitionParams)
	return o
}

// SetHostIgnitionParams adds the hostIgnitionParams to the update host ignition params
func (o *UpdateHostIgnitionParams) SetHostIgnitionParams(hostIgnitionParams *models.HostIgnitionParams) {
	o.HostIgnitionParams = hostIgnitionParams
}

// WithHostID adds the hostID to the update host ignition params
func (o *UpdateHostIgnitionParams) WithHostID(hostID strfmt.UUID) *UpdateHostIgnitionParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the update host ignition params
func (o *UpdateHostIgnitionParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHostIgnitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.HostIgnitionParams != nil {
		if err := r.SetBodyParam(o.HostIgnitionParams); err != nil {
			return err
		}
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
