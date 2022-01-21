package domain

import "errors"

var (
	ErrMissingArgument    	= errors.New("missing argument")
	ErrInvalidMessageType 	= errors.New("invalid message-type")
	ErrNotFound 			= errors.New("not found")
	ErrNotProcessed 		= errors.New("not processed")
	ErrInvalidArgument    	= errors.New("invalid argument")
	ErrParseFailure			= errors.New("parse failure")
)
