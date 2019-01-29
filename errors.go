package erresp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

var message map[string]*MessageBody

/**
	Registering message pack
 */

func Register(pack map[string]*MessageBody){
	message = pack
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
func New(id, lang string, format string, devMessage ...interface{} ) error {
	err := &Error{}
	err.Id = id
	err.Detail, err.Code = GetMessage(id, lang)
	err.Status = http.StatusText(int(err.Code))
	err.DevMessage = fmt.Sprintf(format, devMessage...)
	return err
}


func Parse(message string) (*Error, error) {
	e := new(Error)
	errr := json.Unmarshal([]byte(message), e)
	if errr != nil {
		return nil, errors.Wrap(errr, "json unmarshal failed")
	}
	return e, nil
}

type MessageBody struct{
	ID string
	EN string
	Code int32
}

func GetMessage(id, lang string)(detail string, code int32){

	if lang == ""{
		lang = "ID"
	}
	m := message[id]

	// if error not defined
	if m == nil {
		return "NotDefErrMes", 500
	}

	if lang == "EN"{
		return m.EN, m.Code
	}

	return m.ID, m.Code
}