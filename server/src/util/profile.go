package util

import (
    "os"
    "log"
    "runtime/pprof"
)

func Profile() {
    f, err := os.OpenFile("./tmp/cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

	// 注意，有时候 defer f.Close()， defer pprof.StopCPUProfile() 会执行不到，这时候我们就会看到 prof 文件是空的， 我们需要在自己代码退出的地方，增加上下面两行，确保写文件内容了。
	pprof.StopCPUProfile()
	f.Close()
}