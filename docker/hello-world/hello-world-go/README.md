Write a simple Go program (hello.go):

go
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Compile it statically:


```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hello main.go
```
GOOS=linux: Target Linux as the operating system.
GOARCH=amd64: Target 64-bit architecture.
CGO_ENABLED=0: Ensure a statically linked binary.
Build the Docker image:

Create a Dockerfile with the content you provided:

```
FROM scratch
COPY hello /
CMD ["/hello"]
```
Build the image:

```
docker build -t hello-world .
```
Run the container:

```
docker run --rm hello-world

```
Output:

Hello, World!
