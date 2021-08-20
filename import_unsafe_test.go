package kvstore

import (
	"fmt"
	"runtime"
)

func newActions() []*Action {
	return []*Action{
		NewSetAction("a", NewIntValue(10)),
	}
}

// 利用 Unsafe 模拟成功修改
func ExampleUnsafeApplier_UnsafeMockApplyAndLock() {
	s := NewStore().(UnsafeStore) // 空 store，使用 Unsafe 操作
	_ = s.SetInt64("a", 8)

	// mock goroutine 1 尝试获取 mock 对象
	newStore, err := s.UnsafeMockApplyAndLock(newActions()) // 将 a 设置为 10
	if err != nil {
		panic(err)
	}
	fmt.Println("UnsafeMockApplyAndLock")

	// mock goroutine 2 在 goroutine 1 成功之前先行尝试将 a 的值改变
	// 当然，由于使用的是 UnsafeMockApplyAndLock 因此没有释放锁，这一操作会被阻塞
	ch := make(chan struct{})
	go func() {
		fmt.Println("SetInt64 in goroutine 2")
		_ = s.SetInt64("a", 16)
		fmt.Println("SetInt64 in goroutine 2: Success")
		ch <- struct{}{}
	}()

	// 确保 go 调度器进入上面的 goroutine 中执行
	runtime.Gosched()

	// mock goroutine 1
	fmt.Println("GetInt and ReplaceInner in goroutine 1")
	if v, _ := newStore.GetInt("a"); v != 10 {
		panic(fmt.Sprintf("v = %d != %d", v, 10))
	}
	// ReplaceInner 会上锁，但锁已经在上面的 UnsafeMockApplyAndLock 持有，因此
	// 这里使用不安全的无锁版本 UnsafeReplaceInnerWithoutLock
	s.UnsafeReplaceInnerWithoutLock(newStore.UnsafeUnderlying())
	// 并且需要手动解锁
	s.UnsafeUnlock()

	// 等待之前被阻塞的 goroutine 执行完毕
	_ = <-ch

	// 检查最终结果，一定是 mock goroutine 2 做的后修改，即为 16
	result, _ := s.GetInt("a")
	fmt.Println("a =", result)

	// Output:
	// UnsafeMockApplyAndLock
	// SetInt64 in goroutine 2
	// GetInt and ReplaceInner in goroutine 1
	// SetInt64 in goroutine 2: Success
	// a = 16
}

// 利用 Unsafe 模拟放弃修改
func ExampleUnsafeApplier_UnsafeMockApplyAndLock_fail() {
	s := NewStore().(UnsafeStore) // 空 store，使用 Unsafe 操作
	_ = s.SetInt64("a", 8)

	// mock goroutine 1 尝试获取 mock 对象
	newStore, err := s.UnsafeMockApplyAndLock(newActions()) // 将 a 设置为 10
	if err != nil {
		panic(err)
	}
	fmt.Println("UnsafeMockApplyAndLock")

	// mock goroutine 2 在 goroutine 1 成功之前先行尝试将 a 的值改变
	// 当然，由于使用的是 UnsafeMockApplyAndLock 因此没有释放锁，这一操作会被阻塞
	ch := make(chan struct{})
	go func() {
		fmt.Println("SetInt64 in goroutine 2")
		_ = s.SetInt64("a", 16)
		fmt.Println("SetInt64 in goroutine 2: Success")
		ch <- struct{}{}
	}()

	// 确保 go 调度器进入上面的 goroutine 中执行
	runtime.Gosched()

	// mock goroutine 1
	fmt.Println("GetInt and Discard in goroutine 1")
	if v, _ := newStore.GetInt("a"); v != 10 {
		panic(fmt.Sprintf("v = %d != %d", v, 10))
	}
	// 模拟因对 newStore 的检查不通过而放弃修改
	// 此时对原 store 进行查询得到的结果应当是未修改的 8
	// 因为锁被 UnsafeMockApplyAndLock 持有，因此这里获取原始 store 数据也需要用 Unsafe 版本
	v, _ := s.UnsafeGet("a")
	fmt.Printf("Original a = (%T) %d\n", v, v)
	// 简单手动解锁即可
	s.UnsafeUnlock()

	// 等待之前被阻塞的 goroutine 执行完毕
	_ = <-ch

	// 检查最终结果，一定是 mock goroutine 2 做的后修改，即为 16
	result, _ := s.GetInt("a")
	fmt.Println("a =", result)

	// Output:
	// UnsafeMockApplyAndLock
	// SetInt64 in goroutine 2
	// GetInt and Discard in goroutine 1
	// Original a = (int64) 8
	// SetInt64 in goroutine 2: Success
	// a = 16
}

// 并发修改异常
func ExampleUnsafeApplier_UnsafeMockApplyAndLock_concurrent_edit() {
	s := NewStore() // 空 store
	_ = s.SetInt64("a", 8)

	// mock goroutine 1 尝试获取 mock 对象
	newStore, err := s.MockApply(newActions()) // 将 a 设置为 10
	if err != nil {
		panic(err)
	}

	// mock goroutine 2 在 goroutine 1 成功之前先行将 a 的值改变了
	_ = s.SetInt64("a", 16)

	// mock goroutine 1
	if v, _ := newStore.GetInt("a"); v != 10 {
		panic(fmt.Sprintf("v = %d != %d", v, 10))
	}
	s.ReplaceInner(newStore.UnsafeUnderlying())

	// 这时，因为 goroutine1 应当是「一组操作的集合」，因此它应当是一同执行的
	// goroutine2 对 a 的并发修改发生在 goroutine1 以后
	// 因此最终 a 的值应当是 goroutine2 的 16 而不是 1 的 10
	// 但是因为异常，a 的值错误变成了 10

	result, _ := s.GetInt("a")
	fmt.Println("a =", result)

	// Output: a = 10
}
