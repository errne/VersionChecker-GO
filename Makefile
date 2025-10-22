build:
	go build -o sysinfo_cli
	chmod +x sysinfo_cli

run:
	./sysinfo_cli

clean:
	rm -f sysinfo_cli
