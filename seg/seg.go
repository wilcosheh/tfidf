package seg

// func Seg(text string) []string {
// 	return SegJieba(text)
// }

// func SegJieba(text string) []string {
// 	return jieba.Seg(text)
// }

type Tokenizer interface {
	Seg(text string) []string
	Free()
}
