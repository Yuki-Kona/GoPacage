package filecontrol

import (
	"encoding/json"
	"errors"
	"fmt"
)

/*
[]byteのjsonデータをstructに入れる
送られてきた[]byteデータがjson形式かどうか調べる
[]byteのデータをstructに格納する

**pamar**

	data []byte //Jsonデータを入れる

	value *any //jsonをデコードする先のstruct

**return**

	t any,error
*/
func DecordJson[T any](data []byte, t *T)(*T, error) {
	//送られてきたjson形式が正しいかどうか
	if !json.Valid(data) {
		return t,errors.New("[XoX] invalid fson format received")
	}

	//渡された[]byteのjsonデータをstructにデータ入れる
	err := json.Unmarshal(data, t)
	if err != nil {
		return t,fmt.Errorf("[XoX] json data could not be decoded:%w", err)
	}

	//成功した時
	return t,nil
}

