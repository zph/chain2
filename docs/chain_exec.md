## chain exec

Execute a command in the context of ENV vars fetched from keychain

### Synopsis


	Execute a command in the context of ENV vars fetched from keychain

	eg:
	# Fetch from aws-creds keychain use aws commandline tool in that context
	chain exec aws-creds -- aws s3 ls...


```
chain exec [keychain] -- [execCommand] [execCommandArgs...] [flags]
```

### Options

```
  -h, --help   help for exec
```

### SEE ALSO

* [chain](chain.md)	 - A remake of envchain and generic sibling of aws-vault

###### Auto generated by spf13/cobra on 27-Dec-2022