package forwardproxypolicy

import (
	"strings"
)

type ForwardProxy struct{
	bannedDomains []string 
}

func NewForwardProxy (bannedDomains[] string) *ForwardProxy{
	return &ForwardProxy{
		bannedDomains : bannedDomains,
	}
}

func (f *ForwardProxy) IsBanned(domain string) bool {

	for _, d := range f.bannedDomains{

		if strings.HasPrefix(d,"*") && strings.HasSuffix(domain,d[1:]){
			return true //suffix matching only is d startswith "*"
		}

		if domain == d { //exact match 
			return true 
		}
	}
	
	return false 
}