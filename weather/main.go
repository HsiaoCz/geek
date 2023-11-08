package main

import (
	"flag"
	"log"
	"github.com/go-rod/rod"
)

// view weather
const (
	nationwideUrl = "https://www.tianqi.com/chinacity.html"
)

func main() {
	var city = flag.String("c", "", "the config file")
	flag.Parse()
	url := cityUrl(*city)
	if url == "" {
		log.Fatalf("url is empty")
		return
	}
	weatherData(url)
}

func weatherData(url string) {
	page := rod.New().MustConnect().MustPage(url)
	//直接使用选择器选择七天天气内容体
	projects := page.MustElement("ul.weaul").MustElements("li")
	for _, project := range projects {
		var (
			date       string
			weather    string
			centigrade string
		)
		//获取日期
		if ok, dateEle, _ := project.Has(".weaul_q"); ok {
			if dateEle.MustHas(".fl") {
				date = dateEle.MustElement(".fl").MustText()
			}
		}
		for _, weatherEle := range project.MustElements(".weaul_z") {
			//如果存在span则是温度 否则是天气情况
			if weatherEle.MustHas("span") {
				var (
					startC string
					endC   string
				)
				//获取温度
				for k, cEle := range weatherEle.MustElements("span") {
					if k == 0 {
						startC = cEle.MustText()
					} else {
						endC = cEle.MustText()
					}
				}
				centigrade = startC + "~" + endC
			} else {
				weather = weatherEle.MustText()
			}
		}
		log.Printf("日期：%s｜天气：%s｜摄氏度：%s℃\n", date, weather, centigrade)
	}
}

func cityUrl(city string) string {
	cityMap := make(map[string]string)
	page := rod.New().MustConnect().MustPage(nationwideUrl)
	provinceProjects := page.MustElement(".citybox").MustElements("h2")
	cityProjects := page.MustElement(".citybox").MustElements("span")

	for _, section := range provinceProjects {
		cityMap[section.MustElement("a").MustText()] = section.MustElement("a").MustProperty("href").String()
	}
	for _, section := range cityProjects {
		for _, a := range section.MustElements("a") {
			cityMap[a.MustText()] = a.MustProperty("href").String()
		}
	}
	if href, ok := cityMap[city]; ok {
		return href + "7/"
	}
	return ""
}
