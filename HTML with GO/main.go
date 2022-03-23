package main

import(
	"fmt"
	"html/template"
	"strings"
	"io/ioutil"
	"net/http"
)

type marquee1 struct{
	ScrollAmount string
	ScrollDelay string
	BgColor string
	Direction string
	Loop string
}

func main(){
	fmt.Println("We are up and running!")
	http.HandleFunc("/", OurForm)
	http.ListenAndServe(":8080", nil)
}

func OurForm(w http.ResponseWriter, r *http.Request){
	TplTxt,_ := ioutil.ReadFile("form.html")
	Tpl := template.Must(template.New("form.html").Parse(string(TplTxt)))

	details := marquee1{
		ScrollAmount: r.FormValue("sm"),
		ScrollDelay: r.FormValue("sd"),
		BgColor: r.FormValue("bc"),
		Direction: r.FormValue("direction"),
		Loop: r.FormValue("loop"),
	}
switch {
case strings.EqualFold(details.Direction, "right"):
	fmt.Fprintf(w, "<div><h1 class=\"hello\">Hello there, &hearts;</h1> \n<marquee scrollamount=\"%v\" scrolldelay=\"%v\" bgcolor=\"%v\" direction=\"%v\" loop=\"%v\"><h1>We are heading right!</h1></marquee>\n</div>", details.ScrollAmount, details.ScrollDelay, details.BgColor, details.Direction, details.Loop)
case strings.EqualFold(details.Direction, "left"):
	fmt.Fprintf(w, "<div><h1 class=\"hello\">Hello there, &hearts;</h1> \n<marquee scrollamount=\"%v\" scrolldelay=\"%v\" bgcolor=\"%v\" direction=\"%v\" loop=\"%v\"><h1>We are heading left!</h1></marquee>\n</div>", details.ScrollAmount, details.ScrollDelay, details.BgColor, details.Direction, details.Loop)
case strings.EqualFold(details.Direction, "up"):
	fmt.Fprintf(w, "<div><h1 class=\"hello\">Hello there, &hearts;</h1> \n<marquee scrollamount=\"%v\" scrolldelay=\"%v\" bgcolor=\"%v\" direction=\"%v\" loop=\"%v\"><h1>We are heading Up!</h1></marquee>\n</div>", details.ScrollAmount, details.ScrollDelay, details.BgColor, details.Direction, details.Loop)
case strings.EqualFold(details.Direction, "down"):
	fmt.Fprintf(w, "<div><h1 class=\"hello\">Hello there, &hearts;</h1> \n<marquee scrollamount=\"%v\" scrolldelay=\"%v\" bgcolor=\"%v\" direction=\"%v\" loop=\"%v\"><h1>We are heading Down!</h1></marquee>\n</div>", details.ScrollAmount, details.ScrollDelay, details.BgColor, details.Direction, details.Loop)

default:
	fmt.Fprintf(w, "<div><h1 class=\"hello\">Hello there, &hearts;</h1> \n<marquee scrollamount=\"%v\" scrolldelay=\"%v\" bgcolor=\"%v\" direction=\"%v\" loop=\"%v\"><h1>Where are we heading?</h1></marquee>\n</div>", details.ScrollAmount, details.ScrollDelay, details.BgColor, details.Direction, details.Loop)
}
Tpl.Execute(w, struct{Success bool}{true})
fmt.Println("Values are: %v %v %v %v %v", details.ScrollAmount, details.ScrollDelay, details.BgColor, details.Direction, details.Loop)
}