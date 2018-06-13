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
	"time"
)

var views = 0

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
		views = views + 1
		fmt.Println("views: ", views)
	} else if err == nil {
		//fmt.Println("NOT")
		ctx.SetContentType(s)
		if html {
			ctx.Write(itmpl("static/" + string(ctx.Path())))
			views = views + 1
			fmt.Println("views: ", views)
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

//very insecure template, thus marked with an i infront of it.
//I will use a more secure version for user input, but this is not user input, and can be trusted.
func itmpl(s string) []byte {
	article, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("err")
	}
	t, err := template.New("webpage").Parse(fmt.Sprintf("%s", article))
	if err != nil {
		fmt.Println("err")
	}
	//will load template from file later
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
	//tell the user when the code is done compiling
	fmt.Println("up")
	var tls = true
	var err error
	//async tls server
	go func(ourerr *error, isTls *bool) {
		var rr = fasthttp.ListenAndServeTLS(":443", "/etc/letsencrypt/live/liamnprg.xyz/fullchain.pem", "/etc/letsencrypt/live/liamnprg.xyz/privkey.pem", fastHTTPHandler)
		fmt.Println(rr)
		*ourerr = rr
		fmt.Println(*ourerr)
		*isTls = false
	}(&err, &tls)
	fmt.Println(err)
	time.Sleep(1*time.Second)
	if err != nil {
		fmt.Println(err)
		//if there is an error, tls will be disabled.
		tls = false
	}
	fmt.Println("Is tls on? ", tls)
	//no error, so redirect http to tls
	//else only serve http
	if tls {

		err = fasthttp.ListenAndServe(":80", redirectHandler)
	} else {
		err = fasthttp.ListenAndServe(":80", fastHTTPHandler)
	}
	if err != nil {
		//if http is also broken, print error
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
		//If filename looks like /text or , then it is plain text
		return "text/plain", true
	} else if n[1] == "html" {
		html = true
		end = "html"
	} else if n[1] == "css" {
		end = "css"
	}
	return "text/" + end, html
}
