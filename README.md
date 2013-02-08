GDG Event
=========

A command line tool and a library written in GOLANG for GDG Event.

Install
-------
command line tool:

    go get github.com/dalinaum/gdgevent

library:

    go get github.com/dalinaum/gdgevent/event

Command line tool usage
-----
`gdgevent` `<chapter-id>` `<year>`

    gdgevent 102751345660146384940 2012

Sample
------
~~~~
package main

import (
    "github.com/dalinaum/gdgevent/event"
	"time"
)

func main() {
	start := time.Unix(0, 0)
	end := time.Unix(0, 0)
	// if you set start is equal to end, end will be ignored.

	cid := "102751345660146384940"

	for _, e := range event.GetGDGEvents(cid, start.UTC(), end.UTC()) {
		e.PrintSummary()
	}
}
~~~~

If you want to know more, see [command line tool's source](https://github.com/dalinaum/gdgevent/blob/master/main.go) please.

Authors
---------
 * Homin Lee <homin.lee@suapapa.net>
 * Leonardo YongUk kIm <dalinaum@gmail.com>
