# warnings

It's my aws utility command.
The script is offered for limited cases.
I use it in Elasticsearch Services.


# install (Mac)
```
brew install jq
brew install go
go install github.com/smartystreets/go-aws-auth

mkdir ~/bin
export PATH=$PATH:~/bin

git clone git@github.com:hixi-hyi/aws-bin.git /tmp/aws-bin
cp /tmp/aws-bin/assume-role-env ~/bin
cp /tmp/aws-bin/sigv4-request.go ~/bin
```

# usage
```
(source assume-role-env -p ${your_aws_profile} && sigv4-request.go -m=GET -u=${domain} -b '${json_body}')
```
