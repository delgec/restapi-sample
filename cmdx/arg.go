package cmdx

import (
	"bytes"
	"encoding/csv"
	"strconv"
)

type FlagStrArr []string

func (a *FlagStrArr) String() string {
	buf := bytes.NewBuffer(nil)
	if err := csv.NewWriter(buf).Write(*a); err != nil {
		panic(err)
	}

	return buf.String()
}

func (a *FlagStrArr) Set(value string) error {
	*a = append(*a, value)
	return nil
}

type FlagIntArr []int

func (a *FlagIntArr) String() string {
	buf := bytes.NewBuffer(nil)
	arr := make([]string, len(*a))
	for i, v := range *a {
		arr[i] = strconv.FormatInt(int64(v), 10)
	}

	if err := csv.NewWriter(buf).Write(arr); err != nil {
		panic(err)
	}

	return buf.String()
}

func (a *FlagIntArr) Set(value int) error {
	*a = append(*a, value)
	return nil
}

type FlagBoolArr []bool

func (a *FlagBoolArr) String() string {
	buf := bytes.NewBuffer(nil)
	arr := make([]string, len(*a))
	for i, v := range *a {
		arr[i] = strconv.FormatBool(v)
	}

	if err := csv.NewWriter(buf).Write(arr); err != nil {
		panic(err)
	}

	return buf.String()
}

func (a *FlagBoolArr) Set(value bool) error {
	*a = append(*a, value)
	return nil
}
