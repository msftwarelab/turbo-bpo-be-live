package strings

import "strings"

func TrimObjectChar(i string) string {
	i = strings.ReplaceAll(i, `ObjectID(`, ``)
	i = strings.ReplaceAll(i, `/`, ``)
	i = strings.ReplaceAll(i, `\`, ``)
	i = strings.ReplaceAll(i, `\\`, ``)
	i = strings.ReplaceAll(i, "\\", ``)
	i = strings.ReplaceAll(i, `)`, ``)
	i = strings.ReplaceAll(i, `"`, ``)
	return i
}

func ToObject(i string) *string {
	return &i
}

func ObjectTOString(i *string) string {
	if i != nil {
		return *i
	}
	return ""
}

func ObjectTOArrString(i ...*string) []string {
	var res []string
	for _, i := range i {
		if i != nil {
			res = append(res, *i)
		}
	}
	return res
}
