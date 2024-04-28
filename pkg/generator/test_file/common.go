/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 12:19:02
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-28 12:19:03
 * @FilePath: /dao-generator/pkg/generator/test_file/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package testfile

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

func value(i interface{}) (driver.Value, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func scan(i interface{}, src any) error {
	switch v := src.(type) {
	case string:
		return json.Unmarshal([]byte(v), i)
	case []byte:
		return json.Unmarshal(v, i)
	case driver.Null:
		return nil
	default:
		return fmt.Errorf("invalid-value[%v]", src)
	}
}
