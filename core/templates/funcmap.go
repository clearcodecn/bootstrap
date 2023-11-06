package templates

import "html/template"

func html(s string) template.HTML {
	return template.HTML(s)
}
