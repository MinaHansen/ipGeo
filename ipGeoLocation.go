package main

import (
	"fmt"
	"strings"
)

type IpGeolocation struct {
	IP                  string `json:"ip"`
	ContinentCode       string `json:"continent_code"`
	ContinentName       string `json:"continent_name"`
	CountryCode2        string `json:"country_code2"`
	CountryCode3        string `json:"country_code3"`
	CountryName         string `json:"country_name"`
	CountryNameOfficial string `json:"country_name_official"`
	CountryCapital      string `json:"country_capital"`
	StateProv           string `json:"state_prov"`
	StateCode           string `json:"state_code"`
	District            string `json:"district"`
	City                string `json:"city"`
	Zipcode             string `json:"zipcode"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	IsEu                bool   `json:"is_eu"`
	CallingCode         string `json:"calling_code"`
	CountryTld          string `json:"country_tld"`
	Languages           string `json:"languages"`
	CountryFlag         string `json:"country_flag"`
	GeonameID           string `json:"geoname_id"`
	Isp                 string `json:"isp"`
	ConnectionType      string `json:"connection_type"`
	Organization        string `json:"organization"`
	CountryEmoji        string `json:"country_emoji"`
	Currency            struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currency"`
	TimeZone struct {
		Name            string  `json:"name"`
		Offset          int     `json:"offset"`
		OffsetWithDst   int     `json:"offset_with_dst"`
		CurrentTime     string  `json:"current_time"`
		CurrentTimeUnix float64 `json:"current_time_unix"`
		IsDst           bool    `json:"is_dst"`
		DstSavings      int     `json:"dst_savings"`
		DstExists       bool    `json:"dst_exists"`
		DstStart        struct {
			UtcTime        string `json:"utc_time"`
			Duration       string `json:"duration"`
			Gap            bool   `json:"gap"`
			DateTimeAfter  string `json:"dateTimeAfter"`
			DateTimeBefore string `json:"dateTimeBefore"`
			Overlap        bool   `json:"overlap"`
		} `json:"dst_start"`
		DstEnd struct {
			UtcTime        string `json:"utc_time"`
			Duration       string `json:"duration"`
			Gap            bool   `json:"gap"`
			DateTimeAfter  string `json:"dateTimeAfter"`
			DateTimeBefore string `json:"dateTimeBefore"`
			Overlap        bool   `json:"overlap"`
		} `json:"dst_end"`
	} `json:"time_zone"`
}

