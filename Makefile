INSTALL_PATH = $(DESTDIR)/usr/bin
SRC_DIR = .
DST_DIR = dst

.PHONY : bin

bin:
	mkdir -p $(DST_DIR)
	go build -o $(DST_DIR)/cftool $(SRC_DIR)/*.go

install:
	mkdir -p $(INSTALL_PATH)
	cp -r $(DST_DIR)/* $(INSTALL_PATH)
	chmod +x $(INSTALL_PATH)/cftool

clean:
	rm -rf $(DST_DIR)
