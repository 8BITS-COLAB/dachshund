default: dev

.PHONY: dev 

dev: 
	air

charge-test:
	siege -c255 -t60S -r10  -H 'dachshund-api-key: e39f4051314426334b8fc098e173e859265dddd5' --content-type "application/json" 'http://app.dachshund.local:4000/api/v1/emails POST {"template": "default","context": {"to": {"name": "Lais","email": "laisdarem@gmail.com"},"body": {"title": "Hello World","description": "My name is Dachshund"}}}'