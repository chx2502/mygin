package mygin

import (
	"fmt"
	"reflect"
	"testing"
)

var r = newRouter()
type parsePatternTestCase struct {
	pattern string
	want []string
}

var patterns = [] map[string]string {
	{ "Method": "GET", "Path": "/" },
	{ "Method": "GET", "Path": "/hello/:name" },
	{ "Method": "GET", "Path": "/hello/:language/:name" },
	{ "Method": "GET", "Path": "/hello/b/c" },
	{ "Method": "GET", "Path": "/hi/:name" },
	{ "Method": "GET", "Path": "/assets/*filepath" },
}

func initTestRouter() {
	for _, pattern := range patterns {
		r.addRoute(pattern["Method"], pattern["Path"], nil)
	}
}

func TestRouter_ParsePattern(t *testing.T) {
	cases := []parsePatternTestCase {
		{ pattern: "/p/:name", want: []string { "p", ":name" } },
		{ pattern: "/p/*", want: []string { "p", "*" } },
		{ pattern: "/p/*name/*", want: []string { "p", "*name" } },
	}
	for _, c := range cases {
		ok := reflect.DeepEqual(parsePattern(c.pattern), c.want)
		if !ok {
			t.Fatal("test parsePattern failed\n")
		}
	}
}

func TestRouter_GetRoute(t *testing.T) {
	initTestRouter()
	n, params := r.getRoute("GET", "/hello/chx")

	if n == nil {
		t.Fatal("getRoute test failed, a nil has been returned\n")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("getRoute test failed, wrong pattern\n")
	}

	if params["name"] != "chx" {
		t.Fatal("getRoute test failed, wrong params\n")
	}
	fmt.Printf("matched path: %s, paramas['name']: %s\n", n.pattern, params["name"])
}

