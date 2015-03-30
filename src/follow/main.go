package main

import (
    "fmt"
    "html/template"
    "os"
    "path/filepath"
    "github.com/nathanwinther/go-feedparser"
    "follow/config"
    "follow/feeds"
    "follow/posts"
)

func main() {
    switch len(os.Args) {
        case 1:
            break
        case 2:
            err := os.Chdir(os.Args[1])
            if err != nil {
                panic(err)
            }
            break
        default:
            return
    }

    err := config.Load("./data/data.db")
    if err != nil {
        panic(err)
    }

    _feeds, err := feeds.Load()
    if err != nil {
        panic(err)
    }

    posts.Reset()

    for _, f := range _feeds {
        fmt.Printf("=> %s\n", f.Title)
        items, err := feedparser.Load(f.Feed)
        if err != nil {
            fmt.Println(err.Error())
            continue
        }
        for _, item := range items {
            err = posts.Add(f.Id, item.Title, item.Url, item.Published)
            if err != nil {
                fmt.Println(err.Error())
                continue
            }
        }
    }

    _posts, err := posts.Load()
    if err != nil {
        panic(err)
    }

    m := map[string] interface{} {
        "Feeds": _feeds,
        "Posts": _posts,
    }

    tmpl, err := template.ParseGlob(filepath.Join(config.Get("templates"), "*.*"))
    if err != nil {
        panic(err)
    }

    f, err := os.Create(filepath.Join(config.Get("output"), "index.html"))
    if err != nil {
        panic(err)
    }

    err = tmpl.ExecuteTemplate(f, "index.html", m)
    if err != nil {
        panic(err)
    }

    fmt.Println("OK\n")
}

