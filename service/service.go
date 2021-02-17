package service

import (
	"errors"
)

func new() {

}

func (m *Mail) Process(mailSchema map[string]interface{}) (err error) {
	var mailContent string
	if subject, ok := mailSchema["subject"].(string); ok {
		mailContent += m.Subject(subject)
	}
	for key, value := range mailSchema {
		switch key {
		case "text":
			if text, ok := value.(string); ok {
				mailContent += text
			}
		case "attachment":
			if attach, ok := value.([]string); ok {
				content, err := m.AttachFile(attach[0], attach[1])
				if err != nil {
					return errors.New("Package service Error: Attachment not found")
				}
				mailContent += content
			}
		case "template":
			if template, ok := value.(string); ok {
				content, err := m.Phrase(template)
				if err != nil {
					return errors.New("Package service Error: Cannot phrase text")
				}
				mailContent += content
			}
		default:
			return errors.New("Package service Error: Unkown type")
		}
	}
	if to, ok := mailSchema["to"].([]string); ok {
		m.Send(to, []byte(mailContent))
	} else {
		return errors.New("Package service Error: To list not found")
	}
	return nil
}
