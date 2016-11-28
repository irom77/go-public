package main
//#nohup ./pan2influx -api="" -p="" > /dev/null 2>&1 &
import (
	"log"
	"net/http"
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"io/ioutil"
	"github.com/influxdata/influxdb/client/v2"
	"time"
	"strconv"
	"flag"
	"fmt"
	"os"
)

const testVersion = 1

var (
	IP = flag.String("ip", "10.34.2.21", "PAN firewall IP address")
	SLEEP = flag.Duration("sleep", 60, "Polling time in sec")
	API = flag.String("api", "", "PAN firewall API Key")
	DBNAME = flag.String("d", "firewalls", "InfluxDB name")
	DBADDRESS = flag.String("a", "http://10.34.1.100:8086", "InfluxDB address")
	USERNAME = flag.String("u", "firewall", "InfluxDB username")
	PASSWORD = flag.String("p", "password", "InfluxDB password")
	SITE = flag.String("site", "DC1", "Site name")
	FW = flag.String("fw", "PAN2", "Firewall name")
	NODEID = flag.String("nodeid", "1", "Firewall node-id")
	version = flag.Bool("v", false, "Prints current version")
	//PRINT = flag.Bool("print", true, "print to console")
)
var (
	Version = "No Version Provided"
	BuildTime = ""
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2016 @IrekRomaniuk. All rights reserved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}
}

func main() {
	var API = "&key=" + *API
	var IP = "https://" + *IP + "/esp/restapi.esp?type=op"
	var DSP = []string{"dp0","dp1","dp2"}
	//var AE = []string{"ae1","ae2","ae3"}

	resourceMonitor(DSP, getHTML(IP + "&cmd=<show><running><resource-monitor><second></second></resource-monitor></running></show>" + API))
	//qosThroughput(AE, getHTML(IP + "&cmd=<show><qos><throughput>" + *NODEID + "</throughput><interface>" + "ae1" + "</interface></qos></show>" + API))

}

/*func qosThroughput () {
}*/

func getHTML (url string ) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil { log.Fatal(err) }

	htmlData, err := ioutil.ReadAll(resp.Body)
	/*o := "C:\\Users\\irekromaniuk\\Vagrant\\trusty64\\src\\github.com\\irom77\\go-public\\pan2influx\\output.txt"
	htmlData, err := ioutil.ReadFile(o)*/
	if err != nil { log.Fatal(err) }
	resp.Body.Close()
	return string(htmlData)
}

func resourceMonitor (DSP []string, htmlData string) {
	//parseGoQuery("dp2 pktlog_forwarding", string(htmlData))
	for _, dp := range DSP {
		for i:=0; i<= 11; i++ {
			toInflux(parseGoQuery(dp, "cpu-load-average value", "coreid", i, htmlData, "cpu_load"))
		}
		for i:=0; i<= 3; i++ {
			toInflux(parseGoQuery(dp, "resource-utilization value", "resourceid", i, htmlData, "resource_utilization"))
		}
	}
	//defer resp.Body.Close() // close Body when the function returns
	//fmt.Println(time.Now())
	//time.Sleep(*SLEEP * time.Second)
}

func parseGoQuery(dsp string, t string, id string, i int, b string, p string) (int, string, string, int, string) {
	tag := dsp + " " + t
	htmlCode := strings.NewReader(b)
	doc, err := goquery.NewDocumentFromReader(htmlCode)
	if err != nil { log.Fatal(err) }
	content := []string{}
	doc.Find(tag).Each(func(_ int, s *goquery.Selection) {  //resource-utilization
		content = append(content,s.Text())
	})
	return getMax(strings.Split(content[i],",")), dsp, id, i, p
}

func toInflux(value int, dsp string, id string, i int, p string) {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: *DBADDRESS,
		Username: *USERNAME,
		Password: *PASSWORD,
	})
	if err != nil { log.Fatalln("Error: ", err) }
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  *DBNAME,
		Precision: "s",
	})
	if err != nil { log.Fatalln("Error: ", err) }
	// Create a point and add to batch
	tags := map[string]string{"dsp": dsp, id:strconv.Itoa(i), "site":*SITE,"firewall":*FW }
	fields := map[string]interface{}{
		p:   value,
	}
	pt, err := client.NewPoint(p, tags, fields, time.Now())
	if err != nil { log.Fatalln("Error: ", err) }
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)

}

func getMax(arr []string) int {
	max, _ := strconv.Atoi(arr[0]) // assume first value is the smallest
	for _, value := range arr {
		valueint, _ := strconv.Atoi(value)
		if valueint > max {
			max = valueint // found another smaller value, replace previous value in max
		}
	}
	return max
}





