// update namecheap dyn dns
package main

import (
        "io/ioutil"
        "net/http"
        "log"
        "fmt"
        "github.com/penolver/goncddupdate/app"
)

func main() {

        fmt.Println("\nGo Namecheap DynDNS Updater v0.1\n")

        // load config..
        config := app.LoadConfig("./config.json")

        // grab current IP
        log.Println("Grab current IP from https://api.ipify.org..")
        res, _ := http.Get("https://api.ipify.org")
        ip, _ := ioutil.ReadAll(res.Body)
        log.Println("Current IP: ",string(ip[:]))
        log.Println("Previous/Stored IP: ",config.CurrentIP)

        // if current IP is different from last saved IP, update and save
        if config.CurrentIP != string(ip[:]) {
          log.Println("IP changed, updating IP")
          status := app.NamecheapUpdate(config.Namecheap_Host,config.Namecheap_Domain,config.Namecheap_Pass,string(ip[:]))
          if status == true {
            log.Println("Update successful, saving..")
            config.CurrentIP = string(ip[:])
            app.SaveConfig(config,"./config.json")
          }else {
            log.Println("Update failed, problem updating Namecheap, check password/api key etc.")
          }
        }else{
          log.Println("IP hasnt changed, nothing to do")
        }
        log.Println("Finished.\n")
}
