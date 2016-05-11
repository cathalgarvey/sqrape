## Sqrape - Simple Query Scraping with CSS and Go Reflection
by Cathal Garvey, ©2016, Released under the GNU AGPLv3

### What
When scraping web content, one usually hopes that the content is laid out
logically, and that proper or at least consistent web annotation exists.
This means well-nested HTML, appropriate use of tags, descriptive CSS classes
and unique CSS IDs. Ideally it also means that a given CSS selector will
yield a consistent datatype, also.

In such cases, it's possible to define exactly what you want using only CSS
and a type. For a scraping job, then, it would be ideal to just make a struct
defining the content you want, and to scrape a page directly from that, right?

So, something like this:

```go
type Tweet struct {
	Author  string `csss:"div.original-tweet;attr=data-screen-name"`
	TweetID int64  `csss:"div.original-tweet;attr=data-tweet-id"`
	Content string `csss:"p.js-tweet-text;text"`
}

type TwitterProfile struct {
	Tweets []Tweet `csss:"li.js-stream-item;obj"`
}

func main() {
	resp, _ := http.Get("https://twitter.com/onetruecathal")
	tp := new(TwitterProfile)
	csstostruct.ExtractHTMLReader(resp.Body, tp)
	for _, tweet := range tp.Tweets {
		fmt.Printf("@%s: %s\n", tweet.Author, tweet.Content)
	}
}
```

..well that's Sqrape. In fact, see `examples/tweetgrab.go` for the above
as a CLI tool.

Note; that struct tag is `csss`, not `css`. It's "css selector", because I
didn't want to clobber any preexisting `css` struct tag libs that may exist.

### What's Supported?
Nested structs and array fields containing either basic values or struct values.
This means that, aside from map-fields, most stuff should just work. File an
Issue for cases that fail and I'll try to get it working.

Take a look at the test cases for an example of what works well. Feel free to
volunteer test cases.

A lot of the magic behind this is thanks to [`mapstructure`][ghmapstructure],
which handles "weakly typed" field-filling for structs. Other parts required
reflective magic; right now that part is a bit messy and due re-writing in pure
`reflect` code, but thanks to [`structs`][ghstructs] and
[`reflections`][ghreflections] for tiding me over this much of the project.

[ghmapstructure]: https://github.com/mitchellh/mapstructure
[ghstructs]: https://github.com/fatih/structs
[ghreflections]: https://github.com/oleiade/reflections

### Why?
I scrape content *a lot*. Weekly, sometimes daily, as part of my job or for
personal research. Web scraping is just another way of consuming web content!
I do most of my scraping in the IPython shell, but for something "important"
I'll write something more permanent and use that whenever the need arises.

For this, one typically uses a scraping framework. But, permanence has disadvantages.
If your scraping framework requires a lot of overhead for very basic tasks, then
that means the maintenance burden when things change is also high.

I wanted something where creating and maintaining a scraper could be trivial,
a matter of just defining the data I want and mapping it to the HTML. If or when
the HTML changes, then I only need to change the datatypes or CSS rules and get
back to *using* the data.
