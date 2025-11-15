package handlers

import (
	"fmt"
	"net/http"
)

type pageController struct{}

func PageController() *pageController {
	return &pageController{}
}

const indexPageHtml = `
<a href="/oauth/line/redirect">LINE登入</a>
`

func (c *pageController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, indexPageHtml)
}
