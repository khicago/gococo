# gococo

**gococo** is a tool that helps you parse **commands** from blocks of text.

You can use it to easily retrieve executable commands from documents, code comments, or any text.

## Usage

gococo is very easy to use, just call `gococo.Parse("your awesome text")` and get a list of `gococo.Coco` instances.

Each **Coco** instance is an implementation of the stringer interface, which also has several methods for executing commands.

- `coco.GetCMD() string` - Get the command of type string.
- `coco.GetParams() []gococo.Param` - Acquisition parameters of type `gococo.Param`.

`gococo.Param` is an alias of type string, but has an API to determine if the input parameter is null: `param.IsNil() bool`.

This is because an empty string is often **ambiguous** when parsing text, so it's needed to agree on a specific symbol to represent the "null string".

By default, the null value is a single underscore `_`, and you can change this by changing `gococo.NilParam`.

### The text format of the command

gococo has provided a **default command format** definition, called **coco format**, along with (de)serialization methods.

A standard coco format looks like this:

`<% YOUR_AWWWSOME_CMD: _, PARAM1, PARAM2, PARAM3 %>`

where

- `YOUR_AWWWSOME_CMD` is the command to be executed,
- `_` means that it is the nil string mentioned above.
- `PARAM1, PARAM2, PARAM3 ....` is a continuous input parameter.

The **mark symbols** involved are

- `<%` starting symbol
- `%>`end symbol
- `:` command separator
- `,` parameter separator (computing)
- blank character, supports `' '` or `'\t'`

The coco format also supports some not-entirely-strict write-ins.

- There can be no white space between the marker and the **identifier** (CMD or Param).
- CMD followed by a colon may be omitted.
- Must have CMD, but can have no Params
- The comma between the Params can be omitted.
- The number of split identifier can be any number of more than one.
- No two mark symbols should be directly spaced apart.
- The comma after the last Param can be absent or present. Note that if the last `,` exists, there cannot be a gap between `,` and the last `%>`. It is recommended not to write the last `,`.

Here are some proper demonstrations :

```coco
<%YOUR_AWWWSOME_CMD:_,PARAM1,PARAM2,PARAM3%>
<% YOUR_AWWWSOME_CMD _, PARAM1, PARAM2, PARAM3 %>
<% YOUR_AWWWSOME_CMD %>
<% YOUR_AWWWSOME_CMD _ PARAM1 PARAM2 PARAM3 %>
<%    YOUR_AWWWSOME_CMD    :  _  ,  PARAM1  ,  PARAM2  ,   PARAM3   %>
<% YOUR_AWWWSOME_CMD _ PARAM1 PARAM2 PARAM3 ,%>
```

There are some strict areas to watch out for too

- A coco format command must strictly begin with `<%` and end with `%>`
- Only spaces and tabs can be used to split an identifier. Other symbols, such as `'\n'`, are not allowed.
- Param supports any characters other than `\s`(any symbol standards empty) `,` `:` `%` `~`.  
  All other symbols are supported, including unicode for all languages.

for more examples, see the [gococo_test.go file](./gococo_test.go)

### custom (de)serialization methods

If the default definition doesn't meet your needs, gococo also supports custom (de)serialization methods to support customizable command formats.

By setting `gococo.Matcher`, it is possible to change the behavior of the cmd recognition from text. A Matcher matches the type

```go
type MatcherHandler func(strIn string) (results [][]string, find bool)
```

By setting `gococo.Serializer`, it is possible to change the content of serialized cmd to text. A Serializer matches the type

```go
type SerializeHandler func(co gococo.Coco) string
```

for more information, see the [export.go file](./export.go)

## Contribution

Issues and pull requests are welcome.

<details><summary style="color:royalblue">Contact</summary>
kinghand@foxmail.com
</details>

## License

[MIT](./LICENSE)
