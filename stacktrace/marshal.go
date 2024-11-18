package stacktrace

import (
	"bytes"
	"encoding/json"

	"github.com/DataDog/gostackparse"
)

func Marshal(stack []byte) []byte {
	r := bytes.NewReader(stack)
	g, _ := gostackparse.Parse(r)
	b, _ := json.Marshal(g)
	return b
}
