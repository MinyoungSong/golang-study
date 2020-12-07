# 13. Go Routine


## goroutine(고루틴)  
고루틴은 한 작업을 멈추고 다른 작업을 수행 할 수 있는 동시성을 지원.  
또한 특정 상황에서는 한번에 여러 개의 작업을 동시에 수행 할 수 있는 병렬성도 지원.

> 동시성(Concurrency)과 병렬성(Parallelism)의 차이점은??  
https://seamless.tistory.com/42  
https://blog.lalaworks.com/20200206-concurrency-parallelism/

```golang
func a() {
	for i := 0; i < 10; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Print("b")
	}
}

func main() {

	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("\nend main()")
}
```

> 고루틴은 반환값(return)을 가진 함수를 사용 할 수 없음.

## channel(채널)
고루틴끼리 서로 통신(값 전달) 할수 있는 방법  
고루틴은 반환값(return)을 가진 함수를 호출 할 수 없기 때문에 채널을 사용해서 전달.  
  
채널은 특정 타입의 값만 주고 받을 수 있으므로 채널 선언 시 타입을 지정해야 함.

```golang
var myChannel chan float64
myChannel = make(chan float64)

//myChannel := make(chan float64)
```

채널에서 값의 송/수신에서는 <- 연산자를 사용

```golang
myChannel <- 3.14 // 채널에 값을 송신

returnVal := <-myChannel // 채널에서 값을 수신
```

채널에 값을 송신하면 다른 곳에서 값을 받아 갈때까지 현재 고루틴을 블로킹함.  
값을 수신할 때도 다른 곳에서 보내주기 전까지 현재 고루틴을 블로킹함.

> 동기와 비동기?? 블로킹과 논블로킹??  
https://deveric.tistory.com/99
http://homoefficio.github.io/2017/02/19/Blocking-NonBlocking-Synchronous-Asynchronous/

