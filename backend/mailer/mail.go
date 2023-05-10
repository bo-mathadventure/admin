package mailer

import (
	"bytes"
	"errors"
	"github.com/bo-mathadventure/admin/config"
	gomail "gopkg.in/mail.v2"
	hTemplate "html/template"
	"net/mail"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	tTemplate "text/template"
)

var cfg *config.Config

// MailData type for data of a mail
type MailData map[string]interface{}

// Init create mail instance
func Init(config config.Config) {
	cfg = &config
}

// Send the email
func Send(templatePath string, to string, data MailData) error {
	if cfg.MailHost == "" {
		return errors.New("SMTP server config is empty")
	}
	if cfg.MailPort == 0 {
		return errors.New("SMTP port config is empty")
	}

	if cfg.MailUser == "" {
		return errors.New("SMTP user is empty")
	}

	if cfg.MailFrom == "" {
		return errors.New("SMTP sender email is empty")
	}

	if to == "" {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    cfg.AppName,
		Address: cfg.MailFrom,
	}

	templateRaw, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}
	templateRawString := string(templateRaw)

	tmpl, err := hTemplate.ParseFiles(filepath.Join(filepath.Dir(templatePath), "base.gohtml"), templatePath)
	if err != nil {
		return err
	}

	subjectRegex := regexp.MustCompile(`\[(.*?)\]`)

	tmplSubject, err := tTemplate.New("subject").Parse(strings.Trim(subjectRegex.FindString(templateRawString), "[]"))
	if err != nil {
		return err
	}

	var tmplOutput bytes.Buffer
	if err := tmpl.Execute(&tmplOutput, data); err != nil {
		return err
	}

	// Subject requires {{/* [Test Subject {{.test}}] */}} in first line!
	var tmplSubjectOutput bytes.Buffer
	if err := tmplSubject.Execute(&tmplSubjectOutput, data); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", to)
	m.SetHeader("Subject", tmplSubjectOutput.String())

	m.SetBody("text/html", tmplOutput.String())

	d := gomail.NewDialer(cfg.MailHost, cfg.MailPort, cfg.MailUser, cfg.MailPassword)

	return d.DialAndSend(m)
}
