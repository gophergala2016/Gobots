package main

import "html/template"

type tmpl struct {
	*template.Template
}

type tmplData struct {
	Active      int
	Host        string
	RedirectUri template.URL
	Random      string
	ClientID    string
	Data        map[string]interface{}
	Player      *player
	Scripts     []template.URL
}

func (t *tmpl) ExecuteTemplate(c context, name string, data tmplData) error {
	data.Host = c.r.Host
	data.RedirectUri = template.URL("http://" + c.r.Host + "/auth")
	data.Random = c.magicToken
	data.ClientID = clientId
	data.Player = c.p

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
