package wiki

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	dummyData = `{
		"batchcomplete": "",
		"query": {
			"pages": {
				"21721040": {
					"extract": "月 Ɵɵ, Şş, Əçü ъзчиьбюжэё Stack Overflow is a question and answer site for professional and enthusiast.",
					"ns": 0,
					"pageid": 21721040,
					"title": "Stack Overflow"
				}
			}
		},
		"warnings": {
			"extracts": {
				"*": "\"exlimit\" was too large for a whole article extracts request, lowered to 1."
			}
		}
	}`
	notExpectedPageID = uint32(21721042)
	expectedPageID    = uint32(21721040)
	expectedTitle     = "Stack Overflow"
	expectedExtract   = "月 Ɵɵ, Şş, Əçü ъзчиьбюжэё Stack Overflow is a question and answer site for professional and enthusiast."
)

func TestGetQuery(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(dummyData))
	}))
	defer server.Close()

	// Use Client and URL from the local test server
	c := Client{server.Client(), server.URL}
	body, err := c.getPage(expectedPageID)
	if err != nil {
		t.Fatalf("error in test: %v\n", err)
	}

	expectedPage := Page{expectedExtract, expectedPageID, expectedTitle}
	assert.Equal(t, expectedPage, *body)
}

func TestGetQueryFailure(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(dummyData))
	}))
	defer server.Close()

	// Use Client and URL from the local test server
	c := Client{server.Client(), server.URL}
	body, err := c.getPage(expectedPageID)
	if err != nil {
		t.Fatalf("error in test: %v\n", err)
	}

	notExpectedPage := Page{expectedExtract, notExpectedPageID, expectedTitle}
	assert.NotEqual(t, notExpectedPage, *body)
}
