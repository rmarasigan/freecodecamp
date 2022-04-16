# email-checker-tool

```bash
dev@dev:~/go/src/github.com/development/email-checker-tool$ go mod init
dev@dev:~/go/src/github.com/development/email-checker-tool$ go mod tidy
dev@dev:~/go/src/github.com/development/email-checker-tool$ go build
dev@dev:~/go/src/github.com/development/email-checker-tool$ go run main.go
domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
gmail.com
hasMX true hasSPF true hasDMARC true
spfRecord v=spf1 redirect=_spf.google.com
dmarcRecord v=DMARC1; p=none; sp=quarantine; rua=mailto:mailauth-reports@google.com
mailchimp.com
hasMX true hasSPF true hasDMARC true
spfRecord v=spf1 include:servers.mcsv.net include:mail.zendesk.com include:_spf.google.com include:mailsenders.netsuite.com ip4:199.33.145.1 ip4:199.33.145.32 ip4:148.105.0.14 ip4:35.176.132.251 ip4:52.60.115.116 ~all
dmarcRecord v=DMARC1; p=reject; rua=mailto:19ezfriw@ag.dmarcian.com; ruf=mailto:19ezfriw@fr.dmarcian.com
yahoo.com
hasMX true hasSPF true hasDMARC true
spfRecord v=spf1 redirect=_spf.mail.yahoo.com
dmarcRecord v=DMARC1; p=reject; pct=100; rua=mailto:d@rua.agari.com; ruf=mailto:d@ruf.agari.com;
```