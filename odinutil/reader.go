package odinutil

import (
	"bytes"
	"encoding/json"
	"io"
)

// ReaderUnmarshal
// 读取reader中的内容,并使用json解析
func ReaderUnmarshal(reader io.Reader, info interface{}) error {
	bys, err := io.ReadAll(reader)
	if err != nil {
		return nil
	}
	return json.Unmarshal(bys, info)
}

// ReaderMarshal
/*
使用json编码，并转换成reader
*/
func ReaderMarshal(src interface{}) (io.Reader, error) {
	buf, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}
