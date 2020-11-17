# 11. Interface


## Interface  
인터페이스는 특정 값이 가지고 있기를 기대하는 메소드의 집합.  
인터페이스에 정의된 메소드명, 매개변수 타입, 리턴 값 타입이 모두 일치해야 해당 타입이 인터페이스에 만족.

> 메소드(Method)와 함수(Function)의 차이점은 무엇일까 ???  
> https://c10106.tistory.com/1886

```golang
type Myinterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type Mytype int

func (m Mytype) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m Mytype) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called", f)
}

func (m Mytype) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}
```
#### Interface type, Concrete type
 * Interface type(인터페이스 타입) : 무엇을 할 수 있는지만 정의.
 * Concrete type(구체 타입) : 무엇을 할 수 있는지와 그 값이 무엇 인지 정의.


## Assert
인터페이스 값을 특정 타입으로 변환.
> {인터페이스 값}.({변환 할 타입})

```golang
var noiseMaker NoiseMaker = Robot("Cleaner")

robot, ok := noiseMaker.(Robot)

// ok -> 타입 변환에 성공 할 경우 true
if ok {
    robot.Walk()
}
```

## Error Interface
error 타입은 int, string과 같은 "미리 정의된 식별자".  
유니버스 블록의 일부로 패키지에 상관 없이 어느 곳에서나 사용 가능하므로 import가 필요 없음. 

```golang
type error interface {
	Error() string
}
```

## Stringer Interface
특정 타입의 출력 형식(string)을 직접 지정 가능.  
fmt 패키지에서 Stringer interface에 만족하는지 체크한 후 해당 메소드를 호출 함.

```golang
type Stringer interface {
	String() string
}
```
```golang
type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

func main() {

	fmt.Println(Gallons(12.0945423523)) //12.09 gal
	fmt.Println(Liters(12.0945423523)) //12.09 L

}
```
## Empty Interface
모든 타입의 값을 저장 가능.  
인터페이스를 만족하기 위한 메소드가 없기 때문에 모든 값이 만족 됨.
> 변수나 함수의 매개변수를 정의 할때 어떤 값이 올지 모를 경우 사용.

```golang
type Anything interface {
}
```

