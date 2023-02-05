package gopmat

import "encoding/json"

type _json bool

const JSON _json = true

func (_json) From(v any) (string, error) {
	bytes, err := json.Marshal(v)
	return string(bytes), err
}

func (_json) FromPretty(v any) (string, error) {
	bytes, err := json.MarshalIndent(v, "", "    ")
	return string(bytes), err
}

func (_json) Str(v any) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func (_json) StrPretty(v any) string {
	bytes, _ := json.MarshalIndent(v, "", "    ")
	return string(bytes)
}
