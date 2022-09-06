package gfile

import (
	"bufio"
	"errors"
	"github.com/odinit/global/gtype"
	"golang.org/x/exp/slices"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CountLines 计算文件或文件夹中指定后缀文件的行数
// param
// p 文件或文件夹的路径
// suffix 指定后缀(文件类型)
// return
// n 行数
func CountLines(p string, suffix ...string) (n int, err error) {
	// 判断p是文件还是文件夹
	pState, err := os.Stat(p)
	if err != nil {
		return
	}

	suffixUpper := gtype.SliceStringToUpper(suffix)
	if pState.IsDir() {
		return countLinesInDir(p, suffixUpper...)
	} else {
		return countLinesInFile(p)
	}
}

func countLinesInDir(p string, suffix ...string) (n int, err error) {
	if len(suffix) == 0 {
		err = filepath.WalkDir(p, func(p_ string, d fs.DirEntry, err error) error {
			// 跳过文件夹/跳过文件名后缀不符合要求的文件/以.开头的文件
			if d.IsDir() || strings.Index(d.Name(), ".") == 0 {
				return nil
			}

			// 获取文件的行数
			n_, err := countLinesInFile(p_)
			if err != nil {
				return err
			}
			n += n_
			return nil
		})
	} else {
		err = filepath.WalkDir(p, func(p_ string, d fs.DirEntry, err error) error {
			// 跳过文件夹/跳过文件名后缀不符合要求的文件/以.开头的文件
			if d.IsDir() || !slices.Contains(suffix, strings.ToUpper(filepath.Ext(d.Name()))) || strings.Index(d.Name(), ".") == 0 {
				return nil
			}

			// 获取文件的行数
			n_, err := countLinesInFile(p_)
			if err != nil {
				return err
			}
			n += n_
			return nil
		})
	}

	//err = filepath.Walk(p, func(p_ string, info fs.FileInfo, err error) error {
	//	// 跳过文件夹/跳过文件名后缀不符合要求的文件
	//	if info.IsDir() || !slices.Contains(suffix, strings.ToUpper(filepath.Ext(info.Name()))) {
	//		return nil
	//	}
	//
	//	// 获取文件的行数
	//	n_, err := countLinesInFile(p_)
	//	if err != nil {
	//		return err
	//	}
	//	n += n_
	//	return nil
	//})
	return
}

func countLinesInFile(p string) (n int, err error) {
	// 打开文件
	file_, err := os.Open(p)
	if err != nil {
		return
	}
	defer file_.Close()

	// 遍历文件,\n截断
	fileBuf := bufio.NewReader(file_)
	for {
		_, err = fileBuf.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return n + 1, nil
			} else {
				return
			}
		} else {
			n++
		}
	}
}

func FileSizeFormat(size int64) (s_ string) {
	if s := float64(size) / 1024; s < 1 {
		s_ = strconv.FormatInt(size, 10) + " B"
	} else {
		if ss := s / 1024; ss < 1 {
			s_ = strconv.FormatFloat(s, 'f', 1, 32) + " K"
		} else {
			if sss := ss / 1024; sss < 1 {
				s_ = strconv.FormatFloat(ss, 'f', 1, 32) + " M"
			} else {
				if ssss := sss / 1024; ssss < 1 {
					s_ = strconv.FormatFloat(sss, 'f', 1, 32) + " G"
				} else {
					if sssss := ssss / 1024; sssss < 1 {
						s_ = strconv.FormatFloat(ssss, 'f', 1, 32) + " T"
					}
				}
			}
		}
	}

	return
}
