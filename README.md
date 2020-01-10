# goroutine-example
concurrency example by GoLang

** RUN from bash:
 /path/ go run main.go and go run channels.go

---A goroutine is a function that is capable of running concurrently with other functions.This program consists of two goroutines.

---Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.

        