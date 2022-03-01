# CHAPTER 28. 테스트와 벤치마크

# 28.1 테스트 코드

Go는 테스트 코드 작성과 실행을 언어 차원에서 지원한다. 빠르고 손쉽게 테스트 코드를 작성할 수 있어 사전에 버그를 막는 데 효과적이다.

3가지 표현 규약을 따라 테스트 코드를 작성해야 하며, `go test`  명령으로 실행한다.

1. 파일 이름이 `_test.go`로 끝나야 한다.
→ 테스트 코드는 파일 이름이 `_test.go`로 끝나는 파일 안에 존재해야 한다.
2. testing 패키지를 임포트해야 한다.
→ 테스트 코드를 작성하려면 `import “testing”`으로 testing 패키지를 가져와야 한다.
3. 테스트 코드는 `func TestX(t *testing.T)` 형태여야 한다.
→ 테스트 코드는 함수 내부에 작성되어야 하고 함수명은 반드시 `Test`로 시작해야 한다. 접두사 뒤의 첫 글자는 대문자여야 한다. 또한 매개변수는 `t *teting.T` 하나만 존재해야 한다.

## 28.1.1 테스트 코드 작성하기

```go
package main

import "fmt"

func square(x int) int {
	return 81
}

func main() {
	fmt.Printf("9 * 9 = %d\n", square(9))
}
```

```go
package main

import "testing"

func TestSquare1(t *testing.T) {
	result := square(9)
	if result != 81 {
		t.Errorf("square(9) should be 81 but square9) returns %d.", result)
	}
}

func TestSquare2(t *testing.T) {
	result := square(3)
	if result != 9 {
		t.Errorf("square(3) should be 9 but square9) returns %d.", result)
	}
}
```

실패 메시지 처리 메서드

- `t.Error()` : 테스트가 실패하면 모든 테스트를 중단한다.
- `t.Fail()` : 테스트가 실패해도 다른 테스트를 계속 진행한다.

## 28.1.2 일부 테스트만 실행하기

`go test`를 실행하면 전체 테스트를 모두 실행하지만, `-run` 플래그를 사용해서 특정 테스트만 실행할 수 있다.

```bash
> go test -run [testName]
```

테스트는 다다익선이다. 되도록 많은 테스트 코드를 작성해 빠짐없이 코드를 검증해 버그 없는 프로그램을 만드는 데 적극 활용하자.

## 28.1.3 테스트를 돕는 외부 패키지

