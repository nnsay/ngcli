# NGCLI 

NGCLI providers the ability to use and manage Neural Galaxy data service with the command line style.

## ngcli help
use help to learn the CLI
```
ngcli help
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

# ngcli subeject

### ngcli subject list
```
ngcli subject list
```

### ngcli subject create
```
ngcli subject create --projectId 207 --subjectCustId abc
```

### ngcli subject delete
```
ngcli subject delete --subjectId 3959 --projectId 207
```

## ngcli upload
```
ngcli upload -f ~/Downloads/debug.zip -subjectId 1268 --projectId 207
```