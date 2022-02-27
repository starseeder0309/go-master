# CHAPTER 01. 네트워크 시스템 개요

<aside>
💡 디지털 시대에는 점점 더 많은 수의 장치가 컴퓨터 네트워크를 통해 통신한다. **컴퓨터 네트워크(computer network)**는 두 개 이상의 장치 또는 **노드(node)** 간의 연결을 말하며, 각 노드가 데이터를 공유할 수 있도록 한다. 이러한 연결은 본질적으로 신뢰할 수 없거나 안전하지도 않다. 그러나 감사하게도 Go 언어의 표준 라이브러리와 풍부한 생태계를 활용하면 안전하고 신뢰할 수 있는 네트워크 애플리케이션을 작성할 수 있다.

</aside>

# 네트워크 토폴로지 선택하기

토폴로지(topology) : 네트워크상의 노드 구성

- 유형 : 점대점 연결형, 데이지 체인형, 버스형, 링형, 스타형, 그물형

점대점 연결형(point-to-point)

- 두 노드가 하나의 연결을 공유한다.
- 두 노드 간에 직접 통신이 필요할 때 유용하다.

데이지 체인형(daisy chain)

- 일련의 점대점 연결
- 홉(hop) : 출발지 노드와 목적지 노드 사이의 중간 노드
- 모던 네트워크에서 이 토폴로지를 발견할 가능성의 거의 없다.

버스형(bus)

- 노드들은 공통, 공유 네트워크 링크를 갖는다.
- 유선 네트워크에서 일반적으로 사용되지 않고, 무선 네트워크에서 사용된다.

링형(ring)

- 일부 광섬유 네트워크 배포에 사용된다.
- 데이터가 단일 방향으로 이동하는 폐쇄 루프다.
- 데이터의 속도는 가장 느린 노드의 속도로 제한된다.

스타형(star)

- 중앙 노드는 다른 모든 노드에 개별 점대점 연결을 한다.
- 대부분의 유선 네트워크의 토폴로지
- 중앙 노드는 종종 **네트워크 스위치(network switch)**며, 이는 발신 노드로부터 데이터를 수신하여 우편 서비스와 같이 목적지 노드로 데이터를 재전송하는 장치다.
- 장치를 스위치에 연결하는 것만으로 아주 간단하게 네트워크에 노드를 추가할 수 있다.
- 데이터는 단 하나의 홉만 이동한다.

그물형(mesh)

- 네트워크의 모든 노드는 다른 모든 노드와 직접 연결된다.
- 단일 노드의 장애가 네트워크에 있는 다른 노드 간의 트래픽에 영향을 주지 않기 때문에 단일 장애 지점(single point of failure, SPOF)을 제거한다.
- 노드 수가 증가함에 따라 비용과 복잡성이 증가하여 대규모 네트워크에서는 사용할 수 없다.

하이브리드 네트워크 토폴로지(hybrid network topology) : 둘 이상의 기본적인 토폴로지를 결합한 형태

- 목적 : 각 토폴로지의 장점을 활용하고 단점을 개별 네트워크 세그먼트로 제한하여 안정성, 확장성 및 유연성을 향상시킨다.

하이브리드 네트워크 토폴로지의 종류

- 스타-링(star-ring) 하이브리드 네트워크 : 중앙 노드에 연결된 일련의 링 네트워크
- 스타-버스(star-bus) 하이브리드 네트워크 : 버스와 스타 네트워크 토폴로지의 조합으로 형성된 계층적 구조의 토폴로지

# 대역폭 vs. 레이턴시

대역폭(bandwidth) : 일정 시간 내에 네트워크 연결을 통해 전송할 수 있는 데이터의 양

- [예시] 인터넷 연결 다운로드 속도가 100Mbps인 인터넷 상품
→ 이론적으로 인터넷 연결이 인터넷 서비스 공급자(ISP)에서 모뎀으로 초당 최대 100메가비트를 전송할 수 있다.
- 속도가 빠르다고 해서 반드시 성능이 향상되는 것은 아니다.
→ 대역폭이 낮더라도 레이턴시가 작은 경우 성능이 더 좋을 수 있다.

레이턴시(latency) : 네트워크 리소스 요청을 보내고 응답을 받는 사이에 측정된 시간

- [예시] 웹사이트의 링크를 클릭한 후 결과 페이지가 화면상에 렌더링될 때까지 소요되는 지연시간
- 레이턴시가 크면 나쁜 사용자 경험(UX)을 제공하게 되며, 종종 악의적인 공격자들은 큰 레이턴시를 이용하여 사용자가 소프트웨어나 서비스를 사용할 수 없도록 하는 공격(DDoS)을 할 수 있다.
→ 웹 애플리케이션이나 API 등의 네트워크 소프트웨어를 작성할 때 레이턴시를 최소화하는 것은 UX에서 굉장히 중요하며, 검색 엔진에서 높은 순위를 얻을 수도 있다.

