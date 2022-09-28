package main

import (
	h "html/template"
	t "text/template"
)

func main() {
	t.New("tpl").Parse(`Template TEXT`)
	h.New("tpl").Parse(`Template HTML`)
}
