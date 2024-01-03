package postfix_namespace

import "github.com/statsd/client-interface"
import "time"

type Client struct {
	c      statsd.Client
	postfix string
}

func New(client statsd.Client, postfix string) *Client {
	return &Client{client, postfix}
}

func (c *Client) Gauge(name string, value int) error {
	return c.c.Gauge(name+"."+c.postfix, value)
}

func (c *Client) Incr(name string) error {
	return c.c.Incr(name + c.postfix)
}

func (c *Client) IncrBy(name string, value int) error {
	return c.c.IncrBy(name+"."+c.postfix, value)
}

func (c *Client) Decr(name string) error {
	return c.c.Decr(name + c.postfix)
}

func (c *Client) DecrBy(name string, value int) error {
	return c.c.DecrBy(name+"."+c.postfix, value)
}

func (c *Client) Duration(name string, duration time.Duration) error {
	return c.c.Duration(name+"."+c.postfix, duration)
}

func (c *Client) Histogram(name string, value int) error {
	return c.c.Histogram(name+"."+c.postfix, value)
}

func (c *Client) Annotate(name string, value string, args ...interface{}) error {
	return c.c.Annotate(name, value, args...)
}

func (c *Client) Flush() error {
	return c.c.Flush()
}
