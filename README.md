# Meowvie

Just some random project that attempt to ease and simplify getting movie or
anime download urls. This project divided into 3 repositories

- Backend (this)
- [Frontent](https://github.com/falentio/meowvie-frontend)
- [Crawler](https://github.com/falentio/meowvie-crawler)

## Tech Stacks

- Golang as Backend

  I choose golang beacuse I need good performance, and good search engine
  library (prevent to use elasticsearch or equivalent)

- Deno as Frontend

  It just easy to develop, easy to deploy

- Nodejs as Crawler

  Because my Python skill is poor, so I dont think any other than Nodejs has
  good ecosystem for web scraping.

## Performance

No benchmark was done, but I was tested it with 30 concurrent queries on 256 MB
ram, and so far so good, the slowest response time is arround `500ms`

```
❯ oha "https://meowvie.fly.dev/movie/search?q=borto" -z 10s -c 30 -q 30
Summary:
  Success rate: 1.0000
  Total:        10.0014 secs
  Slowest:      0.5025 secs
  Fastest:      0.0675 secs
  Average:      0.1381 secs
  Requests/sec: 29.7959

  Total data:   0 B
  Size/request: 0 B
  Size/sec:     0 B
```

## Movie/Anime sources

- Otakudesu
- Melongmovie (Partially)
