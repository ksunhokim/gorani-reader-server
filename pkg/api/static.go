package api

import (
	"fmt"
	"net/http"
)

const indexHTML = `
<html>
	<body>
		hello world
	</body>
</html>
`

var indexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, indexHTML)
})
