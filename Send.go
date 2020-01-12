package mailer

import (
	"crypto/tls"
	"github.com/dkuye/helper"
	"github.com/kami-zh/go-capturer"
	"github.com/kataras/iris"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
)

type Template struct {
	Path    string
	Layout  string
	File    string
	Email   string
	Subject string
	Data    map[string]string
}

/*
Mailer Send user ./emailTemplates Folder for any project

*/
func Send(templateData Template) error {
	// Build content
	app := iris.New()
	temp := iris.HTML(templateData.Path, ".html")

	app.RegisterView(temp) //app.RegisterView(iris.HTML("./public/templates/emails", ".html"))

	//template functions. HtmlDisplay() display html content in email templates
	temp.AddFunc("HtmlDisplay", func(str string) template.HTML {
		return template.HTML(str)
	})

	content := capturer.CaptureStdout(func() {
		_ = app.Build()
		writer := os.Stdout
		_ = app.View(writer, templateData.File, templateData.Layout, templateData.Data)
	})

	/// Sender
	host := os.Getenv("MAIL_HOST")
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	port := helper.StringToInt(os.Getenv("MAIL_PORT"))
	fromAddr := os.Getenv("MAIL_FROM_ADDRESS")
	fromAlias := os.Getenv("MAIL_FROM_ALIAS")

	m := gomail.NewMessage()
	m.SetAddressHeader("From", fromAddr, fromAlias)
	m.SetHeader("To", templateData.Email)
	m.SetHeader("Subject", templateData.Subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	/*if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}*/

	err := d.DialAndSend(m)

	return err

}
