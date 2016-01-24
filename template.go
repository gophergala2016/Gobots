package main

import "html/template"

type tmpl struct {
	*template.Template
}

func (t *tmpl) ExecuteTemplate(c context, name string, data map[string]interface{}) error {
	data["Host"] = c.r.Host
	data["Random"] = imAnIdiot
	data["ClientID"] = clientId

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
