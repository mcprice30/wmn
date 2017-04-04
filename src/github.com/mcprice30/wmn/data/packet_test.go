// Testing for package data.
package data

import (
	"reflect"
	"testing"
)

// TestPacketHeaderMarshallUnmarshall ensures that a packet can be marshalled to
// bytes and then successfully unmarshalled back from those bytes.
func TestPacketHeaderMarshallUnmarshall(t *testing.T) {
	original := &PacketHeader{
		SourceAddress:      0x0123,
		DestinationAddress: 0x0045,
		PreviousHop:        0x0678,
		TTL:                25,
		PacketType:         PacketTypeControl,
		SequenceNumber:     0x1111,
		NumBytes:           0x0A,
	}

	bytes := original.ToBytes()
	result := PacketHeaderFromBytes(bytes)
	if !reflect.DeepEqual(result, original) {
		t.Error("Result of ", result, " is different from the original: ", original)
	}
}

// TestPacketMarshallUnmarshall ensures that a packet can be marshalled to bytes
// and then successfully unmarshalled back from those bytes.
func TestPacketMarshallUnmarshall(t *testing.T) {

	original := &DataPacket{
		Header: PacketHeader{
			SourceAddress:      0x0123,
			DestinationAddress: 0x0045,
			PreviousHop:        0x0678,
			TTL:                25,
			PacketType:         PacketTypeControl,
			SequenceNumber:     0x1111,
			NumBytes:           0x0A,
		},
		Body: []SensorData{
			CreateGasData(0, 15.000),
			CreateOxygenData(0, 98.7),
			CreateGasData(1, 16.999),
			CreateLocationData(0, 67.1, 54.3),
			CreateHeartRateData(0, 155.34),
			CreateLocationData(1, 67.1, 54.4),
		},
	}

	bytes := original.ToBytes()
	result := DataPacketFromBytes(bytes)

	if result == nil {
		t.Fatal("Result was nil!")
	}

	if !reflect.DeepEqual(result.Header, original.Header) {
		t.Fatal("Result of ", result.Header, " is different from the original: ",
			original.Header)
	}

	if len(result.Body) != len(original.Body) {
		t.Fatal("Result length", len(result.Body), " differs from original",
			len(original.Body))
	}

	for i, data := range original.Body {
		if result.Body[i].String() != data.String() {
			t.Fatal("Result", i, "data of ", result.Body[i].String(),
				"differs from original", data.String())
		}
	}
}
