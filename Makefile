##
## Simple Makefile for Email-Poller
## 
## Made by axel catusse
##


# vars
GOCMD = go
GOGET = go get
GOBUILD = $(GOCMD) build
GOCLEAN	= $(GOCMD) clean
GOINSTALL = $(GOCMD) install

# dirs
BUILD_DIR = ./build
CONFIG_DIR = ./config
SRC_DIR = ./email-forwarder

# srcs
CONFIG_SRC = $(CONFIG_DIR)/config.go

EMAIL_FORWARDER_SRC = $(SRC_DIR)/email-forwarder.go

# bins
EMAIL_FORWARDER_BIN = email-forwarder

all: get-package
	$(GOBUILD) -o $(BUILD_DIR)/$(EMAIL_FORWARDER_BIN) $(EMAIL_FORWARDER_SRC)

clean:
	rm -rf $(GOBUILD)/*

re: clean all

get-package:
	$(GOGET) github.com/emgenio/email-poller/imap
	$(GOGET) github.com/streadway/amqp
	$(GOGET) github.com/keighl/mandrill
