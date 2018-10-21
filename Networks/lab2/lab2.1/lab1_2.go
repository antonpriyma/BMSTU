package main

import (
	"github.com/mgutz/logxi/v1"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func (i Item) print()  {
	log.Info("")
	log.Info("Name",i.Info.Name)
	log.Info("Link",i.Info.Link)
	if i.Info.Discription!="" {
		log.Info("Discription",i.Info.Discription)
	}
	log.Info("Rating",i.Info.Rating)

	for _,indicator:=range i.ItemIndicators{
		log.Info("Indicator 1",indicator.IndicatorOne)
		log.Info("Indicator 2",indicator.IndicatorTwo)
	}


}

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

func readItem(item *html.Node) *Item {
	var resultItem Item
	//log.Info(getAttr(item.FirstChild.NextSibling,"class"))
	if a := item.FirstChild; isDiv(a, "rating-pamm__info") {
		if cs := getChildren(a); len(cs) == 3 {
			info := ItemInfo{
				cs[0].FirstChild.FirstChild.Data,
				getAttr(cs[0].FirstChild, "href"),
				cs[1].FirstChild.Data,
				strings.Split(getAttr(cs[2].FirstChild, "style"), " ")[1],
			}
			resultItem.Info=info
		}else if cs := getChildren(a); len(cs) == 2 {
				info := ItemInfo{
					cs[0].FirstChild.FirstChild.Data,
					getAttr(cs[0].FirstChild, "href"),
					"",
					strings.Split(getAttr(cs[1].FirstChild, "style"), " ")[1],
				}

			resultItem.Info=info
			}
		}
		if a := item.FirstChild.NextSibling; isDiv(a, "rating-pamm__indicators") {
			cs := getChildren(a)
			var indicators []ItemIndicators
			for _, i := range cs {
				if getAttr(i,"class")=="rating-pamm__histogram" {
					break
				}
				indicator := ItemIndicators{
					i.FirstChild.FirstChild.Data,
					i.FirstChild.NextSibling.FirstChild.Data,
				}
				indicators = append(indicators, indicator)
				resultItem.ItemIndicators=indicators
			}
		}
		return &resultItem
	}


type Item struct {
	Info ItemInfo
	ItemIndicators []ItemIndicators
}

type ItemInfo struct {
	Name,Link, Discription, Rating string
}

type ItemIndicators struct {
	IndicatorOne,IndicatorTwo string
} 


func downloadNews() ([]*Item) {
	log.Info("sending request to alpari.com")
	if response, err := http.Get("https://alpari.com/ru/"); err != nil {
		log.Error("request to alpari.com failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		log.Info("got response from alpari.com", "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				log.Error("invalid HTML from alpari.com", "error", err)
			} else {
				log.Info("HTML from alpari.com parsed successfully")

				return search(doc)
			}
		}
	}
	return nil
}

func search(node *html.Node) []*Item {
	var items []*Item
	if isDiv(node, "rating-pamm") && node.Parent.NextSibling!=nil{
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isDiv(c, "rating-pamm__item"){
					if item := readItem(c); item != nil{
						items = append(items, item)
					}
			}
		}
		return items
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		for _,item:=range search(c){
			items=append(items, item)
			if len(items)==5{
				return items
			}
		}
	}
	if items!=nil {
		return items
	}
	return nil

}






//===================================================================================================



func main() {


	log.Info("Downloader started")
	items:=downloadNews()
	
	for _,item:=range items{
		item.print()
	}





}
