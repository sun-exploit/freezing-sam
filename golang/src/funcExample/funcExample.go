package funcExample

import (
    "fmt"
)

var a = 10
var b int
//can not use a:=10

func TestA(){
    fmt.Printf("%v\n", a)
    fmt.Printf("%v\n", b)
}

func NextInt(b [] byte, i int)(int,int){
    x := 0
    for ;i<len(b);i++ {
        x = x*10 + int(b[i]) - '0'
        fmt.Printf("=======%v\n", int(b[i]))
    }
    return x, i
}

/*
func ReadFull(r Reader, buf [] byte)(n int, err error){
    for len(buf) > 0 && err == nil{
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:len(buf)]
    }
    return
}
*/

func DeferTest1(){
    for i:= 0;i<5;i++{
        defer fmt.Printf("%v\n", i)
    }
}

func DeferTest2()(ret int){
    defer func(x int){
        ret = x
    }(5)
    return 0
}

func MultiArgsTest(args ...int){
    for k, v := range args{
        fmt.Printf("%v:%v\n", k, v)
    }
}

func PassFuncTest(arg string, f func(string)){
    f(arg)
}

func FuncMap(){
    xs := map[int] func()(int) {
        1:func() int {return 10},
        2:func() int {return 20},
        3:func() int {return 30},
    }
    
    for k,v := range xs{
        fmt.Printf("%v, %v\n", k, v())
    }
}

func ThrowsPanic(f func(int))(b bool){
    defer func(){
        if x:=recover();x!=nil{
            println("recovering")
            b = true
        }
    }()
    f(0)
    return
}
