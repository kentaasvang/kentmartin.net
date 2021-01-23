
run:
	open http://localhost:8000;
	php -S localhost:8000;

build-live-site:
	sudo rm -rf /var/www/html/* && sudo cp ./* /var/www/html/
