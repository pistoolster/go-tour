package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type LevyEFrankResult struct {
	AckCode                        string `xml:"ACK_CODE"`
	AckMessage                     string `xml:"ACK_MESSAGE"`
	ReceiptNumber                  string `xml:"RECEIPT_NUMBER"`
	OutBoundFarePaid               string `xml:"OUT_BOUND_FARE_PAID"`
	NameOfTravelerPayer            string `xml:"NAME_OF_TRAVELER_PAYER"`
	DepartureDate                  string `xml:"DEPARTURE_DATE"`
	Last4DigitOfContactPhoneNumber string `xml:"LAST_4_DIGIT_OF_CONTACT_PHONE_NUMBER"`
	FrankDateTime                  string `xml:"FRANK_DATE_TIME"`
	LevyCollectionNumber           string `xml:"LEVY_COLLECTION_NUMBER"`
	LevyAmount                     string `xml:"LEVY_AMOUNT"`
	LowAccountBalanceIndicator     string `xml:"LOW_ACCOUNT_BALANCE_INDICATOR"`
	LevyStamp                      string `xml:"LEVY_STAMP"`
}

type LevyEFrankResponse struct {
	LevyEFrankResult LevyEFrankResult `xml:"LevyEFrankResult"`
}

type LevyEFrankBody struct {
	LevyEFrankResponse LevyEFrankResponse `xml:"LevyEFrankResponse"`
}

type LevyEFrankEnvelope struct {
	LevyEFrankBody LevyEFrankBody `xml:"Body"`
}

type ELevyReqRequest struct {
	ApiKey                         string `xml:"elev1:API_KEY"`
	ReceiptNumber                  string `xml:"elev1:RECEIPT_NUMBER"`
	OutBoundFarePaid               string `xml:"elev1:OUT_BOUND_FARE_PAID"`
	NameOfTravelerPayer            string `xml:"elev1:NAME_OF_TRAVELER_PAYER"`
	DepartureDate                  string `xml:"elev1:DEPARTURE_DATE"`
	Last4DigitOfContactPhoneNumber string `xml:"elev1:LAST_4_DIGIT_OF_CONTACT_PHONE_NUMBER"`
}

type ElevyReqLevyEFrank struct {
	ELevyReqRequest ELevyReqRequest `xml:"elev:request"`
}

type ElevyReqBody struct {
	ElevyReqLevyEFrank ElevyReqLevyEFrank `xml:"elev:LevyEFrank"`
}

type ElevyReqHeader struct {
}

type ElevyReqEnvelope struct {
	XMLName        xml.Name       `xml:"soapenv:Envelope"`
	ElevyReqHeader ElevyReqHeader `xml:"soapenv:Header"`
	ElevyReqBody   ElevyReqBody   `xml:"soapenv:Body"`
	XmlnsSoapenv   string         `xml:"xmlns:soapenv,attr"`
	XmlnsElev      string         `xml:"xmlns:elev,attr"`
	XmlnsElev1     string         `xml:"xmlns:elev1,attr"`
}

func main() {
	url := "https://www.elevy-ticf-train.org.hk:8443/ELevyService.svc"
	soapenv := "http://schemas.xmlsoap.org/soap/envelope/"
	elev := "ELevy.Service"
	elev1 := "ELevy.Schema"

	payload := &ElevyReqEnvelope{
		ElevyReqBody: ElevyReqBody{ElevyReqLevyEFrank: ElevyReqLevyEFrank{ELevyReqRequest: ELevyReqRequest{
			ApiKey:                         " ",
			ReceiptNumber:                  "3125126537252",
			OutBoundFarePaid:               "10",
			NameOfTravelerPayer:            "AFDFSDFSAF",
			DepartureDate:                  "2019-10-01",
			Last4DigitOfContactPhoneNumber: "3234"}}},
		XmlnsSoapenv: soapenv,
		XmlnsElev:    elev,
		XmlnsElev1:   elev1,
	}

	payload_raw := []byte(strings.TrimSpace(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:elev="ELevy.Service" xmlns:elev1="ELevy.Schema">
	<soapenv:Header/>
		<soapenv:Body>
		<elev:LevyEFrank>
		<!--Optional:-->

	`))
	pld, _ := xml.Marshal(payload)
	pldprt, _ := xml.MarshalIndent(payload, "", "\t")
	// fmt.Println("raw---", string(payload_raw))
	// fmt.Println("pld---", string(pldprt))
	fmt.Println(strings.Compare(string(payload_raw), string(pld)))
	fmt.Printf("%+v\n", string(pldprt))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(pld))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Set("SOAPAction", "ELevy.Service/IELevyService/LevyEFrank")

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		Renegotiation:      tls.RenegotiateOnceAsClient,
	// 		InsecureSkipVerify: true},
	// }
	client := http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			Renegotiation:      tls.RenegotiateOnceAsClient,
			InsecureSkipVerify: true},
	}}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Print(string(body))
	// f, err := os.Create("./test_xml")
	// defer f.Close()
	// n, _ := f.Write(body)
	// fmt.Println(n)

	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}

	mr := multipart.NewReader(resp.Body, params["boundary"])
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Header---------", p.Header.Get("Content-Type"))
		// fmt.Println("Body  ---------", string(body))
		if strings.Contains(p.Header.Get("Content-Type"), "xml") {
			fmt.Println("Body  ---------", string(body))
			levy_resp := &LevyEFrankEnvelope{}
			err = xml.Unmarshal(body, levy_resp)
			if err != nil {
				log.Println("Error on unmarshaling xml. ", err.Error())
			}
			fmt.Println(levy_resp)
		}
		if strings.Contains(p.Header.Get("Content-Type"), "octet-stream") {
			f, _ := os.Create("./test_xml.png")
			defer f.Close()
			_, _ = f.Write(body)
		}
	}

	// levy_resp := &LevyEFrankXml{}
	// err = xml.NewDecoder(resp.Body).Decode(levy_resp)
	// if err != nil {
	// 	log.Println("Error on unmarshaling xml. ", err.Error())
	// }
	// fmt.Println(levy_resp)
}
