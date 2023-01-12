# tellme
send mail in one-line script

## build
`./build.sh` or `./build-win.bat` (on windows)

## generate your accesskey
```
./tellme seal --host="127.0.0.1" --port="25" --auth="" --user="" --password="" --name="smtp-mypc-localhost-25"
```

 leave `--auth="" --user="" --password=""` : sendmail anonymously 
 use `--auth="plain" --user="your-username" --password="your-password"` : sendmail with plainAuth
 use `--auth="login" --user="your-mail@hotmail.com" --password="your-password"` : sendmail with startTLS, for hotmail.com or outlook.com, etc.

 then you will get your --accesskey="...", keep it.
 i.e.(for localhost 25 smtp server):
 --accesskey="Zt5BmhsR9C8C029xblTUkhR0JJazlduCVdKZIIX2aPqN7HQzd9/Bq4HPp6qU0DkBQ7H243gI2akWljyv1Jpsh6sAg/4vTE+IpMbSTV6LkC+IVR2zS1B6z+XOWqkBrSNOj/4hm0DedQPxGcZ434LVbQ=="
	

## send mail
```
 sendmail  --accesskey="Zt5BmhsR9C8C029xblTUkhR0JJazlduCVdKZIIX2aPqN7HQzd9/Bq4HPp6qU0DkBQ7H243gI2akWljyv1Jpsh6sAg/4vTE+IpMbSTV6LkC+IVR2zS1B6z+XOWqkBrSNOj/4hm0DedQPxGcZ434LVbQ==" --from="your-mail@address.com"   --to="receiver@address.com" --subject="test" --file="/home/harryzhu/email.html"
```
