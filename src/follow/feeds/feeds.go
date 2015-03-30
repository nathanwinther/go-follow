package feeds

import (
    "database/sql"
    "html/template"
    "follow/dao"
)

type Feed struct {
    Id int
    Title template.HTML
    Url template.HTML
    Feed string
}

type helper struct {
    feeds []*Feed
    index int
    size int
    growth int
}

func Load() ([]*Feed, error) {
    size := 100
    h := &helper {
        make([]*Feed, size),
        0,
        size,
        size,
    }

    q := `
        SELECT
            id
            , title
            , url
            , feed
        FROM feed
        ORDER BY title;
    `

    err := dao.Query(q, nil, h.Add)
    if err != nil {
        return nil, err
    }

    return h.List(), nil
}

func (h *helper) List() []*Feed {
    return h.feeds[:h.index]
}

func (h *helper) Add(rows *sql.Rows) error {
    if h.index == h.size {
        h.size = h.size + h.growth
        _h := make([]*Feed, h.size)
        copy(_h, h.feeds)
        h.feeds = _h
    }

    f := new(Feed)
    var title string
    var url string

    err := rows.Scan(&f.Id, &title, &url, &f.Feed)
    if err != nil {
        return err
    }

    f.Title = template.HTML(title)
    f.Url = template.HTML(url)

    h.feeds[h.index] = f

    h.index = h.index + 1

    return nil
}

