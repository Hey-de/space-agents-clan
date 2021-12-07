package main

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-session/session"
)

func userProcessor(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/note" {
		err := r.ParseForm()
		if err != nil {
			Error500(err, w, r)
			return
		}
	} else {
		err := r.ParseMultipartForm(2080)
		if err != nil {
			Error500(err, w, r)
			return
		}
	}
	switch r.URL.Path {
	case "/users/login":
		if r.Method != "POST" {
			ErrorBadRequest(w, r)
			return
		}
		if _, ok := r.PostForm["doGo"]; !ok {
			ErrorBadRequest(w, r)
			return
		}
		fmt.Println(r.PostForm["login"][0])
		login := r.PostForm["login"][0]
		hasLogin, err := findUser(login)
		fmt.Println(hasLogin)
		if err != nil {
			Error500(err, w, r)
			return
		}
		if !hasLogin {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		fmt.Printf("%x\n", sha1.Sum([]byte(r.PostForm["password"][0])))
		password := r.PostForm["password"][0]
		ok, err := checkPassword(login, password)
		if err != nil {
			Error500(err, w, r)
			return
		}
		if ok {
			store, err := session.Start(context.Background(), w, r)
			if err != nil {
				Error500(err, w, r)
				return
			}
			store.Set("login", login)
			store.Set("password", sha1.Sum([]byte(password)))
			err = store.Save()
			if err != nil {
				Error500(err, w, r)
				return
			}
			http.Redirect(w, r, "/", http.StatusFound)
			return
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	case "/users/singup":
		if r.Method != "POST" {
			fmt.Println("Singup")
			ErrorBadRequest(w, r)
			return
		}
		fmt.Println(r.PostForm)
		// if _, ok := r.PostForm["doGo"]; !ok {
		// 	fmt.Println("Singup")
		// 	ErrorBadRequest(w, r)
		// 	return
		// }
		fmt.Println(r.PostForm["login"][0])
		login := r.PostForm["login"][0]
		hasLogin, err := findUser(login)
		fmt.Println(hasLogin)
		if err != nil {
			Error500(err, w, r)
			return
		}
		if hasLogin {
			http.Redirect(w, r, "/singup", http.StatusSeeOther)
			return
		}
		store, err := session.Start(context.Background(), w, r)
		if err != nil {
			Error500(err, w, r)
			return
		}
		password := r.PostForm["password"][0]
		err = register(login, password)
		if err != nil {
			Error500(err, w, r)
			return
		}
		os.Mkdir(login, 0777)
		store.Set("login", login)
		store.Set("password", sha1.Sum([]byte(password)))
		err = store.Save()
		if err != nil {
			Error500(err, w, r)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	case "/users/registered":
		if r.Method != "GET" {
			ErrorBadRequest(w, r)
			return
		}
		if login, ok := r.Form["name"]; ok {
			ok, err := findUser(login[0])
			if err != nil {
				Error500(err, w, r)
				return
			}
			if ok {
				w.Write([]byte("yes"))
			} else {
				w.Write([]byte("no"))
			}
		} else {
			ErrorBadRequest(w, r)
		}
	case "/users/note":
		if r.Method != "POST" {
			fmt.Println("Singup")
			ErrorBadRequest(w, r)
			return
		}
		fmt.Println(r.PostForm)
		// if _, ok := r.PostForm["doGo"]; !ok {
		// 	fmt.Println("Singup")
		// 	ErrorBadRequest(w, r)
		// 	return
		// }
		store, err := session.Start(context.Background(), w, r)
		if err != nil {
			Error500(err, w, r)
			return
		}
		fmt.Println(r.PostForm["title"][0])
		if login, ok := store.Get("login"); ok {
			if password, ok := store.Get("password"); ok {
				title := r.PostForm["title"][0]
				textarea := r.PostForm["content"][0]
				files := r.MultipartForm.File["pictures"]
				var fileNames []string
				for _, f := range files {
					file, _ := f.Open()
					fileNames = append(fileNames, f.Filename)
					tmpfile, err := os.Create("./userdata/" + login.(string) + "/" + f.Filename)
					if err != nil {
						Error500(err, w, r)
					}
					io.Copy(tmpfile, file)
					tmpfile.Close()
					file.Close()
				}
				err = createPost(login.(string), password, title, textarea, fileNames)
				if err != nil {
					Error500(err, w, r)
					return
				}
			} else {
				Error500(errors.New("error when gotting password"), w, r)
				return
			}
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	default:
		Error404(w, r)
	}
}