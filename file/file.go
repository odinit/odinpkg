package file

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"io/fs"
	"os"
	"path"
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

	if pState.IsDir() {
		return countLinesInDir(p, suffix...)
	} else {
		return countLinesInFile(p)
	}
}

func countLinesInDir(p string, suffix ...string) (n int, err error) {
	if len(suffix) == 0 {
		err = filepath.WalkDir(p, func(p_ string, d fs.DirEntry, err error) error {
			// 跳过文件夹/跳过以.开头的文件/这里忽略后缀的限制
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
			// 跳过文件夹/跳过文件名后缀不符合要求的文件/跳过以.开头的文件
			if d.IsDir() || !slices.Contains(suffix, filepath.Ext(d.Name())) || strings.Index(d.Name(), ".") == 0 {
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

// ExtSum 获取给定路径下所有文件的扩展名列表
func ExtSum(ps ...string) (exts []string, err error) {
	if len(ps) == 0 {
		ps = []string{"."}
	}
	var ext string
	for _, p := range ps {
		pInfo, err := os.Stat(p)
		if err != nil {
			fmt.Printf("%s 路径出错: %s", p, err.Error())
			continue
		}
		if pInfo.IsDir() {
			_ = filepath.Walk(p, func(p_ string, info fs.FileInfo, err error) error {
				if info.IsDir() || strings.HasPrefix(info.Name(), ".") {
					return nil
				}
				ext = filepath.Ext(info.Name())
				if !slices.Contains(exts, ext) {
					exts = append(exts, ext)
				}
				return nil
			})
		} else {
			ext = filepath.Ext(path.Base(p))
			if !slices.Contains(exts, ext) {
				exts = append(exts, ext)
			}
		}
	}
	return
}

// DeleteByExt 通过扩展名删除给定路径下所有文件
func DeleteByExt(p string, ext ...string) (err error) {
	if len(ext) == 0 {
		err = errors.New("请输入要删除的文件扩展名")
		return
	}
	if p == "" {
		p = "."
	}

	pInfo, err := os.Stat(p)
	if err != nil {
		fmt.Printf("%s 路径出错: %s", p, err.Error())
		return
	}
	if pInfo.IsDir() {
		err = filepath.Walk(p, func(p_ string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if slices.Contains(ext, filepath.Ext(p_)) {
				return os.Remove(p_)
			}
			return nil
		})
	} else {
		if slices.Contains(ext, filepath.Ext(p)) {
			err = os.Remove(p)
		}
	}

	return
}

// FindByExt 通过扩展名查找给定路径下所有文件
func FindByExt(p string, ext ...string) (files []string, err error) {
	if len(ext) == 0 {
		err = errors.New("请输入要查找的文件扩展名")
		return
	}
	if p == "" {
		p = "."
	}

	pInfo, err := os.Stat(p)
	if err != nil {
		fmt.Printf("%s 路径出错: %s", p, err.Error())
		return
	}
	if pInfo.IsDir() {
		err = filepath.Walk(p, func(p_ string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if slices.Contains(ext, filepath.Ext(p_)) {
				files = append(files, p_)
			}
			return nil
		})
	} else {
		if slices.Contains(ext, filepath.Ext(p)) {
			files = append(files, p)
		}
	}

	return
}
