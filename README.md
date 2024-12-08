# groxy
Building a simple proxy server in go.

## Security
A self-signed certificate is generated for the server, just for testing purposes. It is not recommended to use this in production.
```
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 3650 -nodes -subj "/C=XX/ST=StateName/L=CityName/O=CompanyName/OU=CompanySectionName/CN=CommonNameOrHostname"

```
