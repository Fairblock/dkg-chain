package keeper

import (
	"bytes"
	"context"
	"encoding/binary"

	"fmt"
	"io"

	types1 "github.com/cosmos/cosmos-sdk/codec/types"

	"strconv"

	"dkg/x/dkg/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RefundMsgRequest(goCtx context.Context, msg *types.MsgRefundMsgRequest) (*types.MsgRefundMsgRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	b, _ := msg.Marshal()

	count := k.IncreaseCounter(ctx, 1)
	str_count := strconv.FormatUint(count, 10)

	_ = ctx
	event := sdk.NewEvent(
		types.EventTypeKeygen,
		sdk.NewAttribute(types.AttributeValueMsg, string(b)),
		sdk.NewAttribute("module", "dkg"),
		sdk.NewAttribute("index", str_count),
	)
	ctx.EventManager().EmitEvent(event)

	msgBack := Unmarshal(b)

	message := new(types.ProcessKeygenTrafficRequest)

	message.Unmarshal(msgBack.InnerMessage.Value)

	if message.Payload.RoundNum == "1" {
		if message.Payload.IsBroadcast {

			bcast := new(Bcast)
			bcast.UnmarshalBinary(message.Payload.Payload)

			k.AddPk(ctx, bcast.UIVssCommit.CoeffCommits[0], uint64(bcast.ID))

		}
	}

	return &types.MsgRefundMsgRequestResponse{}, nil
}

func Unmarshal(dAtA []byte) (m types.MsgRefundMsgRequest) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				panic("error")
			}
			if iNdEx >= l {
				panic("error")
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			panic("error")
		}
		if fieldNum <= 0 {
			panic("error")
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				panic("error")
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					panic("error")
				}
				if iNdEx >= l {
					panic("error")
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				panic("error")
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				panic("error")
			}
			if postIndex > l {
				panic("error")
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				panic("error")
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					panic("error")
				}
				if iNdEx >= l {
					panic("error")
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				panic("error")
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				panic("error")
			}
			if postIndex > l {
				panic("error")
			}
			if m.InnerMessage == nil {
				m.InnerMessage = &types1.Any{}
			}
			if err := m.InnerMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				panic("error")
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				panic("error")
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				panic("error")
			}
			if (iNdEx + skippy) > l {
				panic("error")
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		panic("error")
	}
	return m
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

// UnmarshalBinary for the Bcast type.
func (b *Bcast) UnmarshalBinary(data []byte) error {
	idData := data[len(data)-2:]
	id := binary.LittleEndian.Uint16(idData)
	b.ID = uint(id)

	commitData := data[:len(data)-2]
	var commit vssCommit
	err := commit.UnmarshalBinary(commitData)
	if err != nil {
		return err
	}
	b.UIVssCommit = commit

	return nil
}

// UnmarshalBinary for the vssCommit type.
func (c *vssCommit) UnmarshalBinary(data []byte) error {

	commitSize := 48
	index := 0
	for i := 0; i < len(data); i++ {
		if data[i] == 48 {
			if data[i-1] != 0 {
				index = i + 1
				break
			}
		}
	}
	data = data[index:]
	var commits [][]byte
	reader := bytes.NewReader(data)
	for {
		commit := make([]byte, commitSize)
		_, err := reader.Read(commit)
		if err != nil {
			break
		}
		commits = append(commits, commit)
	}
	c.CoeffCommits = commits

	return nil
}
