package resty

type Client struct {
	// ... existing fields
	DisableMultipartRetry bool
}

func (c *Client) SetDisableMultipartRetry(d bool) *Client {
	c.DisableMultipartRetry = d
	return c
}

// Inside the execution loop where retry logic resides:
func (c *Client) execute(req *Request) (*Response, error) {
	// ...
	for i := 0; i <= c.RetryCount; i++ {
		resp, err := c.do(req)
		if err != nil && c.shouldRetry(req, err) {
			if req.isMultipart() && c.DisableMultipartRetry {
				return resp, err
			}
			// proceed with retry
		}
		// ...
	}
}