package experiment

import (
    "log"
    "runtime"
    "fmt"
)

// Go 메모리 관리 테스트
func RunMemoryTest() {
    log.Println("== Memory test started ==")

    log.Println("Initial memory usage:")
    printMemUsage()

    // make([]T, len, cap)
    // T: 슬라이스의 타입
    // len: 슬라이스의 초기 길이
    // cap: 슬라이스의 초기 용량 (optional, 생략 시 len의 값)
    log.Println("#####  Creating large slice")
    mySlice := make([]int, 0, 1<<20) // 비트 연산자로 2의 20승, 슬라이스 초기용량 100만 설정
    for i := 0; i < cap(mySlice); i++ {
        mySlice = append(mySlice, i)
    }

    log.Println("Memory usage after creating large slice:")
    printMemUsage()

    log.Println("#####  Releasing slice reference")
    mySlice = nil // 슬라이스 참조 해제

    log.Println("#####  Running GC")
    runtime.GC() // runtime.GC() 로 GC 강제 실행

    log.Println("Memory usage after GC:")
    printMemUsage()

    log.Println("== Memory test finished ==")
}



// 메모리 사용량 출력
func printMemUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    // 메모리 사용량을 KB 단위로 출력
    fmt.Printf("현재 사용중인 힙 메모리 Alloc = %v KiB", m.Alloc / 1024)
    fmt.Println("프로그램 시작 이후 할당 메모리 총량 TotalAlloc = %v KiB", m.TotalAlloc / 1024)
    fmt.Println("Go 런타임 시스템 메모리 총량 Sys = %v KiB", m.Sys / 1024)
    fmt.Println("GC 발생 횟수 NumGC = %v\n", m.NumGC)
}