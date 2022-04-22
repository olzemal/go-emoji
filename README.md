# go-emoji
CLI tools to interact with the [Full Emoji List](https://unicode.org/emoji/charts/full-emoji-list.html).

## lsemoji
`lsemoji` queries or lists emojis from the list.  
Example:

```console
$ lsemoji dog
hot_dog=🌭
guide_dog=🦮
service_dog=🐕‍🦺
dog=🐕
dog_face=🐶
```

## toemoji
`toemoji` replaces all emojis formatted like `:name:` from `stdin` and
writes them back to `stdout`.  
Example:

```console
$ echo "hello :sun:" | toemoji
hello ☀
```
