package payload

import (
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
)

type Payload struct {
	Data      any                   `json:"data"`
	Publisher gogen.ApplicationData `json:"publisher"`
	TraceID   string                `json:"traceId"`
}
