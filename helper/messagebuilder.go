package helper

import (
	"strconv"
	"strings"
)

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		fields: make(map[string]string),
		separator:"~",
		headerLength: 4,
	}
}

type MessageBuilder struct {
	fields map[string]string
	separator string
	headerLength int
}

func(mb *MessageBuilder) SetSeparator(sep string){
	mb.separator = sep
}

func(mb *MessageBuilder) SetHeaderLengt(length int){
	mb.headerLength = length
}

func(mb *MessageBuilder) Add(key string, value string){
	mb.fields[strings.ToLower(key)] = value
}

func(mb *MessageBuilder) Compose() string{
	var buff = []string{}
	for key, value := range mb.fields {
		field := []string{key, "=", value}
		buff = append(buff, strings.Join(field, ""))
	}

	var msg string
	if len(buff) > 0 {
		msg = strings.Join(buff,mb.separator)
	}
	return msg
}

func(mb *MessageBuilder) Parse(msg string){
	flds := strings.Split(msg, mb.separator)
	for _, field := range flds{
		keyVal := strings.Split(field, "=")
		mb.fields[strings.ToLower(keyVal[0])]=keyVal[1]
	}
}

func(mb *MessageBuilder) Get(key string) string{
	return mb.fields[strings.ToLower(key)]
}

func(mb *MessageBuilder) SetHeader(msg string) string{
	n := len(msg)
	header := LeftPad(strconv.Itoa(n), mb.headerLength, "0")
	return strings.Join([]string{header, msg}, "")
}


