package main

import (
	"passwd"
	"fmt"
	"time"
	"bytes"
	"encoding/json"
	"strings"
    "io/ioutil"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

type Manufacturer struct {
	ID   string `json:"ID"`
	Desc string `json:"Description"`
	Devs []Device `json:"Devices,omitempty" bson:",omitempty"`
}

type Device struct {
	ID   string `json:"ID"`
	Desc string `json:"Desc"`
	Subs []Sub `json:"Sub,omitempty" bson:",omitempty"`
}

type Sub struct {
	ID   string `json:"ID"`
	Desc string `json:"Desc"`
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
func readFileAndUpdate(){
	fmt.Println("Update Initiated")
	url := "https://pci-ids.ucw.cz/v2.2/pci.ids"

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

	    db, err := sql.Open("mysql", "pci:"+ passwd.GetSQLPassword() +"@tcp(127.0.0.1:3306)/pci")
		if err != nil {
			println(err)
		}
		defer db.Close()

	    strs := strings.Split(str, "\n")

		topID := ""
		grpID := ""
		manufIns, err := db.Prepare("REPLACE INTO manuf_ids VALUES( ?, ? )")
		if err != nil {

			panic(err.Error())
		}
		defer manufIns.Close()

		devIns, err := db.Prepare("REPLACE INTO dev_ids VALUES( ?, ?, ? )")
		if err != nil {

			panic(err.Error())
		}
		defer devIns.Close()

		subDevIns, err := db.Prepare("REPLACE INTO sub_dev_ids VALUES( ?, ?, ?, ? )")
		if err != nil {

			panic(err.Error())
		}
		defer subDevIns.Close()

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
	    		//This is a top level ID, usually the Manufacturer
	        	strParts := strings.Split(newstring, ",")
	        	topID = strParts[0]

	        	fmt.Println("REPLACE INTO manuf_ids VALUES( '" + strParts[0] + "', '" + strParts[1] + "' )")
				_, err = manufIns.Exec(strParts[0], strParts[1])
				if err != nil {
					fmt.Println(err)
				}

	    	} else if strings.Contains(newstring, "\t") && !strings.Contains(newstring, "\t\t") && !strings.HasPrefix(newstring, "#") && len(newstring) > 0{
	    		//This is a second level ID, usually a device or group of devices.
	    		newstring = strings.Replace(newstring, "\t", "", -1)
	    		strParts := strings.Split(newstring, ",")
	    		grpID = strParts[0]

	    		fmt.Println("REPLACE INTO dev_ids VALUES( '" + strParts[0] + "', '" + topID + "', '" + strParts[1] + "' )")
	    		_, err = devIns.Exec(strParts[0], topID, strParts[1])
				if err != nil {
					fmt.Println(err)
				}

	    	} else if strings.Contains(newstring, "\t\t") && !strings.HasPrefix(newstring, "#") && len(newstring) > 0{
	    		//This is a third level ID, usually a more specific device.
	    		newstring = strings.Replace(newstring, "\t", "", -1)
	    		strParts := strings.Split(newstring, ",")

	    		fmt.Println("REPLACE INTO sub_dev_ids VALUES( '" + strParts[0] + "', '" + topID + "', '" + grpID + "', '" + strParts[1] + "' )")
	    		_, err = subDevIns.Exec(strParts[0], topID, grpID, strParts[1])
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
func searchForDevices(manuf, dev, sub string) ([]Manufacturer){

	results := []Manufacturer{}

	db, err := sql.Open("mysql", "pci:"+ passwd.GetSQLPassword() +"@tcp(127.0.0.1:3306)/pci")
	if err != nil {
		println(err)
	}
	defer db.Close()

	if !(len(manuf) > 0) && !(len(dev) > 0) && !(len(sub) > 0){
		//No Search Criteria
		results = []Manufacturer{}
	} else if len(manuf) > 0 && !(len(dev) > 0) && !(len(sub) > 0){
		//Vendor search
		reply, err := db.Query("SELECT * FROM manuf_ids WHERE id='" + manuf + "';")
		if err != nil {
			panic(err.Error())
		}

		for reply.Next() {
			var manufacturer Manufacturer
			err = reply.Scan(&manufacturer.ID, &manufacturer.Desc)
			if err != nil {
				panic(err.Error())
			}

			replytwo, err := db.Query("SELECT id,descrip FROM dev_ids WHERE manuf='" + manuf + "';")
			if err != nil {
				panic(err.Error())
			}

			for replytwo.Next() {
				var devices Device
				err = replytwo.Scan(&devices.ID, &devices.Desc)
				if err != nil {
					panic(err.Error())
				}

				//Get subdevices
				replythree, err := db.Query("SELECT id,descrip FROM sub_dev_ids WHERE parent='" + devices.ID + "';")
				if err != nil {
					panic(err.Error())
				}

				for replythree.Next() {
					var sub_devices Sub
					err = replythree.Scan(&sub_devices.ID, &sub_devices.Desc)
					if err != nil {
						panic(err.Error())
					}
			        devices.Subs = append(devices.Subs, sub_devices)
				}
				replythree.Close()
		        manufacturer.Devs = append(manufacturer.Devs, devices)
			}
			replytwo.Close()
	        results = append(results, manufacturer)
		}
		reply.Close()

	} else if !(len(manuf) > 0) && len(dev) > 0 && !(len(sub) > 0) {
		var manufacturer Manufacturer
		manuf_id := ""
		manuf_set := false

		reply, err := db.Query("SELECT id,descrip,manuf FROM dev_ids WHERE id='" + dev + "';")
		if err != nil {
			panic(err.Error())
		}

		for reply.Next() {

			var devices Device
			err = reply.Scan(&devices.ID, &devices.Desc, &manuf_id)
			if err != nil {
				panic(err.Error())
			}
			if manuf_set == false {
				reply, err := db.Query("SELECT * FROM manuf_ids WHERE id='" + manuf_id + "';")
				if err != nil {
					panic(err.Error())
				}

				for reply.Next() {
					err = reply.Scan(&manufacturer.ID, &manufacturer.Desc)
					if err != nil {
						panic(err.Error())
					}
					manuf_set = true
				}
				reply.Close()
			}

			//Get subdevices
			replytwo, errtwo := db.Query("SELECT id,descrip FROM sub_dev_ids WHERE parent='" + devices.ID + "';")
			if errtwo != nil {
				panic(errtwo.Error())
			}
			for replytwo.Next() {
				var sub_devices Sub
				err = replytwo.Scan(&sub_devices.ID, &sub_devices.Desc)
				if err != nil {
					panic(err.Error())
				}
		        devices.Subs = append(devices.Subs, sub_devices)
			}
			replytwo.Close()
	        manufacturer.Devs = append(manufacturer.Devs, devices)
		}
		reply.Close()
		results = append(results, manufacturer)

	} else if !(len(manuf) > 0) && !(len(dev) > 0) && len(sub) > 0 {
		//Sub only search
		var manufacturer Manufacturer
		var devices Device
		parent := ""
		manuf := ""
		manuf_set := false
		parent_set := false

		reply, err := db.Query("SELECT id,descrip,parent FROM sub_dev_ids WHERE id='" + sub + "';")
		if err != nil {
			panic(err.Error())
		}

		for reply.Next() {
			var sub_devices Sub
			err = reply.Scan(&sub_devices.ID, &sub_devices.Desc, &parent)
			if err != nil {
				panic(err.Error())
			}

			if parent_set == false{
				//Get Parent
				replytwo, errtwo := db.Query("SELECT id,descrip,manuf FROM dev_ids WHERE id='" + parent + "';")
				if errtwo != nil {
					panic(err.Error())
				}
				for replytwo.Next() {
					errtwo = replytwo.Scan(&devices.ID, &devices.Desc, &manuf)
					if errtwo != nil {
						panic(errtwo.Error())
					}

					if manuf_set == false{
						//Get vendor
						replythree, errthree := db.Query("SELECT id,descrip FROM manuf_ids WHERE id='" + manuf + "';")
						if errthree != nil {
							panic(err.Error())
						}
						for replythree.Next() {
							errthree = replythree.Scan(&manufacturer.ID, &manufacturer.Desc)
							if errthree != nil {
								panic(errthree.Error())
							}
						}
						replythree.Close()
					}
				}
				replytwo.Close()
			}
	        devices.Subs = append(devices.Subs, sub_devices)
	        manufacturer.Devs = append(manufacturer.Devs, devices)
		}
		reply.Close()
		results = append(results, manufacturer)
	} else if len(manuf) > 0 && len(dev) > 0 && !(len(sub) > 0) {
		//Device search with manufacturer
		var manufacturer Manufacturer
		manuf_set := false

		reply, err := db.Query("SELECT id,descrip FROM dev_ids WHERE id='" + dev + "' AND manuf='" + manuf + "';")
		if err != nil {
			panic(err.Error())
		}

		for reply.Next() {

			var devices Device
			err = reply.Scan(&devices.ID, &devices.Desc)
			if err != nil {
				panic(err.Error())
			}

			if manuf_set == false {
				reply, err := db.Query("SELECT * FROM manuf_ids WHERE id='" + manuf + "';")
				if err != nil {
					panic(err.Error())
				}

				for reply.Next() {
					err = reply.Scan(&manufacturer.ID, &manufacturer.Desc)
					if err != nil {
						panic(err.Error())
					}
					manuf_set = true
				}
				reply.Close()
			}

			//Get subdevices
			replytwo, errtwo := db.Query("SELECT id,descrip FROM sub_dev_ids WHERE parent='" + devices.ID + "';")
			if errtwo != nil {
				panic(errtwo.Error())
			}

			for replytwo.Next() {
				var sub_devices Sub
				err = replytwo.Scan(&sub_devices.ID, &sub_devices.Desc)
				if err != nil {
					panic(err.Error())
				}
		        devices.Subs = append(devices.Subs, sub_devices)
			}
			replytwo.Close()
	        manufacturer.Devs = append(manufacturer.Devs, devices)
		}
		reply.Close()
		results = append(results, manufacturer)

	} else if !(len(manuf) > 0) && len(dev) > 0 && len(sub) > 0 {
		//Sub device search with parent
		var manufacturer Manufacturer
		var devices Device
		parent := dev
		manuf := ""
		manuf_set := false
		parent_set := false

		reply, err := db.Query("SELECT id,descrip FROM sub_dev_ids WHERE id='" + sub + "' AND parent='" + dev + "';")
		if err != nil {
			panic(err.Error())
		}

		for reply.Next() {
			var sub_devices Sub
			err = reply.Scan(&sub_devices.ID, &sub_devices.Desc)
			if err != nil {
				panic(err.Error())
			}

			if parent_set == false{
				//Get Parent
				replytwo, errtwo := db.Query("SELECT id,descrip,manuf FROM dev_ids WHERE id='" + parent + "';")
				if errtwo != nil {
					panic(err.Error())
				}
				for replytwo.Next() {
					errtwo = replytwo.Scan(&devices.ID, &devices.Desc, &manuf)
					if errtwo != nil {
						panic(errtwo.Error())
					}

					if manuf_set == false{
						//Get vendor
						replythree, errthree := db.Query("SELECT id,descrip FROM manuf_ids WHERE id='" + manuf + "';")
						if errthree != nil {
							panic(err.Error())
						}
						for replythree.Next() {
							errthree = replythree.Scan(&manufacturer.ID, &manufacturer.Desc)
							if errthree != nil {
								panic(errthree.Error())
							}
						}
						replythree.Close()
					}
				}
				replytwo.Close()
			}
	        devices.Subs = append(devices.Subs, sub_devices)
	        manufacturer.Devs = append(manufacturer.Devs, devices)
		}
		reply.Close()
		results = append(results, manufacturer)

	} else if len(manuf) > 0 && !(len(dev) > 0) && len(sub) > 0{
		//Sub search with manuf
		var manufacturer Manufacturer
		var devices Device
		parent := ""
		manuf_set := false
		parent_set := false

		reply, err := db.Query("SELECT id,descrip,parent FROM sub_dev_ids WHERE id='" + sub + "';")
		if err != nil {
			panic(err.Error())
		}

		for reply.Next() {
			var sub_devices Sub
			err = reply.Scan(&sub_devices.ID, &sub_devices.Desc, &parent)
			if err != nil {
				panic(err.Error())
			}

			if parent_set == false{
				//Get Parent
				replytwo, errtwo := db.Query("SELECT id,descrip FROM dev_ids WHERE id='" + parent + "' AND manuf='" + manuf + "';")
				if errtwo != nil {
					panic(err.Error())
				}
				for replytwo.Next() {
					errtwo = replytwo.Scan(&devices.ID, &devices.Desc)
					if errtwo != nil {
						panic(errtwo.Error())
					}

					if manuf_set == false{
						//Get vendor
						replythree, errthree := db.Query("SELECT id,descrip FROM manuf_ids WHERE id='" + manuf + "';")
						if errthree != nil {
							panic(err.Error())
						}
						for replythree.Next() {
							errthree = replythree.Scan(&manufacturer.ID, &manufacturer.Desc)
							if errthree != nil {
								panic(errthree.Error())
							}
						}
						replythree.Close()
					}
				}
				replytwo.Close()
			}
	        devices.Subs = append(devices.Subs, sub_devices)
	        manufacturer.Devs = append(manufacturer.Devs, devices)
		}
		reply.Close()
		results = append(results, manufacturer)
	} else if len(manuf) > 0 && len(dev) > 0 && len(sub) > 0{
		//Full search
		var manufacturer Manufacturer
		var devices Device
		manuf_set := false
		parent_set := false

		reply, err := db.Query("SELECT id,descrip FROM sub_dev_ids WHERE id='" + sub + "' AND parent='" + dev + "';")
		if err != nil {
			panic(err.Error())
		}

		for reply.Next() {
			var sub_devices Sub
			err = reply.Scan(&sub_devices.ID, &sub_devices.Desc)
			if err != nil {
				panic(err.Error())
			}

			if parent_set == false{
				//Get Parent
				replytwo, errtwo := db.Query("SELECT id,descrip FROM dev_ids WHERE id='" + dev + "' AND manuf='" + manuf + "';")
				if errtwo != nil {
					panic(err.Error())
				}
				for replytwo.Next() {
					errtwo = replytwo.Scan(&devices.ID, &devices.Desc)
					if errtwo != nil {
						panic(errtwo.Error())
					}

					if manuf_set == false{
						//Get vendor
						replythree, errthree := db.Query("SELECT id,descrip FROM manuf_ids WHERE id='" + manuf + "';")
						if errthree != nil {
							panic(err.Error())
						}
						for replythree.Next() {
							errthree = replythree.Scan(&manufacturer.ID, &manufacturer.Desc)
							if errthree != nil {
								panic(errthree.Error())
							}
						}
						replythree.Close()
					}
				}
				replytwo.Close()
			}
	        devices.Subs = append(devices.Subs, sub_devices)
	        manufacturer.Devs = append(manufacturer.Devs, devices)
		}
		reply.Close()
		results = append(results, manufacturer)
	}

	return results
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//If input matches the format of the IDs allow it. Otherwise log an inject attempt. 
//(This is to prevent SQL injection and also for science)
//////////////////////////////////////////////////////////////////////////////////////////////////
func cleanText(in string) string{

	match := false
	clean := ""

	if strings.HasPrefix(in, "0x") || strings.HasPrefix(in, "VEN_") || strings.HasPrefix(in, "DEV_"){
		in = in[len(in)-4:]
	}

	if len(in) == 4{
		match, _ = regexp.MatchString("[A-Fa-f0-9]{4}", in)
	} else if len(in) == 9{
		match, _ = regexp.MatchString("[A-Fa-f0-9]{4}\\W[A-Fa-f0-9]{4}", in)
	} else if len(in) == 0{
		match = true
	}

	if match {
		clean = in
	}

	return clean
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
	results := []Manufacturer{}

	if action == "update" && password == passwd.GetAPIPassword(){
		readFileAndUpdate()
	} else if action == "search" {
		manuf := r.URL.Query().Get("vendor")
		dev := r.URL.Query().Get("device")
		sub := r.URL.Query().Get("subdevice")

		results = searchForDevices(cleanText(manuf), cleanText(dev), cleanText(sub))
		json.NewEncoder(w).Encode(results)
	} else {
		help := Help{"to PCILookup!", "search", "vendor, device, subdevice", "rafikithegrouch@gmail.com"}
		json.NewEncoder(w).Encode(help)
	}
}

func loader(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/root/pci/bin/loaderio-48f6d9c8cf565ce68f5e0d90b1704656.txt")
	if err != nil {
		fmt.Println(err)
	}

	http.ServeContent(w, r,"loaderio-48f6d9c8cf565ce68f5e0d90b1704656", time.Now(), bytes.NewReader(data))
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Listen on port 8000 for calls and run function.
//////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/loaderio-48f6d9c8cf565ce68f5e0d90b1704656/", loader)
	http.ListenAndServe(":8000", nil)
}
