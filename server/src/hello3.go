package main

import(
    // "fmt"
	// "sync"
	"module"
)

func main() {
    // p := &sync.Pool{
    //     New: func() interface{} {
	// 		fmt.Println("xxxxxxxxxxxxxxxx")
    //         return module.RUOKAsk{false}
    //     },
    // }

	// a := p.Get().(module.RUOKAsk)
	// // p.Put(a)
	// a = module.RUOKAsk{true}
	// p.Put(a)
	// p.Put(a)
	// // aa := &a
    // // p.Put(a)
	// b := p.Get().(module.RUOKAsk)
	// // // b = module.RUOKAsk{false}
	// // bb := &b
	// // p.Put(b)
	// c := p.Get().(module.RUOKAsk)
	// // cc := &c
	// // // p.Put(c)
	// // // d := p.Get().(module.RUOKAsk)
	// // fmt.Println(a, b, c, aa==bb, aa==cc, bb==cc)
	// fmt.Println(a, b, c)
	
	a := module.RUOKAsk{true}
	b := a.(module.UserMsg)
	//   p := &sync.Pool{
    //     New: func() interface{} {
	// 		fmt.Println("xxxxxxxxxxxxxxxx")
    //         return 1
    //     },
    // }

	// a := p.Get().(int)
	// // a = module.RUOKAsk{true}
	// aa := &a
    // p.Put(a)
	// b := p.Get().(int)
	// bb := &b
    // fmt.Println(a, b, aa==bb)
}