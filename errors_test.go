package erresp

import (
	"testing"
)

var pack = map[string]*MessageBody{
	"ErrServInvalidErrorCode": &MessageBody{
		ID:   "Hubungi developer kami, Error Code tidak terdefinisi",
		EN:   "Call our developer, Error Code Not Defined yet",
		Code: 500,
	},
}

func TestRegister(t *testing.T) {

	Register(nil)

	if message != nil {
		t.Error("message must be nil")
	}

	Register(pack)

	if message == nil {
		t.Error("message must not be nil, register failed")
	}

}

func TestGetMessage(t *testing.T) {

	if message == nil {
		t.Error("message must not be nil")
	}

	detail, _ := GetMessage("ErrServInvalidErrorCode", "ID")

	if detail != pack["ErrServInvalidErrorCode"].ID {
		t.Error("Cannot get lang ID ")
	}

	detail, _ = GetMessage("ErrServInvalidErrorCode", "EN")

	if detail != pack["ErrServInvalidErrorCode"].EN {
		t.Error("Cannot get lang EN ")
	}

}


func TestNew(t *testing.T){
	
}