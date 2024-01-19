package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type FileLogger struct {
	file *os.File
	once *sync.Once
}

func NewFileLogger(path string) FileLogger {
	dir, err := os.MkdirTemp(path, "dir-log")
	if err != nil {
		panic(err)
	}
	file, fErr := os.CreateTemp(dir, "log")
	if fErr != nil {
		panic(fErr)
	}
	return FileLogger{
		file: file,
		once: &sync.Once{},
	}
}

func (f FileLogger) Log(s string) error {
	_, err := fmt.Fprintf(f.file, "%s - %s \n", time.Now(), s)
	return err
}

/***
@see :
Please note that if the once field in the FileLogger struct is not a pointer,
since the receivers are not of pointer type, when the close method is called twice in the main,
two new copies of FileLogger will be created, and a new copy of once will also be created.
In this case, the close method seems to be of no use and gets executed twice.
To address this issue, either the receivers should be pointers or the once field should be a pointer.
*/

func (f FileLogger) Close() (err error) {
	f.once.Do(func() {
		err = f.file.Close()
	})
	return
}

func main() {
	f := NewFileLogger("./05-once")
	err := f.Log("go")
	if err != nil {
		panic(err)
	}
	err = f.Log("rust")
	if err != nil {
		panic(err)
	}

	cErr := f.Close()
	if cErr != nil {
		panic(cErr)
	}

	cErr2 := f.Close()
	if cErr2 != nil {
		panic(cErr2)
	}
}
