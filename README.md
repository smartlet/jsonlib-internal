# json-iterator修改
- 基于github.com/json-iterator/go@v1.1.2(实现是master)
- 保留"*.go"
- 去除"*_test.go"
- 去除"extra"子目录内容
- 修改Marshal()接口添加ignoreOmitempty参数.
- 修改NewDecoder().Decode()接口添加ignoreOmitempty参数.
- 修改Stream对象添加ignoreOmitempty配置,并在Reset()后设置为false.
- 修改StructFieldEncoder

# reflect2修改
- 基于github.com/modern-go/reflect2@v1.0.2
- 保留"*.go"与"*.s"