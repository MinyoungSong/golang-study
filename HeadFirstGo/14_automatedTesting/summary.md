# 14. Automated Testing


## Testing 코드 작성
{파일명}_test.go 생성  
"testing" package import  
"Test"를 접두어, 파마미터 타입을 (*testing.T) 로 테스트 함수 생성


```golang
package prose

import (
	"fmt"
	"testing"
)

func TestJoinWithCommas(t *testing.T) {

	input := []string{"apple", "orange", "grape"}
	want := "apple, orange and grape"

	get := JoinWithCommas(input)
	// fmt.Println(get)

	if get != want {
		t.Errorf("Input :: %#v", input)
		t.Errorf("Wanted :: %#v", want)
		t.Errorf("Output :: %#v", get)
	}

}
```

go test {패키지 경로} 로 test 수행

```golang
❯ go test ./prose
--- FAIL: TestJoinWithCommas (0.00s)
    prose_test.go:16: Input :: []string{"apple", "orange", "grape"}
    prose_test.go:17: Wanted :: "apple, orange and grape"
    prose_test.go:18: Output :: "apple- orange and grape"
FAIL
FAIL    automated/testing/prose 0.347s
FAIL
```

## TDD, BDD ??

> https://m.blog.naver.com/rkdudwl/221973507455
https://velog.io/@y_dragonrise/Agile-TDD-BDD