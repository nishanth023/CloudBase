package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"text/template"
)

type Mail struct {
	emailHost     string
	emailFrom     string
	emailPassword string
	emailPort     int32
	emailAuth     smtp.Auth
}

func (m *Mail) Load() (err error) {

}

func (m *Mail) Auth() {
	m.emailAuth = smtp.PlainAuth("", m.emailFrom, m.emailPassword, m.emailHost)
}

func (m *Mail) Send(to []string, msg []byte) (err error) {
	return smtp.SendMail(m.emailHost+":"+string(m.emailPort), m.emailAuth, m.emailFrom, to, msg)
}

func (m *Mail) Phrase(templateStruct interface{}, templateData string) (content string, err error) {
	t, err := template.New("text").Parse(templateData)
	var doc bytes.Buffer
	err = t.Execute(&doc, templateStruct)
	content = doc.String()
	return
}
func (m *Mail) Subject(subject string) string {
	return fmt.Sprintf("subject:%s\n", subject)
}
func (m *Mail) AttachFile(fileName string, filePath string) (attachment string, err error) {
	mime := "MIME-version: 1.0;" +
		"\nContent-Type: text/plain;" +
		"charset=\"UTF-8\";" +
		"Content-Transfer-Encoding: base64\r\n" +
		"Content-Disposition: attachment;filename=\"" +
		fileName + "\"\r\n" +
		"\n\n"

	rawFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	attachment = mime + "\r\n" + base64.StdEncoding.EncodeToString(rawFile)
	return
}
