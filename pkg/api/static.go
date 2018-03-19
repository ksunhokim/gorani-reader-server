package api

import (
	"fmt"
	"net/http"
)

const indexHTML = `
<html>
  <head>
    <title>asdf</title>
  </head>
  <body>
	<div id="root"></div>
	<script src="http://localhost:3333/bundle.js"></script>
  </body>
</html>
`

var indexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, indexHTML)
})
