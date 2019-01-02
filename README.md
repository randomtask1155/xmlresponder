

## Simple app that returns a valid xml response

## run locally

```
export PORT=8080
./xmlresponder
```

## run in cloud foundry

```
cf push
```


## get an html response

```
curl http://127.0.0.1:8080/
```


## get an xml response

```
curl http://127.0.0.1:8080/xml
```