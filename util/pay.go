package util

import(
	"sync"
	"encoding/xml"
	"encoding/json"
	"io"
	"strings"
	"sort"
	"errors"
	// "fmt"
)

const (
	NULL       = ""
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
	SUCCESS    = "SUCCESS"
	FAIL       = "FAIL"
	OK         = "OK"
	Version    = "1.5.9"
)

type BodyMap map[string]interface{}

var mu sync.RWMutex

// 设置参数
func (bm BodyMap) Set(key string, value interface{}) {
	mu.Lock()
	bm[key] = value
	mu.Unlock()
}

// 获取参数
func (bm BodyMap) Get(key string) string {
	if bm == nil {
		return NULL
	}
	mu.RLock()
	defer mu.RUnlock()
	value, ok := bm[key]
	if !ok {
		return NULL
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return NULL
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return NULL
	}
	str = string(bs)
	return
}

// 删除参数
func (bm BodyMap) Remove(key string) {
	mu.Lock()
	delete(bm, key)
	mu.Unlock()
}

type xmlMapMarshal struct {
	XMLName xml.Name
	Value   interface{} `xml:",cdata"`
}

type xmlMapUnmarshal struct {
	XMLName xml.Name
	Value   string `xml:",cdata"`
}

func (bm BodyMap) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if len(bm) == 0 {
		return nil
	}
	start.Name = xml.Name{NULL, "xml"}
	if err = e.EncodeToken(start); err != nil {
		return
	}
	for k := range bm {
		if v := bm.Get(k); v != NULL {
			e.Encode(xmlMapMarshal{XMLName: xml.Name{Local: k}, Value: v})
		}
	}
	return e.EncodeToken(start.End())
}

func (bm *BodyMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for {
		var e xmlMapUnmarshal
		err = d.Decode(&e)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		bm.Set(e.XMLName.Local, e.Value)
	}
}

// ("bar=baz&foo=quux") sorted by key.
func (bm BodyMap) EncodeWeChatSignParams(apiKey string) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	mu.RLock()
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	mu.RUnlock()
	for _, k := range keyList {
		if v := bm.Get(k); v != NULL {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	buf.WriteString("key")
	buf.WriteByte('=')
	buf.WriteString(apiKey)
	return buf.String()
}

// ("bar=baz&foo=quux") sorted by key.
func (bm BodyMap) EncodeAliPaySignParams() string {
	var (
		buf     strings.Builder
		keyList []string
	)
	mu.RLock()
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	mu.RUnlock()
	for _, k := range keyList {
		if v := bm.Get(k); v != NULL {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return NULL
	}
	return buf.String()[:buf.Len()-1]
}

func (bm BodyMap) EncodeGetParams() string {
	var (
		buf strings.Builder
	)
	for k, _ := range bm {
		if v := bm.Get(k); v != NULL {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return NULL
	}
	return buf.String()[:buf.Len()-1]
}

func (bm BodyMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.Get(k); v == NULL {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return errors.New(strings.Join(emptyKeys, ", ") + " : cannot be empty")
	}
	return nil
}

func weixin(){
	// url := "https://api.mch.weixin.qq.com/pay/unifiedorder"

}