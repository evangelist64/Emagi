package data

import (
	"Emagi/log"
	"bytes"
	"encoding/binary"
	"errors"
	fmt "fmt"
	"hash/crc32"
	"io"

	proto "github.com/golang/protobuf/proto"
)

type DataProcessor interface {
	Serialize(data interface{}, w io.Writer) error
	Unserialize(r io.Reader) (interface{}, error)
}

//protobuf
type PBProcessor struct {
	id2Func map[uint32]func(proto.Message)
}

func (p *PBProcessor) init() {
	p.id2Func = make(map[uint32]func(proto.Message))
}

func (p *PBProcessor) Register(msg interface{}, cb func(proto.Message)) {
	pb := msg.(proto.Message)
	msgName := proto.MessageName(pb)
	msgId := GetPBMsgId(msgName)
	for {
		if _, ok := p.id2Func[msgId]; ok {
			log.Info(fmt.Sprintf("duplicate msgid, msgName=%s", msgName))
			msgId++
		} else {
			break
		}
	}
	p.id2Func[msgId] = cb
}

func (p *PBProcessor) UnRegister(msg interface{}) {
	pb := msg.(proto.Message)
	msgName := proto.MessageName(pb)
	delete(p.id2Func, GetPBMsgId(msgName))
}

func (p *PBProcessor) Serialize(data interface{}, w io.Writer) error {
	//创建一个临时buffer存放相关数据，保证写入内容的完整性
	var b bytes.Buffer
	pb, ok := data.(proto.Message)
	if !ok {
		return errors.New("msg and processor not match")
	}
	//id
	err := binary.Write(&b, binary.BigEndian, GetPBMsgId(proto.MessageName(pb)))
	if err != nil {
		return err
	}
	msgBytes, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	//length
	err = binary.Write(&b, binary.BigEndian, (uint32)(len(msgBytes)))
	if err != nil {
		return err
	}
	//data
	err = binary.Write(&b, binary.BigEndian, msgBytes)
	if err != nil {
		return err
	}
	return binary.Write(w, binary.BigEndian, b.Bytes())
}

func (p *PBProcessor) Unserialize(r io.Reader) (interface{}, error) {
	return nil, nil
}

func GetPBMsgId(msgName string) uint32 {
	return crc32.ChecksumIEEE([]byte(msgName))
}
