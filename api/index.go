package api

import (
	"jokyo3/probingai/api/helper"
	"jokyo3/probingai/common"
	"net/http"
	"strings"
//	"os"
//	"net/url" 
)

//var bingURL = os.Getenv("BING_PROXY_DM")
//	if bingURL == "" {
//		bingURL = common.BING_URL.String()
//	}
//	proxyurl, _ := url.Parse(bingURL)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, common.PROXY_WEB_PAGE_PATH, http.StatusFound)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/turing/captcha/challenge") {
		if !helper.CheckAuth(r) {
			helper.UnauthorizedResult(w)
			return
		}
		common.NewSingleHostReverseProxy(common.BingURL).ServeHTTP(w, r)
	} else {
		common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)
		}
}

