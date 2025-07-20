package filecontrol

//Json形式にするPacage

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	/*
	   Jsonエンコードで使えるPrefixesで使えるList
	*/
	allowedPrefixes = []string{"", " ", "// "}

	/*
	   Jsonエンコードで使えるPrefixesで使えるList
	*/
	allowedIndents = []string{"", " ", "  ", "\t"}
	/*
		jsonエンコードで使えるPrefixesの初期値
	*/
	prefix = ""
	/*
		jsonエンコードで使えるindentの初期値
	*/
	indent = "  "
)

/*
structをJson形式にコンパイルする
フィールド名の戦闘は大文字にする必要があります
送られてきたstructをjsonファイルにコンパイルしファイルに保存する

**param**

	value any //struct型を受け取ります

	properties ...[]string

//json.MarshalIndentの第2、第3引数を決める、index0:第2引数、index1:第３引数

//propertiesに値を渡す場合は[変数名...]にする必要があります

**properties whitelist**

//properties[0] -> ""," ","// "

//properties[1] -> ""," ","  ","\t"

**return**

	[]byte , error
*/
func CompileStructToJson(value any, properties ...string) ([]byte, error) {
	//valueがnilなら弾く
	if value == nil {
		return nil, errors.New("[XoX] struct is nil")
	}

	//propertiesに値が入っている場合その値が使用できるか調べる
	if properties != nil {
		//errorを返す変数
		var err error
		//prefixのプロパティ値が使えるか調べる
		fmt.Println(properties[0])
		prefix, err = isAllowedIndent(properties[0], allowedPrefixes)
		if err != nil {
			return nil, fmt.Errorf("[XoX] invalid value provided for the prefix parameter :%w", err)
		}

		//indentのプロパティ値が使えるか調べる
		indent, err = isAllowedIndent(properties[1], allowedIndents)
		if err != nil {
			return nil, fmt.Errorf("[XoX] invalid value provided for the indent parameter :%w", err)
		}
	}

	//受け取ったstruct型をjsonにエコードする
	bytes, err := json.MarshalIndent(value, prefix, indent)
	if err != nil {
		return nil, fmt.Errorf("[XoX] an error occuurred during json conversion: %w", err)
	}

	return bytes, nil
}

/*
AllowedIndentに入れることのできる値かどうか確認する

**param**

	s stirng //使用するプロパティ

properties []string //使用できるプロパティ配列

**return**

	string , error
*/
func isAllowedIndent(s string, properties []string) (string, error) {
	for _, v := range properties {
		if s == v {
			return v, nil
		}
	}

	return "", fmt.Errorf(" [%s]", s)
}
