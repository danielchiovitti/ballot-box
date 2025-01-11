package model

import "github.com/danielchiovitti/ballot-box/pkg/shared"

type JsonErrorMessage struct {
	Message shared.ErrorMessage
	Code    shared.ErrorCode
}
