# moonshot
Go sdk for moonshot


## 正确请求

```json
{
    "id": "cmpl-df0e829246bf457494ab0d13fb5441a3",
    "object": "chat.completion",
    "created": 4325361,
    "model": "moonshot-v1-8k",
    "choices": [
        {
            "index": 0,
            "message": {
                "role": "assistant",
                "content": " 你好，李雷！1+1 等于 2。如果你有其他问题或需要帮助，随时告诉我！"
            },
            "finish_reason": "stop"
        }
    ],
    "usage": {
        "prompt_tokens": 81,
        "completion_tokens": 26,
        "total_tokens": 107
    }
}
```

## 请求错误

```json
{
    "code": 5,
    "error": "url.not_found",
    "message": "没找到对象",
    "method": "GET",
    "scode": "0x5",
    "status": false,
    "ua": "Go-http-client/2.0",
    "url": "/v1/chat/completions"
}
```