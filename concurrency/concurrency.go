package concurrency

import (
    "fmt"
)

func Concurrency() {
    // concurrency vs paralleslism
    // Mike van sickle - concurrent programming in Go
    // concurrency is abput independent executing process
    // Robe pike http://blog.golang.org/concurrency-is-not-paralleslism
    // Go routines are scheduled by go run time
    // Gor routine just layers on tops of threads
    // Why??
    // - Lightweight
    // go manages all go routines
    // Less switching - switching multiple routines against a single thread
    // Faster start up times
    // safe communication - channels
    // Go concurrency model - CSP Communicating Sequential Processes - Actor model

}