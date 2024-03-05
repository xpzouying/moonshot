# moonshot
Go sdk for moonshot


## Chat

Result of chat completions.

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


Error chat completions.

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


## List Models

Result of list models.

```json
{
    "data": [
        {
            "created": 1709149158,
            "id": "moonshot-v1-8k",
            "object": "model",
            "owned_by": "moonshot",
            "permission": [
                {
                    "created": 0,
                    "id": "",
                    "object": "",
                    "allow_create_engine": false,
                    "allow_sampling": false,
                    "allow_logprobs": false,
                    "allow_search_indices": false,
                    "allow_view": false,
                    "allow_fine_tuning": false,
                    "organization": "public",
                    "group": "public",
                    "is_blocking": false
                }
            ],
            "root": "",
            "parent": ""
        },
        {
            "created": 1709149158,
            "id": "moonshot-v1-32k",
            "object": "model",
            "owned_by": "moonshot",
            "permission": [
                {
                    "created": 0,
                    "id": "",
                    "object": "",
                    "allow_create_engine": false,
                    "allow_sampling": false,
                    "allow_logprobs": false,
                    "allow_search_indices": false,
                    "allow_view": false,
                    "allow_fine_tuning": false,
                    "organization": "public",
                    "group": "public",
                    "is_blocking": false
                }
            ],
            "root": "",
            "parent": ""
        },
        {
            "created": 1709149158,
            "id": "moonshot-v1-128k",
            "object": "model",
            "owned_by": "moonshot",
            "permission": [
                {
                    "created": 0,
                    "id": "",
                    "object": "",
                    "allow_create_engine": false,
                    "allow_sampling": false,
                    "allow_logprobs": false,
                    "allow_search_indices": false,
                    "allow_view": false,
                    "allow_fine_tuning": false,
                    "organization": "public",
                    "group": "public",
                    "is_blocking": false
                }
            ],
            "root": "",
            "parent": ""
        }
    ]
}
```