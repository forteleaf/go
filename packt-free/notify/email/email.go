package email

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/JSila/packtFree/info"
	"github.com/JSila/packtFree/notify"
	"github.com/JSila/packtFree/notify/email/options"
	"github.com/kennygrant/sanitize"
	"github.com/mailgun/mailgun-go"
)

const (
	// Subject is email's Subject line
	Subject = "%s - free today from Packt Publishing"
	// EmailTmpl is path to email template
	EmailTmpl = "email.tmpl"
)

// Mail contains data for client initializing and holds email data
type Mail struct {
	data    *info.Info
	email   string
	options *options.Options
}

// Notify sends email to mailgun to specified user
func (m *Mail) Notify(to *notify.User) error {
	if err := m.LoadEmail(); err != nil {
		return err
	}
	gun := mailgun.NewMailgun(
		m.options.Domain,
		m.options.APIKey,
		m.options.PublicKey,
	)

	mg := mailgun.NewMessage(
		fmt.Sprintf("%s <%s>", m.options.From.Name, m.options.From.Email),
		fmt.Sprintf(Subject, m.data.Title),
		sanitize.HTML(m.email),
		fmt.Sprintf("%s <%s>", to.Name, to.Email),
	)
	mg.SetHtml(m.email)

	if _, _, err := gun.Send(mg); err != nil {
		return err
	}
	return nil
}

// LoadEmail loads email template and executes it
func (m *Mail) LoadEmail() error {
	funcMap := template.FuncMap{
		"formatTime": func(t time.Time) string {
			return fmt.Sprintf("%s %d, %d at %d:%d", t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
		},
	}
	t, err := template.New("email.tmpl").Funcs(funcMap).ParseFiles(EmailTmpl)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	if err = t.Execute(&b, m.data); err != nil {
		return err
	}
	m.email = b.String()
	return nil
}

// Name return email notifier's id
func (m *Mail) Name() string {
	return "email"
}

// New creates new Mail instance given info and mail's options
func New(data *info.Info, options *options.Options) (*Mail, error) {
	if options.APIKey == "" {
		return nil, errors.New("email.New: options not provided")
	}
	m := &Mail{
		data:    data,
		options: options,
	}
	return m, nil
}
