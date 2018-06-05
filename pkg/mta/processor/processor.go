package processor

import (
	"io/ioutil"

	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
)

// YulmailsProcessor checks if
var YulmailsProcessor = func() backends.Decorator {
	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				if task == backends.TaskSaveMail {
					ioutil.WriteFile("/tmp/mail", []byte(e.RcptTo[0].String()), 0644)
					return p.Process(e, task)
				}
				return p.Process(e, task)
			},
		)
	}
}
