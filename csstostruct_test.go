package sqrape

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

type Simple struct {
	Heading     string   `csss:"h1;text"`
	Paragraphs  []string `csss:"p;text"`
	ImageSrc    string   `csss:"img;attr=src"`
	ImageNumber int      `csss:"img;attr=data-some-number"`
}

type Complex struct {
	Heading    string   `csss:"h1;text"`
	Paragraphs []string `csss:"p;text"`
	ParaHTML   []string `csss:"p;html"` // Will not be filled.
	Paracount  int
	Image      struct {
		Src    string `csss:"img;attr=src"`
		Number int    `csss:"img;attr=data-some-number"`
	} `csss:"div;obj"`
}

// SqrapePostFlight fulfils the PostFlighter interface.
// It counts the number of scraped paragraphs.
func (cm *Complex) SqrapePostFlight(context ...interface{}) error {
	if len(context) == 0 {
		return fmt.Errorf("No context args passed to Complex testcase.")
	}
	if context[0].(string) != "foo" || context[1].(int) != 2 {
		return fmt.Errorf("Context args incorrect: %+v", context)
	}
	cm.Paracount = len(cm.Paragraphs)
	return nil
}

// SqrapeFieldSelect fulfiles the FieldSelecter interface.
// It makes the testcase skip ParaHTML for whatever reason.
// It errors if it doesn't receive the expected context.
func (cm *Complex) SqrapeFieldSelect(fieldName string, context ...interface{}) (bool, error) {
	if len(context) == 0 {
		return false, fmt.Errorf("No context args passed to Complex testcase.")
	}
	if context[0].(string) != "foo" || context[1].(int) != 2 {
		return false, fmt.Errorf("Context args incorrect: %+v", context)
	}
	if fieldName == "ParaHTML" {
		return false, nil
	}
	return true, nil
}

type Peep struct {
	Name    string `csss:";text"`
	Website string `csss:"a;attr=href"`
}

type MoreComplex struct {
	Heading string `csss:"h1;text"`
	Entries []Peep `csss:"li.name;obj"`
}

type Tweet struct {
	Author  string `csss:"div.original-tweet;attr=data-screen-name"`
	TweetID int64  `csss:"div.original-tweet;attr=data-tweet-id"`
	Content string `csss:"p.js-tweet-text;text"`
}

type TwitterProfile struct {
	Tweets []Tweet `csss:"li.js-stream-item;obj"`
}

var (
	twitterProfileExpected = []Tweet{
		Tweet{Author: "onetruecathal", TweetID: 637274908041605120, Content: "Join us Thursday next at 6pm for @FormaBiolabs' official (but relaxed!) opening party: http://www.eventbrite.ie/e/forma-labs-opening-ceremony-tickets-18264954972 … #science #cork #scicomm"},
		Tweet{Author: "csoghoian", TweetID: 642042919466180608, Content: "Shorter FBI Director: I love strong encryption. I want strong encryption to protect my data. I just don't want you to have strong encryption"},
		Tweet{Author: "onetruecathal", TweetID: 642045936995336193, Content: "Awesome news for #celiac peeps: soon you could be enjoying 4-5 slices of decent bread! https://twitter.com/franknfoode/status/642042436240429057 …"},
	}
	moreComplexHTML = `<html>
  <head></head>
  <body>
  <h1>Blogroll!</h1>
  <div class="peeps">
  <ul>
  <li class="name" data-favourite-colour="blue"><a href="http://john.doe">John Doe</a></li>
  <li class="name" data-age="50"><a href="http://jane.doe">Jane Doe</a></li>
  </ul>
  </div>
  </body>
  </html>`

	simpleHTML = `<html><head></head><body>
  <h1>This is a title for my super simple blogpost</h1>
  <p>I have experienced lots in my time but nothing as awesome as scraping</p>
  <p>Let me show you more..</p>
  <div><img data-some-number="2" src="https://face.ly/totesawes" /></div>
  </body></html>`

	testGetValueStr = cssStructer{
		targetStruct: &Simple{},
		collectedFieldValues: map[string]interface{}{
			"Heading":     "This is a heading",
			"Paragraphs":  []string{"This is a paragraph", "So is this"},
			"ImageSrc":    "https://foo.bar",
			"ImageNumber": 2,
		},
	}

	testGetComplexValueStr = cssStructer{
		targetStruct: &Complex{},
		collectedFieldValues: map[string]interface{}{
			"Heading":    "This is a heading",
			"Paragraphs": []string{"This is a paragraph", "So is this"},
			"Image": struct {
				Src    string
				Number int
			}{"https://foo.bar", 2},
		},
	}
)

