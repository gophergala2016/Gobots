package main

import "html/template"

type tmpl struct {
	*template.Template
}

func (t *tmpl) ExecuteTemplate(c context, name string, data interface{}) error {
	if err := t.Template.ExecuteTemplate(c.w, "head.html", data); err != nil {
		return err
	}

	if err := t.Template.ExecuteTemplate(c.w, name, data); err != nil {
		return err
	}

	if err := t.Template.ExecuteTemplate(c.w, "foot.html", data); err != nil {
		return err
	}

	return nil
}
