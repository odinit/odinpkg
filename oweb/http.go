package gweb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Http struct {
	IP     string // ip
	Port   int    // port
	Router string // 路由

	Client *http.Client // 客户端

	Request *http.Request     // 请求
	Method  string            // method
	URL     string            // url
	Body    io.Reader         // 请求body
	Header  map[string]string // 请求头

	Field map[string]interface{} // 字段信息,用于请求参数
	File  map[string]io.Reader   // 文件信息,用于发送文件formdata

	Response *http.Response // 请求返回
}

// Get
/*
执行Get请求
*/
func (h *Http) Get() error {
	if h.URL == "" {
		h.NewURL()
	}

	h.Method = "GET"

	if h.Field != nil {
		p := make([]string, 0)
		for k, v := range h.Field {
			if v_, ok := v.(string); ok {
				p = append(p, k+"="+v_)
				continue
			}
			if v_, ok := v.(int); ok {
				p = append(p, k+"="+strconv.Itoa(v_))
				continue
			}
		}
		if len(p) > 0 {
			h.URL += "?" + strings.Join(p, "&")
		}
		if strings.Contains(h.URL, " ") {
			return errors.New("参数中value值包含特殊字符,请使用post请求")
		}
	}

	h.Body = nil

	if err := h.NewRequest(); err != nil {
		return err
	}
	return h.Do()
}

// Post
/*
执行Post请求
*/
func (h *Http) Post() error {
	if h.URL == "" {
		h.NewURL()
	}

	h.Method = "POST"

	if h.Body == nil {
		if h.Field != nil {
			jsonB, err := json.Marshal(h.Field)
			if err != nil {
				return err
			}
			h.Body = bytes.NewReader(jsonB)
		}
	}

	h.Header["Content-Type"] = "application/json"
	if err := h.NewRequest(); err != nil {
		return err
	}
	h.SetHeader()

	return h.Do()
}

// FormData
/*
执行FormData请求
*/
func (h *Http) FormData() error {
	if h.URL == "" {
		h.NewURL()
	}

	h.Method = "POST"

	if h.Body == nil {
		buf := &bytes.Buffer{}
		writer := multipart.NewWriter(buf)

		if h.Field != nil {
			for k, v := range h.Field {
				fieldWriter, err := writer.CreateFormField(k)
				if err != nil {
					return err
				}
				if value, ok := v.(string); !ok {
					continue
				} else {
					_, err = fieldWriter.Write([]byte(value))
					if err != nil {
						return err
					}
				}
			}
			//for k, v := range h.Field {
			//	if value, ok := v.(string); !ok {
			//		continue
			//	} else {
			//		err := writer.WriteField(k, value)
			//		if err != nil {
			//			return err
			//		}
			//	}
			//}
		}
		if h.File != nil {
			for k, v := range h.File {
				fileWriter, err := writer.CreateFormFile("file", k)
				if err != nil {
					return err
				}
				fContent, err := io.ReadAll(v)
				if err != nil {
					return err
				}
				_, err = fileWriter.Write(fContent)
				if err != nil {
					return err
				}
			}
		}
		writer.Close()
		h.Body = buf
		h.Header["Content-Type"] = writer.FormDataContentType()
	}

	if err := h.NewRequest(); err != nil {
		return err
	}
	h.SetHeader()

	return h.Do()
}

// NewURL
/*
创建新的URL
*/
func (h *Http) NewURL() {
	if h.IP == "" {
		h.IP = "127.0.0.1"
	}
	if h.Port == 0 {
		h.Port = 80
	}
	if strings.Index(h.Router, "/") != 0 {
		h.Router = "/" + h.Router
	}

	h.URL = fmt.Sprintf("http://%s:%d%s", h.IP, h.Port, h.Router)

}

// NewBody
/*
创建请求body
*/
func (h *Http) NewBody() error {
	if h.Method == "" {
		return errors.New("请设置method")
	}

	switch strings.ToUpper(h.Method) {
	case "GET":
		h.Body = nil
	case "POST":

	}

	return nil
}

// NewRequest
/*
创建新的请求
*/
func (h *Http) NewRequest() error {
	req, err := http.NewRequest(h.Method, h.URL, h.Body)
	h.Request = req
	return err
}

// SetHeader
/*
设置Header
*/
func (h *Http) SetHeader() {
	if h.Header != nil {
		for k, v := range h.Header {
			h.Request.Header.Set(k, v)
		}
	}
}

// Do
/*
执行请求
*/
func (h *Http) Do() error {
	res, err := h.Client.Do(h.Request)
	h.Response = res
	return err
}

// NewClient
/*
创建新的客户端
*/
func (h *Http) NewClient(t ...int) {
	if len(t) == 0 {
		h.Client = http.DefaultClient
		return
	}

	h.Client = &http.Client{
		Timeout: time.Millisecond * time.Duration(t[0]),
	}
}

// NewHttp
/*
初始化一个新的Http工具
*/
func NewHttp(ip string, port int, router string) Http {
	http_ := Http{
		IP:     ip,
		Port:   port,
		Router: router,
		Header: map[string]string{},
		Client: http.DefaultClient,
	}
	http_.NewURL()
	return http_
}
