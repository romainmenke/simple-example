package public

//go:generate simple-bundle -source ./assets/css -out ./assets/css/bundle .go styleguide DS_Store
//go:generate simple-mini -source ./assets/css/bundle -out ./assets/css/min .go styleguide DS_Store
//go:generate simple-gzip -source ./assets/css/min -out ./assets/css/gzip -level 9 .go styleguide DS_Store

//go:generate simple-bundle -source ./assets/js -out ./assets/js/bundle .go DS_Store
//go:generate simple-mini -source ./assets/js/bundle -out ./assets/js/min .go DS_Store
//go:generate simple-gzip -source ./assets/js/min -out ./assets/js/gzip -level 9 .go DS_Store

//go:generate simple-template -source ./assets/templates -out ./assets/html .go DS_Store
//go:generate simple-mini -source ./assets/html -out ./assets/html/min .go DS_Store
//go:generate simple-gzip -source ./assets/html/min -out ./assets/html/gzip -level 9 .go DS_Store

//go:generate simple-gzip -source ./assets/img -out ./assets/img/gzip -level 9 .go DS_Store

//go:generate simple-gzip -source ./assets/fonts -out ./assets/fonts/gzip -level 9 .go DS_Store

//go:generate go-bindata -pkg public -o ./bindata.go -ignore=.DS_Store -ignore=zzz_generate.go -nocompress ./assets/...
