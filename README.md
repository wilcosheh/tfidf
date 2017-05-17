# TFIDF

## Introduction

+ tokenizer support, contains english and jieba Chinese Tokenizer.
+ TFIDF, calculate tfidf value of giving document.
+ Cosine, calculate Cosine value of giving documents pair.
+ glide is used to manage go packages.

## Guide

```
go get github.com/yanyiwu/gojieba
glide i
```


```
package main

import (
	"fmt"
	"strings"

	"github.com/yanyiwu/gojieba"
)

func main() {
	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

}
```

```

```

See example in [jieba_test](jieba_test.go), [extractor_test](extractor_test.go)
