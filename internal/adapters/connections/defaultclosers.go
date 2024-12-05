package connections

type Closer struct {
	closeFunc func() error
}

func (c *Closer) Close() error {
	return c.closeFunc()
}

var DefaulCloser = &Closer{closeFunc: func() error { return nil }}
