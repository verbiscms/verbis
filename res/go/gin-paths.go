package models

fmt.Println("------")
fmt.Printf("%+v\n", g.Request)
fmt.Printf("%+v\n", g.Request.Header.Get("origin"))
fmt.Println("------")


@Ainsley syscall.Signal(12) seems to be "user defined signal 2" on windows: https://golang.org/src/syscall/types_windows.go#L50