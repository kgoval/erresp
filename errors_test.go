package erresp

import (
	"testing"
)

var pack = map[string]*MessageBody{
	"ErrServInvalidErrorCode": &MessageBody{
		Lang: map[string]string{
			"ID": "Hubungi developer kami, Error Code tidak terdefinisi",
			"EN": "Call our developer, Error Code Not Defined yet",
		},
		Code: 500,
	},
}

var lang = []string{"ID", "EN",}

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

	for _, l := range lang {
		detail, _ := GetMessage("ErrServInvalidErrorCode", l)

		if detail != pack["ErrServInvalidErrorCode"].Lang[l] {
			t.Errorf("Cannot get lang %s", l)
		}
	}

}

func TestNew(t *testing.T) {

	devMessage := "Error bro!"
	id := "ErrServInvalidErrorCode"
	for _, l := range lang {

		err := New(id, l, devMessage)

		// check if err has implement err
		if err.Error() == "" {
			t.Error("generated error failed")
		}

		// test error body
		conv := err.(*Error)
		if conv.Id != id {
			t.Error("invalid id")
		}

		if conv.DevMessage != devMessage {
			t.Error("invalid id")
		}

		if conv.Detail != pack[id].Lang[l] {
			t.Errorf("failed generated detail error for %s language", l)
		}

	}

}

func TestNewf(t *testing.T) {
	format := "error bro %s %s"
	devMessage1 := "1"
	devMessage2 := "2"
	want := "error bro 1 2"
	id := "ErrServInvalidErrorCode"
	for _, l := range lang {

		err := Newf(id, l, format, devMessage1, devMessage2)

		// check if err has implement err
		if err.Error() == "" {
			t.Error("generated error failed")
		}

		// test error body
		conv := err.(*Error)
		if conv.Id != id {
			t.Error("invalid id")
		}

		if conv.DevMessage != want {
			t.Error("invalid id")
		}

		if conv.Detail != pack[id].Lang[l] {
			t.Errorf("failed generated detail error for %s language", l)
		}
	}
}

func TestParse(t *testing.T) {
	json := `{"id":"ErrServInvalidErrorCode","code":500,"detail":"Call our developer, Error Code Not Defined yet","status":"Internal Server Error","dev_message":"Error bro!"}`
	jsonErr := `{asasas"}`
	errResp, err := Parse(json)
	if err != nil {
		t.Error(err)
	}

	if errResp.Code != 500 {
		t.Error("invalid code to parse")
	}

	// must error
	_, err = Parse(jsonErr)
	if err == nil {
		t.Error("parse string must be error")
	}
}
