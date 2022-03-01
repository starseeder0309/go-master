# CHAPTER 27. 객체지향 설계 원칙 SOLID

# 27.1 객체지향 설계 5가지 원칙 SOLID

- 단일 책임 원칙(Single Responsibility Principle, SRP)
- 개방-폐쇄 원칙(Open-Closed Principle, OCP)
- 리스코프 치환 원칙(Liskov Substitution PRinciple, LSP)
- 인터페이스 분리 원칙(Interface Segregation Principle, ISP)
- 의존 관계 역전 원칙(Dependency Inversion Principle, DIP)

SOLID 5가지 원칙은 반드시 지켜야 하는 의무사항은 아니지만 이 원칙들에 입각해서 더 좋은 소프트웨어를 설계할 수 있다.

SOLID 원칙은 상호 연결되어 있어서 하나만 잘지켜져도 나머지가 저절로 지켜지는 경우가 많다.

## 왜 설계를 잘해야 하는가?

설계 : 프로그램 코드를 이루는 각 모듈 간 의존 관계를 정의하는 것

→ 현대의 프로그래밍은 과거에 비해서 매우 복잡하다.

## 나쁜 설계

- 경직성(rigidity) : 모듈 간의 결합도(coupling)가 너무 높아서 코드를 변경하기 매우 어려운 구조
- 부서기기 쉬움(fragility) : 한 부분을 건드렸더니 다른 부분까지 망가지는 경우
- 부동성(imoobility) : 코드 일부분을 분리해서 다른 프로젝트에도 사용하고 싶지만 모듈 간 결합도가 너무 높아서 옮길 수 없는 경우

상호 결합도가 매우 높고 응집도 낮다.

## 좋은 설계

상호 결합도가 낮고 응집도가 높은 설계

# 27.2 단일 책임 원칙

정의

- 모든 객체는 책임을 하나만 져야 한다.
- 객체나 모듈은 변경하려는 단 하나 이유만을 가져야 한다.

이점

- 코드 재사용성을 높여준다.

```go
type Report interface {
	Report() string
}

type FinanceReport struct {
	report string
}

func (r *financeReport) Report() string {
	return r.report
}

type ReportSender struct {
	// ...
}

func (s *ReportSender) SendReport(report Report) {
	// ...
}
```

# 27.3 개방-폐쇄 원칙

정의

- 확장에는 열려 있고, 변경에는 닫혀 있다.
- 프로그램에 기능을 추가할 때 기존 코드의 변경을 최소화해야 한다.

이점

- 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 된다.

```go
type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct {
	// ...
}

func (s *emailSender) Send(r *Report) {
	// 이메일 전송
}

type FaxSender struct {
	// 팩스 전송
}
```

# 27.4 리스코프 치환 원칙

정의

- “q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라 하자. 그렇다면 S가 T의 하위 타입이라면 q(y)는 타입 S의 객체 y에 대해 증명할 수 있어야 한다.”
- 함수 계약 관계를 준수하는 코드를 작성해야 한다.

이점

- 예상치 못한 작동을 예방할 수 있다.

```go
type T interface {
	Something()
}

type S struct {
}

// 구조체 S는 인터페이스 T를 구현함.
func (s *S) Something() {
}

type U struct {
}

// 구조체 U는 인터페이스 T를 구현함.
func (u *U) Something() {
}

func q(t T) {
}

var s = &S{} // S 타입 변수 s
var u = &U{} // U 타입 변수 u

q(y)
q(u)
```

함수 정의 → 함수를 호출하는 호출자와 함수 구현 간의 계약 관계가 발생한다.

# 27.5 인터페이스 분리 원칙

정의

- 클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.
- 많은 메서드를 포함하는 커다란 인터페이스보다는 적은 수의 메서드를 가진 인터페이스가 여러 개로 이뤄진 객체를 지향해야 한다.

이점

- 인터페이스를 분리하면 불필요한 메서드들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있다.

```go
type Report interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}
```

# 27.6 의존 관계 역전 원칙

정의

- 상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.

원칙

1. 상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야 한다.
2. 추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다.

이점

- 구체화된 모듈이 아닌 추상 모듈에 의존함으로써 확장성이 증가한다.
- 상호 결합도가 낮아져서 다른 프로그램으로의 이식성이 증가한다.

## 27.6.1 원칙 1 뜯어보기

> 상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야 한다.
> 

각 모듈의 의존 관계를 떨어뜨리면 모듈은 본연의 기능에 충실할 수있다.

## 27.6.2 원칙 2 뜯어보기

> 추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다.
> 

```go
type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnReceive() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("알람!")
}

var mail = &Mail{}
var listener EventListener = &Alam{}

mail.Register(listener)
mail.OnReceive()
```

# 27.7 학습 마무리

결합도는 낮게, 응집도는 높게.