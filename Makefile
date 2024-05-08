idl:
	./kitex_gen.sh
env:
	docker-compose up -d
	docker start kibana
	docker start elasticsearch
build:
	cd cmd/api && go build -o api
	cd cmd/user && go build -o user
	cd cmd/video && go build -o video
	cd cmd/publish && go build -o publish
	cd cmd/favorite && go build -o favorite
	cd cmd/relation && go build -o relation
api:
	cd cmd/api && ./api
user:	
	cd cmd/user &&./user
video:
	cd cmd/video &&./video
publish:
	cd cmd/publish && ./publish
favorite:
	cd cmd/favorite && ./favorite
relation:
	cd cmd/relation && ./relation		

go:env build

clean:
	rm -f test