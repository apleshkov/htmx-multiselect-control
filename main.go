package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"htmxtest/musel"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func renderTemplate(wr io.Writer, filename string, data any) error {
	tpl := template.Must(template.ParseFiles(filename))
	return tpl.Execute(wr, data)
}

type randomUser struct {
	Email string `json:"email"`
	Name  struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
}

type randomUserData struct {
	Results []randomUser `json:"results"`
}

func readUserData(filename string) (*randomUserData, error) {
	if bytes, err := os.ReadFile(filename); err != nil {
		return nil, err
	} else {
		data := new(randomUserData)
		err := json.Unmarshal(bytes, data)
		return data, err
	}
}

type indexPageData struct {
	Submit        bool
	SelectedUsers string
	UsersControl  template.HTML
}

func newUsersControl(keys []string) *musel.Control {
	return &musel.Control{
		Name:         "users",
		SearchURL:    "/user-search",
		UpdateURL:    "/users-control",
		SelectedKeys: keys,
		Placeholder:  "Search users...",
	}
}

func main() {
	userData, err := readUserData("./users.json")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		path := req.URL.String()
		log.Println("[http.request]", req.Method, path)
		if path == "/" {
			pageData := indexPageData{}
			if req.Method == http.MethodPost {
				req.ParseForm()
				pageData.SelectedUsers = req.Form.Get("users")
				pageData.Submit = true
			}
			uc := newUsersControl(
				musel.ControlSelectedKeysFromString(pageData.SelectedUsers),
			)
			buf := new(strings.Builder)
			renderTemplate(buf, "./b-musel-control.html", uc)
			pageData.UsersControl = template.HTML(buf.String())
			renderTemplate(wr, "./index.html", pageData)
		} else if strings.HasPrefix(path, "/user-search") {
			req.ParseForm()
			sq := strings.ToLower(req.Form.Get("users-query"))
			selected := req.Form.Get("users")
			opts := musel.Options{
				SearchQuery: sq,
				ControlName: "users",
				SelectURL:   "/users-control",
				EmptyText:   "No results",
			}
			if len(sq) > 0 {
				all := userData.Results
				filtered := make([]musel.Option, 0, len(all))
				for _, entry := range all {
					if strings.Contains(selected, entry.Email) {
						continue
					}
					name := entry.Name
					lwf := strings.ToLower(name.First)
					lwl := strings.ToLower(name.Last)
					if strings.HasPrefix(lwf, sq) || strings.HasPrefix(lwl, sq) {
						o := musel.Option{
							Key:   entry.Email,
							Title: name.First + " " + name.Last,
						}
						filtered = append(filtered, o)
						if len(filtered) >= 5 {
							break
						}
					}
				}
				if len(filtered) > 0 {
					opts.List = filtered
				}
				renderTemplate(wr, "./b-musel-options.html", opts)
				time.Sleep(time.Millisecond * 400)
			} else {
				renderTemplate(wr, "./b-musel-options.html", opts)
			}
		} else if path == "/users-control" && req.Method == http.MethodPost {
			req.ParseForm()
			form := req.Form
			uc := newUsersControl(
				musel.ControlSelectedKeysFromString(form.Get("users")),
			)
			action := form.Get("action")
			if action == "select" {
				key := form.Get("key")
				uc.SelectedKeys = append(uc.SelectedKeys, key)
			} else if action == "remove" {
				key := form.Get("key")
				uc.RemoveKey(key)
			}
			renderTemplate(wr, "./b-musel-control.html", uc)
		} else {
			http.NotFound(wr, req)
		}
	})
	port := ":1234"
	fmt.Println("Now listening: http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
