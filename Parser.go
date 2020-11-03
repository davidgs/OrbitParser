package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// ORBIT is your Orbit API_KEY here
const ORBIT = "YOUR_API_KEY_HERE"

// ORG is your Orbit Organization
const ORG = "YOUR_ORBIT_ORG"

// ITEMS is how many records you want to fetch each time
const ITEMS = "100"

// OrbitData is the struct for the JSON. if there are fields missing, you can add them
type OrbitData struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			ActivitiesCount         int           `json:"activities_count"`
			Bio                     string        `json:"bio"`
			Company                 interface{}   `json:"company"`
			Devto                   interface{}   `json:"devto"`
			FirstActivityOccurredAt time.Time     `json:"first_activity_occurred_at"`
			LastActivityOccurredAt  time.Time     `json:"last_activity_occurred_at"`
			Location                string        `json:"location"`
			Linkedin                interface{}   `json:"linkedin"`
			Name                    string        `json:"name"`
			OrbitLevel              int           `json:"orbit_level"`
			NewOrbitLevel           int           `json:"new_orbit_level"`
			Reach                   int           `json:"reach"`
			TagList                 []interface{} `json:"tag_list"`
			Teammate                bool          `json:"teammate"`
			URL                     string        `json:"url"`
			WebURL                  string        `json:"web_url"`
			Twitter                 string        `json:"twitter"`
			Github                  string        `json:"github"`
			Discourse               interface{}   `json:"discourse"`
			Email                   string        `json:"email"`
			Love                    float64       `json:"love"`
		} `json:"attributes"`
	} `json:"data"`
}

var outFile *os.File
var err error

func main() {

	useFilePtr := flag.String("file", "in.json", "a file")
	outFilePtr := flag.String("out", "out.csv", "an output file.")
	flag.Parse()


	if *outFilePtr == "" {
		outFile, err = os.Create("allOrbit.csv")
		if err != nil {
			log.Fatal("Could not open file ", err)
		}
	} else {
		outFile, err = os.Create(*outFilePtr)
		if err != nil {
			log.Fatal("Open File Error", err)
		}
	}
	defer outFile.Close()

	_, err := outFile.WriteString("ActivitiesCount,Bio,Company,DevTo,FirstActivity,LastActivity,Location,LinkedIn,Name,OrbitLevel,TagList,Reach,URL,WebURL,Twitter,GitHub,Discourse,Email,Love\n")

	var DefaultClient = &http.Client{}
	const urlPlace = "https://app.orbit.love/api/v1/" + ORG + "/members?items=" + ITEMS + "?page="

	var x = 1
	var fileUsed = false
	for {
		if fileUsed {
			break
		}
		var oData = OrbitData{}
		if *useFilePtr != "" {
			f, err := os.Open(*useFilePtr)
			defer f.Close()
			if err != nil {
				log.Fatal("Failed to open file ", err)
			}
			data, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatal("Got no Data")
			}
			_ = json.Unmarshal(data, &oData)
			fileUsed = true
			if len(oData.Data) == 0 {
				break
			}
		} else {

			thisURL := urlPlace + strconv.Itoa(x) +  "&" + ORBIT
			x++
			res, err := DefaultClient.Get(thisURL)
			if err != nil {
				log.Fatal("Can't GET ", err)
			}
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal("Got no Data")
			}
			res.Body.Close()
			_ = json.Unmarshal(data, &oData)
			if len(oData.Data) == 0 {
				break
			}
			l := len(oData.Data)
			fmt.Printf("\nRequest Number:\t\t%d\nNumber of records: %d\n", x-1, l)
		}
		for i := 0; i < len(oData.Data); i++ {
			var sb strings.Builder
			d := oData.Data[i]
			if !d.Attributes.Teammate {
				sb.WriteString(strconv.Itoa(d.Attributes.ActivitiesCount) + ",")

				if d.Attributes.Bio != "" {
					sb.WriteString("\"" + d.Attributes.Bio + "\",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Company != nil {
					sb.WriteString("\"" + fmt.Sprintf("%v,", d.Attributes.Company) + "\",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Devto != nil {
					sb.WriteString(fmt.Sprintf("%v,", d.Attributes.Devto) + ",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.FirstActivityOccurredAt.String() != "" {
					sb.WriteString("\"" + d.Attributes.FirstActivityOccurredAt.String() + "\",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.LastActivityOccurredAt.String() != "" {
					sb.WriteString("\"" + d.Attributes.LastActivityOccurredAt.String() + "\",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Location != "" {
					sb.WriteString("\"" + d.Attributes.Location + "\",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Linkedin != nil {
					sb.WriteString(fmt.Sprintf("%v,", d.Attributes.Linkedin))
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Name != "" {
					sb.WriteString("\"" + d.Attributes.Name + "\",")
				} else {
					sb.WriteString(",")
				}
				sb.WriteString(strconv.Itoa(d.Attributes.OrbitLevel) + ",")
				if d.Attributes.TagList != nil {
					sb.WriteString(fmt.Sprintf("%v,", d.Attributes.TagList))
				} else {
					sb.WriteString(",")
				}
				sb.WriteString(strconv.Itoa(d.Attributes.Reach) + ",")
				if d.Attributes.URL != "" {
					sb.WriteString(d.Attributes.URL + ",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.WebURL != "" {
					sb.WriteString(d.Attributes.WebURL + ",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Twitter != "" {
					sb.WriteString(d.Attributes.Twitter + ",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Github != "" {
					sb.WriteString(d.Attributes.Github + ",")
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Discourse != nil {
					sb.WriteString(fmt.Sprintf("%v,", d.Attributes.Discourse))
				} else {
					sb.WriteString(",")
				}
				if d.Attributes.Email != "" {
					sb.WriteString(d.Attributes.Email + ",")
				} else {
					sb.WriteString(",")
				}
				sb.WriteString(fmt.Sprintf("%f", d.Attributes.Love) + "\n")
			}
			_, err = outFile.WriteString(sb.String())
			if err != nil {
				log.Fatal("Failed to write file ", err)
			}
		}

	}
}
