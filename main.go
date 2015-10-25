package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

/*
type StockRequestObject struct {
	Name       [5]string
	Percentage [5]int
	Budget     float32
	TradeId    int
}

type StockResponseObject struct {
	TradeId            int
	Name               [5]string
	NumberOfStocks     [5]int
	StockValue         [5]float64
	UnvestedAmount     float64
	CurrentMarketValue float64
	ProfitLoss         [5]string
}

var names [5]string
var val [5]float64
var tradeId int

func (this *Server) Receive(Sr1 StockRequestObject, Sresp *StockResponseObject) error {

	for index := 0; index < len(Sr1.Name); index++ {
		if Sr1.Name[index] != "" {

			selectQuery := "https://query.yahooapis.com/v1/public/yql?q=select%20LastTradePriceOnly%2C%20Symbol%20from%20yahoo.finance.quote%20"
			whereQuery := "where%20symbol%20in%20("
			endQuery := ")&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback="

			whereQuery = whereQuery + "%27" + Sr1.Name[index] + "%27"
			finalQuery := selectQuery + whereQuery + endQuery
			res, err := http.Get(finalQuery)

			if err != nil {
				log.Fatal(err)
			}
			robots, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			var myjson MyJsonName
			err = json.Unmarshal(robots, &myjson)

			names[index] = myjson.Query.Results.Quote.Name
			val[index], err = strconv.ParseFloat(myjson.Query.Results.Quote.LastTradePriceOnly, 64)

			Sresp.Name[index] = names[index]
			fmt.Println("Stock name: ", names[index])
			fmt.Println("Stock value: ", val[index])
		}
	}
	var amountLeft float64
	amountLeft = 0

	for NewIndex := 0; NewIndex < len(Sr1.Name); NewIndex++ {
		if Sr1.Name[NewIndex] != "" {

			AllocatedAmount := float64((Sr1.Budget * float32(Sr1.Percentage[NewIndex])) / 100)

			Sresp.NumberOfStocks[NewIndex] = int(AllocatedAmount / val[NewIndex])
			var tempSum float64

			var stValue float64
			stValue = float64(val[NewIndex]) * float64(Sresp.NumberOfStocks[NewIndex])
			tempSum = float64(AllocatedAmount - stValue)
			Sresp.StockValue[NewIndex] = stValue
			amountLeft += tempSum
		}
	}
	Sresp.UnvestedAmount = amountLeft

	tradeId += 1
	Sresp.TradeId = tradeId
	Portfolio[Sresp.TradeId] = *Sresp

	return nil
}

func (this *Server) GetPortfolio(Sr1 StockRequestObject, Sresp *StockResponseObject) error {
	Test := Portfolio[Sr1.TradeId]

	var StockNames [5]string
	var NumberOfStocks [5]int
	var StockValues [5]float64

	for index := 0; index < len(Test.Name); index++ {
		StockNames[index] = Test.Name[index]
		NumberOfStocks[index] = Test.NumberOfStocks[index]
		StockValues[index] = Test.StockValue[index]
	}

	for index := 0; index < len(StockNames); index++ {
		if StockNames[index] != "" {

			selectQuery2 := "https://query.yahooapis.com/v1/public/yql?q=select%20LastTradePriceOnly%2C%20Symbol%20from%20yahoo.finance.quote%20"
			whereQuery2 := "where%20symbol%20in%20("
			endQuery2 := ")&format=json&diagnostics=true&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback="

			whereQuery2 = whereQuery2 + "%27" + StockNames[index] + "%27"
			finalQuery2 := selectQuery2 + whereQuery2 + endQuery2
			res, err := http.Get(finalQuery2)

			if err != nil {
				log.Fatal(err)
			}
			robots2, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			var myjson MyJsonName
			err = json.Unmarshal(robots2, &myjson)

			CurrentStockNames[index] = myjson.Query.Results.Quote.Name
			CurrentStockValues[index], err = strconv.ParseFloat(myjson.Query.Results.Quote.LastTradePriceOnly, 64)

			Sresp.Name[index] = names[index]
			Sresp.TradeId = Sr1.TradeId
			fmt.Println()
			fmt.Println("New Stock name: ", CurrentStockNames[index])
			fmt.Println("New Stock value: ", CurrentStockValues[index])

		}
	}

	Sresp.CurrentMarketValue = 0
	for NewIndex2 := 0; NewIndex2 < len(CurrentStockNames); NewIndex2++ {
		if CurrentStockNames[NewIndex2] != "" {
			Sresp.CurrentMarketValue += (CurrentStockValues[NewIndex2] * float64(Test.NumberOfStocks[NewIndex2]))
			Sresp.NumberOfStocks[NewIndex2] = Test.NumberOfStocks[NewIndex2]
			Sresp.StockValue[NewIndex2] = (CurrentStockValues[NewIndex2] * float64(Test.NumberOfStocks[NewIndex2]))
			var TestProfitLoss float64
			TestProfitLoss = (CurrentStockValues[NewIndex2] * float64(Test.NumberOfStocks[NewIndex2])) - (StockValues[NewIndex2] * float64(Test.NumberOfStocks[NewIndex2]))
			if TestProfitLoss > 0 {
				Sresp.ProfitLoss[NewIndex2] = " +{Profit} "
			} else if TestProfitLoss > 0 {
				Sresp.ProfitLoss[NewIndex2] = " -{Loss} "
			} else {
				Sresp.ProfitLoss[NewIndex2] = " {NoChange} "
			}
		}
	}
	Sresp.UnvestedAmount = Test.UnvestedAmount

	fmt.Println("Current total value of stocks: ", Sresp.CurrentMarketValue)

	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
*/

