package seg

import "github.com/yanyiwu/gojieba"

type JiebaTokenizer struct {
	x *gojieba.Jieba
}

func NewJieba() *JiebaTokenizer {
	return &JiebaTokenizer{
		x: gojieba.NewJieba(),
	}
}

func (j *JiebaTokenizer) Seg(text string) []string {
	// x := gojieba.NewJieba()
	// defer x.Free()
	// fmt.Println(x.ExtractWithWeight(text, 5))
	// return x.Cut(text, true)
	return j.x.Cut(text, true)

}

func (j *JiebaTokenizer) Free() {
	if j.x != nil {
		j.x.Free()
	}
}
