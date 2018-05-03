package main

import (
	"passwd"
	"fmt"
	"encoding/json"
	"strings"
    "io/ioutil"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Vendor struct {
	ID   string `json:"ID"`
	Desc string `json:"Description"`
	Devs []Device `json:"Devices,omitempty" bson:",omitempty"`
}

type Device struct {
	ID   string `json:"ID"`
	Desc string `json:"Desc"`
	VendorID string `json:"Vendor ID"`
	VendorDesc string `json:"Vendor Description"`
}

type Help struct{
	Welcome string `json:"Welcome"`
	Actions string `json:"Actions"`
	Args    string `json:"Arguments"`
	Contact string `json:"Contact"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Read PCI.IDS File, parse contents, put into respective tables in database.
//////////////////////////////////////////////////////////////////////////////////////////////////
func readFileAndUpdate(url string){
	fmt.Println("Update Initiated")

    var client http.Client
	resp, err := client.Get(url)
	if err != nil {
	    panic(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 { // OK
	    bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	    str := string(bodyBytes)
		if err2 != nil {
		    panic(err2.Error())
		}

	    db, err := sql.Open("mysql", "pci:"+ passwd.GetSQLPassword() +"@tcp(127.0.0.1:3307)/pci-dev")
		if err != nil {
			println(err)
		}
		defer db.Close()

		devIns, err := db.Prepare("REPLACE INTO dev_ids VALUES( ?, ?, ?, ? )")
		if err != nil {

			panic(err.Error())
		}
		defer devIns.Close()

	    strs := strings.Split(str, "\n")

    	vendorID := ""
    	vendorDescrip := ""

	    for i := range strs {

	    	tx, err := db.Begin()
			if err != nil {
				fmt.Println(err)
			}

	    	newstring := strings.Replace(strs[i], "  ", ",", -1)

	    	if strings.Contains(strs[i], "'"){
	    		newstring = strings.Replace(newstring, "'", "\\'", -1)
	    	}

	    	if !strings.Contains(newstring, "\t") && !strings.HasPrefix(newstring, "#") && len(newstring) > 0{
	    		//This is a top level ID, usually the Vendor
	        	strParts := strings.Split(newstring, ",")
	        	vendorID = strParts[0]
	        	vendorDescrip = strParts[1]

	    	} else if strings.Contains(newstring, "\t") && !strings.Contains(newstring, "\t\t") && !strings.HasPrefix(newstring, "#") && len(newstring) > 0{
	    		//This is a second level ID, usually a device or group of devices.
	    		newstring = strings.Replace(newstring, "\t", "", -1)
	    		strParts := strings.Split(newstring, ",")
	    		devID := strParts[0]
	    		devDescrip := strParts[1]

	    		fmt.Println(devID, devDescrip, vendorID, vendorDescrip)
	    		_, err = devIns.Exec(devID, devDescrip, vendorID, vendorDescrip)
				if err != nil {
					fmt.Println(err)
				}

	    	} else if strings.Contains(newstring, "\t\t") && !strings.HasPrefix(newstring, "#") && len(newstring) > 0{
	    		//This is a third level ID, usually a more specific device.
	    		newstring = strings.Replace(newstring, "\t", "", -1)
	    		strParts := strings.Split(newstring, ",")
	    		devID := strParts[0]
	    		devDescrip := strParts[1]

	    		fmt.Println(devID, devDescrip, vendorID, vendorDescrip)
	    		_, err = devIns.Exec(devID, devDescrip, vendorID, vendorDescrip)
				if err != nil {
					fmt.Println(err)
				}
	    	}
	    	tx.Commit()
	    }
	    fmt.Println("Updated Database")
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Take user search parameters and query database, return results.
//////////////////////////////////////////////////////////////////////////////////////////////////
func searchForDevices(vendor, dev string) ([]Device){
	if vendor == "" {vendor = "%"}
	if dev == "" {dev = "%"}

	devices := []Device{}

	db, err := sql.Open("mysql", "pci:"+ passwd.GetSQLPassword() +"@tcp(127.0.0.1:3307)/pci-dev")
	if err != nil {
		println(err)
	}
	defer db.Close()

	deviceReply, deviceErr := db.Query("SELECT deviceID,deviceDescrip,vendorID,vendorDescrip FROM dev_ids WHERE (vendorID LIKE ? OR vendorDescrip LIKE ?) AND (deviceID LIKE ? OR deviceDescrip LIKE ?)", "%" + vendor + "%", "%" + vendor + "%", "%" + dev + "%", "%" + dev + "%")
	if deviceErr != nil {
		panic(deviceErr.Error())
	}

	for deviceReply.Next() {
		var device Device
		deviceErr = deviceReply.Scan(&device.ID, &device.Desc, &device.VendorID, &device.VendorDesc)
		if deviceErr != nil {
			panic(deviceErr.Error())
		}
		devices = append(devices, device)
	}
	return devices
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Get query parameters and run the necessary function.
//////////////////////////////////////////////////////////////////////////////////////////////////
func hello(w http.ResponseWriter, r *http.Request) {

    // allow cross domain AJAX requests
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
    //Tell the browser we are sending JSON
    w.Header().Set("Content-Type", "application/json")

	action := r.URL.Query().Get("action")
	password := r.URL.Query().Get("password")
	results := []Device{}

	if action == "update" && password == passwd.GetAPIPassword(){
		readFileAndUpdate("https://pci-ids.ucw.cz/v2.2/pci.ids")
		readFileAndUpdate("http://www.linux-usb.org/usb.ids")
	} else if action == "search" {
		vendor := r.URL.Query().Get("vendor")
		dev := r.URL.Query().Get("device")

		results = searchForDevices(vendor, dev)
		json.NewEncoder(w).Encode(results)
	} else {
		help := Help{"to PCILookup!", "search", "vendor, device", "rafikithegrouch@gmail.com"}
		json.NewEncoder(w).Encode(help)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Listen on port 8000 for calls and run function.
//////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}