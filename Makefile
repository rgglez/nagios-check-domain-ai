build:
	cd ./src && go build -o ../dist/check_domain_ai *.go

install:
	cp -v ./dist/check_domain_ai /usr/local/nagios/libexec/
