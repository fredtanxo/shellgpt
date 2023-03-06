# ShellGPT

The ChatGPT in shell.

## Usage
```
Usage of ./shellgpt:
  -api-key string
    	OpenAI API key
```

### Run in Docker(recommend)
```shell
export MY_OPENAI_API_KEY="sk-ThisIsNotTheRealSecretKey"
docker run -it --rm fredtan/shellgpt -api-key="$MY_OPENAI_API_KEY"
```

### Build from source
```shell
git clone https://github.com/fredtanxo/shellgpt.git
cd shellgpt
make build
```

## License
GPLv3
