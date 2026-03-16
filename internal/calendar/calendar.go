package calendar

// CalDAV / iCloud calendar integration — to be implemented in step 2.

type Event struct {
	Title    string
	Start    string
	End      string
	Location string
}

type Client struct{}

func NewClient(rawURL, username, password string) *Client {
	return &Client{}
}

func (c *Client) EventsRange(days int) ([]Event, error) {
	// TODO: implement CalDAV fetch via go-webdav
	return nil, nil
}
