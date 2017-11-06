#Install beego (Requires git)
go get github.com/astaxie/beego

#In case https validation not working
#git config --global http.sslVerify false

#Installing bee
go get github.com/beego/bee

#Govender simulates package.json and helps in installing the required packages
go get -u github.com/kardianos/govendor

#Maps and context for getting locations
go get googlemaps.github.io\maps
go get golang.org/x/net/context
