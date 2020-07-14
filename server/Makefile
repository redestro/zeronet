BIN := zeronet

clean:
	docker ps -q --filter "name=${BIN}" | xargs docker container stop
	rm -f ${BIN}
	rm -rf build/cache/*

dev:
	docker-compose -f docker-compose.yml up --build -d && docker logs -f server_zeronet_1