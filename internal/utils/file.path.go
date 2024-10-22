package utils

import (
	"os"
)

// 判断文件夹是否存在
func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MkDir(path string) error {
	if ok, err := IsExists(path); !ok {
		if err != nil {
			return err
		}
		// os.MkdirAll(dname,os.ModeDir|os.ModePerm)
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}

	}
	return nil
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// 区分目录和文件
// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(f string) bool {
	fi, err := os.Stat(f)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

// 判断文件/目录是否存在
// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