type Person struct {
	Name  string
	Phone string
}
type MyJsonResult struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

type MyJsonName struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name"`
	Address    string        `json:"address"`
	City       string        `json:"city"`
	State      string        `json:"state"`
	Zip        string        `json:"zip"`
	Coordinate struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"coordinate"`
	//Id2   bson.ObjectId. `json:"id" bson:"_id,omitempty"`
}

func Getlocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

	session, err := mgo.Dial("mongodb://user1:pass1@ds045054.mongolab.com:45054/mydatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydatabase").C("people")
	id := p.ByName("name")
	//id := "562975762b877b26f82bf8c5"
	oid := bson.ObjectIdHex(id)
	//od := "562972862b877b2968b8a1a0"
	var result MyJsonName
	c.FindId(oid).One(&result)
	//fmt.Fprintf(rw, result)

	//	err = c.Find(bson.M{"id": oid}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Searched Name:", result.Name)
	fmt.Println("Searched Address:", result.Address)
	fmt.Println("Searched City:", result.City)
	fmt.Println("Searched State:", result.State)
	fmt.Println("Searched Zip:", result.Zip)
	fmt.Println("Searched latitude:", result.Coordinate.Lat)
	fmt.Println("Searched longitude:", result.Coordinate.Lng)

	fmt.Println("Id2:", result.Id.String())
	oid = bson.ObjectId(result.Id)

	b2, err := json.Marshal(result)
	if err != nil {
	}
	// fmt.Fprintf(rw, "New response")

	fmt.Fprintf(rw, string(b2))
	fmt.Println("Method Name: " + req.Method)
}

func Postlocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	var myjson3 MyJsonName
	s3 := json.NewDecoder(req.Body)
	err := s3.Decode(&myjson3)
	StartQuery := "http://maps.google.com/maps/api/geocode/json?address="
	WhereQuery := myjson3.Address + " " + myjson3.City + " " + myjson3.State
	WhereQuery = strings.Replace(WhereQuery, " ", "+", -1)
	EndQuery := "&sensor=false"
	Url1 := StartQuery + WhereQuery + EndQuery
	fmt.Println("Published URL: " + Url1)
	//Url1 := "http://maps.google.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&sensor=false"
	res, err := http.Get(Url1)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	//fmt.Println(robots)
	res.Body.Close()
	//fmt.Println(robots)
	if err != nil {
		log.Fatal(err)
	}
	var myjsonresult1 MyJsonResult
	err = json.Unmarshal(robots, &myjsonresult1)
	if err != nil {
		log.Fatal(err)
	}
	//myjsonresult1.Results[0].Geometry
	fmt.Println(myjsonresult1.Results[0].Geometry.Location.Lat)
	fmt.Println(myjsonresult1.Results[0].Geometry.Location.Lng)

	myjson3.Id = bson.NewObjectId()
	//bson.NewObjectId().
	fmt.Println("Check1")
	fmt.Println(string(myjson3.Id))
	fmt.Println(myjson3.Id.Hex())
	fmt.Println(myjson3.Id.String())
	fmt.Println(myjson3.Id.Pid())

	myjson3.Coordinate.Lat = myjsonresult1.Results[0].Geometry.Location.Lat
	myjson3.Coordinate.Lng = myjsonresult1.Results[0].Geometry.Location.Lng

	fmt.Println("Name " + myjson3.Name)
	fmt.Println("\nAddress: " + myjson3.Address)
	fmt.Println("\nCity:  " + myjson3.City)
	fmt.Println("\nState: " + myjson3.State)
	fmt.Println("\nLat and long : ")

	fmt.Println(myjson3.Coordinate.Lat)
	fmt.Println(myjson3.Coordinate.Lng)

	//test1 := http.Get(Url1)
	// fmt.Fprintf(rw, "\nAddress: "+myjson3.Address)
	if err != nil {
	}

	session, err := mgo.Dial("mongodb://user1:pass1@ds045054.mongolab.com:45054/mydatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mydatabase").C("people")

	//uc.session.DB("go_rest_tutorial").C("users").FindId(oid).One(&u)
	err = c.Insert(myjson3)
	if err != nil {
		log.Fatal(err)
	}

	/*
		result := myjson3
		err = c.FindId("ObjectID()")
	*/
	result := MyJsonName{}
	fmt.Println()
	//	c.session.DB("test").C("people").FindId(oid).One(&result)
	//id := "56293c0597eca8fc493d5434"
	id := myjson3.Id.Hex()
	oid := bson.ObjectIdHex(id)
	//od := "562972862b877b2968b8a1a0"
	c.FindId(oid).One(&result)
	//	err = c.Find(bson.M{"id": oid}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Name:", result.Name)
	fmt.Println("Address:", result.Address)
	fmt.Println("Id2:", result.Id.String())
	oid = bson.ObjectId(result.Id)

	b2, err := json.Marshal(result)
	if err != nil {
	}

	fmt.Fprintf(rw, string(b2))
}

func PutLocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	//fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

	var myjson3 MyJsonName
	s3 := json.NewDecoder(req.Body)
	err := s3.Decode(&myjson3)

	session, err := mgo.Dial("mongodb://user1:pass1@ds045054.mongolab.com:45054/mydatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydatabase").C("people")

	id := p.ByName("name")
	//id := "562975762b877b26f82bf8c5"
	oid := bson.ObjectIdHex(id)
	//od := "562972862b877b2968b8a1a0"
	var result MyJsonName
	//c.UpdateId(, update interface{})
	c.FindId(oid).One(&result)
	//fmt.Fprintf(rw, result)
	fmt.Println("Old Name:", result.Name)
	fmt.Println("Old Address:", result.Address)
	fmt.Println("Old City:", result.City)
	fmt.Println("Old State:", result.State)
	fmt.Println("Old Zip:", result.Zip)
	fmt.Println("Old latitude:", result.Coordinate.Lat)
	fmt.Println("Old longitude:", result.Coordinate.Lng)

	if myjson3.Name != "" {
		result.Name = myjson3.Name
	}
	if myjson3.Address != "" {
		result.Address = myjson3.Address
	}
	if myjson3.City != "" {
		result.City = myjson3.City
	}
	if myjson3.State != "" {
		result.State = myjson3.State
	}
	if myjson3.Zip != "" {
		result.Zip = myjson3.Zip
	}

	StartQuery := "http://maps.google.com/maps/api/geocode/json?address="
	WhereQuery := result.Address + " " + result.City + " " + result.State
	WhereQuery = strings.Replace(WhereQuery, " ", "+", -1)

	EndQuery := "&sensor=false"
	Url1 := StartQuery + WhereQuery + EndQuery
	fmt.Println("Published URL: " + Url1)

	res, err := http.Get(Url1)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	//fmt.Println(robots)
	res.Body.Close()
	//fmt.Println(robots)
	if err != nil {
		log.Fatal(err)
	}
	var myjsonresult1 MyJsonResult
	err = json.Unmarshal(robots, &myjsonresult1)
	if err != nil {
		log.Fatal(err)
	}
	//myjsonresult1.Results[0].Geometry
	fmt.Println("New latitude longiteude")

	fmt.Println(myjsonresult1.Results[0].Geometry.Location.Lat)
	fmt.Println(myjsonresult1.Results[0].Geometry.Location.Lng)

	//myjson3.Id = bson.NewObjectId()
	result.Coordinate.Lat = myjsonresult1.Results[0].Geometry.Location.Lat
	result.Coordinate.Lng = myjsonresult1.Results[0].Geometry.Location.Lng
	//end
	c.UpdateId(oid, result)
	//	err = c.Find(bson.M{"id": oid}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\nNew Name:", result.Name)
	fmt.Println("New Address:", result.Address)
	fmt.Println("New City:", result.City)
	fmt.Println("New State:", result.State)
	fmt.Println("New Zip:", result.Zip)
	fmt.Println("New latitude:", result.Coordinate.Lat)
	fmt.Println("New longitude:", result.Coordinate.Lng)

	fmt.Println("Id2:", result.Id.String())
	oid = bson.ObjectId(result.Id)

	b2, err := json.Marshal(result)
	if err != nil {
	}
	// fmt.Fprintf(rw, "New response")

	fmt.Fprintf(rw, string(b2))
	fmt.Println("Method Name: " + req.Method)
}

func DeleteLocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	fmt.Fprintf(rw, "Deleting Id, %s!\n", p.ByName("name"))

	session, err := mgo.Dial("mongodb://user1:pass1@ds045054.mongolab.com:45054/mydatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydatabase").C("people")
	id := p.ByName("name")
	//id := "562975762b877b26f82bf8c5"
	oid := bson.ObjectIdHex(id)
	//od := "562972862b877b2968b8a1a0"
	//var result MyJsonName
	c.RemoveId(oid)
	//c.FindId(oid).One(&result)
	//fmt.Fprintf(rw, result)
	//	err = c.Find(bson.M{"id": oid}).One(&result)
	fmt.Fprintf(rw, "Deleted, %s!\n", p.ByName("name"))

}

func main() {

	mux := httprouter.New()
	mux.GET("/locations/:name", Getlocations)
	mux.POST("/locations", Postlocations)
	mux.PUT("/locations/:name", PutLocations)
	mux.DELETE("/locations/:name", DeleteLocations)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
