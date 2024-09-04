# ImageToASCII
Golang module to convert images to ASCII art
## How to install
Install:
```
go get github.com/pkondratev/ImageToASCII
```
In code:
```
import "github.com/pkondratev/ImageToASCII"
```
## How to use it
At first you need to get "ImageASCII" object using one of the next functions:  
- func LoadFromImage(img image.Image) (*ImageASCII, error)
- func LoadFromStream(r io.Reader) (*ImageASCII, error)
- func LoadFromFile(file_name string) (*ImageASCII, error)

## Example
```go
package main
import (
	ascii "github.com/bobbykitten/ImageToASCII"
	"os"
	"fmt"
)
func main() {
	img, err := ascii.LoadFromFile("image.jpg")
	if err != nil {
		panic(err)
	}
	out, err := os.Create("out.txt")
	defer out.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(out, img.ToString())
}
```
