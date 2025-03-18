package errcode

import "net/http"

var (
	OK = New(http.StatusOK, "", "OK")

	ErrInternal = New(http.StatusInternalServerError, "InternalError", "Internal Server Error")

	ErrNotFound = New(http.StatusNotFound, "NotEound", "Not Found")

	ErrDBRead = New(http.StatusInternalServerError, "Internal.DBReadError", "DB Read Error")

	ErrDBWrite = New(http.StatusInternalServerError, "Internal.DBWriteError", "DB Write Error")
)
