# NGCLI 

NGCLI providers the ability to use and manage Neural Galaxy data service with the command line style.

# Installation
Go to the `Release` page, find the one version which you like. Current every version have two kinds of binary:
- for mac amd64
- for linux amd64

Download the binary to your **PATH** directory rename or alias it.
```
# assume you download the darwin-amd64-ngcli at ~/Download
...

# move it to /usr/local/bin
mv darwin-amd64-ngcli /usr/local/bin/ngcli

# or you can make alias in your shell
alias ngcli="~/Download/darwin-amd64-ngcli"
```

# Usage
## ngcli help
use help to learn the CLI
```
ngcli help
```

## ngcli generate config
```
ngcli --endpoint test2-api.neuralgalaxy.cn --username $username --password $password --applicationType 1
```
## ngcli login
if `ngcli` has no any sub-command, it will auto login once and save the token in the config file.
```
ngcli
```
Of cause, you also can run login sub-command:
```
ngcli auth login
```

## ngcli project 

### ngcli project list
```
ngcli project list
```

## ngcli subeject

### ngcli subject list
```
ngcli subject list
```
or select subject with `jq`
```
ngcli subject list | jq -r '.[]|select(.subjectCustId=="testsub4")'
```

### ngcli subject create
```
ngcli subject create --projectId 207 --subjectCustId ngcliSubj
```

### ngcli subject delete
```
ngcli subject delete --subjectId 3959 --projectId 207
```

## ngcli upload
```
ngcli upload -f ~/Downloads/debug.zip --subjectId 1268 --projectId 207
```