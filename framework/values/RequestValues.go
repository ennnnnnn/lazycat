package values

import (
	"strconv"
	"strings"
	"errors"
	"net/url"
	"regexp"
)

type RequestValues struct {
	values *url.Values
}

func (this *RequestValues) Get(name string) RequestValue {
	return RequestValue{parent: this, name: name}
}

type RequestValue struct {
	parent *RequestValues
	name   string
}

/*
 * 这个值是否存在
 */
func (this *RequestValue) IsNull() bool {
	_, ok := (*(*this.parent).values)[this.name]
	return ok
}

/*
 * 这个值是否为空
 */
func (this *RequestValue) IsEmpty() bool {
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		return len(value) == 0 || len(strings.TrimSpace(value[0])) == 0
	}
	return true
}

/*
 * 获取第一个字符
 */
func (this *RequestValue) StringOutError() (string, error) {
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		return value[0], nil
	}
	return "", errors.New("值不纯在")
}

/*
 * 获取第一个字符
 */
func (this *RequestValue) String() string {
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		return value[0]
	}
	return ""
}

/*
 * 获取所以值
 */
func (this *RequestValue) Strings() []string {
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		return value
	}
	return []string{}
}

/*
 * 不存在使用默认值
 */
func (this *RequestValue) NullDefault(v string) string {
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		return value[0]
	}
	return v
}

/*
 * 值为空使用默认值
 */
func (this *RequestValue) EmptyDefault(v string) string {
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		if len(value) == 0 || len(strings.TrimSpace(value[0])) == 0 {
			return v
		}
		return value[0]
	}
	return v
}

/*
 * 获取64整形,输出错误
 */
func (this *RequestValue) Int64OutError() (int64, error) {
	var tmp string
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		if len(value) == 0 || len(strings.TrimSpace(value[0])) == 0 {
			return 0, errors.New("参数为空")
		}
		tmp = value[0]
	}
	vt, err := strconv.ParseInt(tmp, 10, 64)
	if err == nil {
		return vt, nil
	} else {
		return 0, errors.New("不是整型参数")
	}
}

/*
 * 获取64整形，使用默认值
 */
func (this *RequestValue) Int64Default(v int64) int64 {
	tmp, err := this.Int64OutError()
	if err != nil {
		return v
	}
	return tmp
}

/*
 * 获取64整形，默认为-1
 */
func (this *RequestValue) Int64() int64 {
	tmp, err := this.Int64OutError()
	if err != nil {
		return -1
	}
	return tmp
}

/*
 * 输出64位整形组，错误输出
 */
type Int64sValue struct {
	Value int64
	Error error
}

func (this *RequestValue) Int64sOutError() ([]Int64sValue, error) {
	var tmp []string
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		if len(value) == 0 || len(strings.TrimSpace(value[0])) == 0 {
			return []Int64sValue{}, errors.New("值为空")
		}
		tmp = value
	}
	var rtmp []Int64sValue
	haserr := false
	for _, v := range tmp {
		vt, err := strconv.ParseInt(v, 10, 64)
		rtmp = append(rtmp, Int64sValue{Value: vt, Error: err})
		if err != nil {
			haserr = true
		}
	}
	if haserr {
		return rtmp, errors.New("存在错误值")
	} else {
		return rtmp, nil
	}
}

/*
 * 输出32位整形，错误输出
 */
func (this *RequestValue) IntOutError() (int, error) {

	var tmp string
	if value, ok := (*(*this.parent).values)[this.name]; ok {
		if len(value) == 0 || len(strings.TrimSpace(value[0])) == 0 {
			return 0, errors.New("参数为空")
		}
		tmp = value[0]
	}
	vt, err := strconv.Atoi(tmp)
	if err == nil {
		return vt, nil
	} else {
		return 0, errors.New("不是整型参数")
	}
}

/*
 * 获取32整形，使用默认值
 */
func (this *RequestValue) IntDefault(v int) int {
	tmp, err := this.IntOutError()
	if err != nil {
		return v
	}
	return tmp
}

/*
 * 获取32整形，默认为-1
 */
func (this *RequestValue) Int() int {
	tmp, err := this.IntOutError()
	if err != nil {
		return -1
	}
	return tmp
}

/*
 * 匹配一个值
 */
func (this *RequestValue) Match(str string) bool {
	tmp, err := this.StringOutError()
	if err != nil {
		return false
	}
	v, err := regexp.MatchString(str, tmp)
	if err != nil {
		return false
	}
	return v
}
