package filecontrol

//ファイルを読み取るパッケージ

import (
	"fmt"
	"io"
	"os"
)

/*
指定したpathのファイルを開くことができるか調べる

**param**

	path string // file path

**return**

	bool , os.FileInfo , error
*/
func FileExists(path string) (bool, os.FileInfo, error) {
	info, err := os.Stat(path)
	if err == nil {
		//ファイルを開れる
		return true, info, nil
	}

	//ファイルが無いことによるエラーなのか判定する
	if os.IsNotExist(err) {
		return false, nil, fmt.Errorf("[XoX] not find :%w", err)
	}

	//ファイルへのアクセス権限がない時
	if os.IsPermission(err) {
		return false, nil, fmt.Errorf("[XoX] not have permission :%w", err)
	}

	//その他のエラー
	return false, nil, fmt.Errorf("[XoX] stat error :%w", err)
}


/*
引数に渡したパスのテキストファイルを読み取る

**param**

	path string //ファイルPath

**returns**

	[]byte , error
*/
func ReadFile(path string) ([]byte, error) {
	//ファイルが存在しているか
	_, _, err := FileExists(path)
	if err != nil {
		return nil, err
	}

	// ファイルを開く
	file, err := os.Open(path)
	// エラー処理
	if err != nil {
		return nil, err
	}
	// このメソッドが終了するときにファイルを閉じる。<!>エラーを確認してから実行
	defer file.Close()
	// ファイルの内容を読み込む
	return io.ReadAll(file)
}

/*
ファイルのjsonデータを読み取って指定したstructに値を格納する

**param**

	path string //file path

	t *T //struct

**return**

	*T , error
*/
func ReadJsonFile[T any](path string, t T) (T, error) {
	json, err := ReadFile(path)
	if err != nil {
		return t, err
	}

	t, jsonDecordErr := DecordJson(json, t)
	if jsonDecordErr != nil {
		return t, err
	}
	return t, nil
}
