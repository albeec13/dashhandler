# dashhandler
Dashhandler is a webservice, written in [Go](https://golang.org/), and utilizing [gin](https://github.com/gin-gonic/gin) for route handling, intended to be used as a target endpoint for a router running [dnsmasq](http://www.thekelleys.org.uk/dnsmasq/docs/dnsmasq-man.html).  Dnsmasq can be configured, using the `--dhcp-script` switch, to run a script on DHCP events passing along information about the event type, assigned IP address, MAC address, and hostname (if available) of any devices that request DHCP addresses.  Dashhandler can then be accessed via a GET request, passing along that information, where it can be processed for arbitrary actions.  In this case, dashhandler is specifically intended to handle Amazon Dash button presses.
