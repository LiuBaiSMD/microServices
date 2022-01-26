cd /root/microServices/service-web2
go run /root/microServices/service-web2/main.go &
haproxy -f /root/haproxy.cfg
