package curlify

import (
	"log"
	"net/http"
	"testing"
)

func TestGetNoHeaders(t *testing.T) {
	req, err := http.NewRequest("POST", "http://google/com", nil)
	if err != nil {
		t.Errorf("the request should be built")
	}

	log.Printf("%s\n", DumpToCurl(req))
}