일반적인 레이턴시 해결 방법

- 콘텐츠 전송 네트워크(content delivery network, CDN)나 클라우드를 사용하여 사용자와 서비스 사이의 거리와 홉 수를 모두 줄인다.
- 요청 및 응답 크기를 최적화한다.
- 캐싱 전략을 수립하여 성능을 개선한다.
- Go 언어의 고루틴(goroutine)을 통해 동시성과 병행성의 이점을 활용하여 서버 측 응답이 블로킹되는 것을 최소화한다.

# 개방형 시스템 상호 연결 참조 모델

개방형 시스템 상호 연결(Open Systems Interconnection, OSI) 참조 모델

- 프로토콜의 개발과 통신을 위한 프레임워크
- 프로토콜(protocol) : 네트워크를 통해 전송되는 데이터의 포맷과 순서를 결정하는 규칙 혹은 절차

## OSI 참조 모델의 계층 구조 레이어

네트워크 내의 모든 활동을 7개의 계층으로 구성된 엄격한 계층 구조로 구분한다.

각각의 레이어를 스택으로 배열하며, 7계층은 맨 위에, 1계층은 맨 아래에 있다.

![OSI 참조 모델의 7개의 계층](CHAPTER%2001%209d73c/Untitled.png)

OSI 참조 모델의 7개의 계층

계층을 나누는 것 → 추상화(abstraction)

**7계층 - 애플리케이션 계층**

- 네트워크 애플리케이션과 라이브러리는 대부분 이 계층과 상호작용한다.
- 호스트를 식별하고 리소스를 검색한다.
- [예시] 웹 브라우저, 비트토렌트 클라이언트

**6계층 - 프레젠테이션 계층**

- 데이터가 아래 계층으로 이동할 때 네트워크 계층에 대한 데이터를 준비하고, 데이터가 스택 위로 이동할 때 애플리케이션 계층에 대한 데이터를 제공한다.
- [예시] 암호화 및 복호화, 데이터 인코딩 및 디코딩

**5계층 - 세션 계층**

- 네트워크 노드 간의 연결 수명 주기를 관리한다.
- 일부 7계층의 프로토콜은 5계층에서 제공하는 서비스에 의존한다.
- [예시] 연결 수립, 연결 시간 초과 관리, 작동 모드 조정, 연결 종료

**4계층 - 전송 계층**

- 전송의 안정성을 유지하면서 두 노드 간 데이터의 전송을 제어하고 조정한다.
- 수신자가 데이터 수신을 인정하지 않는 경우 이 계층의 프로토콜이 데이터를 재전송한다.
- [예시] 에러 수정, 데이터 전송 속도 제어, 데이터 청킹(chunking) 및 분할, 누락된 데이터 재전송, 수신 데이터 확인

**3계층 - 네트워크 계층**

- 노드 간에 데이터를 전송한다.
- 원격 노드에 직접 점대점으로 연결하지 않고도 네트워크 주소로 데이터를 전송할 수 있다.
- 전송의 신뢰성이나 전송 에러 등을 전송자에게 알려 주기 위한 별도의 프로토콜이 필요하지 않다.
- [예시] 라우팅, 주소 지정, 멀티캐스팅, 트래픽 제어

**2계층 - 데이터 링크 계층**

- 직접 연결된 두 노드 간의 데이터 전송을 처리한다.
- 시스템에서 스위치로, 스위치에서 다른 시스템으로 데이털르 쉽게 전송할 수 있도록 지원한다.
- 물리 계층의 에러를 식별하여 수정을 시도한다.
- 재전송 및 흐름 젱 ㅓ기능은 하위의 물리적 매체에 따라 달라진다.

**1층 - 물리 계층**

- 네트워크 스택에서 발생한 비트를 하위의 물리적 매체가 제어할 수 있는 전기, 광학 또는 무선 신호로 변환하고, 타 노드의 물리적 매체에서 받은 신호를 다시 비트로 변환한다.
- 비트 전송률을 제어한다.

## 데이터 캡슐화를 사용한 트래픽 전송

캡슐화(encapsulation) : 구현의 세부 정보를 숨기고 관련한 기타 세부 정보만 사용하는 방법

