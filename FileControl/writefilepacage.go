package filecontrol

import (
	"errors"
	"fmt"
	"os"
)

/*
引数pathに保存する先のpathを入力

**param**

	path string //ファイルPath

	value []byte //保存する[]byteデータ

**returns**

	error
*/
func WriteFile(path string, value []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("file create error: %w", err)
	}
	defer f.Close()

	_, err = f.Write(value)
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}
	return nil
}

/*
structをjsonにコンパイルし引数で受け取ったpathに保存する

**parame**

	pats string //ファイルpath
*/
func SaveAsJson(path string, value any, properties ...string) error {
	//pathが指定されているか調べる
	if path == "" {
		//指定されていない
		return errors.New("[XoX] file path nil")
	}

	//json形式にコンパイルする
	json, err := CompileStructToJson(value, properties...)
	if err != nil {
		//コンパイルできなかった
		return err
	}

	err = WriteFile(path, json)
	if err != nil {
		return err
	}

	return nil
}
