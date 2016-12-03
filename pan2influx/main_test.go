package main


import (
	"testing"
	"github.com/influxdata/influxdb/client/v2"
	"fmt"
)
const targetTestVersion = 1
var (
	max int
	dsp string
)

var GoQueryTests = []struct {
	tag string
	dsp string
	t string
	id int
	max int
}{
	{"dsp", "dp2", "cpu-load-average value", 0, 89},
	{"dsp", "dp1", "resource-utilization value", 1, 20},
	{"dsp", "dp1", "resource-utilization value", 0, 12},
	{"dsp", "dp0", "cpu-load-average value", 2, 57},

}
func TestParseInfoSession (t *testing.T) {
	if con, _ := parseSessionInfo("num-active",HTMLinfosession,"whatever"); con != "102448" {
		t.Fatalf("Active session count: got %s, want %s",con, "102448")
	}
}

func TestParseGoQuery (t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, test := range GoQueryTests {
		max, _, dsp, _, _, _ = parseResourceMonitor(test.tag, test.dsp, test.t, "whatever", test.id, htmlData, "whatever")
		if  max != test.max || dsp != test.dsp {
			t.Fatalf("%s %s %s %d: got %s %d, want %s %d",test.tag, test.dsp, test.t, test.id, dsp, max, test.dsp, test.max )
		}
	}

}

func TestQosThroughput (t *testing.T) {
	class, _ := parseThroughput("result",HTMLThroughput2,"whatever")
	//fmt.Println(class)
	if class[3] != "130784" {
		t.Fatalf("got %s, want %s",class[3],"130784"  )
	}
}

var GoInfluxTests = []struct {
	field string
	value int
	tag string
	id int
	limit int
}{
	{"TestMeasurement",44,"TESTid",99,1},

}

func TestToInflux (t *testing.T) {
	//value int, tagid string, tag string, id string, i int, p string
	toInflux(44,"TEST","TESTdsp","TESTid",199,"TestMeasurement") //"TestToInfluxValue":44,"TEST":"dpTEST","id":0,"site":*SITE,"firewall":*FW
	//select * from Test_Measurement where TESTid='99' limit 1
	q := fmt.Sprintf("SELECT * FROM %s WHERE %s AND time > now() - 3s LIMIT %d", "TestMeasurement","TESTid='199'", 1)
	//fmt.Println(q)
	// Make client
	client, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: *DBADDRESS,
		Username: *USERNAME,
		Password: *PASSWORD,
	})
	if err != nil { t.Fatalf("Error: ", err) }
	defer client.Close()
	res, err := queryDB(client, q)
	if err != nil { t.Fatalf("Error: ", err) }
	read := res[0].Series[0].Values[0][2]
	if read != "199" {
		t.Fatalf("Influx: read %s, wrote %s", read,"199")
	}

}

func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: *DBNAME,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

