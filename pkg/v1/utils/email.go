package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/k3a/html2text"
	gomail "gopkg.in/gomail.v2"
)

type VerificationEmailData struct {
	Url     string
	Name    string
	Subject string
}

func SendVerificationEmail(email string, data *VerificationEmailData) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading content from env file")
	}

	smtp_user := os.Getenv("SMTP_USER")
	smtp_pass := os.Getenv("SMTP_PASS")
	smtp_host := os.Getenv("SMTP_HOST")
	smtp_port := os.Getenv("SMTP_PORT")
	from := os.Getenv("EMAIL_FROM")
	to := email

	// Define email template data
	var body bytes.Buffer

	template, err := ParseTemplateDir("pkg/v1/templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	num, _ := strconv.Atoi(smtp_port)
	fmt.Println(smtp_host, smtp_pass, smtp_port, smtp_user, num)
	d := gomail.NewDialer(smtp_host, num, smtp_user, smtp_pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		// log.Fatal("Could not send email: ", err)
		return err
	}
	return err
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
