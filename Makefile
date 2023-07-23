all:
	cd src/ && go build && mv massmail ..

clean:
	rm -rf massmail massmail_release

release:
	mkdir -p massmail_release

	mkdir -p massmail_release/win_x32 && \
	 cp README.md massmail_release/win_x32 && \
	 cp message.html massmail_release/win_x32 && \
	 cp mail_list.txt massmail_release/win_x32 && \
	 cd src && \
	 CGO_ENABLED=0 GOARCH=386 GOOS=windows go build && \
	 mv massmail.exe ../massmail_release/win_x32 && \
	 cd ../massmail_release && \
	 zip -r win_x32.zip win_x32 && \
	 rm -rf win_x32

	mkdir -p massmail_release/win_x64 && \
	 cp README.md massmail_release/win_x64 && \
	 cp message.html massmail_release/win_x64 && \
	 cp mail_list.txt massmail_release/win_x64 && \
	 cd src && \
	 CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build && \
	 mv massmail.exe ../massmail_release/win_x64 && \
	 cd ../massmail_release && \
	 zip -r win_x64.zip win_x64 && \
	 rm -rf win_x64
	 
	mkdir -p massmail_release/linux_x64 && \
	 cp README.md massmail_release/linux_x64 && \
	 cp message.html massmail_release/linux_x64 && \
	 cp mail_list.txt massmail_release/linux_x64 && \
	 cd src && \
	 CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build && \
	 mv massmail ../massmail_release/linux_x64 && \
	 cd ../massmail_release && \
	 zip -r linux_x64.zip linux_x64 && \
	 rm -rf linux_x64
	
	mkdir -p massmail_release/darwin_x64 && \
	 cp README.md massmail_release/darwin_x64 && \
	 cp message.html massmail_release/darwin_x64 && \
	 cp mail_list.txt massmail_release/darwin_x64 && \
	 cd src && \
	 CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build && \
	 mv massmail ../massmail_release/darwin_x64 && \
	 cd ../massmail_release && \
	 zip -r darwin_x64.zip darwin_x64 && \
	 rm -rf darwin_x64
	
	mkdir -p massmail_release/source && \
	 cp -r src/* massmail_release/source \
	 && cp Makefile massmail_release/source \
	 && cp README.md massmail_release/source \
	 && cp message.html massmail_release/source \
	 && cp message.txt massmail_release/source \
	 && cp mail_list.txt massmail_release/source



