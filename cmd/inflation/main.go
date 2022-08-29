package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"sort"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly/v2"
	"github.com/jszwec/csvutil"

	"git.sr.ht/~mendelmaleh/amazon"
)

const (
	baseurl  = "https://amazon.com/dp/"
	selector = "#corePriceDisplay_desktop_feature_div span.a-price.reinventPricePriceToPayMargin span.a-offscreen"
)

func fmtdate(t time.Time) string {
	return t.Format("2006-01-02")
}

func fmtcents(c int, sign bool) string {
	whole := c / 100
	cents := c % 100

	if cents < 0 {
		cents *= -1
	}

	if sign {
		return fmt.Sprintf("%+d.%.2d", whole, cents)
	}

	return fmt.Sprintf("%d.%.2d", whole, cents)
}

func Currency(s string) (c amazon.Currency) {
	c.UnmarshalText([]byte(s))
	return
}

type Item struct {
	amazon.Item

	Done     bool
	NewPrice amazon.Currency
}

type Result struct {
	URL *url.URL

	Price amazon.Currency
	Error error
}

func (r Result) ID() string {
	return path.Base(r.URL.Path)
}

func (r Result) String() string {
	return fmt.Sprintf("%s, %d, %v", r.ID(), r.Price.Cents, r.Error)
}

func main() {
	// flags
	opts := struct {
		year int
		file string
	}{}

	flag.IntVar(&opts.year, "y", 2021, "year filter")
	flag.StringVar(&opts.file, "f", "items.csv", "items file path")
	flag.Parse()

	// track result processing
	var wg sync.WaitGroup
	defer wg.Wait()

	// scraper setup
	c := colly.NewCollector(
		colly.Async(),
		colly.CacheDir("cache"),
	)
	defer c.Wait()

	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       time.Second / 5,
		RandomDelay: time.Second / 2,
	}); err != nil {
		log.Fatal("error setting collector limit:", err)
	}

	// errors
	errors := make(chan error)

	go func() {
		for i := 0; true; i++ {
			err := <-errors
			log.Println(i, err)
		}
	}()

	// results
	results := make(chan Result)

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		wg.Add(1)
		results <- Result{URL: e.Request.URL, Price: Currency(e.Text)}
	})

	c.OnError(func(resp *colly.Response, err error) {
		wg.Add(1)
		results <- Result{URL: resp.Request.URL, Error: err}
	})

	var mx sync.Mutex
	items := make(map[string]Item)

	bar := pb.StartNew(0)

	go func() {
		for {
			res := <-results
			id := res.ID()

			if res.Error != nil {
				log.Printf("error visiting %q: %s", id, res.Error.Error())

				bar.Add(1)
				wg.Done()
				continue
			}

			mx.Lock()

			// only add new items, skip duplicate results
			if item := items[id]; !item.Done {
				bar.Add(1)

				item.Done = true
				item.NewPrice = res.Price
				items[id] = item
			}

			mx.Unlock()
			wg.Done()
		}
	}()

	// parse export
	file, err := os.ReadFile(opts.file)
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	var raw []amazon.Item
	if err := csvutil.Unmarshal(file, &raw); err != nil {
		log.Fatal("error unmarshaling items:", err)
	}

	// dispatch
	var count int
	for _, v := range raw {
		// only 2021 items
		if y, _, _ := v.Date.Time.Date(); y != opts.year {
			continue
		}

		// skip non-new items and gift cards
		if v.Condition != "new" || v.Category == "GIFT_CARD" {
			continue
		}

		// this isn't fair
		if v.Category == "COMPUTER_PROCESSOR" {
			continue
		}

		mx.Lock()

		// only add new items, skip duplicates
		if _, ok := items[v.AsinIsbn]; !ok {
			items[v.AsinIsbn] = Item{Item: v}
		}

		mx.Unlock()

		count++
		c.Visit(baseurl + v.AsinIsbn)
	}

	bar.SetTotal(int64(count))

	c.Wait()  // scraper
	wg.Wait() // processor

	// statistics
	var fails int
	for k, v := range items {
		if !v.Done {
			fails++
			delete(items, k)
		}
	}

	fmt.Printf("%d successful items, %d failed items\n", len(items), fails)

	sorted, i := make([]Item, len(items)), 0
	for _, v := range items {
		sorted[i] = v
		i++
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Date.Before(sorted[j].Date.Time)
	})

	var prev, now int
	for _, v := range sorted {
		old := v.PurchasePricePerUnit.Cents
		cur := v.NewPrice.Cents

		prev += old
		now += cur

		/*
			fmt.Printf(
				"%d:\t%s\t%s\told: %s\tnew: %s\tdiff: %s\t%s\n",
				i,
				v.AsinIsbn,
				fmtdate(v.Date.Time),
				fmtcents(old, false),
				fmtcents(cur, false),
				fmtcents(cur-old, true),
				v.Category,
			)
		*/
	}

	fmt.Printf(
		"%s total %d\n%s total now\n%s (%+d%%) difference\n",
		fmtcents(prev, false), opts.year,
		fmtcents(now, false),
		fmtcents(now-prev, true), (now-prev)*100/prev,
	)
}
