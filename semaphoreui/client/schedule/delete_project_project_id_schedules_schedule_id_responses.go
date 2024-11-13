// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProjectProjectIDSchedulesScheduleIDReader is a Reader for the DeleteProjectProjectIDSchedulesScheduleID structure.
type DeleteProjectProjectIDSchedulesScheduleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectProjectIDSchedulesScheduleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectProjectIDSchedulesScheduleIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /project/{project_id}/schedules/{schedule_id}] DeleteProjectProjectIDSchedulesScheduleID", response, response.Code())
	}
}

// NewDeleteProjectProjectIDSchedulesScheduleIDNoContent creates a DeleteProjectProjectIDSchedulesScheduleIDNoContent with default headers values
func NewDeleteProjectProjectIDSchedulesScheduleIDNoContent() *DeleteProjectProjectIDSchedulesScheduleIDNoContent {
	return &DeleteProjectProjectIDSchedulesScheduleIDNoContent{}
}

/*
DeleteProjectProjectIDSchedulesScheduleIDNoContent describes a response with status code 204, with default header values.

schedule deleted
*/
type DeleteProjectProjectIDSchedulesScheduleIDNoContent struct {
}

// IsSuccess returns true when this delete project project Id schedules schedule Id no content response has a 2xx status code
func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project project Id schedules schedule Id no content response has a 3xx status code
func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project project Id schedules schedule Id no content response has a 4xx status code
func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project project Id schedules schedule Id no content response has a 5xx status code
func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project project Id schedules schedule Id no content response a status code equal to that given
func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project project Id schedules schedule Id no content response
func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) Code() int {
	return 204
}

func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/schedules/{schedule_id}][%d] deleteProjectProjectIdSchedulesScheduleIdNoContent", 204)
}

func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /project/{project_id}/schedules/{schedule_id}][%d] deleteProjectProjectIdSchedulesScheduleIdNoContent", 204)
}

func (o *DeleteProjectProjectIDSchedulesScheduleIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}