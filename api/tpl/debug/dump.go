package debug

type dump struct {
	options string
	v       interface{}
}

func Dump(i interface{}) string {
	return newFormatState(i).Get()
}

func DumpHTML(i interface{}) string {
	return newFormatState(i).Get()
}
