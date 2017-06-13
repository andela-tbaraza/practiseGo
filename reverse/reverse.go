package reverse

func Reverse(str string) string  {
	newString := ""
	length := len(str) - 1
    for i := length; i >= 0; i-- {
        newString += string(str[i])
    }
	return newString
}
