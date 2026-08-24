PKG_NAME=WatchYourLAN
USR_NAME=aceberg

build:
	npm run build && \
	rm -rf ../backend/internal/web/public/assets/* && \
	cp -r dist/assets ../backend/internal/web/public

replace:
	cd ../backend/internal/web/public/assets/ && \
	cat index.js | sed 's/assets/fs\/public\/assets/g' > tmp && \
	mv tmp index.js && \
	cat index.css | sed 's|url(/assets/|url(/fs/public/assets/|g' > tmp && \
	mv tmp index.css

all: build replace