var htmlData = `<response status="success"><result>
<resource-monitor>
<data-processors>
<dp2>
<second>
<cpu-load-average>
<entry>
<coreid>0</coreid>
<value>0,0,0,89,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<coreid>1</coreid>
<value>5,5,10,6,5,5,6,6,5,6,6,5,5,5,5,4,5,5,5,5,5,5,5,6,5,5,5,5,5,5,5,5,5,5,4,5,5,4,4,5,5,4,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4</value>
</entry>
<entry>
<coreid>2</coreid>
<value>7,7,7,8,7,6,7,8,6,7,7,7,6,6,6,6,6,7,6,6,6,6,7,7,7,6,6,6,6,6,6,6,7,6,6,6,6,5,6,6,6,6,6,6,5,5,6,5,5,5,5,5,6,6,5,5,5,6,6,5</value>
</entry>
<entry>
<coreid>3</coreid>
<value>11,10,10,12,12,11,13,13,11,13,13,12,11,11,10,10,11,12,10,10,10,10,11,12,11,10,9,9,10,10,11,11,11,11,10,10,9,8,9,10,10,10,9,9,9,9,9,8,8,8,8,9,9,8,7,7,8,8,8,9</value>
</entry>
<entry>
<coreid>4</coreid>
<value>11,10,10,11,11,11,13,13,10,12,13,12,11,10,10,9,10,11,9,9,10,9,10,11,10,10,9,10,10,9,10,10,11,10,9,10,9,8,8,10,10,9,9,9,8,8,9,8,8,8,8,9,9,8,7,7,8,8,8,9</value>
</entry>
<entry>
<coreid>5</coreid>
<value>17,16,17,19,18,17,20,20,16,19,20,18,17,16,16,16,17,17,15,14,14,14,17,17,16,16,15,15,15,14,16,17,17,16,14,14,13,11,13,16,18,15,14,14,13,13,14,13,12,12,12,13,14,13,11,11,13,13,13,13</value>
</entry>
<entry>
<coreid>6</coreid>
<value>16,15,16,18,18,17,20,19,15,17,18,16,15,15,14,14,15,16,14,13,13,13,15,15,14,14,13,13,14,14,15,16,16,15,13,13,13,11,12,14,16,14,12,14,13,12,13,11,10,11,11,12,13,12,10,11,12,12,12,12</value>
</entry>
<entry>
<coreid>7</coreid>
<value>15,15,15,16,16,15,19,18,14,16,17,16,15,14,13,14,15,16,13,13,12,13,15,14,14,13,12,12,13,12,14,15,15,14,12,12,12,11,12,14,15,13,11,12,12,11,12,11,10,10,11,12,12,11,10,10,10,11,11,11</value>
</entry>
<entry>
<coreid>8</coreid>
<value>15,14,14,15,15,15,18,18,13,16,17,15,14,13,13,13,14,15,12,12,12,12,14,14,13,13,12,12,13,12,13,14,14,13,12,12,12,9,10,13,14,12,11,12,11,11,12,11,9,10,11,11,12,11,9,9,10,10,10,10</value>
</entry>
<entry>
<coreid>9</coreid>
<value>14,13,13,15,15,14,17,17,13,15,16,14,14,13,12,12,14,15,12,11,11,11,13,13,12,12,11,11,12,11,13,14,14,13,11,11,11,9,10,13,14,11,11,12,10,11,11,10,9,9,9,10,11,10,9,9,9,10,10,10</value>
</entry>
<entry>
<coreid>10</coreid>
<value>13,13,13,15,14,14,17,16,13,15,15,13,13,13,12,11,13,13,11,11,11,10,13,13,12,12,11,11,11,11,12,13,13,12,11,11,11,8,9,12,13,10,9,10,10,10,11,9,9,9,9,10,11,10,9,8,9,9,10,10</value>
</entry>
<entry>
<coreid>11</coreid>
<value>5,5,5,5,5,5,6,6,4,5,5,5,5,4,4,4,4,5,4,4,4,4,5,5,5,4,4,4,4,4,4,5,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4</value>
</entry>
</cpu-load-average>
<task>
<pktlog_forwarding>5%</pktlog_forwarding>
<flow_lookup>12%</flow_lookup>
<flow_fastpath>12%</flow_fastpath>
<flow_np>12%</flow_np>
<aho_result>15%</aho_result>
<zip_result>14%</zip_result>
<flow_host>7%</flow_host>
<flow_forwarding>12%</flow_forwarding>
<module_internal>12%</module_internal>
<flow_ctrl>7%</flow_ctrl>
<lwm>0%</lwm>
<flow_slowpath>12%</flow_slowpath>
<dfa_result>14%</dfa_result>
<nac_result>14%</nac_result>
<flow_mgmt>5%</flow_mgmt>
</task>
<cpu-load-maximum>
<entry>
<coreid>0</coreid>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<coreid>1</coreid>
<value>5,5,5,6,5,5,6,6,5,6,6,5,5,5,5,4,5,5,5,5,5,5,5,6,5,5,5,5,5,5,5,5,5,5,4,5,5,4,4,5,5,4,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4</value>
</entry>
<entry>
<coreid>2</coreid>
<value>7,7,7,8,7,6,7,8,6,7,7,7,6,6,6,6,6,7,6,6,6,6,7,7,7,6,6,6,6,6,6,6,7,6,6,6,6,5,6,6,6,6,6,6,5,5,6,5,5,5,5,5,6,6,5,5,5,6,6,5</value>
</entry>
<entry>
<coreid>3</coreid>
<value>11,10,10,12,12,11,13,13,11,13,13,12,11,11,10,10,11,12,10,10,10,10,11,12,11,10,9,9,10,10,11,11,11,11,10,10,9,8,9,10,10,10,9,9,9,9,9,8,8,8,8,9,9,8,7,7,8,8,8,9</value>
</entry>
<entry>
<coreid>4</coreid>
<value>11,10,10,11,11,11,13,13,10,12,13,12,11,10,10,9,10,11,9,9,10,9,10,11,10,10,9,10,10,9,10,10,11,10,9,10,9,8,8,10,10,9,9,9,8,8,9,8,8,8,8,9,9,8,7,7,8,8,8,9</value>
</entry>
<entry>
<coreid>5</coreid>
<value>17,16,17,19,18,17,20,20,16,19,20,18,17,16,16,16,17,17,15,14,14,14,17,17,16,16,15,15,15,14,16,17,17,16,14,14,13,11,13,16,18,15,14,14,13,13,14,13,12,12,12,13,14,13,11,11,13,13,13,13</value>
</entry>
<entry>
<coreid>6</coreid>
<value>16,15,16,18,18,17,20,19,15,17,18,16,15,15,14,14,15,16,14,13,13,13,15,15,14,14,13,13,14,14,15,16,16,15,13,13,13,11,12,14,16,14,12,14,13,12,13,11,10,11,11,12,13,12,10,11,12,12,12,12</value>
</entry>
<entry>
<coreid>7</coreid>
<value>15,15,15,16,16,15,19,18,14,16,17,16,15,14,13,14,15,16,13,13,12,13,15,14,14,13,12,12,13,12,14,15,15,14,12,12,12,11,12,14,15,13,11,12,12,11,12,11,10,10,11,12,12,11,10,10,10,11,11,11</value>
</entry>
<entry>
<coreid>8</coreid>
<value>15,14,14,15,15,15,18,18,13,16,17,15,14,13,13,13,14,15,12,12,12,12,14,14,13,13,12,12,13,12,13,14,14,13,12,12,12,9,10,13,14,12,11,12,11,11,12,11,9,10,11,11,12,11,9,9,10,10,10,10</value>
</entry>
<entry>
<coreid>9</coreid>
<value>14,13,13,15,15,14,17,17,13,15,16,14,14,13,12,12,14,15,12,11,11,11,13,13,12,12,11,11,12,11,13,14,14,13,11,11,11,9,10,13,14,11,11,12,10,11,11,10,9,9,9,10,11,10,9,9,9,10,10,10</value>
</entry>
<entry>
<coreid>10</coreid>
<value>13,13,13,15,14,14,17,16,13,15,15,13,13,13,12,11,13,13,11,11,11,10,13,13,12,12,11,11,11,11,12,13,13,12,11,11,11,8,9,12,13,10,9,10,10,10,11,9,9,9,9,10,11,10,9,8,9,9,10,10</value>
</entry>
<entry>
<coreid>11</coreid>
<value>5,5,5,5,5,5,6,6,4,5,5,5,5,4,4,4,4,5,4,4,4,4,5,5,5,4,4,4,4,4,4,5,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4</value>
</entry>
</cpu-load-maximum>
<resource-utilization>
<entry>
<name>session</name>
<value>2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2</value>
</entry>
<entry>
<name>packet buffer</name>
<value>1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1</value>
</entry>
<entry>
<name>packet descriptor</name>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<name>packet descriptor (on-chip)</name>
<value>1,1,1,1,1,1,1,1,1,1,1,2,2,1,2,2,1,1,1,1,1,1,1,1,1,1,1,1,2,1,1,1,1,1,1,1,1,2,2,1,1,2,2,2,2,2,1,1,1,1,1,2,1,2,2,2,2,2,2,1</value>
</entry>
</resource-utilization>
</second>
</dp2>
<dp1>
<second>
<cpu-load-average>
<entry>
<coreid>0</coreid>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<coreid>1</coreid>
<value>4,4,4,4,4,4,4,4,4,4,4,4,4,4,5,4,4,4,5,5,5,5,5,4,4,5,4,4,4,4,4,4,4,5,5,4,4,4,4,4,4,4,3,4,3,3,4,4,3,4,4,8,9,8,6,10,10,6,4,4</value>
</entry>
<entry>
<coreid>2</coreid>
<value>5,6,5,5,6,5,6,5,5,6,6,5,5,5,6,5,5,6,7,6,6,6,6,6,6,6,5,5,5,5,5,5,6,6,6,5,6,5,5,5,5,5,4,5,5,5,5,5,5,5,6,9,10,9,7,10,11,7,5,5</value>
</entry>
<entry>
<coreid>3</coreid>
<value>9,9,9,9,10,9,10,9,7,11,10,10,9,9,10,8,9,11,11,10,10,10,11,10,9,10,9,8,10,8,8,8,9,10,10,9,10,9,7,7,9,7,7,8,7,7,8,10,9,9,10,11,15,12,9,12,12,9,7,8</value>
</entry>
<entry>
<coreid>4</coreid>
<value>9,9,8,9,9,9,9,8,8,10,10,10,8,9,9,9,10,10,10,9,10,10,10,9,9,11,9,8,9,7,8,8,9,9,9,9,10,8,7,8,9,7,7,8,7,7,8,11,9,9,9,11,13,11,9,11,11,9,7,8</value>
</entry>
<entry>
<coreid>5</coreid>
<value>13,12,12,13,14,12,13,12,11,14,14,14,12,13,14,12,14,14,14,12,14,14,15,13,12,15,13,11,13,11,11,12,12,13,14,12,13,11,10,10,12,10,9,11,10,10,11,13,12,11,13,14,17,13,12,14,14,12,10,10</value>
</entry>
<entry>
<coreid>6</coreid>
<value>12,11,11,11,13,11,12,11,10,13,13,13,11,12,12,11,12,14,14,12,13,14,14,13,12,13,12,11,12,10,11,11,12,12,13,12,13,11,10,10,12,9,9,11,10,9,10,11,11,10,12,13,15,13,11,13,13,11,9,10</value>
</entry>
<entry>
<coreid>7</coreid>
<value>11,11,11,11,12,11,11,11,10,12,12,12,10,11,12,10,12,13,13,11,13,13,12,11,11,13,11,10,11,9,10,10,10,11,12,11,12,10,9,10,11,9,9,9,9,9,9,11,11,10,11,11,16,12,10,12,12,10,9,9</value>
</entry>
<entry>
<coreid>8</coreid>
<value>10,10,10,10,10,10,11,10,9,11,12,12,10,10,11,10,11,12,12,10,12,12,12,11,10,13,11,9,10,8,9,10,10,11,11,10,11,10,9,10,10,8,8,9,9,9,9,11,10,10,11,11,13,11,10,11,11,10,8,9</value>
</entry>
<entry>
<coreid>9</coreid>
<value>10,10,10,9,11,9,10,9,8,11,12,11,9,10,11,9,11,12,12,10,11,11,12,10,10,12,9,9,10,8,9,9,10,10,10,10,10,9,8,8,10,8,8,9,8,8,9,10,9,9,10,10,12,11,9,10,10,9,8,8</value>
</entry>
<entry>
<coreid>10</coreid>
<value>9,9,9,9,10,9,10,9,8,11,11,10,9,9,10,9,10,11,11,9,11,11,11,10,9,11,9,8,9,8,8,9,9,10,10,9,10,9,8,8,10,7,7,8,8,7,8,10,9,9,9,10,13,10,8,10,10,9,7,7</value>
</entry>
<entry>
<coreid>11</coreid>
<value>3,3,3,3,4,4,4,4,3,4,4,4,3,4,4,3,4,4,4,4,4,4,4,4,4,4,3,3,4,3,3,3,4,4,4,4,4,4,3,3,3,3,3,3,3,3,3,4,3,4,4,5,6,5,4,5,6,4,3,3</value>
</entry>
</cpu-load-average>
<task>
<pktlog_forwarding>3%</pktlog_forwarding>
<flow_lookup>9%</flow_lookup>
<flow_fastpath>9%</flow_fastpath>
<flow_np>9%</flow_np>
<aho_result>11%</aho_result>
<zip_result>10%</zip_result>
<flow_host>6%</flow_host>
<flow_forwarding>9%</flow_forwarding>
<module_internal>9%</module_internal>
<flow_ctrl>6%</flow_ctrl>
<lwm>0%</lwm>
<flow_slowpath>9%</flow_slowpath>
<dfa_result>10%</dfa_result>
<nac_result>10%</nac_result>
<flow_mgmt>4%</flow_mgmt>
</task>
<cpu-load-maximum>
<entry>
<coreid>0</coreid>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<coreid>1</coreid>
<value>4,4,4,4,4,4,4,4,4,4,4,4,4,4,5,4,4,4,5,5,5,5,5,4,4,5,4,4,4,4,4,4,4,5,5,4,4,4,4,4,4,4,3,4,3,3,4,4,3,4,4,8,9,8,6,10,10,6,4,4</value>
</entry>
<entry>
<coreid>2</coreid>
<value>5,6,5,5,6,5,6,5,5,6,6,5,5,5,6,5,5,6,7,6,6,6,6,6,6,6,5,5,5,5,5,5,6,6,6,5,6,5,5,5,5,5,4,5,5,5,5,5,5,5,6,9,10,9,7,10,11,7,5,5</value>
</entry>
<entry>
<coreid>3</coreid>
<value>9,9,9,9,10,9,10,9,7,11,10,10,9,9,10,8,9,11,11,10,10,10,11,10,9,10,9,8,10,8,8,8,9,10,10,9,10,9,7,7,9,7,7,8,7,7,8,10,9,9,10,11,15,12,9,12,12,9,7,8</value>
</entry>
<entry>
<coreid>4</coreid>
<value>9,9,8,9,9,9,9,8,8,10,10,10,8,9,9,9,10,10,10,9,10,10,10,9,9,11,9,8,9,7,8,8,9,9,9,9,10,8,7,8,9,7,7,8,7,7,8,11,9,9,9,11,13,11,9,11,11,9,7,8</value>
</entry>
<entry>
<coreid>5</coreid>
<value>13,12,12,13,14,12,13,12,11,14,14,14,12,13,14,12,14,14,14,12,14,14,15,13,12,15,13,11,13,11,11,12,12,13,14,12,13,11,10,10,12,10,9,11,10,10,11,13,12,11,13,14,17,13,12,14,14,12,10,10</value>
</entry>
<entry>
<coreid>6</coreid>
<value>12,11,11,11,13,11,12,11,10,13,13,13,11,12,12,11,12,14,14,12,13,14,14,13,12,13,12,11,12,10,11,11,12,12,13,12,13,11,10,10,12,9,9,11,10,9,10,11,11,10,12,13,15,13,11,13,13,11,9,10</value>
</entry>
<entry>
<coreid>7</coreid>
<value>11,11,11,11,12,11,11,11,10,12,12,12,10,11,12,10,12,13,13,11,13,13,12,11,11,13,11,10,11,9,10,10,10,11,12,11,12,10,9,10,11,9,9,9,9,9,9,11,11,10,11,11,16,12,10,12,12,10,9,9</value>
</entry>
<entry>
<coreid>8</coreid>
<value>10,10,10,10,10,10,11,10,9,11,12,12,10,10,11,10,11,12,12,10,12,12,12,11,10,13,11,9,10,8,9,10,10,11,11,10,11,10,9,10,10,8,8,9,9,9,9,11,10,10,11,11,13,11,10,11,11,10,8,9</value>
</entry>
<entry>
<coreid>9</coreid>
<value>10,10,10,9,11,9,10,9,8,11,12,11,9,10,11,9,11,12,12,10,11,11,12,10,10,12,9,9,10,8,9,9,10,10,10,10,10,9,8,8,10,8,8,9,8,8,9,10,9,9,10,10,12,11,9,10,10,9,8,8</value>
</entry>
<entry>
<coreid>10</coreid>
<value>9,9,9,9,10,9,10,9,8,11,11,10,9,9,10,9,10,11,11,9,11,11,11,10,9,11,9,8,9,8,8,9,9,10,10,9,10,9,8,8,10,7,7,8,8,7,8,10,9,9,9,10,13,10,8,10,10,9,7,7</value>
</entry>
<entry>
<coreid>11</coreid>
<value>3,3,3,3,4,4,4,4,3,4,4,4,3,4,4,3,4,4,4,4,4,4,4,4,4,4,3,3,4,3,3,3,4,4,4,4,4,4,3,3,3,3,3,3,3,3,3,4,3,4,4,5,6,5,4,5,6,4,3,3</value>
</entry>
</cpu-load-maximum>
<resource-utilization>
<entry>
<name>session</name>
<value>2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,12,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2</value>
</entry>
<entry>
<name>packet buffer</name>
<value>1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,20,1</value>
</entry>
<entry>
<name>packet descriptor</name>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,1,0,0</value>
</entry>
<entry>
<name>packet descriptor (on-chip)</name>
<value>2,2,1,2,1,1,1,2,2,1,1,1,1,1,1,1,1,1,2,2,2,1,1,1,1,2,1,1,1,1,1,1,1,1,1,2,1,2,1,1,1,2,2,2,1,1,1,2,1,2,1,1,1,1,2,2,1,2,1,1</value>
</entry>
</resource-utilization>
</second>
</dp1>
<dp0>
<second>
<cpu-load-average>
<entry>
<coreid>0</coreid>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<coreid>1</coreid>
<value>3,4,3,3,3,3,5,3,3,3,3,3,4,4,3,3,3,4,3,4,4,3,4,3,3,3,3,3,3,3,3,4,4,4,3,4,3,3,3,3,3,3,3,3,3,2,3,3,3,3,2,3,3,3,3,3,3,3,3,3</value>
</entry>
<entry>
<coreid>2</coreid>
<value>6,7,6,6,7,6,8,7,6,7,6,7,7,57,6,6,6,7,7,7,7,6,7,6,7,6,6,6,6,6,6,7,7,7,6,7,7,7,6,6,6,6,6,6,6,6,6,6,6,6,5,6,6,6,6,6,7,7,6,6</value>
</entry>
<entry>
<coreid>3</coreid>
<value>8,10,7,8,9,8,11,9,8,9,8,9,8,10,8,7,7,9,8,8,8,8,9,8,8,8,8,8,8,7,8,9,9,9,8,9,9,8,8,8,8,7,8,8,7,6,8,8,7,7,7,8,8,7,8,8,9,8,8,7</value>
</entry>
<entry>
<coreid>4</coreid>
<value>7,9,8,8,8,8,11,8,7,9,8,9,9,10,8,8,7,9,9,8,8,7,9,8,9,8,7,8,7,7,7,9,9,9,8,9,9,8,8,8,8,7,8,8,7,7,8,8,7,7,6,8,7,7,7,7,9,8,8,7</value>
</entry>
<entry>
<coreid>5</coreid>
<value>9,15,11,12,13,10,15,12,10,13,12,13,12,14,12,10,9,12,12,11,11,11,13,11,12,12,10,12,11,10,11,13,12,12,11,13,12,10,11,11,12,10,11,11,9,9,12,12,9,9,8,11,10,9,10,10,13,12,11,9</value>
</entry>
<entry>
<coreid>6</coreid>
<value>9,14,10,11,12,10,13,11,9,12,11,12,11,13,11,9,9,11,10,9,10,10,12,10,11,10,10,10,10,9,10,12,11,11,10,13,12,10,10,9,11,9,9,10,9,8,10,11,9,8,7,10,8,9,10,9,12,11,10,9</value>
</entry>
<entry>
<coreid>7</coreid>
<value>8,13,9,10,11,10,13,10,9,12,9,12,11,13,10,9,8,11,10,9,10,9,11,9,11,11,9,10,10,9,9,11,11,11,10,12,11,9,10,9,10,8,9,9,8,7,9,10,9,8,7,10,9,9,9,9,11,10,10,8</value>
</entry>
<entry>
<coreid>8</coreid>
<value>8,12,9,9,11,9,13,10,8,10,10,11,10,12,9,9,7,10,10,9,9,9,11,9,10,9,8,9,9,8,9,11,10,11,9,11,10,9,9,9,10,7,9,9,7,7,10,10,8,7,7,8,8,8,9,8,10,10,9,8</value>
</entry>
<entry>
<coreid>9</coreid>
<value>7,12,7,9,11,9,12,9,8,11,9,11,10,11,9,8,7,10,9,8,9,9,10,9,9,9,8,8,8,7,9,10,10,10,8,11,10,9,9,9,9,8,8,8,8,7,9,9,7,7,7,8,8,7,8,8,10,9,9,8</value>
</entry>
<entry>
<coreid>10</coreid>
<value>7,12,8,8,10,8,12,9,8,10,9,10,9,11,9,7,7,9,9,8,9,9,10,9,9,8,8,9,8,7,8,10,9,9,8,10,9,8,9,8,9,7,8,8,7,7,9,9,7,7,6,8,7,8,7,8,10,9,9,7</value>
</entry>
<entry>
<coreid>11</coreid>
<value>4,5,4,4,4,4,6,5,4,5,4,4,4,4,4,4,4,5,4,4,4,4,5,4,5,4,4,5,4,4,4,5,5,5,4,5,5,4,4,4,4,3,4,4,4,3,4,4,4,4,3,4,4,4,4,4,5,5,4,4</value>
</entry>
</cpu-load-average>
<task>
<pktlog_forwarding>5%</pktlog_forwarding>
<flow_lookup>10%</flow_lookup>
<flow_fastpath>9%</flow_fastpath>
<flow_np>9%</flow_np>
<aho_result>12%</aho_result>
<zip_result>11%</zip_result>
<flow_host>7%</flow_host>
<flow_forwarding>10%</flow_forwarding>
<module_internal>10%</module_internal>
<flow_ctrl>7%</flow_ctrl>
<lwm>0%</lwm>
<flow_slowpath>10%</flow_slowpath>
<dfa_result>11%</dfa_result>
<nac_result>11%</nac_result>
<flow_mgmt>4%</flow_mgmt>
</task>
<cpu-load-maximum>
<entry>
<coreid>0</coreid>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<coreid>1</coreid>
<value>3,4,3,3,3,3,5,3,3,3,3,3,4,4,3,3,3,4,3,4,4,3,4,3,3,3,3,3,3,3,3,4,4,4,3,4,3,3,3,3,3,3,3,3,3,2,3,3,3,3,2,3,3,3,3,3,3,3,3,3</value>
</entry>
<entry>
<coreid>2</coreid>
<value>6,7,6,6,7,6,8,7,6,7,6,7,7,7,6,6,6,7,7,7,7,6,7,6,7,6,6,6,6,6,6,7,7,7,6,7,7,7,6,6,6,6,6,6,6,6,6,6,6,6,5,6,6,6,6,6,7,7,6,6</value>
</entry>
<entry>
<coreid>3</coreid>
<value>8,10,7,8,9,8,11,9,8,9,8,9,8,10,8,7,7,9,8,8,8,8,9,8,8,8,8,8,8,7,8,9,9,9,8,9,9,8,8,8,8,7,8,8,7,6,8,8,7,7,7,8,8,7,8,8,9,8,8,7</value>
</entry>
<entry>
<coreid>4</coreid>
<value>7,9,8,8,8,8,11,8,7,9,8,9,9,10,8,8,7,9,9,8,8,7,9,8,9,8,7,8,7,7,7,9,9,9,8,9,9,8,8,8,8,7,8,8,7,7,8,8,7,7,6,8,7,7,7,7,9,8,8,7</value>
</entry>
<entry>
<coreid>5</coreid>
<value>9,15,11,12,13,10,15,12,10,13,12,13,12,14,12,10,9,12,12,11,11,11,13,11,12,12,10,12,11,10,11,13,12,12,11,13,12,10,11,11,12,10,11,11,9,9,12,12,9,9,8,11,10,9,10,10,13,12,11,9</value>
</entry>
<entry>
<coreid>6</coreid>
<value>9,14,10,11,12,10,13,11,9,12,11,12,11,13,11,9,9,11,10,9,10,10,12,10,11,10,10,10,10,9,10,12,11,11,10,13,12,10,10,9,11,9,9,10,9,8,10,11,9,8,7,10,8,9,10,9,12,11,10,9</value>
</entry>
<entry>
<coreid>7</coreid>
<value>8,13,9,10,11,10,13,10,9,12,9,12,11,13,10,9,8,11,10,9,10,9,11,9,11,11,9,10,10,9,9,11,11,11,10,12,11,9,10,9,10,8,9,9,8,7,9,10,9,8,7,10,9,9,9,9,11,10,10,8</value>
</entry>
<entry>
<coreid>8</coreid>
<value>8,12,9,9,11,9,13,10,8,10,10,11,10,12,9,9,7,10,10,9,9,9,11,9,10,9,8,9,9,8,9,11,10,11,9,11,10,9,9,9,10,7,9,9,7,7,10,10,8,7,7,8,8,8,9,8,10,10,9,8</value>
</entry>
<entry>
<coreid>9</coreid>
<value>7,12,7,9,11,9,12,9,8,11,9,11,10,11,9,8,7,10,9,8,9,9,10,9,9,9,8,8,8,7,9,10,10,10,8,11,10,9,9,9,9,8,8,8,8,7,9,9,7,7,7,8,8,7,8,8,10,9,9,8</value>
</entry>
<entry>
<coreid>10</coreid>
<value>7,12,8,8,10,8,12,9,8,10,9,10,9,11,9,7,7,9,9,8,9,9,10,9,9,8,8,9,8,7,8,10,9,9,8,10,9,8,9,8,9,7,8,8,7,7,9,9,7,7,6,8,7,8,7,8,10,9,9,7</value>
</entry>
<entry>
<coreid>11</coreid>
<value>4,5,4,4,4,4,6,5,4,5,4,4,4,4,4,4,4,5,4,4,4,4,5,4,5,4,4,5,4,4,4,5,5,5,4,5,5,4,4,4,4,3,4,4,4,3,4,4,4,4,3,4,4,4,4,4,5,5,4,4</value>
</entry>
</cpu-load-maximum>
<resource-utilization>
<entry>
<name>session</name>
<value>7,7,7,7,7,87,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7</value>
</entry>
<entry>
<name>packet buffer</name>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,20</value>
</entry>
<entry>
<name>packet descriptor</name>
<value>0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0</value>
</entry>
<entry>
<name>packet descriptor (on-chip)</name>
<value>2,1,1,1,1,1,2,2,1,1,1,1,1,1,1,1,1,1,1,1,1,2,2,1,1,2,1,1,2,2,1,1,1,1,1,2,1,1,2,1,1,1,1,1,1,2,2,1,2,2,1,1,1,1,2,1,1,1,1,1</value>
</entry>
</resource-utilization>
</second>
</dp0>
</data-processors>
</resource-monitor>
</result></response>`

