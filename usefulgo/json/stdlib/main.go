package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

type MyTime time.Time

func (t *MyTime) MarshalJSON() ([]byte, error) {
	ts := time.Now().UnixNano()
	value := float64(ts) * 1e-9
	return []byte(strconv.FormatFloat(value, 'f', 10, 64)), nil
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	sec := int64(value)
	nsec := int64((value - float64(sec)) * 1e9)
	*t = MyTime(time.Unix(sec, nsec))
	return nil
}

func (t MyTime) String() string {
	_t := time.Time(t)
	return _t.String()
}

// func (t *MyTime) String() string {
// 	_t := time.Time(*t)
// 	return _t.String()
// }

type T struct {
	Int        int64     `json:"int,omitempty"`
	String     string    `json:"string,omitempty"`
	Float0     float64   `json:"float0,omitempty"`
	Float1     float64   `json:"float1,string,omitempty"`
	Bool0      bool      `json:"bool0,omitempty"`
	Bool1      bool      `json:"bool1"`
	unexported int64     `json:"unexported,omitempty"`
	Time0      time.Time `json:"time0",omitempty`
	Time1      MyTime    `json:"time1",omitempty`
	NoTag      int
}

var jsondata = `{
	"int": 9223372036854775807,
	"string": "golang",
	"float0": 900000500000010002345600231,
	"float1": "900000500000010002345600231",
	"time0": "2018-09-19T23:57:00.985061+08:00",
	"time1": 1537372620.9853050709,
	"interface": 900000500000010002345600231,
	"NoTag": 1
  }`

func main() {
	foo()
	bar()
}

func foo() {
	fmt.Println("foo...")
	t := T{
		Int:        math.MaxInt64,
		String:     "golang",
		Float0:     900000500000010002345600231,
		Float1:     900000500000010002345600231,
		Time0:      time.Now(),
		Time1:      MyTime(time.Now()),
		Bool0:      true,
		unexported: 1,
		NoTag:      1,
	}
	data, err := json.MarshalIndent(&t, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")
	encoder.Encode(&t)
	fmt.Println(buf.String())
}

func bar() {
	fmt.Println("bar...")
	data := []byte(jsondata)
	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		panic(err)
	}
	fmt.Println(t)

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&t); err != nil {
		panic(err)
	}
	fmt.Println(t)
}
