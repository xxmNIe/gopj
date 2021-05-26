package server

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/disk"
)

const timeFormat = "2006-01-02 15:04:05"

func nowFormat() string {
	return time.Now().Format(timeFormat)
}

func makeFilePart(name, part string) string {
	return fmt.Sprintf("%s.part%s", name, part)
}

// 文件 md5 值
func fileMD5(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", errors.New(fmt.Sprintf("%s is a dir", filename))
	}

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func splitPath(dir string) []string {
	paths := strings.Split(dir, "/")
	l := []string{}
	for _, v := range paths {
		if v != "" {
			l = append(l, v)
		}
	}
	return l
}

func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}

// total, used , usedPercent
func diskUsed() (uint64, uint64, float64) {
	mDisk, _ := disk.Usage("/")
	return mDisk.Total, mDisk.Total - mDisk.Free, float64(mDisk.Total-mDisk.Free) * 100 / float64(mDisk.Total)
}

func WriteFile(filename string, reader io.Reader) (written int64, err error) {
	newFile, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer newFile.Close()

	return io.Copy(newFile, reader)
}

func CopyFile(src, dest string) (written int64, err error) {
	srcF, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcF.Close()

	return WriteFile(dest, srcF)
}
