package generator

import (
	"encoding/json"
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/go-redis/redis"
	"gitlab.com/charliemaiors/shinobi-param-generator/parser"
)

type redisJenniferGen struct {
	addr     string
	password string
	db       int
}

//RedisPath represent the subpath plus a flag for indication if is secure
type RedisPath = parser.Path

//RedisCodec represent a subset of paths for a given codec
type RedisCodec struct {
	CodecType string      `json:"codecType"`
	Paths     []RedisPath `json:"paths"`
}

//RedisProtocol is the structure which represent a particular IPCam with it's vendor, protocol, codec and relative subpath
type RedisProtocol struct {
	Models []string     `json:"models"`
	Codecs []RedisCodec `json:"codecs"`
}

//NewRedisGenerator instantiate a generator which deploy all data on redis and generate a client file
func NewRedisGenerator(addr, password string, db int) ShinobiGenerator {
	return &redisJenniferGen{
		addr:     addr,
		password: password,
		db:       db,
	}
}

func (codec RedisCodec) MarshalBinary() ([]byte, error) {
	return json.Marshal(&codec)

}

func (protocol RedisProtocol) MarshalBinary() ([]byte, error) {
	return json.Marshal(&protocol)
}

func (gen *redisJenniferGen) Generate(pkg, dest string, params []parser.ShinobiParam) error {
	fmt.Printf("Dest is %s", dest)
	client := redis.NewClient(&redis.Options{
		Addr:     gen.addr,
		Password: gen.password,
		DB:       gen.db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		return err
	}

	for _, param := range params {
		value := gen.generateMap(param.Protocols)
		if len(value) == 0 {
			value["Not supported"] = RedisProtocol{}
		}
		status := client.HMSet(param.Vendor, value) //consider implementing encoding.BinaryMarshaler
		if err = status.Err(); err != nil {
			return err
		}
	}
	err = gen.generateRedisFile(pkg, dest)
	return err
}

func (gen *redisJenniferGen) generateMap(params []parser.Protocol) map[string]interface{} {
	res := make(map[string]interface{})

	for _, protocol := range params {
		redisProt := gen.generateRedisProtocol(protocol)
		res[protocol.Protocol.String()] = redisProt
	}

	return res
}

func (gen *redisJenniferGen) generateRedisProtocol(protocol parser.Protocol) interface{} {
	response := RedisProtocol{
		Codecs: gen.convertCodecs(protocol.Codecs),
		Models: protocol.Models,
	}
	return response
}

func (gen *redisJenniferGen) convertCodecs(codecs []parser.Codec) []RedisCodec {
	res := make([]RedisCodec, 0, 0)
	for _, codec := range codecs {
		tmp := RedisCodec{
			CodecType: codec.Type.String(),
			Paths:     codec.Paths,
		}
		res = append(res, tmp)
	}
	return res
}

func (gen *redisJenniferGen) generateRedisFile(pkg, dest string) error {
	file := jen.NewFilePathName(dest, pkg)
	file.Comment("This is a generated file for mapping shinobi params from redis source, DO NOT EDIT!!!")

	file.ImportName("github.com/go-redis/redis", "redis")

	file.Comment("Path represent a path with given default values if present")
	file.Type().Id("Path").Struct(
		jen.Id("IsSecure").Bool().Tag(map[string]string{"json": "isSecure"}),
		jen.Id("SubPath").String().Tag(map[string]string{"json": "subpath"}),
	)

	file.Comment("//Codec represent a subset of paths for a given codec")
	file.Type().Id("Codec").Struct(
		jen.Id("CodecType").String().Tag(map[string]string{"json": "codecType"}),
		jen.Id("Paths").Index().Id("Path").Tag(map[string]string{"json": "paths"}),
	)

	file.Comment("//Protocol is the structure which represent a particular IPCam with it's vendor, protocol, codec and relative subpath")
	file.Type().Id("Protocol").Struct(
		jen.Id("Connection").String().Tag(map[string]string{"json": "connection,omitempty"}),
		jen.Id("Models").Index().String().Tag(map[string]string{"json": "models"}),
		jen.Id("Codecs").Id("[]Codec").Tag(map[string]string{"json": "codecs"}),
	)

	file.Type().Id("RedisParamsExtractor").Struct(
		jen.Id("internalRedisClient").Op("*").Id("redis.Client"),
	)

	file.Func().Id("NewParamsExtractor").Params(jen.Id("address").Op(",").Id("password").String(), jen.Id("db").Int()).Id("*RedisParamsExtractor").Block(
		jen.Return(
			jen.Op("&").Id("RedisParamsExtractor").Values(jen.Dict{
				jen.Id("internalRedisClient"): jen.Qual("github.com/go-redis/redis", "NewClient").Call(jen.Op("&").Id("redis").Op(".").Id("Options").Values(jen.Dict{
					jen.Id("Addr"):     jen.Id("address"),
					jen.Id("Password"): jen.Id("password"),
					jen.Id("DB"):       jen.Id("db"),
				})),
			}),
		),
	)

	file.Func().Parens(jen.Id("extractor").Op("*").Id("RedisParamsExtractor")).Id("TakeAllPathsForVendorWithProtocol").Params(jen.Id("vendor").String(), jen.Id("protocol").String()).Parens(jen.Id("Protocol").Op(",").Error()).Block(
		jen.Id("status").Op(":=").Id("extractor").Op(".").Id("internalRedisClient").Op(".").Id("HGet").Call(jen.Id("vendor"), jen.Id("protocol")),
		jen.If(jen.Id("status").Op(".").Id("Err").Call().Op("!=").Nil()).Block(
			jen.Return(jen.Id("Protocol{}"), jen.Id("status").Op(".").Id("Err").Call()),
		),
		jen.Id("res").Op(":=").Id("Protocol{}"),
		jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Index().Byte().Call(jen.Id("status").Op(".").Id("Val").Call()), jen.Op("&").Id("res")),
		jen.Id("res").Op(".").Id("Connection").Op("=").Id("protocol"),
		jen.Return(jen.Id("res"), jen.Id("err")),
	)

	fmt.Printf("Saving file to %s", dest+"redis.go")
	err := file.Save(dest + "redis.go")
	fmt.Printf("Error? %v", err)
	return err
}
