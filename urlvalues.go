package urlvalues

import (
	"net/url"
	"strconv"
)

func ParseQuery(query string) (*Values, error) {
	v, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	return &Values{Values: v}, nil
}

func New(v url.Values) *Values {
	return &Values{Values: v}
}

func FromURL(u url.URL) *Values {
	return &Values{Values: u.Query()}
}

type Values struct {
	url.Values
}

func (v *Values) Has(key string) bool {
	if v.Values == nil {
		return false
	}
	_, ok := v.Values[key]
	return ok
}

func (v *Values) SetString(key string, value string) {
	v.Set(key, value)
}

func (v *Values) SetBool(key string, value bool) {
	v.Set(key, strconv.FormatBool(value))
}

func (v *Values) SetInt(key string, value int) {
	v.Set(key, strconv.Itoa(value))
}

func (v *Values) SetInt32(key string, value int32) {
	v.Set(key, strconv.FormatInt(int64(value), 10))
}

func (v *Values) SetInt64(key string, value int64) {
	v.Set(key, strconv.FormatInt(value, 10))
}

func (v *Values) SetUint(key string, value uint) {
	v.Set(key, strconv.FormatUint(uint64(value), 10))
}

func (v *Values) SetUint32(key string, value uint32) {
	v.Set(key, strconv.FormatUint(uint64(value), 10))
}

func (v *Values) SetUint64(key string, value uint64) {
	v.Set(key, strconv.FormatUint(value, 10))
}

func (v *Values) SetFloat32(key string, value float32) {
	v.Set(key, strconv.FormatFloat(float64(value), 'f', 0, 32))
}

func (v *Values) SetFloat64(key string, value float64) {
	v.Set(key, strconv.FormatFloat(value, 'f', 0, 64))
}

func (v *Values) GetString(key string) string {
	return v.Get(key)
}

func (v *Values) GetBool(key string) (bool, error) {
	return strconv.ParseBool(v.Get(key))
}

func (v *Values) GetInt(key string) (int, error) {
	return strconv.Atoi(v.Get(key))
}

func (v *Values) GetInt32(key string) (int32, error) {
	value, err := strconv.ParseInt(v.Get(key), 10, 32)
	return int32(value), err
}

func (v *Values) GetInt64(key string) (int64, error) {
	return strconv.ParseInt(v.Get(key), 10, 64)
}

func (v *Values) GetUint(key string) (uint, error) {
	value, err := strconv.ParseUint(v.Get(key), 10, 0)
	return uint(value), err
}

func (v *Values) GetUint32(key string) (uint32, error) {
	value, err := strconv.ParseUint(v.Get(key), 10, 32)
	return uint32(value), err
}

func (v *Values) GetUint64(key string) (uint64, error) {
	return strconv.ParseUint(v.Get(key), 10, 64)
}

func (v *Values) GetFloat64(key string) (float64, error) {
	return strconv.ParseFloat(v.Get(key), 64)
}
