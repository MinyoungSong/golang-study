# 12. Recovery


## defer  
defer 키워드는 호출하고 있는 함수가 return 키워드로 인해 일찍 종료되더라도 특정함수의 호출이 수행됨을 보장함.

```golang
func Socialize() error {

	defer fmt.Println("Goodbye !!")
	fmt.Println("Hello")
	return fmt.Errorf("I don't want to talk")
	fmt.Println("Nice weather, eh?")
	return nil
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
```
```
Hello
Goodbye !!
2020/11/18 01:53:10 I don't want to talk
exit status 1
```

## panic
패닉이 발생하면 현재 함수는 실행을 중단하고 에러메세지를 출력한 뒤 크래쉬를 발생.
> 패닉은 사용자의 실수나 의도와는 무관한 프로그램 자체의 버그를 나타내는 "제어 불가능한" 상황에서만 사용

```golang
func one() {
	defer fmt.Println("defered in one()")
	two()
}

func two() {
	defer fmt.Println("defered in two()")
	panic("Let's see what's been deferred")
}

func main() {
	one()
}
```

```
defered in two()
defered in one()
panic: Let's see what's been deferred

goroutine 1 [running]:
main.two()
        /Users/smy/Documents/golang_practice/12_recovery/main-p366.go:14 +0x95
main.one()
        /Users/smy/Documents/golang_practice/12_recovery/main-p366.go:9 +0x85
main.main()
        /Users/smy/Documents/golang_practice/12_recovery/main-p366.go:18 +0x25
exit status 2
```

## recover
panic 상태에 빠진 프로그램을 복구 할 수 있는 내장 함수.
> recover() 를 직접 defer에 적용 할 수는 없음.  
별도 함수 안에 감싸고 그 함수를 defer 호출 해야함.

```golang
func calmDown() {
	recover()
}

func freakOut() {
	defer calmDown()
	panic("Oh no")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}

```

### recover 에서 panic 값 사용하기

```golang
func calmDown() {
	p := recover()

	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there is an error")
	panic(err)
}
```
```
there is an error
```

## Golang에서 Exception 처리는 없는가...??
>Exception 기능은 의도적으로 생략 되었음.  
Go 메인테이너는 Exception 기능이 프로그램의 흐름을 복잡하게 만들 뿐 아니라 나중에 제대로 처리를 못한다는 것을 발견함.  
에러가 발생 될 부분에 직접적으로 처리하는 것이 더 명확 하다고 생각함.  
그래서 panic과 recover를 exception 처리 하듯이 사용하지 말 것을 권장함.