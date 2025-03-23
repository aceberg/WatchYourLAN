PKG_NAME=WatchYourLAN
USR_NAME=aceberg

build:
	npm run build && \
	rm ../internal/web/public/assets/* && \
	cp -r dist/assets ../internal/web/public

replace:
	cd ../internal/web/public/assets/ && \
	cat index-*.js | sed 's/assets/fs\/public\/assets/g;s/http:\/\/0.0.0.0:8840//' > index.js && \
	cat index-*.css | sed 's/assets/fs\/public\/assets/g' > index.css && \
	cat Config-*.js | sed 's/index-.*\.js/index.js/g' > tmp && \
	cp tmp Config-*.js && \
	cat History-*.js | sed 's/index-.*\.js/index.js/g' > tmp && \
	cp tmp History-*.js && \
	cat HostPage-*.js | sed 's/index-.*\.js/index.js/g' > tmp && \
	cp tmp HostPage-*.js && \
	cat MacHistory-*.js | sed 's/index-.*\.js/index.js/g' > tmp && \
	cp tmp MacHistory-*.js && \
	rm tmp index-*

all: build replace