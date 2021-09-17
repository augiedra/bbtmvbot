package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	parseLinkAlio		 = "https://www.alio.lt/paieska/?category_id=1393&city_id=228822&search_block=1&search[eq][adresas_1]=228822&search[in][adresas_2][]=250583&search[in][adresas_2][]=250589&search[in][adresas_2][]=250577&search[in][adresas_2][]=353724&search[in][adresas_2][]=250580&search[in][adresas_2][]=353731&search[in][adresas_2][]=250591&search[in][adresas_2][]=250593&order=ad_id"
	parseLinkAruodas     = "https://m.aruodas.lt/?obj=4&FRegion=43&FDistrict=6&FOrder=AddDate&from_search=1&detailed_search=1&FShowOnly=FOwnerDbId0%2CFOwnerDbId1&FQuartal=25%2C100%2C26%2C139%2C36%2C40%2C37&act=search"
	parseLinkDomoplius   = "https://m.domoplius.lt/skelbimai/butai?action_type=3&address_1=43&category_id=1&address_2%5B3_1%5D=3_1&address_2%5B3_4%5D=3_4&address_2%5B3_5%5D=3_5&address_2%5B3_84%5D=3_84&address_2%5B3_93%5D=3_93&address_2%5B3_12%5D=3_12&address_2%5B3_13%5D=3_13&address_2%5B3_95%5D=3_95&qt="
	parseLinkKampas      = "https://www.kampas.lt/api/classifieds/search-new?query={%22municipality%22%3A%2215%22%2C%22settlement%22%3A4188%2C%22page%22%3A1%2C%22sort%22%3A%22new%22%2C%22section%22%3A%22bustas-nuomai%22%2C%22type%22%3A%22flat%22}"
	parseLinkNuomininkai = "https://nuomininkai.lt/paieska/?propery_type=butu-nuoma&propery_contract_type=&renter_type=&propery_location=43&imic_property_district=6&new_quartals=,25,100,26,139,36,39,37,40,&imic_property_quartals=&imic_property_streets=&min_price=&max_price=&min_price_meter=&max_price_meter=&min_area=&max_area=&floor_type=&building_type=&zm_skaicius=&rooms_from=&rooms_to=&irengimas=&house_year_from=&house_year_to=&lot_size_from=&lot_size_to=&by_date="
	parseLinkRinka       = "https://www.rinka.lt/nekilnojamojo-turto-skelbimai/butu-nuoma?filter[KainaForAll][min]=&filter[KainaForAll][max]=&filter[NTnuomakambariuskaiciusButai][min]=&filter[NTnuomakambariuskaiciusButai][max]=&filter[NTnuomabendrasplotas][min]=&filter[NTnuomabendrasplotas][max]=&filter[NTnuomastatybosmetai][min]=&filter[NTnuomastatybosmetai][max]=&filter[NTnuomaaukstuskaicius][min]=&filter[NTnuomaaukstuskaicius][max]=&filter[NTnuomaaukstas][min]=&filter[NTnuomaaukstas][max]=&cities[0]=21"
	parseLinkSkelbiu     = "https://www.skelbiu.lt/skelbimai/?import=2&district=6&quarter=25,100,26,139,36,39,40,37&cities=43&mainCity=1&search=1&category_id=322&type=1&visited_page=1&orderBy=1&detailsSearch=1"
)

func compileAddressWithStreet(district, street, houseNumber string) (address string) {
	if district == "" {
		address = "Kaunas"
	} else if street == "" {
		address = "Kaunas, " + district
	} else if houseNumber == "" {
		address = "Kaunas, " + district + ", " + street
	} else {
		address = "Kaunas, " + district + ", " + street + " " + houseNumber
	}
	return
}

func compileAddress(district, street string) (address string) {
	if district == "" {
		address = "Kaunas"
	} else if street == "" {
		address = "Kaunas, " + district
	} else {
		address = "Kaunas, " + district + ", " + street
	}
	return
}

var httpClient = &http.Client{Timeout: time.Second * 30}

func fetch(link string) ([]byte, error) {
	res, err := httpClient.Get(link)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		u, _ := url.Parse(link)
		return nil, fmt.Errorf("%s returned: %s %s", u.Host, res.Status, string(content))
	}

	return content, nil
}

func fetchDocument(link string) (*goquery.Document, error) {
	res, err := httpClient.Get(link)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		u, _ := url.Parse(link)
		return nil, fmt.Errorf("%s returned: %s", u.Host, res.Status)
	}

	return goquery.NewDocumentFromReader(res.Body)
}
