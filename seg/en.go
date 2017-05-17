package seg

import (
	"strings"
)

type EnTokenizer struct {
}

func (s *EnTokenizer) Seg(text string) []string {
	return strings.Fields(text)
}

func (s *EnTokenizer) Free() {

}
