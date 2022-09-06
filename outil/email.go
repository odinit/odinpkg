package outil

import (
	"errors"
	"github.com/go-gomail/gomail"
)

const (
	smtp163server = "smtp.163.com"
	smtp163port   = 25
)

func EmailNew() *Email {
	e := &Email{}
	e.Message = gomail.NewMessage()
	return e
}

func EmailNew163() *Email {
	e := &Email{
		Host: smtp163server,
		Port: smtp163port,
	}
	e.Message = gomail.NewMessage()
	return e
}

type Email struct {
	Dialer   *gomail.Dialer
	Host     string
	Port     int
	From     string
	Name     string
	Password string
	To       []string

	Message       *gomail.Message
	Header        map[string][]string
	AddressHeader map[string][]string
	ContentType   string
	Body          string

	isChecked bool
	isOK      bool
}

func (e *Email) Send() (err error) {
	err = e.Prepare()
	return e.Dialer.DialAndSend(e.Message)
}

func (e *Email) Prepare() (err error) {
	if e.Dialer == nil {
		err = e.NewDialer()
		if err != nil {
			return
		}
	}
	if e.Message == nil {
		e.NewMessage()
		e.SetHeader()
		e.SetAddressHeader()
		e.SetBody()
	}
	return nil
}

func (e *Email) NewDialer() (err error) {
	err = e.Check()
	if err != nil {
		return
	}
	e.Dialer = gomail.NewDialer(e.Host, e.Port, e.From, e.Password)
	return nil
}

func (e *Email) Check() error {
	if e.Host == "" {
		e.Host = "127.0.0.1"
	}
	if e.Port == 0 {
		e.Port = 465
	}
	if e.From == "" {
		return errors.New("请设置发件邮箱")
	}
	if e.Password == "" {
		return errors.New("请设置发件邮箱密码")
	}
	if len(e.To) == 0 {
		return errors.New("请设置收件邮箱")
	}
	return nil
}

func (e *Email) NewMessage() {
	e.Message = gomail.NewMessage()
}

func (e *Email) SetHeader() {
	if e.Header != nil {
		e.Message.SetHeaders(e.Header)
	}
}

func (e *Email) SetAddressHeader() {
	if e.AddressHeader != nil {
		for k, v := range e.AddressHeader {
			switch len(v) {
			case 0:
				continue
			case 1:
				e.Message.SetAddressHeader(k, v[0], "")
			default:
				e.Message.SetAddressHeader(k, v[0], v[1])
			}
		}
	}
	e.Message.SetAddressHeader("From", e.From, e.Name)
}

func (e *Email) SetBody() {
	if e.Body != "" {
		if e.ContentType == "" {
			e.ContentType = "text/html"
		}
		e.Message.SetBody(e.ContentType, e.Body)
	}
}
