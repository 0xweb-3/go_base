package main

import (
	"fmt"
	"sync"
	"time"
)

// 不可重入性
func NonReentrant() {
	rw := sync.RWMutex{}
	rw.RLock()
	fmt.Println("获取读锁")
	rw.Lock() // 读锁未解锁的情况获取写锁会导致panic
	fmt.Println("获取写锁")
	//rw.RUnlock()
	rw.Unlock()
}

// 写锁只有在读锁和写锁都处于未加锁的状态下才能成功加锁
func OnlyOneWriteLock() {
	lock := sync.RWMutex{}
	//lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("获取写锁")
	lock.Lock() //写锁未解锁的情况下会阻塞
	fmt.Println("获取写锁")
	defer lock.Unlock()
}

func LockandUnlockIndiffGoroutine(lock *sync.Mutex) {
	time.Sleep(time.Second)
	lock.Unlock()
	fmt.Println("LockandUnlockIndiffGoroutine:释放写锁")
}

func main() {
	//NonReentrant()
	//OnlyOneWriteLock()

	lock := sync.Mutex{}
	for {
		// 未解锁将在此处阻塞
		lock.Lock()
		fmt.Println("LockandUnlockIndiffGoroutine-main:获取写锁")
		LockandUnlockIndiffGoroutine(&lock)
		time.Sleep(time.Second * 1)
	}
}