- 데이터가 스택을 따라 이동할 때 하위 계층에 의해 캡슐화된다.
- 단일 OSI 계층에서 작동하는 프로토콜도 데이터 캡슐화를 사용한다.

페이로드(payload) : 스택을 따라 이동하는 데이터

- 메시지 본문(message body), 서비스 데이터 단위(service data unit, SDU)
- 페이로드가 스택 위로 이동하면 각 레이어는 이전 레이어에서의 헤더 정보를 제거한다.
- 클라이언트의 네트워크 스택을 타고 물리적인 매체를 통해 서버로 이동하여, 서버의 네트워크 스택에서 클라이언트가 페이로드를 전송한 동일한 계층으로 이동한다.
→ 출발지 노드의 어떤 계층에서 전송된 데이터는 목적지 노드의 동일한 계층에 도달한다.

수평 통신(horizontal communication) : 동일 계층 내의 클라이언트와 서버 간의 통신

- 실제로 데이터는 클라이언트의 네트워크 스택을 따라 하위 계층으로 이동한 후에 서버의 네트워크 스택을 따라 상위 계층으로 이동한다.

# TCP/IP 모델

종단 간의 연결 원칙(end-to-end principle)

- 각 네트워크의 세그먼트는 비트를 적절하게 전송하고 라우팅하기에 필요한 기능만을 포함한다.
- 그 외의 모든 기능은 엔드포인트 혹은 송신자와 수신자의 네트워크 스택에 속한다.
- 견고한 구현을 권장한다.
→ 요구 사항이 제대로 구성된 패킷을 전송하되, 패킷이 기술 사양을 준수하는지의 여부에 관계없이 수신된 패킷은 모두 수락해야만 한다.

TCP/IP 모델 구성

- 애플리케이션 계층(애플리케이션, 프레젠테이션, 세션 계층)
- 전송 계층
- 인터넷 계층
- 링크 계층(데이터 링크, 물리 계층)

![4개의 계층으로 이루어진 TCP/IP 모델과 7개의 계층으로 이루어진 OSI 참조 모델](CHAPTER%2001%209d73c/Untitled%201.png)

4개의 계층으로 이루어진 TCP/IP 모델과 7개의 계층으로 이루어진 OSI 참조 모델

## 애플리케이션 계층

애플리케이션 계층(Application Layer)

- 소프트웨어 애플리케이션과 직접 상호작용한다.
- 특정한 프레젠테이션 기능이나 세션 기능을 정의하지 않기 때문에 일반적으로 세 개의 OSI 계층을 포함한다.
→ 특정 애플리케이션에서 프로토콜을 구현할 때에는 OSI 모델의 프레젠테이션 계층, 세션 계층에서 필요로 하는 세부 정보가 중요하다.

애플리케이션 계층의 주요 프로토콜

- HTTP(Hyper Text Transfer Protocol)
- FTP(File Transfer Protocol)
- SMTP(Simple Mail Transfer Protocol)
- DHCP(Dynamic Hosts Configuration Protocol)
- DNS(Domain name System)

## 전송 계층

전송 계층(Transport Layer)

- 두 노드 간의 데이터 전송을 처리한다.
- 출발지에서 전송된 모든 데이터가 목적지로 **데이터 무결성(integrity)**을 보장하며 오나전하고도 올바르게 전송되도록 한다.
→ 중복 또는 누락된 데이터 없이 목적지에서 수신한 데이터의 순서가 정확하다.
- 페이로드로 세그먼트(segment) 또는 데이터그램(datagram)을 처리한다.

전송 계층의 주요 프로토콜

- TCP(Transmission Control Protocol)
- UDP(User Datagram Protocol)

## 인터넷 계층

인터넷 계층(Internet Layer)

- 출발지 노드와 목적지 노드 사이의 상위 계층에서 데이터 패킷을 라우팅한다.
- 종종 이기종 물리적 매체를 사용하는 여러 네트워크를 통해 라우팅한다.

인터넷 계층의 주요 프로토콜

- IPv4
- IPv6
- BGP(Border Gateway Protocol)
- ICMP(Inter Control Message Protocol)
- IGMP(Inter Group Management Protocol)
- IPsec

## 링크 계층

링크 계층(Link Layer)

- TCP/IP 프로토콜의 핵심적인 부분과 물리적인 매체 사이의 인터페이스

링크 계층의 주요 프로토콜

- ARP(Address Resolution Protocol)
    - 주소 결정 프로토콜
    - 노드의 IP 주소를 네트워크 인터페이스의 MAC 주소로 변환한다.
    - 프레임을 물리적 네트워크에 전달하기 전에 각 프레임의 헤더에 MAC 주소를 포함한다.