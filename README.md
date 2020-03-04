# bddgo

Test your go web application using
[bddrest](https://github.com/pylover/bddrest) python package.


## Install

```bash
sudo apt install python3-venv
go get -U github.com/alibabaco/bddgo
```


## How to use

Navigate to your go package and create a directory named `tests`
```bash
cd path/to/go/package
bddgo init
```

### Initialize

The `init` command creates a python virtual environment and install the 
[bddrest](https://github.com/pylover/bddrest) inside it to run tests.


Add a file named `bdd.go` containing a function named `getMainHandler`:

```go
package mywebapp

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getMainHandler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld)
	return r
}
```

### Writing tests

A simple hello world test file: 

`tests/test_helloworld.py`

```python
from bddrest import when, status


def test_helloworld(story):
    with story('/'):
        assert status == 200
        assert response == 'Hello World'

```

### Running tests

```bash
bddgo test
```
