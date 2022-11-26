package fileuploader

import "io"

type IStrategy interface {
	Save(reader_in io.Reader, result_out any) error
}

type FileUploader struct {
	strategy IStrategy
}

func (my *FileUploader) SetStrategy(strategy IStrategy) *FileUploader {
	my.strategy = strategy
	return my
}

func (my *FileUploader) Save(reader_in io.Reader, result_out any) error {
	return my.strategy.Save(reader_in, result_out)
}
