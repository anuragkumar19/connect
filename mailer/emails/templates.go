package emails

import (
	"fmt"
	"text/template"
)

type Templates struct {
	verifyEmail *template.Template
	welcome     *template.Template
}

func NewTemplates() (Templates, error) {
	verifyEmail, err := parseTemplate("verify-email")
	if err != nil {
		return Templates{}, err
	}
	welcome, err := parseTemplate("welcome")
	if err != nil {
		return Templates{}, err
	}
	return Templates{
		verifyEmail: verifyEmail,
		welcome:     welcome,
	}, nil
}

func parseTemplate(name string) (*template.Template, error) {
	path := fmt.Sprintf("generated/%s.html", name)
	templ, err := template.ParseFS(generateEmails, path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s email templated, make sure you have generated email templates and they have valid syntax: %w", path, err)
	}
	return templ, nil
}