func (ip IpGeolocation) ToString() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("IP: %s\n", ip.IP))
	sb.WriteString(fmt.Sprintf("ContinentCode: %s\n", ip.ContinentCode))
	sb.WriteString(fmt.Sprintf("ContinentName: %s\n", ip.ContinentName))
	sb.WriteString(fmt.Sprintf("CountryCode2: %s\n", ip.CountryCode2))
	sb.WriteString(fmt.Sprintf("CountryCode3: %s\n", ip.CountryCode3))
	sb.WriteString(fmt.Sprintf("CountryName: %s\n", ip.CountryName))
	sb.WriteString(fmt.Sprintf("CountryNameOfficial: %s\n", ip.CountryNameOfficial))
	sb.WriteString(fmt.Sprintf("CountryCapital: %s\n", ip.CountryCapital))
	sb.WriteString(fmt.Sprintf("StateProv: %s\n", ip.StateProv))
	sb.WriteString(fmt.Sprintf("StateCode: %s\n", ip.StateCode)) // Corrected typo here
	sb.WriteString(fmt.Sprintf("District: %s\n", ip.District))
	sb.WriteString(fmt.Sprintf("City: %s\n", ip.City))
	sb.WriteString(fmt.Sprintf("Zipcode: %s\n", ip.Zipcode))
	sb.WriteString(fmt.Sprintf("Latitude: %s\n", ip.Latitude))
	sb.WriteString(fmt.Sprintf("Longitude: %s\n", ip.Longitude))
	sb.WriteString(fmt.Sprintf("IsEu: %t\n", ip.IsEu))
	sb.WriteString(fmt.Sprintf("CallingCode: %s\n", ip.CallingCode))
	sb.WriteString(fmt.Sprintf("CountryTld: %s\n", ip.CountryTld))
	sb.WriteString(fmt.Sprintf("Languages: %s\n", ip.Languages))
	sb.WriteString(fmt.Sprintf("CountryFlag: %s\n", ip.CountryFlag))
	sb.WriteString(fmt.Sprintf("GeonameID: %s\n", ip.GeonameID))
	sb.WriteString(fmt.Sprintf("Isp: %s\n", ip.Isp))
	sb.WriteString(fmt.Sprintf("ConnectionType: %s\n", ip.ConnectionType))
	sb.WriteString(fmt.Sprintf("Organization: %s\n", ip.Organization))
	sb.WriteString(fmt.Sprintf("CountryEmoji: %s\n", ip.CountryEmoji))
	sb.WriteString("Currency:\n")
	sb.WriteString(fmt.Sprintf("\tCode: %s\n", ip.Currency.Code))
	sb.WriteString(fmt.Sprintf("\tName: %s\n", ip.Currency.Name))
	sb.WriteString(fmt.Sprintf("\tSymbol: %s\n", ip.Currency.Symbol))
	sb.WriteString("TimeZone:\n")
	sb.WriteString(fmt.Sprintf("\tName: %s\n", ip.TimeZone.Name))
	sb.WriteString(fmt.Sprintf("\tOffset: %d\n", ip.TimeZone.Offset))
	sb.WriteString(fmt.Sprintf("\tOffsetWithDst: %d\n", ip.TimeZone.OffsetWithDst))
	sb.WriteString(fmt.Sprintf("\tCurrentTime: %s\n", ip.TimeZone.CurrentTime))
	sb.WriteString(fmt.Sprintf("\tCurrentTimeUnix: %f\n", ip.TimeZone.CurrentTimeUnix))
	sb.WriteString(fmt.Sprintf("\tIsDst: %t\n", ip.TimeZone.IsDst))
	sb.WriteString(fmt.Sprintf("\tDstSavings: %d\n", ip.TimeZone.DstSavings))
	sb.WriteString(fmt.Sprintf("\tDstExists: %t\n", ip.TimeZone.DstExists))
	sb.WriteString("DstStart:\n")
	sb.WriteString(fmt.Sprintf("\tUtcTime: %s\n", ip.TimeZone.DstStart.UtcTime))
	sb.WriteString(fmt.Sprintf("\tDuration: %s\n", ip.TimeZone.DstStart.Duration))
	sb.WriteString(fmt.Sprintf("\tGap: %t\n", ip.TimeZone.DstStart.Gap))
	sb.WriteString(fmt.Sprintf("\tDateTimeAfter: %s\n", ip.TimeZone.DstStart.DateTimeAfter))
	sb.WriteString(fmt.Sprintf("\tDateTimeBefore: %s\n", ip.TimeZone.DstStart.DateTimeBefore))
	sb.WriteString(fmt.Sprintf("\tOverlap: %t\n", ip.TimeZone.DstStart.Overlap))
	sb.WriteString("DstEnd:\n")
	sb.WriteString(fmt.Sprintf("\tUtcTime: %s\n", ip.TimeZone.DstEnd.UtcTime))
	sb.WriteString(fmt.Sprintf("\tDuration: %s\n", ip.TimeZone.DstEnd.Duration))
	sb.WriteString(fmt.Sprintf("\tGap: %t\n", ip.TimeZone.DstEnd.Gap))
	sb.WriteString(fmt.Sprintf("\tDateTimeAfter: %s\n", ip.TimeZone.DstEnd.DateTimeAfter))
	sb.WriteString(fmt.Sprintf("\tDateTimeBefore: %s\n", ip.TimeZone.DstEnd.DateTimeBefore))
	sb.WriteString(fmt.Sprintf("\tOverlap: %t\n", ip.TimeZone.DstEnd.Overlap))
	return sb.String()
}
