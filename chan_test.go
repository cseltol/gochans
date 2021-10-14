package main_test

import (
	"testing"
	"time"
)

//checkReadOk
//checkReadHang
//checkReadPanic


func TestReadEmpty(t *testing.T) {
	ch := make(chan int, 5)
	checkReadHang(t, ch)
}

func TestReadClosedEmpty(t *testing.T) {
	ch := make(chan int, 5)
	close(ch)
	checkReadOk(t, ch)
}

func TestReadClosed(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 5
	close(ch)
	checkReadOk(t, ch)
}

//checkWriteOk
//checkWriteHang
//checkWritePanic

func TestWriteFull(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 5
	checkWriteHang(t, ch)
}

func TestWriteClosed(t *testing.T) {
	ch := make(chan int, 5)
	close(ch)
	checkWritePanic(t, ch)
}

//
// Там дальше хэлперы к тестам.
// Их пока смотреть не надо.
// Можно ознакомиться после занятия.
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// Я серьезно. Не нужно на них сейчас смотреть :)
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

func checkReadOk(t *testing.T, ch <-chan int) {
	t.Helper()
	select {
	case <-ch:
	case <-time.NewTimer(100 * time.Millisecond).C:
		t.Errorf("should read")
	}
}

func checkReadHang(t *testing.T, ch <-chan int) {
	t.Helper()
	select {
	case <-ch:
		t.Errorf("should not read")
	case <-time.NewTimer(100 * time.Millisecond).C:
	}
}

func checkReadPanic(t *testing.T, ch <-chan int) {
	t.Helper()
	t.Errorf("read never panic")
}

func checkWriteOk(t *testing.T, ch chan<- int) {
	t.Helper()
	defer func() {
		t.Helper()
		if v := recover(); v != nil {
			t.Errorf("shouldn't panic")
		}
	}()
	select {
	case ch <- 1:
	case <-time.NewTimer(100 * time.Millisecond).C:
		t.Errorf("should write")
	}
}

func checkWriteHang(t *testing.T, ch chan<- int) {
	t.Helper()
	defer func() {
		t.Helper()
		if v := recover(); v != nil {
			t.Errorf("shouldn't panic")
		}
	}()
	select {
	case ch <- 1:
		t.Errorf("should not write")
	case <-time.NewTimer(100 * time.Millisecond).C:
	}
}

func checkWritePanic(t *testing.T, ch chan<- int) {
	t.Helper()
	defer func() {
		recover()
	}()
	select {
	case ch <- 1:
		t.Errorf("should panic")
	case <-time.NewTimer(100 * time.Millisecond).C:
		t.Errorf("should panic")
	}
}
