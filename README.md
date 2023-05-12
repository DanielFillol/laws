# Laws
This package provides functionality to retrieve legal typification's related to a given subject. The package includes functions to read and parse a CSV file containing legal information and then store it in memory to provide quick access.

## Usage
Import the laws package and call the GetLawsBySubject function to retrieve legal typifications for a given subject.
```go
package main

import (
	"fmt"
	"github.com/DanielFillol/laws"
)

func main() {
	typifications := laws.GetLawsBySubject("Crimes contra a vida / Homicídio Simples")
	fmt.Println(typifications)
}
```
This will be the output:
```
[{[{Artigo 121 do Decreto Lei nº 2.848 de 07 de Dezembro de 1940}] Homicídio} {[{Parágrafo 1 Artigo 205 do Decreto Lei nº 1.001 de 21 de Outubro de 1969}] Homicídio simples}]
```

## Testing
This package includes a test file getLawsBySubject_test.go that tests the GetLawsBySubject function with various inputs. To run the tests, use go test:
```go
go test
```

## Contributing
Contributions are welcome! If you find a bug or would like to add a feature, please open an issue or pull request.
