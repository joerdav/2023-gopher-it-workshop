package helloworld

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", uuid.NewString())
}
