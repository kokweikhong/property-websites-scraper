package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const (
	PROPERTY_GURU_URL = "https://www.propertyguru.com.my/property-for-sale?listing_type=sale&market=residential&district_code=JH016&region_code=MY01&freetext=Johor+Bahru&newProject=all"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	nodes := []*cdp.Node{}
	err := chromedp.Run(ctx,
		chromedp.Navigate(PROPERTY_GURU_URL),
		chromedp.WaitVisible(`#search-results-container`, chromedp.ByID),
		chromedp.Sleep(2*time.Second),
		chromedp.Nodes(`div.listing-card.listing-card-sale`, &nodes, chromedp.ByQueryAll),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of nodes: ", len(nodes))

	for _, node := range nodes {
		// fmt.Println("node: ", node)
		var nodeText string
		err = chromedp.Run(ctx,
			chromedp.Text(`div.header-container a.nav-link`, &nodeText, chromedp.FromNode(node), chromedp.ByQuery),
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("Property Name/Location: ", nodeText)
	}

}
