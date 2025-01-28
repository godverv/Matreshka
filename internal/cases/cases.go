package cases

import (
	"strings"

	"go.redsock.ru/evon"
)

func SnakeToPascal(v string) string {
	parts := strings.Split(v, evon.ObjectSplitter)
	for i := range parts {
		if _, ok := initialisms[parts[i]]; ok {
			parts[i] = strings.ToUpper(parts[i])
		} else {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}

	return strings.Join(parts, "")
}

var initialisms = map[string]struct{}{
	"acl": {}, "api": {}, "ascii": {}, "cpu": {}, "css": {}, "dns": {}, "eof": {}, "guid": {},
	"html": {}, "http": {}, "https": {}, "id": {}, "ip": {}, "json": {}, "qps": {}, "ram": {},
	"rpc": {}, "sla": {}, "smtp": {}, "sql": {}, "ssh": {}, "tcp": {}, "tls": {}, "ttl": {},
	"udp": {}, "ui": {}, "gid": {}, "uid": {}, "uuid": {}, "uri": {}, "url": {}, "utf8": {},
	"vm": {}, "xml": {}, "xmpp": {}, "xsrf": {}, "xss": {}, "sip": {}, "rtp": {}, "amqp": {},
	"db": {}, "ts": {},
}
