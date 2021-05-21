// Created by vinson on 2021/5/21.

package milkomeda

import (
	"fmt"
	"milkomeda.org/pi"
	"net/http"
	"time"
)

var stat = &pi.InterposerStat{
	Name:      "DefaultRouter",
	Author:    "vinson",
	CreatedAt: time.Date(2021, 5, 21, 0, 0, 0, 0, &time.Location{}),
}

type DefaultRouter struct {
}

func (engine *DefaultRouter) Stat() *pi.InterposerStat {
	return stat
}

func (engine DefaultRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintln(w, "hello pi!")
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
