## Yomu

notes
- i will be solely focusing on the backend for now. want to make it scalable.
- use redis for caching frequently requested manga metadata. seperate api and scraper
- into microservices with grpc. probably use queue for the scraper to avoid rate limiting.

- then, we can focus on the frontend. most of the functinoality already done. have to add 
- library, history, and progress. ok, these are hard. make the ui responsive with dark/light 
- mode enabled. clean the react code: hooks, api, lib, etc.

schema

* postgres:

users {
    id: int [pk]
    first_name: string [not null]
    last_name?: string 
    username?: string
    created_at: datetime
}

library {
    id: int
    user_id: int [fk]
    manga_url: string
    added_at: datetime
    unique(user_id, manga_url)
}

history {
    id: int
    user_id: int [fk]
    manga_url: string 
    read_at: datetime
    index(user_id, manga_url)
}

progress {
    id: int
    user_id: int [fk]
    manga_url: string
    chapter: int
    page: int
    updated_at: datetime
    index(user_id, manga_url)
}

* redis:

manga {
    - key:
        manga_url: string 
    - val:
        title: string
        author: string
        genre: []string
        chapters: {name: string -> url: string}
        cover_url: string
        description: string
}
