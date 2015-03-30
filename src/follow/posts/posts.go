package posts

import (
    "database/sql"
    "html/template"
    "strconv"
    "time"
    "follow/config"
    "follow/dao"
)

type Post struct {
    Site template.HTML
    Title template.HTML
    Url template.HTML
    Published *Date
}

type Date struct {
    Unix int64
    D int
    DD string
    DDD string
    DDDD string
    M int
    MM string
    MMM string
    MMMM string
    YY string
    YYYY int
    HR int
    HR24 string
    MIN string
    AMPM string
}

type helper struct {
    posts []*Post
    index int
    size int
    growth int
}

func Add(feedId int, title string, url string, published int64) error {
    q := `
        INSERT INTO post VALUES(
            NULL
            , ?
            , ?
            , ?
            , ?
        );
    `

    params := []interface{} {
        feedId,
        title,
        url,
        published,
    }

    _, err := dao.Exec(q, params)

    return err
}

func Load() ([]*Post, error) {
    size := 100
    h := &helper {
        make([]*Post, size),
        0,
        size,
        size,
    }

    q := `
        SELECT
            f.title site
            , p.title title
            , p.url url
            , p.published published
        FROM post p, feed f
        WHERE p.feed_id = f.id
        ORDER BY p.published DESC
        LIMIT 100;
    `

    err := dao.Query(q, nil, h.Add)
    if err != nil {
        return nil, err
    }

    return h.List(), nil
}

func NewDate(unix int64) *Date {
    loc, _ := time.LoadLocation(config.Get("timezone"))

    t := time.Unix(unix, 0).In(loc)

    d := new(Date)
    d.Unix = unix
    d.D, _ = strconv.Atoi(t.Format("02"))
    d.DD = t.Format("02")
    d.DDD = t.Format("Mon")
    d.DDDD = t.Format("Monday")
    d.M, _ = strconv.Atoi(t.Format("01"))
    d.MM = t.Format("01")
    d.MMM = t.Format("Jan")
    d.MMMM = t.Format("January")
    d.YY = t.Format("06")
    d.YYYY, _ = strconv.Atoi(t.Format("2006"))
    d.HR, _ = strconv.Atoi(t.Format("03"))
    d.HR24 = t.Format("15")
    d.MIN = t.Format("04")
    d.AMPM = t.Format("PM")

    return d
}

func Reset() error {
    _, err := dao.Exec("DELETE FROM post;", nil)
    return err
}

func (h *helper) List() []*Post {
    return h.posts[:h.index]
}

func (h *helper) Add(rows *sql.Rows) error {
    if h.index == h.size {
        h.size = h.size + h.growth
        _h := make([]*Post, h.size)
        copy(_h, h.posts)
        h.posts = _h
    }

    var site string
    var title string
    var url string
    var published int64

    err := rows.Scan(&site, &title, &url, &published)
    if err != nil {
        return err
    }

    h.posts[h.index] = &Post {
        template.HTML(site),
        template.HTML(title),
        template.HTML(url),
        NewDate(published),
    }

    h.index = h.index + 1

    return nil
}

