// interact with Namecheap DynDNS API
package app

import (
    "net/http"
    "crypto/tls"
		"io/ioutil"
		"log"
    "strings"
    "time"

)
// Update Namecheap Dyn DNS API
func NamecheapUpdate(host string,domain string,password string,ip string) bool {
  //setup connection options
  tr := &http.Transport{
      TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }
  timeout := time.Duration(60 * time.Second)
  client := &http.Client{Timeout: timeout,Transport: tr}

  // make connection
  req, err := http.NewRequest("GET", `https://dynamicdns.park-your-domain.com/update?host=`+host+`&domain=`+domain+`&password=`+password+`&ip=`+ip, nil)
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal("Error connecting to Namecheap: ", err)
  }
  defer resp.Body.Close()

  // Read response
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
      log.Fatal("Error reading response from Namecheap. ", err)
  }

  // if response is OK, return true, else false
  if resp.StatusCode == 200 && err == nil && strings.Contains(string(data[:]), "<ErrCount>0</ErrCount>"){
      return true
  } else { return false}

  // if all else fails, return false
	return false

} // NamecheapUpdate
