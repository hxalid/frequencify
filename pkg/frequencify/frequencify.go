package wiki

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/hxalid/frequencify/pkg/util"
	"github.com/pkg/errors"
)

// Query - query can have multiple pages
type Query struct {
	Query Pages
}

// Pages - a map of page id and Page
type Pages struct {
	Pages map[uint32]Page
}

// Page - a single page
type Page struct {
	Extract string
	PageID  uint32
	Title   string
}

// Client - HTTP client object with endpoint
type Client struct {
	httpClient *http.Client
	url        string
}

// NewClient - creates a Client object
func NewClient(timeout time.Duration, url string) *Client {
	cli := Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		url: url,
	}

	return &cli
}

// getPages - extracts content of the page identified by pid
func (cli *Client) getPage(pid uint32) (*Page, error) {
	req, err := http.NewRequest(http.MethodGet, cli.url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response body failed")
	}
	bodyString := string(bodyBytes)

	var query Query
	err = json.Unmarshal(bodyBytes, &query)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unmarshaling of %s failed", bodyString))
	}

	if query.Query.Pages == nil {
		return nil, errors.Wrap(err, "pages is nil")
	}

	p, ok := query.Query.Pages[pid]
	if !ok {
		return nil, errors.Wrap(err, fmt.Sprintf("could not extract page for id %d\n", pid))
	}

	return &p, nil
}

// Frequencify - prints top count frequencies on page pid
func (cli *Client) Frequencify(count uint32, pid uint32) {
	//send http request
	p, err := cli.getPage(pid)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("URL: %s\n\n", cli.url)
	fmt.Printf("Title: %s\n\n", p.Title)
	fmt.Printf("Top %d words:\n\n", count)

	wc := util.ReverseMap(util.WordCount(p.Extract))
	keys := make([]int, 0, len(wc))
	for key := range wc {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	c := uint32(0)
	for _, key := range keys {
		fmt.Printf("- %d, %s\n\n", key, wc[key])
		c++
		if c == count {
			break
		}
	}
}
