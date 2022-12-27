package golangsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type DNSMethods struct {
	Client *Client
}

func (dns *DNSMethods) GetZoneByID(id string) ([]DNSZone, error) {
	body, err := dns.Client.doRequest("GET", "/dns/zones?id="+id, nil)
	if err != nil {
		return nil, err
	}

	zones := []DNSZone{}
	err = json.Unmarshal(body, &zones)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

//	client.DNS.GetZones(&map[string]string{
//		"offset": "2",
//		"limit": "10",
//	})
func (dns *DNSMethods) GetZones(filters *map[string]string) ([]DNSZone, error) {
	var fquery string
	if filters != nil {
		values := url.Values{}
		for key, value := range *filters {
			values.Set(key, value)
		}

		fquery = "?" + values.Encode()
	}
	body, err := dns.Client.doRequest("GET", "/dns/zones"+fquery, nil)
	if err != nil {
		return nil, err
	}

	zones := []DNSZone{}
	err = json.Unmarshal(body, &zones)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (dns *DNSMethods) AddZone(zone DNSZone) (*DNSZone, error) {
	reqbody, err := json.Marshal(zone)
	if err != nil {
		return nil, err
	}

	zone_type := "forward"
	if zone.ZoneType == "r" {
		zone_type = "reverse"
	}

	body, err := dns.Client.doRequest("POST", "/dns/zones/"+zone_type, bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_record DNSZone
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

func (dns *DNSMethods) UpdateZone(zone DNSZone) (*DNSZone, error) {
	reqbody, err := json.Marshal(zone)
	if err != nil {
		return nil, err
	}

	body, err := dns.Client.doRequest("PATCH", "/dns/zones/"+string(zone.ID), bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_record DNSZone
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

func (dns *DNSMethods) DeleteZone(zone DNSZone) (*DNSZone, error) {
	return dns.Client.DNS.DeleteZoneByID(string(zone.ID))
}

func (dns *DNSMethods) DeleteZoneByID(zoneId string) (*DNSZone, error) {

	body, err := dns.Client.doRequest("DELETE", "/dns/zones/"+zoneId, nil)
	if err != nil {
		return nil, err
	}

	var resp_record DNSZone
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

//	client.DNS.GetZoneRecords("428964", &map[string]string{
//		"offset": "2",
//		"limit": "10",
//	})
func (dns *DNSMethods) GetZoneRecords(zone_id string, filters *map[string]string) ([]DNSRecord, error) {
	var fquery string
	if filters != nil {
		values := url.Values{}
		for key, value := range *filters {
			values.Set(key, value)
		}

		fquery = "?" + values.Encode()
	}

	body, err := dns.Client.doRequest("GET", "/dns/zones/"+zone_id+"/records"+fquery, nil)
	if err != nil {
		return nil, err
	}

	records := []DNSRecord{}
	err = json.Unmarshal(body, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

//	newRecord := golangsdk.DNSRecord{
//		ParentID:    "428964",
//		Name:        "Golang TXT Record",
//		RecordType:  "TXT",
//		RecordHost:  "golang.example.com.",
//		RecordValue: "Created from GOLANGSDK 13",
//		RecordTTL:   3600,
//	}
//
// records, err := client.DNS.AddZoneRecord(newRecord)
func (dns *DNSMethods) AddZoneRecord(record DNSRecord) (*DNSRecord, error) {
	reqbody, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	body, err := dns.Client.doRequest("POST", "/dns/zones/"+string(record.ParentID)+"/records", bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_record DNSRecord
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

func (dns *DNSMethods) UpdateZoneRecord(record DNSRecord) (*DNSRecord, error) {
	reqbody, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	body, err := dns.Client.doRequest("PATCH", "/dns/zones/"+string(record.ParentID)+"/records/"+string(record.ID), bytes.NewBuffer(reqbody))
	if err != nil {
		return nil, err
	}

	var resp_record DNSRecord
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

func (dns *DNSMethods) DeleteZoneRecord(record DNSRecord) (*DNSRecord, error) {
	return dns.Client.DNS.DeleteZoneRecordByID(string(record.ParentID), string(record.ID))
}

func (dns *DNSMethods) DeleteZoneRecordByID(zoneId, recordId string) (*DNSRecord, error) {

	body, err := dns.Client.doRequest("DELETE", "/dns/zones/"+zoneId+"/records/"+recordId, nil)
	if err != nil {
		return nil, err
	}

	var resp_record DNSRecord
	err = json.Unmarshal(body, &resp_record)
	if err != nil {
		return nil, err
	}

	return &resp_record, nil
}

//Push Calls

func (dns *DNSMethods) PushZoneByID(zoneId string) (*string, error) {

	body, err := dns.Client.doRequest("POST", "/dns/zones/"+zoneId+"/push", nil)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return nil, err
	}

	var pid string
	if val, ok := jsonMap["pid"]; ok {
		pid = fmt.Sprintf("%v", val)
	}

	return &pid, nil
}

func (dns *DNSMethods) PushGroupByID(groupId string) (*string, error) {

	body, err := dns.Client.doRequest("POST", "/dns/groups/"+groupId+"/push", nil)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return nil, err
	}

	var pid string
	if val, ok := jsonMap["pid"]; ok {
		pid = fmt.Sprintf("%v", val)
	}

	return &pid, nil
}

func (dns *DNSMethods) PushServerByID(serverId string) (*string, error) {

	body, err := dns.Client.doRequest("POST", "/dns/servers/"+serverId+"/push", nil)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return nil, err
	}

	var pid string
	if val, ok := jsonMap["pid"]; ok {
		pid = fmt.Sprintf("%v", val)
	}

	return &pid, nil
}