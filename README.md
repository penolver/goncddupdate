# goncddupdate #

super simple go(lang) code to update namecheap dynamic dns

* no external dependencies
* use crontab to run it at whatever interval you want
* edit the config.json file with your details (should be obvious)
* uses this method to update namecheap: https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-do-i-use-a-browser-to-dynamically-update-the-hosts-ip

## crontab example - every 5 mins run..
crontab -e
*/5       *       *       *       *       cd /location; /location/goncddupdate >/dev/null 2>&1
