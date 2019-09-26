package main

import (
    "testing"
    "math/rand"
    "time"
)

// run cmd in terminal for test
// go test heap.go heap_test.go -v

func is_heap(nodes []Node) bool {
    for i := len(nodes)-1; i > 0; i-- {
        if nodes[i].Value < nodes[(i-1)/2].Value {
            return false
        }
    }
    return true
}


// test empty array
func Test_empty(t *testing.T) {
    a := []Node{}
    Init(a)
    Pop(a)
    n := Node{1}
    Push(n, a)
    Remove(a, n)
    // fmt.Println(a)
    if !is_heap(a) {
        t.Error("Error")
    }
}

func Test_Init(t *testing.T) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    a := [10000]Node{}
    for i:= 0; i < 100; i++ {
        r = rand.New(rand.NewSource(int64(r.Int())))
        for j, _ := range a {
            a[j].Value = r.Int()
        }
        Init(a[:])
        if !is_heap(a[:]) {
            t.Error("Error")
        }
    }
}

func Test_Push(t *testing.T) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    a := []Node{}
    Init(a)
    for i:= 0; i < 10000; i++ {
        Push(Node{r.Int()}, a)
        if !is_heap(a) {
            t.Error("Error")
        }
    }
}

func Test_Pop(t *testing.T) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    a := [10000]Node{}
    for j, _ := range a {
        a[j].Value = r.Int()
    }
    b := a[:]
    Init(b)
    for i := 0; i < 10000; i++ {
        _, b = Pop(b)
        if !is_heap(b) {
            t.Error("Error")
        }
    }
    if len(b) > 0 {
        t.Error("Error2")
    }
}

func Test_Remove(t *testing.T) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    a := [10000]Node{}
    for j, _ := range a {
        a[j].Value = r.Int()
    }
    b := a[:]
    Init(b)
    for i := 0; i < 100; i++ {
        b = Remove(b, b[r.Int()%len(b)])
        if !is_heap(b) {
            t.Error("Error")
        }
    }
}

