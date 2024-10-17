package rest

import "net/http"

func Ok(body any) Response {
	return Response{
		Status: http.StatusOK,
		Body:   body,
	}
}

func Created(body any) Response {
	return Response{
		Status: http.StatusCreated,
		Body:   body,
	}
}

func BadRequest(err error) Response {
	return Response{
		Status: http.StatusBadRequest,
		Error:  err,
	}
}

func InternalServerError(err error) Response {
	return Response{
		Status: http.StatusInternalServerError,
		Error:  err,
	}
}
