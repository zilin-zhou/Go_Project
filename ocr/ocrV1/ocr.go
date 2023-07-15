package ocrV1

type icor interface {
	Ocr(input map[string]string) (string, error)
}

// 可以使用通用的执行接口口进行调用
func ExecIocr(ocr icor, input map[string]string) (string, error) {
	return ocr.Ocr(input)
}
