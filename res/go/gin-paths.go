package res

fmt.Println("------")
fmt.Printf("%+v\n", g.Request)
fmt.Printf("%+v\n", g.Request.Header.Get("origin"))
fmt.Println("------")

.go#L50