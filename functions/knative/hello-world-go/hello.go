package hello

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "%s\t, Hello, %s!\n", hostname, r.URL.Path[1:])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	go func() {
		for {
			fmt.Printf("Hello, %s!\n\n", hostname)
			time.Sleep(3 * time.Second)
		}
	}()
	<-ctx.Done()
}
