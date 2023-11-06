.PHONY: build restart install_supervisor install


build:
	@go build -o bin/tools

restart:
	@supervisorctl restart tools

install_supervisor:
	@apt install supervisor

install:
	@mkdir -p bin/
	@cp -r web bin/
	@cp -r conf/config.yaml bin/
	@cp conf/supervisor.conf /etc/supervisor/conf.d
