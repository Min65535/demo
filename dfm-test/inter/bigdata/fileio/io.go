package fileio

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type FileIo struct {
	first bool
	fn    *os.File
	lock  sync.Mutex
}

func fileIsExistOrCreate(name string) (*os.File, bool, bool) {
	var firstCreate bool
	path, err := filepath.Abs(name)
	if err != nil {
		// fmt.Println("Abs err: ", err)
		return nil, firstCreate, false
	}

	// for unix
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return nil, firstCreate, false
	}

	_, err1 := os.Stat(name)
	if err1 != nil {
		if os.IsNotExist(err1) {
			f, err2 := os.Create(name)
			if err2 != nil {
				return nil, firstCreate, false
			} else {
				firstCreate = true
				return f, firstCreate, true
			}
		}
	} else {
		f, err3 := os.OpenFile(name, os.O_APPEND|os.O_RDWR, os.ModeAppend)
		if err3 != nil {
			return nil, firstCreate, false
		}
		return f, firstCreate, true
	}
	return nil, firstCreate, false
}

func NewFileIo(newFile string) *FileIo {
	fn, first, sign := fileIsExistOrCreate(newFile)
	if !sign {
		return nil
	}
	return &FileIo{
		first: first,
		fn:    fn,
	}
}

func (fi *FileIo) Close() {
	fi.fn.Close()
}

func (fi *FileIo) FileWrite(ll []byte) {
	fi.lock.Lock()
	defer fi.lock.Unlock()
	if fi.first {
		fi.fn.WriteString(fmt.Sprintf("%s", string(ll)))
		fi.first = false
	} else {
		fi.fn.WriteString(fmt.Sprintf("\n%s", string(ll)))
	}
}

func (fi *FileIo) FileLineReader() *bufio.Reader {
	return bufio.NewReader(fi.fn)
}
