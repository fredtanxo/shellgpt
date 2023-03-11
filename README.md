# ShellGPT

![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/fredtan/shellgpt)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/fredtanxo/shellgpt/ci.yml)
![GitHub](https://img.shields.io/github/license/fredtanxo/shellgpt)

The ChatGPT in shell.

[![asciicast](https://asciinema.org/a/KGgUJ30JY23dnT9z9xcOAsq1p.svg)](https://asciinema.org/a/KGgUJ30JY23dnT9z9xcOAsq1p)

## Usage

```
Usage of ./shellgpt:
  -api-key string
        OpenAI API key
  -api-url string
        OpenAI API base url (default "https://api.openai.com/v1")
```

### Run in Docker (recommend)

```shell
export MY_OPENAI_API_KEY="sk-ThisIsNotTheRealSecretKey"
docker run -it --rm fredtan/shellgpt -api-key="$MY_OPENAI_API_KEY"
# or
echo "OPENAI_API_KEY=sk-ThisIsNotTheRealSecretKey" > ./shellgpt.env
docker run -it --rm --env-file ./shellgpt.env fredtan/shellgpt
```

### Build from source

```shell
git clone https://github.com/fredtanxo/shellgpt.git
cd shellgpt
make
```

## License

GPL-3.0
