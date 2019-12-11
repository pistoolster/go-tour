package main

import (
	// "encoding/json"
	"encoding/xml"
	"fmt"
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
	AckCode             string `xml:"elev1:API_KEY"`
	AckMessage          string `xml:"elev1:RECEIPT_NUMBER"`
	ReceiptNumber       string `xml:"elev1:OUT_BOUND_FARE_PAID"`
	OutBoundFarePaid    string `xml:"elev1:NAME_OF_TRAVELER_PAYER"`
	NameOfTravelerPayer string `xml:"elev1:DEPARTURE_DATE"`
	DepartureDate       string `xml:"elev1:LAST_4_DIGIT_OF_CONTACT_PHONE_NUMBER"`
}

type ElevyReqLevyEFrank struct {
	ELevyReqRequest ELevyReqRequest `xml:"request"`
}

type ElevyReqBody struct {
	ElevyReqLevyEFrank ElevyReqLevyEFrank `xml:"LevyEFrank"`
}

type ElevyReqHeader struct {
	ElevyReqBody ElevyReqBody `xml:"soapenv:Body"`
}

type ElevyReqEnvelope struct {
	XMLName        xml.Name       `xml:"soapenv:Envelope"`
	ElevyReqHeader ElevyReqHeader `xml:"soapenv:Header"`
	XmlnsSoapenv   string         `xml:"xmlns:soapenv,attr"`
	XmlnsElev      string         `xml:"xmlns:elev,attr"`
	XmlnsElev1     string         `xml:"xmlns:elev1,attr"`
}

/*
<Persons>
    <Person>
        <Name>studygolang</Name>
        <Age>27</Age>
        <Career>码农</Career>
        <Interests>
            <Interest>编程</Interest>
            <Interest>下棋</Interest>
        </Interests>
    </Person>
</Persons>
*/

type Result struct {
	Person Person
}
type Person struct {
	Name      string
	Age       int
	Career    string
	Interests Interests
}
type Interests struct {
	Interest []string
}

func main() {
	payload := `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"><s:Body><LevyEFrankResponse xmlns="ELevy.Service"><LevyEFrankResult xmlns:a="ELevy.Schema" xmlns:i="http://www.w3.org/2001/XMLSchema-instance"><a:ACK_CODE>S</a:ACK_CODE><a:ACK_MESSAGE>OK</a:ACK_MESSAGE><a:RECEIPT_NUMBER>1234ss5ss90</a:RECEIPT_NUMBER><a:OUT_BOUND_FARE_PAID>1.11</a:OUT_BOUND_FARE_PAID><a:NAME_OF_TRAVELER_PAYER>AFDFSDFSAF SFADS, FADF  DSA, FASFASFAS</a:NAME_OF_TRAVELER_PAYER><a:DEPARTURE_DATE>2019-10-01</a:DEPARTURE_DATE><a:LAST_4_DIGIT_OF_CONTACT_PHONE_NUMBER>0601</a:LAST_4_DIGIT_OF_CONTACT_PHONE_NUMBER><a:FRANK_DATE_TIME>2019-10-08 14:20:57</a:FRANK_DATE_TIME><a:LEVY_COLLECTION_NUMBER>21021857161913</a:LEVY_COLLECTION_NUMBER><a:LEVY_AMOUNT>0.01</a:LEVY_AMOUNT><a:LOW_ACCOUNT_BALANCE_INDICATOR>0</a:LOW_ACCOUNT_BALANCE_INDICATOR><a:LEVY_STAMP><xop:Include href="cid:http://tempuri.org/1/637061542823396714" xmlns:xop="http://www.w3.org/2004/08/xop/include"/></a:LEVY_STAMP></LevyEFrankResult></LevyEFrankResponse></s:Body></s:Envelope>`
	resp := &LevyEFrankEnvelope{}
	err := xml.Unmarshal([]byte(payload), resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
	// s, _ := json.MarshalIndent(resp, "", "\t")
	// fmt.Println(string(s))
	payload = `<Persons><Person><Name>polaris</Name><Age>28</Age><Career>无业游民</Career><Interests><Interest>编程</Interest><Interest>下棋</Interest></Interests></Person></Persons>`
	res := &Result{}
	err = xml.Unmarshal([]byte(payload), res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", res)

	envelope := &ElevyReqEnvelope{}
	s, _ := xml.MarshalIndent(envelope, "", "\t")
	fmt.Printf("%+v\n", string(s))
}