func TestGetValue(t *testing.T) {
	err := testGetValueStr.GetValue()
	assert.Nil(t, err)
	targ := testGetValueStr.targetStruct.(*Simple)
	assert.Equal(t, "This is a heading", targ.Heading)
	assert.Equal(t, []string{"This is a paragraph", "So is this"}, targ.Paragraphs)
	assert.Equal(t, "https://foo.bar", targ.ImageSrc)
	assert.Equal(t, 2, targ.ImageNumber)
}

func TestSimpleExtraction(t *testing.T) {
	s := new(Simple)
	simpleResp, err := goquery.NewDocumentFromReader(strings.NewReader(simpleHTML))
	assert.Nil(t, err)
	err = extractByTags(simpleResp.Selection, s)
	assert.Nil(t, err)
	assert.Equal(t, "This is a title for my super simple blogpost", s.Heading)
	assert.Equal(t, []string{
		"I have experienced lots in my time but nothing as awesome as scraping",
		"Let me show you more..",
	}, s.Paragraphs)
	assert.Equal(t, "https://face.ly/totesawes", s.ImageSrc)
	assert.Equal(t, 2, s.ImageNumber)
}

func TestComplexExtraction(t *testing.T) {
	c := new(Complex)
	simpleResp, err := goquery.NewDocumentFromReader(strings.NewReader(simpleHTML))
	assert.Nil(t, err)
	err = extractByTags(simpleResp.Selection, c, "foo", 2)
	assert.Nil(t, err)
	assert.Equal(t, "This is a title for my super simple blogpost", c.Heading)
	assert.Equal(t, []string{
		"I have experienced lots in my time but nothing as awesome as scraping",
		"Let me show you more..",
	}, c.Paragraphs)
	assert.Equal(t, "https://face.ly/totesawes", c.Image.Src)
	assert.Equal(t, 2, c.Image.Number)
}

func TestMoreComplexExtraction(t *testing.T) {
	mc := new(MoreComplex)
	moreComplexResp, err := goquery.NewDocumentFromReader(strings.NewReader(moreComplexHTML))
	assert.Nil(t, err)
	err = extractByTags(moreComplexResp.Selection, mc)
	assert.Nil(t, err)
	assert.Equal(t, "Blogroll!", mc.Heading)
	assert.Equal(t, []Peep{
		Peep{"John Doe", "http://john.doe"},
		Peep{"Jane Doe", "http://jane.doe"},
	}, mc.Entries)
	colour := new(struct {
		Favourite []string `csss:"li[data-favourite-colour];attr=data-favourite-colour"`
	})
	err = extractByTags(moreComplexResp.Selection, colour)
	assert.Nil(t, err)
	assert.Equal(t, []string{"blue"}, colour.Favourite)
}

func TestParseTag(t *testing.T) {
	// parseTag(tag string) (selector, valueType, attrName string, err error)
	sel, val, att, err := parseTag("h1>.foo;text")
	assertTagBits(t, "h1>.foo", sel, "text", val, "", att, err)

	sel, val, att, err = parseTag("p;text")
	assertTagBits(t, "p", sel, "text", val, "", att, err)

	sel, val, att, err = parseTag("img;attr=src")
	assertTagBits(t, "img", sel, "attr", val, "src", att, err)
}

func TestTweetParse(t *testing.T) {
	tw := new(Tweet)
	resp, err := goquery.NewDocumentFromReader(strings.NewReader(tweetRawHTML))
	assert.Nil(t, err)
	err = extractByTags(resp.Selection, tw)
	assert.Nil(t, err)
	assert.Equal(t, tw.Author, "nuts4ag")
	assert.Equal(t, tw.Content, "@HomeDepot best talk of the UCDavis #bee and #Neonic conference was Ray Jarvis. Doing good work with your sustainability program")
	assert.EqualValues(t, tw.TweetID, 641830475854536704)
}

func TestTwitterStreamParse(t *testing.T) {
	tp := new(TwitterProfile)
	resp2, err := goquery.NewDocumentFromReader(strings.NewReader(streamRawHTML))
	assert.Nil(t, err)
	err = extractByTags(resp2.Selection, tp)
	assert.Nil(t, err)
	assert.Equal(t, twitterProfileExpected, tp.Tweets)
}

func assertTagBits(t *testing.T, s, sel, v, val, a, att string, err error) {
	assert.Nil(t, err)
	assert.Equal(t, s, sel)
	assert.Equal(t, v, val)
	assert.Equal(t, a, att)
}
