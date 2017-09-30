package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

// IssuesSearchResult 问题搜索结果
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	//templ1()
	//templ2()
	templ3()
}

func templ1() {
	const templ = `{{.TotalCount}} issues:`
	report, err := template.New("report").Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	var result = IssuesSearchResult{
		TotalCount: 3,
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func templ2() {
	const templ = `{{.TotalCount}} issues:`
	var report = template.Must(template.New("issuelist").Parse(templ))
	var result = IssuesSearchResult{
		TotalCount: 3,
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func templ3() {
	const templ = `{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.10s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}`
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	var result = IssuesSearchResult{
		TotalCount: 3,
		Items: []*Issue{
			&Issue{
				Number:    10,
				User:      &User{Login: "xiaoming"},
				Title:     "习主席带领我们强军  第一集逐梦  第二集铸魂",
				CreatedAt: time.Date(2017, 9, 28, 20, 34, 58, 651387237, time.UTC),
			},
			&Issue{
				Number:    20,
				User:      &User{Login: "xiaohong"},
				Title:     "欢度国庆  中秋晚会  “五个一工程”获奖作品巡礼",
				CreatedAt: time.Date(2017, 9, 29, 20, 34, 58, 651387237, time.UTC),
			},
			&Issue{
				Number:    20,
				User:      &User{Login: "Jack"},
				Title:     "喜迎十九大  理上网来  砥砺奋进的五年  《榜样》",
				CreatedAt: time.Date(2017, 9, 30, 20, 34, 58, 651387237, time.UTC),
			},
		},
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
