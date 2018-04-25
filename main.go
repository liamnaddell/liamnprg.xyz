package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println(string(ctx.Path()))
	f, err := os.Stat("static" + string(ctx.Path()))
	var s string
	var html bool
	if err == nil {
		s, html = getEnding(f)
	}
	if err == nil && f.IsDir() {
		ctx.SetContentType("text/html")
		ctx.Write(itmpl("static" + string(ctx.Path()) + "index.html"))
	} else if err == nil {
		//fmt.Println("NOT")
		ctx.SetContentType(s)
		if html {
			ctx.Write(itmpl("static/" + string(ctx.Path())))
		} else {
			fasthttp.ServeFile(ctx, "static"+string(ctx.Path()))
		}
	} else {
		ctx.NotFound()
		fasthttp.ServeFile(ctx, "static/404.html")
	}
}

func redirectHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println("REDIRECT")
	ctx.Redirect("https://liamnprg.xyz"+string(ctx.Path()), 302)
}

func itmpl(s string) []byte {
	article, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("err")
	}
	t, err := template.New("webpage").Parse(fmt.Sprintf("%s", article))
	if err != nil {
		fmt.Println("err")
	}
	data := struct {
		Title string
	}{
		Title: "Welcome to liamnprg.xyz",
	}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err = t.Execute(writer, data)
	writer.Flush()
	return []byte(b.String())
}
func main() {
	fmt.Println("up")
	//TODO: Make it so that it redirects http to https unless https is unavailable
	var tls = true
	var err error
	go func(err *error) {
		*err = fasthttp.ListenAndServeTLS(":443", "/etc/letsencrypt/live/liamnprg.xyz/fullchain.pem", "/etc/letsencrypt/live/liamnprg.xyz/privkey.pem", fastHTTPHandler)
	}(&err)
	fmt.Println("HELLO", tls)
	if err != nil {
		fmt.Println(err)
		tls = false
	}
	fmt.Println(tls)
	if tls {

		err = fasthttp.ListenAndServe(":80", redirectHandler)
	} else {
		err = fasthttp.ListenAndServe(":80", fastHTTPHandler)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func getEnding(f os.FileInfo) (string, bool) {
	n := strings.Split(f.Name(), ".")
	var zerolen /*threelen,*/, html bool
	var /*mimetype,*/ end string
	if len(n) == 1 {
		zerolen = true
	} else if len(n) > 2 {
		/*threelen = true*/
	}
	if zerolen {
		return "text/plain", true
	} else if n[1] == "html" {
		html = true
		end = "html"
	} else if n[1] == "css" {
		end = "css"
	}
	return "text/" + end, html
}
