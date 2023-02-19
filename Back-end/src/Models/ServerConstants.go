package models

import (
    "net"
  	"os"
)

type CodeTable struct {
	OK						int
	Created					int
	Accepted				int
	NoContent				int
	MovedPermanently		int
	MovedTemporarily		int
	NotModified				int
	BadRequest				int
	Unauthorized			int
	Forbidden				int
	NotFound				int
	InternalServerError		int
	NotImplemented	    	int
	BadGateway	        	int
	ServiceUnavailable		int
}

var ServerCodes CodeTable = CodeTable{
	OK: 200,
	Created: 201,
	Accepted: 202,
	NoContent: 204,
	MovedPermanently: 301,
	MovedTemporarily: 302,
	NotModified: 304,
	BadRequest: 400,
	Unauthorized: 401,
	Forbidden: 403,
	NotFound: 404,
	InternalServerError: 500,
	NotImplemented: 501,
	BadGateway: 502,
	ServiceUnavailable: 503,
}

func GetCurrentMacAddress() string {
    
    host, err := os.Hostname()
    
    if err != nil {
		return ""
    } 

    addr, err := net.LookupIP(host)
    
    if err != nil {
		return ""
    } 
    
    return addr[2].String()
}


func GetCurrentMachineIp() string {
    
    host, err := os.Hostname()
    
    if err != nil {
		return ""
    } 

    addr, err := net.LookupIP(host)
    
    if err != nil {
	return ""
    } 
    
    return addr[1].String()
}