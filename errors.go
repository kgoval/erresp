package erresp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var message map[string]*MessageBody

/**
	Registering message pack
 */

func Register(pack map[string]*MessageBody){
	message = pack
}

type MessageBody struct{
	Lang map[string]string
	Code int32
}

// Error implements the error interface.
type Error struct {
	Id     string `json:"id"`
	Code   int32  `json:"code"` // http code
	Detail string `json:"detail"`
	Status string `json:"status"`
	DevMessage string `json:"dev_message"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// New generates a custom error.
func New(id, lang, devMessage string) error {
	err := &Error{}
	err.Id = id
	err.Detail, err.Code = GetMessage(id, lang)
	err.Status = http.StatusText(int(err.Code))
	err.DevMessage = devMessage
	return err
}

func Newf(id, lang string, format string, devMessage ...interface{} ) error {
	err := &Error{}
	err.Id = id
	err.Detail, err.Code = GetMessage(id, lang)
	err.Status = http.StatusText(int(err.Code))
	err.DevMessage = fmt.Sprintf(format, devMessage...)
	return err
}


func Parse(message string) *Error {
	e := new(Error)
	errr := json.Unmarshal([]byte(message), e)
	if errr != nil {
		 e.Detail  = errr.Error()
	}
	return e
}

func GetMessage(id, lang string)(detail string, code int32){

	return message[id].Lang[lang], message[id].Code
}