var HTMLThroughput1 = `<response status="success">
<result>
Class 1 0 kbps Class 2 0 kbps Class 3 0 kbps Class 4 130784 kbps Class 5 0 kbps Class 6 0 kbps Class 7 20 kbps Class 8 12 kbps
</result>
</response>`

var HTMLThroughput2 = `<response status="success"><result>Class 1              0 kbps
Class 2              0 kbps
Class 3              0 kbps
Class 4         130784 kbps
Class 5              0 kbps
Class 6              0 kbps
Class 7             20 kbps
Class 8             12 kbps
</result></response>`

var HTMLinfosession = `<response status="success"><result>
<tmo-udp>30</tmo-udp>
<tcp-nonsyn-rej>False</tcp-nonsyn-rej>
<tmo-tcp>3600</tmo-tcp>
<pps>1124</pps>
<num-max>4194302</num-max>
<age-scan-thresh>80</age-scan-thresh>
<tmo-tcphalfclosed>120</tmo-tcphalfclosed>
<num-active>102448</num-active>
<dis-def>60</dis-def>
<num-mcast>0</num-mcast>
<icmp-unreachable-rate>200</icmp-unreachable-rate>
<tmo-tcptimewait>15</tmo-tcptimewait>
<age-scan-ssf>8</age-scan-ssf>
<vardata-rate>10485760</vardata-rate>
<age-scan-tmo>10</age-scan-tmo>
<tmo-tcpinit>5</tmo-tcpinit>
<dp>*.dp0</dp>
<dis-tcp>90</dis-tcp>
<num-udp>29570</num-udp>
<tmo-icmp>6</tmo-icmp>
<max-pending-mcast>0</max-pending-mcast>
<age-accel-thresh>80</age-accel-thresh>
<tmo-tcphandshake>10</tmo-tcphandshake>
<oor-action>drop</oor-action>
<tmo-def>30</tmo-def>
<age-accel-en>True</age-accel-en>
<age-accel-tsf>2</age-accel-tsf>
<hw-offload>True</hw-offload>
<num-icmp>114</num-icmp>
<num-predict>0</num-predict>
<tmo-cp>30</tmo-cp>
<strict-checksum>True</strict-checksum>
<tmo-tcp-unverif-rst>30</tmo-tcp-unverif-rst>
<num-bcast>0</num-bcast>
<ipv6-fw>True</ipv6-fw>
<num-installed>4142667796</num-installed>
<num-tcp>72139</num-tcp>
<dis-udp>60</dis-udp>
<cps>750</cps>
<kbps>1519</kbps>
</result></response>`




