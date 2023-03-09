package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if (t == JSONTime{}) {
		formatted := fmt.Sprintf("\"%s\"", "")
		return []byte(formatted), nil
	} else {
		formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
		return []byte(formatted), nil
	}
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	ti, err := time.Parse("2006-01-02 15:04:05", value) //parse time
	if err != nil {
		return err
	}
	*t = JSONTime{Time: ti} //set result using the pointer
	return nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type JsonNumberList []uint64

func (c *JsonNumberList) Value() (driver.Value, error) {
	b, err := json.Marshal(*c)
	return string(b), err
}

func (c JsonNumberList) Sort() JsonNumberList {
	sort.Slice(c, func(i, j int) bool {
		return (c)[i] < (c)[j]
	})
	return c
}

func (c *JsonNumberList) Scan(input interface{}) error {
	bs := input.([]byte)
	if len(bs) == 0 {
		return nil
	}
	if bs[0] != '[' {
		for _, a := range strings.Split(string(bs), ",") {
			v, _ := strconv.ParseUint(a, 10, 64)
			*c = append(*c, v)
		}
		return nil
	}
	return json.Unmarshal(bs, c)
}

func DistinctJsonNumberLists(lst ...JsonNumberList) (ret []uint64) {
	var all []uint64
	for _, l := range lst {
		all = append(all, l...)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})

	for i, n := range all {
		if i == 0 || n != all[i-1] {
			ret = append(ret, n)
		}
	}
	return
}
