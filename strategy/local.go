package strategy

import (
	"io"
	"os"
	"path/filepath"
)

// 文件本地到保存本地
type LocalStrategy struct {
	Dir  string
	Name string
	Ext  string
}

// result_out *string
func (my *LocalStrategy) Save(reader_in io.Reader, result_out any) error {
	result := result_out.(*string)
	var err error

	if err = os.MkdirAll(my.Dir, 0750); err != nil {
		return err
	}

	name := filepath.Join(my.Dir, my.Name+my.Ext)

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	f.ReadFrom(reader_in)

	*result = name
	return nil
}
