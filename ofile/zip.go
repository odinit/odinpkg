package ofile

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path"
	"path/filepath"
)

// Zip 压缩文件为zip文件
func Zip(w io.Writer, p string) (err error) {
	writer := zip.NewWriter(w)
	defer writer.Close()

	return zipWriter(p, path.Base(p), writer)
}

// Unzip 解压缩zip文件
func Unzip(zipFile, dstDir string) (err error) {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return
	}

	err = os.MkdirAll(dstDir, 0755)
	if err != nil {
		return
	}

	for fIndex := range reader.File {
		fPath := filepath.Join(dstDir, reader.File[fIndex].Name)
		if reader.File[fIndex].FileInfo().IsDir() {
			err = os.MkdirAll(fPath, reader.File[fIndex].Mode())
			if err != nil {
				return
			}
			continue
		}

		fReader, err := reader.File[fIndex].Open()
		if err != nil {
			return err
		}
		fDst, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, reader.File[fIndex].Mode())
		if err != nil {
			return err
		}
		if _, err := io.Copy(fDst, fReader); err != nil {
			return err
		}

		fReader.Close()
		fDst.Close()
	}

	return nil
}

// IsZip 判断文件是否为zip文件
func IsZip(p string) bool {
	f, err := os.Open(p)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

func zipWriter(p, pName string, writer *zip.Writer) (err error) {
	pStat, err := os.Stat(p)
	if err != nil {
		return
	}
	if pStat.IsDir() {
		_, _ = writer.Create(pName + "/")

		ps, err := os.ReadDir(p)
		if err != nil {
			return err
		}

		for pIndex := range ps {
			err = zipWriter(path.Join(p, ps[pIndex].Name()), path.Join(pName, ps[pIndex].Name()), writer)
			if err != nil {
				return err
			}
		}
	} else {
		f_, _ := os.Open(p)
		w_, _ := writer.Create(pName)
		_, _ = io.Copy(w_, f_)
		_ = f_.Close()
	}

	return nil
}
