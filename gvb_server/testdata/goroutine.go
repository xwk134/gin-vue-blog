package main

import (
	"fmt"
	"sync"
)
	

/**生产者消费者的例子*/
func ProductAndConsumer() {        
    wg := sync.WaitGroup{}        
    wg.Add(1)        
    //带有缓冲区的通道
    cint := make(chan int, 10)        
    go func() {                
        //product  ，循环往通道中写入一个元素              
        for i := 0; i < 100; i++ {                       
            cint <- i                        
        }        
        //关闭通道
        close(cint)        
     }()        
    go func() {                
        defer wg.Done()                
        //consumer   遍历通道消费元素并打印        
        for temp := range cint {                        
            fmt.Println(temp) 
            //len函数可以查看当前通道元素个数
            fmt.Println("当前通道元素个数",len(cint))
        }        
    }()        
    wg.Wait()
}

func main() {
	ProductAndConsumer()
}