package myerrors

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

type error interface{
	Error() string
}

// 1.0 -- 原本error实现方式；只要其实现了error interface即可
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

/* 2.0 -- 优化：自定义一个错误类型 */
type MyError struct {
	ErrorMessage string
	ErrorLine    int
}

func (m *MyError) Error() string {
	return fmt.Sprintf("%s: %d", m.ErrorMessage, m.ErrorLine)
}

func Test() error {
	return &MyError{ErrorMessage: "test error", ErrorLine: 1}
}
// 自定义error类型的问题：如果想使用类型断言和switch，需要让error变为public类型，从而导致与调用者之间产生强耦合；
// 尽管自定义的error类型 比sentinel errors携带了更多信息，但强耦合还是 尽量避免使用自定义类型作为公共API的一部分；

/* 3.0 -- 使用 opaque errors 方式进行处理 */
// 只在意错误的行为 而非错误的具体值或类型
type opaqueError interface {
	error //opaqueError 需要实现error接口的基本方法
	Timeout() bool // if current error timeout?
	Temporary() bool // if current error an temporary err?
}

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	temp, ok := err.(temporary)
	return ok && temp.Temporary()
}

// example: count lines version-1.0
func CountLines(r io.Reader) (int, error){
	var (
		br = bufio.NewReader(r)
		lines int
		err error
	)
	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

// count lines version-2.0 直接使用scanner即可完成操作
func CountLines2(r io.Reader) (int, error){
	scanner := bufio.NewScanner(r)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

// wrap errors : both errors.Wrap (errors.Wrapf)
func ReadFile(path string)([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	// or I could also do:
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open: %v", path)
	}
	defer f.Close()
	return nil, nil
}

func ReadConfig()([]byte, error){
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home,".settings.xml"))
	return config, errors.WithMessage(err, "could not read config")
}

// construct a new error raised from the logic process
func parseArgs(args []string) error {
	if len(args) < 3 {
		return errors.New("the length of args is smaller than three")
	}
	if len(args) < 5 {
		return errors.Errorf("the length of args is smaller than five")
	}
	return nil
}

// if calling the method belongs to other pkg, just return the error directly


