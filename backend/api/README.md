# Proto开发

## Buf使用

当然，在使用前，您需要安装buf
接下来 在backend\api
```bash
go install github.com/bufbuild/buf/cmd/buf@latest

buf --version
```

### 更新buf.lock
```bash
buf dep update
```

### 检查proto语法正确
```bash
buf lint
```

### 生成proto代码
buf generate