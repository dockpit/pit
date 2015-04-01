package model

import (
	"bytes"
	"encoding/base64"
	"regexp"
	"strconv"
	"time"
)

// A JSON Serializable regexp pattern
type RegexpPattern regexp.Regexp

func (r RegexpPattern) MarshalJSON() ([]byte, error) {
	exp := regexp.Regexp(r)

	return []byte(strconv.Quote(exp.String())), nil
}

func (r *RegexpPattern) UnmarshalJSON(data []byte) error {
	patt, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	exp, err := regexp.Compile(patt)
	if err != nil {
		return err
	}

	*r = RegexpPattern(*exp)
	return nil
}

// A JSON Serializable time duration
type Duration time.Duration

func (d *Duration) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	nd, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	*d = Duration(nd)
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	duration := time.Duration(d)

	return []byte(strconv.Quote(duration.String())), nil
}

// A JSON Serializable output buffer
type Output struct {
	Buffer *bytes.Buffer
}

func NewOutput() *Output {
	return &Output{bytes.NewBuffer(nil)}
}

func (o *Output) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(base64.StdEncoding.EncodeToString(o.Buffer.Bytes()))), nil
}
