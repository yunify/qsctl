package constants

import "errors"

var (
	// ErrorFlowInvalid returned while flow is invalid.
	ErrorFlowInvalid = errors.New("flow is invalid")
	// ErrorActionNotImplemented returned while current is not implemented.
	ErrorActionNotImplemented = errors.New("action not implemented")

	// ErrorExpectSizeRequired returned while expect size is required but not given.
	ErrorExpectSizeRequired = errors.New("expect-size is required")

	// ErrorQsPathInvalid returned while qs-path is invalid.
	ErrorQsPathInvalid = errors.New("qingstor path invalid")
	// ErrorQsPathObjectKeyRequired returned while object key is required but not given.
	ErrorQsPathObjectKeyRequired = errors.New("qingstor path object key is required")

	// ErrorFileTooLarge returned while file is too large.
	ErrorFileTooLarge = errors.New("file too large")
	// ErrorFileNotExist returned while file is not found.
	ErrorFileNotExist = errors.New("file not exist")
)
