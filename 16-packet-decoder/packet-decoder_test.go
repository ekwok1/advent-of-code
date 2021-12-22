package main

import (
	"testing"
)

func TestParsePacketVersions_8A004A801A8002F478(t *testing.T) {
	hexadecimal := "8A004A801A8002F478"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	versionCounter := parsePacketVersions(binary, -1)
	if versionCounter != 16 {
		t.Errorf("parsePacketVersions(binary, -1) = %d; want 16", versionCounter)
	}
}

func TestParsePacketVersions_620080001611562C8802118E34(t *testing.T) {
	hexadecimal := "620080001611562C8802118E34"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	versionCounter := parsePacketVersions(binary, -1)
	if versionCounter != 12 {
		t.Errorf("parsePacketVersions(binary, -1) = %d; want 12", versionCounter)
	}
}

func TestParsePacketVersions_C0015000016115A2E0802F182340(t *testing.T) {
	hexadecimal := "C0015000016115A2E0802F182340"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	versionCounter := parsePacketVersions(binary, -1)
	if versionCounter != 23 {
		t.Errorf("parsePacketVersions(binary, -1) = %d; want 23", versionCounter)
	}
}

func TestParsePacketVersions_A0016C880162017C3686B18A3D4780(t *testing.T) {
	hexadecimal := "A0016C880162017C3686B18A3D4780"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	versionCounter := parsePacketVersions(binary, -1)
	if versionCounter != 31 {
		t.Errorf("parsePacketVersions(binary, -1) = %d; want 31", versionCounter)
	}
}

func TestParseAndOperate_C200B40A82(t *testing.T) {
	hexadecimal := "C200B40A82"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 3 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 3", value)
	}
}

func TestParseAndOperate_04005AC33890(t *testing.T) {
	hexadecimal := "04005AC33890"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 54 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 54", value)
	}
}

func TestParseAndOperate_880086C3E88112(t *testing.T) {
	hexadecimal := "880086C3E88112"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 7 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 7", value)
	}
}

func TestParseAndOperate_CE00C43D881120(t *testing.T) {
	hexadecimal := "CE00C43D881120"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 9 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 9", value)
	}
}

func TestParseAndOperate_D8005AC2A8F0(t *testing.T) {
	hexadecimal := "D8005AC2A8F0"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 1 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 1", value)
	}
}

func TestParseAndOperate_F600BC2D8F(t *testing.T) {
	hexadecimal := "F600BC2D8F"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 0 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 0", value)
	}
}

func TestParseAndOperate_9C005AC2F8F0(t *testing.T) {
	hexadecimal := "9C005AC2F8F0"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 0 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 0", value)
	}
}

func TestParseAndOperate_9C0141080250320F1802104A08(t *testing.T) {
	hexadecimal := "9C0141080250320F1802104A08"

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal, &hexadecimalMap)

	value, _ := parseAndOperate(binary, 0)
	if value != 1 {
		t.Errorf("parseAndOperate(binary, 0) = %d; want 1", value)
	}
}
