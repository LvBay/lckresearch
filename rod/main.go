// Package main ...
package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	// flag.Parse()

	// // get the commandline arguments
	// source := strings.TrimSpace(strings.Join(flag.Args(), " "))
	// if source == "" {
	// 	log.Fatal("usage: go run main.go -- 'This is the phrase to translate to Spanish.'")
	// }

	browser := rod.New().MustConnect()

	page := browser.MustPage("https://lol.ps/ranking/?lane=-1&page=1&region=kr")

	el := page.MustElementX(`//*[@id="content-container"]/div[4]/div[2]/div[1]/a[1]`)
	// //*[@id="content-container"]/div[4]/div[2]/div[1]/a[1]
	// //*[@id="content-container"]/div[4]/div[3]/div[1]/a[1]

	// wait := page.MustWaitRequestIdle("https://accounts.google.com")
	// el.MustInput(source)
	// wait()

	id := el.MustAttribute("href")
	u, _ := url.Parse(*id)
	sid := u.Query().Get("sid")

	fmt.Println("id:", sid)
	matchList(browser, "https://lol.ps"+*id)
}

func matchList(b *rod.Browser, a string) {
	fmt.Println("a:", a)
	page := b.MustPage(a)
	log.Println("page load finish")
	// {
	// 	el := page.MustElementX(`//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]/div/div[1]/div/div[2]/div[1]/div[2]/div[2]/div/div/img`)
	// 	src := el.MustAttribute("src")
	// 	fmt.Println("img src:", *src)
	// }

	// //*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]/div/div[1]/div/div[2]/div[1]/div[2]/div[1]
	// //*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]/div/div[1]/div/div[2]/div[1]/div[2]/div[2]

	// #content-container > div:nth-child(3) > section > div > div.flex.flex-col.gap-4.bg-psgr-light.px-4.pt-6.dark\:bg-dark-psgr-darkest.desktop\:w-\[693px\].desktop\:bg-transparent.desktop\:px-0.desktop\:pt-0.dark\:desktop\:bg-transparent > div.-mx-4.flex.w-\[calc\(100\%_\+_32px\)\].flex-col.gap-px.bg-white.dark\:bg-dark-psgr-darkest.desktop\:mx-0.desktop\:w-auto.desktop\:gap-2 > div:nth-child(1) > div > div.transition-color.flex.w-full.flex-none.flex-col.border-l-8.bg-opacity-10.duration-300.dark\:bg-opacity-\[0\.15\].desktop\:h-30.desktop\:w-\[95\%\].desktop\:border-l-4.desktop\:py-3.desktop\:pl-4.desktop\:pr-4.border-psbl-deep.bg-ui-bl-base.dark\:bg-psbl-deep > div > div.flex.flex-row.gap-2 > div.flex.flex-col.justify-center.gap-1.p-2.pr-0.desktop\:w-\[210px\].desktop\:p-0 > div.flex.flex-row.gap-1 > div:nth-child(2) > div > div > img
	// el := page.MustElement(`#content-container > div:nth-child(3) > section > div > div.flex.flex-col.gap-4.bg-psgr-light.px-4.pt-6.dark\:bg-dark-psgr-darkest.desktop\:w-\[693px\].desktop\:bg-transparent.desktop\:px-0.desktop\:pt-0.dark\:desktop\:bg-transparent > div.-mx-4.flex.w-\[calc\(100\%_\+_32px\)\].flex-col.gap-px.bg-white.dark\:bg-dark-psgr-darkest.desktop\:mx-0.desktop\:w-auto.desktop\:gap-2 > div:nth-child(1) > div > div.transition-color.flex.w-full.flex-none.flex-col.border-l-8.bg-opacity-10.duration-300.dark\:bg-opacity-\[0\.15\].desktop\:h-30.desktop\:w-\[95\%\].desktop\:border-l-4.desktop\:py-3.desktop\:pl-4.desktop\:pr-4.border-psbl-deep.bg-ui-bl-base.dark\:bg-psbl-deep > div > div.flex.flex-row.gap-2 > div.flex.flex-col.justify-center.gap-1.p-2.pr-0.desktop\:w-\[210px\].desktop\:p-0 > div.flex.flex-row.gap-1 > div:nth-child(1) > div > div > img`)
	// src := el.MustAttribute("src")
	// fmt.Println("img src:", *src)

	// 查询有几个子元素
	// //*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]/div/div[1]/div/div[2]/div[1]/div[2]
	{
		x := `//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]`
		els := page.MustElementX(x).MustParent().MustElements(".relative")
		fmt.Println("el:", len(els))
	}
	{
		x := `//*[@id="content-container"]/div[2]/section/div/div[2]/div[3]/button`
		el := page.MustElementX(x)
		h := el.MustHTML()
		log.Println("html:", h)
		el.MustClick()
	}
	time.Sleep(5 * time.Second)
	log.Println("sleep wait next")
	{
		x := `//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]`
		els := page.MustElementX(x).MustParent().MustElements(".relative")
		fmt.Println("el:", len(els))
	}
	{
		x := `//*[@id="content-container"]/div[2]/section/div/div[2]/div[3]/button`
		el := page.MustElementX(x)
		h := el.MustHTML()
		log.Println("html:", h)
		el.MustClick()
	}
	time.Sleep(5 * time.Second)
	log.Println("sleep wait next")
	{
		x := `//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]`
		els := page.MustElementX(x).MustParent().MustElements(".relative")
		fmt.Println("el:", len(els))
	}
	// els := page.MustElementsX(x)
	// fmt.Println("els:", len(els))

	// 遍历装备
	// for i := 1; i < 4; i++ {
	// 	x := fmt.Sprintf(`//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]/div/div[1]/div/div[2]/div[1]/div[2]/div[%d]/div/div/img`, i)
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// 	pageWithCancel := page.Context(ctx)
	// 	go func() {
	// 		time.Sleep(2 * time.Second)
	// 		cancel()
	// 	}()
	// 	el, err := pageWithCancel.ElementX(x)
	// 	if err != nil {
	// 		fmt.Println("err", i, err)
	// 		continue
	// 	}
	// 	// el = el.MustElementX(`//*/div/div/img`)

	// 	// log.Println("x:", i)
	// 	// // tmpPage := page.Context(ctx)
	// 	// // el := tmpPage.MustElementX(x)
	// 	src := el.MustAttribute("src")
	// 	fmt.Println("img src:", *src)
	// }
	{

	}
	// el := page.MustElementX(`//*[@id="content-container"]/div[2]/section/div/div[2]/div[2]/div[1]/div/div[1]/div/div[2]/div[1]/div[2]/div[3]/div/div/img`)
	// src := el.MustAttribute("src")
	// fmt.Println("img src:", *src)
}

func git() {
	page := rod.New().MustConnect().MustPage("https://github.com/avelino/awesome-go")

	section := page.MustElementR("h3", "Selenium and browser control tools").MustParent().MustNext()
	fmt.Println(section.MustHTML())

	// get children elements of an element
	projects := section.MustElements("ul>li")
	log.Println("as", len(projects))

	for _, project := range projects {
		link := project.MustElement("a")
		log.Printf(
			"project %s (%s): '%s'",
			link.MustText(),
			link.MustProperty("href"),
			project.MustText(),
		)
	}
}
