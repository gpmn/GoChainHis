/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"GoChainHis/cmd"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// handleException :
func handleException() {
	c := make(chan os.Signal, 10)
	signal.Notify(c, syscall.SIGABRT)
	signal.Notify(c, syscall.SIGSTOP)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGKILL)
	signal.Notify(c, syscall.SIGSEGV)
	signal.Notify(c, syscall.SIGQUIT)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.Signal(30)) // custom signal
	go func(c chan os.Signal) {
		for {
			sig := <-c
			log.Println("get signal : ", sig)
			buf := make([]byte, 1024*200)
			cnt := runtime.Stack(buf, true)
			buf = buf[:cnt]

			log.Printf(`=== BEGIN goroutine stack dump ===
	%s
	=== END goroutine stack dump ===
	`, string(buf))
			switch sig {
			case syscall.SIGTERM:
				fallthrough
			case syscall.SIGSTOP:
				fallthrough
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGABRT:
				fallthrough
			case syscall.SIGKILL:
				fallthrough
			case syscall.SIGQUIT:
				fallthrough
			case syscall.SIGSEGV:
				fallthrough
			default:
				log.Printf("got signal %v, exiting ...", sig)
				time.Sleep(3 * time.Second)
				os.Exit(1)
			}
		}
	}(c)
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Ldate)
	handleException()

	cmd.Execute()
}
