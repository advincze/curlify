package curlify

import (
	"bytes"
	"net/http"
)

func DumpToCurl(req *http.Request) string {
	var buffer bytes.Buffer
	buffer.WriteString("curl")

	if req.Method != "GET" {
		buffer.WriteString(" -X ")
		buffer.WriteString(req.Method)
	}

	for key, values := range req.Header {
		if key == "" {
			continue
		}
		for _, value := range values {
			buffer.WriteString(" -H ")
			buffer.WriteString(`"`)
			buffer.WriteString(key)
			buffer.WriteString(":")
			buffer.WriteString(value)
			buffer.WriteString(`"`)
		}
	}

	buffer.WriteString(" ")
	buffer.WriteString(req.URL.String())

	return buffer.String()
}
