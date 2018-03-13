# warnings

It's my aws utility command.
The script is offered for limited cases.
I use it in Elasticsearch Services.


# install (Mac)
```
brew install jq
brew install go
go install github.com/smartystreets/go-aws-auth
```

# usage
```
(source assume-role-env -p ${your_aws_profile} && sigv4-request.go -m=GET -u=${domain} -b '${json_body}')
```
