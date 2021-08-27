package customerror

import (
	"fmt"
	"strings"
)

var ErrInvalidRequest = New("400", "01710001", "ERR_INVALID_REQUEST", "You have invalid request, please check your request again.")
var ErrEntityNotFound = New("404", "01710002", "ERR_ENTITY_NOT_FOUND", "Sorry, you are looking for nothing.")
var ErrInternalServerError = New("500", "01710003", "ERR_INTERNAL_SERVER_ERROR", "Uh, oh! There is problems, we will fix it. Come back later.")
var ErrCouldNotConnectDB = New("500", "01710004", "ERR_COULD_NOT_CONNECT_DATABASE", "could not connect to database.")
var ErrCouldNotConnectRedis = New("500", "01710005", "ERR_COULD_NOT_CONNECT_REDIS", "could not connect to redis.")
var ErrFailedCommunicateWithRepository = New("500", "01710006", "ERR_FAILED_COMMUNICATE_WITH_REPOSITORY", "Uh, oh! There is problems, we will fix it. Come back later.")

type Error struct {
	Status string       `json:"status"`
	Code   string       `json:"code"`
	Title  string       `json:"title"`
	Detail string       `json:"detail"`
	Source *ErrorSource `json:"source,omitempty"`
	Op     string       `json:"-"`
	Err    error        `json:"-"`
}

type ErrorSource struct {
	Parameter string `json:"parameter,omitempty"`
	Header    string `json:"header,omitempty"`
}

func New(status, code, title, detail string) *Error {
	return &Error{
		Status: status,
		Code:   code,
		Title:  title,
		Detail: detail,
	}
}

func NewOpError(op string, err error) *Error {
	return &Error{
		Op:  op,
		Err: err,
	}
}

func (e *Error) WithSource(parameter, header string) *Error {
	clone := e.clone()
	clone.Source = &ErrorSource{
		Parameter: parameter,
		Header:    header,
	}
	return clone
}

func (e *Error) clone() *Error {
	clone := *e
	return &clone
}

func (e *Error) Error() string {
	var buf strings.Builder

	if e.Err != nil {

		if e.Op != "" {
			fmt.Fprintf(&buf, "%s: ", e.Op)
		}
		buf.WriteString(e.Err.Error())
	} else {
		fmt.Fprintf(
			&buf,
			`{"status":"%s","code":"%s","title":"%s","detail":"%s"`,
			e.Status, e.Code, e.Title, e.Detail,
		)
		if e.Source != nil {
			fmt.Fprintf(
				&buf,
				`,"source":{"parameter":"%s","header":"%s"}`,
				e.Source.Parameter, e.Source.Header,
			)
		}
		buf.WriteRune('}')
	}

	return buf.String()
}

func (e *Error) ErrorDetail() string {
	return errorDetail(e)
}

func errorDetail(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Detail != "" {
		return e.Detail
	} else if ok && e.Err != nil {
		return errorDetail(e.Err)
	} else if !ok {
		return err.Error()
	}
	return "An internal error has occurred."
}
