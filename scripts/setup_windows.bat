REM Install beego (Requires git)
go get github.com/astaxie/beego

REM In case https validation not working
REM git config --global http.sslVerify false

REM Installing bee
go get github.com/beego/bee

REM Govender simulates package.json and helps in installing the required packages
go get -u github.com/kardianos/govendor

REM Maps and context for getting locations
go get googlemaps.github.io\maps
go get golang.org/x/net/context