[stretchr/testify](https://github.com/stretchr/testify)

[https://github.com/stretchr/testify](https://github.com/stretchr/testify)

```bash
> go get -u github.com/stretchr/testify
```

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81.")
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9.")
}
```

**stretchr/testify/assert 패키지에서 제공하는 유용한 함수**

- Equal() : expected와 actual 값을 비교하여 다를 경우 테스트를 실패시키고 메시지를 출력한다.
- Greater() : e1이 e2보다 크지 않으면 테스트를 실패시키고 메시지를 출력한다.
- Len() : object의 항목 개수가 length가 아니면 테스트를 실패시키고 메시지를 출력한다.
- NotNilf() : object가 nil이면 테스트를 실패시키고 메시지를 출력한다.
- NotEqualf() : expected와 actual 값이 비교하여 같을 경우 테스트를 실패시키고 메시지를 출력한다.

**stretchr/testify 패키지에서 제공하는 그외 유용한 기능**

- mock 패키지 : 모듈의 행동을 가장하는 목업(mockup) 객체를 제공한다. 네트워크 기능 테스트를 할 때 유용하다.
- suite 패키지 : 테스트 준비 작업이나 테스트 종료 후 뒤처리 작업을 쉽게 할 수 있도록 도와준다

# 28.2 테스트 주도 개발

테스트의 중요성이 점차 커지고 있는 이유

1. 과거에 비해서 프로그램 규모가 커졌기 때문이다.
2. 과거에 비해 고가용성(high availability)에 대한 요구사항이 높아졌다.

**블랙박스 테스트**

- 제품 내부를 오픈하지 않은 상태에서 진행되는 테스트
- 사용성 테스트(usability test)
- 프로그램 내부 코드를 직접 검증하는 게 아니라 프로그램을 실행한 상태로 실행 동작을 검사한다.
- 프로그래머보다는 전문 테스터, QA, QV 직군에서 주로 담당한다.

**화이트박스 테스트**

- 프로그램 내부 코드를 직접 검증하는 테스트
- 단위 테스트(unit test)
- 프로그래머가 직접 테스트 코드를 작성해서 내부 테스트를 검사한다.

**테스트 주도 개발(Test Driven Development, TDD)**

- 화이트박스 테스트에서 테스트 코드 작성 시기를 과감하게 코드 작성 이전으로 옮긴 방식
- 과정
    1. 테스트 코드를 작성한다.
    2. 테스트를 진행한다. 구현하기 전에 테스트 코드부터 작성했기 때문에 당연히 테스트가 실패한다.
    3. 작성한 테스트를 성공시키는 가장 간단한 코드를 작성해 테스트를 성공시킨다.
    4. SOLID 원칙에 입각해 코드를 리팩토링(refectoring)한다.
    5. ...
- 장점
    - 테스트 코드가 자연적으로 촘촘해진다. 허술한 테스트 코드 없이 다양한 경우에 대해서 검증할 수 있다.
    - ‘작은 목표 설정 → 실패 → 달성 → 달성 강화 → 새로운 작은 목표 설정' 절차를 따르게 때문에 개발 자체가 재밌어진다.

# 28.3 벤치마크

Go는 코드 성능을 검사하는 벤치마크 기능을 언어 차원에서 지원한다.

3가지 표현 규약을 따라 테스트 코드를 작성해야 하며, `go test` 명령으로 실행한다.

1. 파일 이름이 `_test.go`로 끝나야 한다.
→ 테스트 코드는 파일 이름이 `_test.go`로 끝나는 파일 안에 존재해야 한다.
2. testing 패키지를 임포트해야 한다.
→ 테스트 코드를 작성하려면 `import “testing”`으로 testing 패키지를 가져와야 한다.
3. 벤치마크 코드는 `func BenchmarkX(b *testing.B)` 형태여야 한다.
→ 테스트 코드는 함수 내부에 작성되어야 하고 함수명은 반드시 `Benchmark`로 시작해야 한다. 접두사 뒤의 첫 글자는 대문자여야 한다. 또한 매개변수는 `b *teting.B` 하나만 존재해야 한다.

```go
package main

import "fmt"

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	first := 1
	second := 0
	result := 0
	for i := 2; i <= n; i++ {
		result = first + second
		second = first
		first = result
	}

	return result
}

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}
```

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci(-1) should be 0.")
	assert.Equal(0, fibonacci1(0), "fibonacci(0) should be 0.")
	assert.Equal(1, fibonacci1(1), "fibonacci(1) should be 1.")
	assert.Equal(2, fibonacci1(3), "fibonacci(3) should be 2.")
	assert.Equal(233, fibonacci1(13), "fibonacci(223) should be 233")
}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci(-1) should be 0.")
	assert.Equal(0, fibonacci2(0), "fibonacci(0) should be 0.")
	assert.Equal(1, fibonacci2(1), "fibonacci(1) should be 1.")
	assert.Equal(2, fibonacci2(3), "fibonacci(3) should be 2.")
	assert.Equal(233, fibonacci2(13), "fibonacci(223) should be 233")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}
```

```bash
goos: darwin
goarch: amd64
pkg: tucker-go-language-programming/ch28-test-and-benchmark/ex28-02
cpu: VirtualApple @ 2.50GHz
BenchmarkFibonacci1-10             38095             31915 ns/op
BenchmarkFibonacci2-10          163231354                7.321 ns/op
PASS
ok      tucker-go-language-programming/ch28-test-and-benchmark/ex28-02  3.960s
```

```bash
> go test -bench .
```
