rmdir /Q /S ..\base
call openapi-generator-cli generate -i openapi.yaml -g go-experimental -t template/go-experimental -c config/go-experimental.yaml --type-mappings=ecpay-date-time=ECPayDateTime  --type-mappings=integer=int -o ../base
rem call openapi-generator generate -i openapi.yaml -g go -t template/go-experimental -c config/go-experimental.yaml --type-mappings=ecpay-date-time=ECPayDateTime -o ../base
cd ../base
del go.mod
cd ..
go fmt
cd api