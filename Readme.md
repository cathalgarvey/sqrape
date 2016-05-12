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

### How?
#### Basics
Sqrape uses struct tags to figure out how to access and extract data. These
tags consist of two portions; a CSS selector, and a data extractor, separated
by a semicolon.. The former are an exercise for the reader and are well
documented. CSS selectors are passed to [`goquery`][ghgoquery], under the hood,
so consult docs there if in doubt.

One difference from `goquery`: Empty selectors are OK, and indicate "extract
data from the entire selection"; these are more commonly useful for embedded
structs or slices, where the passed data may be ready for extraction and require
no further CSS searching.

The second portion simply indicates what part or form of the selected data is
desired, and can take four forms, three of which are trivial:

* `text`: The text contents of matched data are returned.
* `html`: The HTML contents of matched data are returned
* `attr=<attribute name>`: Extract the value of an attribute on the matched selection.
* `obj`: This indicates a struct or array field that is parsed recursively.

Therefore, to extract the `data-foo` value from a div, use `csss:"div[data-foo];attr=data-foo"`: this selects any `div` with a `data-foo`
attribute, and extracts the value of that attribute.

To extract values other than strings, simply set the field type in your struct
to the desired type; this magic is handled by [`mapstructure`][ghmapstructure]!
So, if `data-foo` is a number, then the field the above tag annotates can be an
`int` or `int64`.

If your field is a singleton, then the first value will be extracted in the case
of attributes, and the concatenation of all values in the case of text or HTML.
If your field is a slice, then the values will be added iteratively from the
goquery selection.

If your field is a struct or slice of structs, then the extractor portion of
the tag should be `obj`, to indicate that parsing data from extracted structs
should be deferred to the embedded struct fields. See the Twitter example, above.

#### More Advanced: Optional Methods
Sometimes a datatype needs to be filled from multiple sources, or has fields
that should only be filled under certain other conditions, or should have
conditional or context-aware behaviour... for this, you can define optional
methods that alter Sqrape's behaviour and allow you to selectively fill fields,
or to perform post-processing or post-scrape data validation on your struct.

The methods supported so far include:

* `SqrapeFieldSelect(fieldName string, context...interface{}) (doField bool, cancelScrape error)`
* `SqrapePostFlight(context... interface{}) error`

The `context` argument in either case is a variadic list of arbitrary datatypes
which are passed by you to the entrypoint functions when operating a scrape.

So, for example, you could implement multi-page scraping by passing the current
URL to your scrape and defining a `SqrapeFieldSelect` method that fills fields
only for relevant URLs.

Or, you could perform data validation on your expected output
with a `SqrapePostFlight` method, either with hardcoded regex/validation or
by passing per-job regex or callbacks. Any error you raise in PostFlight will
be returned from the job to you.

### What's Supported?
Nested structs and array fields containing either basic values or struct values.
This means that, aside from map-fields, most stuff should just work. File an
Issue for cases that fail and I'll try to get it working.

Take a look at the test cases for an example of what works well. Feel free to
volunteer test cases.

### Credits Reel
Obviously, [`GoQuery`][ghgoquery] deserves a huge slice of the credit for this.

A lot of the magic behind field-filling is thanks to [`mapstructure`][ghmapstructure],
which handles "weakly typed" field-filling for structs.

There's a lot of reflective magic in this code; right now that's predictably
messy and due re-writing in pure `reflect` code. Meanwhile, thanks to
[`structs`][ghstructs] and [`reflections`][ghreflections] for tiding me over
this much of the project, by offering handy higher-level abstractions for `reflect`.

Reflection may give you the shivers; you're right, this code is potentially
explosive right now! Caveat emptor. However, the entry point functions do have
a blanket-recover deferred, so this code shouldn't panic, merely return an error
on panicky behaviour. **Please report any panic errors you encounter, to help me
make this more stable**.

[ghgoquery]: https://github.com/PuerkitoBio/goquery
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
