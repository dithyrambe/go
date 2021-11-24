package util

import (
	"embed"
	"errors"
	"log"
	"net"
	"strings"
)

var (
	ErrEmailNotValid       = errors.New("email not valid")
	ErrEmailDomainNotValid = errors.New("email domain not valid")
	ErrEmailBlackListed    = errors.New("email black listed")
)

var (
	blackListEmailList []string
)

//go:embed data/*
var blacklistFiles embed.FS

func init() {
	data, err := blacklistFiles.ReadFile("data/email-domain-blacklist.txt")
	if err != nil {
		panic(err)
	}
	blackListEmailList = append(blackListEmailList, strings.Split(string(data), "\n")...)
}

func ValidEmail(str string) error {

	if strings.Count(str, "@") != 1 {
		return ErrEmailNotValid
	}

	strList := strings.Split(str, "@")
	domainName := strList[1]
	if strings.Count(domainName, ".") < 1 {
		return ErrEmailNotValid
	}

	for k := range blackListEmailList {
		if domainName == blackListEmailList[k] {
			return ErrEmailBlackListed
		}
	}

	mxs, err := net.LookupMX(domainName)

	if err != nil {
		log.Println(err)
		return ErrEmailDomainNotValid
	}

	if len(mxs) == 0 {
		return ErrEmailDomainNotValid
	}

	return nil
}
