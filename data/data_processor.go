package data

import (
	"Emagi/log"
	"bytes"
	"encoding/binary"
	"errors"
	fmt "fmt"
	"hash/crc32"
	"io"
	"reflect"

	proto "github.com/golang/protobuf/proto"
)

type DataProcessor interface {
	//todo 加密解密(AES?)
	Serialize(data interface{}, w io.Writer) error
	Unserialize(r io.Reader) error
}

//protobuf
type PBProcessor struct {
	id2Func map[uint32]func(proto.Message)
	id2Type map[uint32]reflect.Type
}

func (p *PBProcessor) Init() {
	p.id2Func = make(map[uint32]func(proto.Message))
	p.id2Type = make(map[uint32]reflect.Type)
}

func (p *PBProcessor) Register(msg proto.Message, cb func(proto.Message)) {
	msgName := proto.MessageName(msg)
	//直接用crc32转换成id，保证双端规则一致即可
	msgId := getPBMsgId(msgName)
	for {
		if _, ok := p.id2Func[msgId]; ok {
			log.Info(fmt.Sprintf("duplicate msgid, msgName=%s", msgName))
			msgId++
		} else {
			break
		}
	}
	p.id2Func[msgId] = cb
	p.id2Type[msgId] = reflect.TypeOf(msg)
}

func (p *PBProcessor) UnRegister(msg proto.Message) {
	msgName := proto.MessageName(msg)
	delete(p.id2Func, getPBMsgId(msgName))
}

func (p *PBProcessor) Serialize(data interface{}, w io.Writer) error {
	//创建一个临时buffer存放相关数据，保证写入内容的完整性
	b := new(bytes.Buffer)
	pb, ok := data.(proto.Message)
	if !ok {
		return errors.New("msg and processor not match")
	}
	//id
	err := binary.Write(b, binary.BigEndian, getPBMsgId(proto.MessageName(pb)))
	if err != nil {
		return err
	}
	dataBytes, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	//length
	err = binary.Write(b, binary.BigEndian, (uint32)(len(dataBytes)))
	if err != nil {
		return err
	}
	//data
	err = binary.Write(b, binary.BigEndian, dataBytes)
	if err != nil {
		return err
	}
	_, err = w.Write(b.Bytes())

	return err
}

func (p *PBProcessor) Unserialize(r io.Reader) error {
	b := make([]byte, 4)
	//id
	_, err := io.ReadFull(r, b)
	if err != nil {
		return err
	}
	id := binary.BigEndian.Uint32(b)
	//length
	_, err = io.ReadFull(r, b)
	if err != nil {
		return err
	}
	dataLen := binary.BigEndian.Uint32(b)

	//data
	dataBuf := make([]byte, dataLen)
	_, err = io.ReadFull(r, dataBuf)
	if err != nil {
		return err
	}

	//转成对应数据结构
	//todo 把反序列化和消息处理分开，消息处理扔到主线程调用
	t, exist := p.id2Type[id]
	if !exist {
		return errors.New("invalid msgid")
	}
	f, exist := p.id2Func[id]
	if !exist {
		return errors.New("invalid msgid")
	}

	pb := reflect.New(t.Elem()).Interface()
	pbMsg := pb.(proto.Message)
	err = proto.Unmarshal(dataBuf, pbMsg)
	if err != nil {
		return err
	}

	f(pbMsg)
	return nil
}

func getPBMsgId(msgName string) uint32 {
	return crc32.ChecksumIEEE([]byte(msgName))
}
