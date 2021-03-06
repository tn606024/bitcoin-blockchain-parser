package main

import (
	"bytes"
	"testing"
)

var (
	transactionTestBytes = []byte{
	0x01,0x00,0x00,0x00,0x00,0x01,0x01,0x9b,
	0x34,0xe2,0x9c,0x56,0x2d,0xe2,0x0e,0xf3,
	0x26,0x68,0x2f,0xa5,0x82,0x67,0x98,0x47,
	0xa4,0xdc,0x4a,0xc7,0xd5,0x26,0x4d,0x32,
	0x28,0xa1,0x7c,0x8f,0xef,0x91,0x0c,0x01,
	0x00,0x00,0x00,0x00,0xff,0xff,0xff,0xff,
	0x01,0x16,0xaa,0x0f,0x00,0x00,0x00,0x00,
	0x00,0x16,0x00,0x14,0x52,0x15,0x83,0x08,
	0xca,0x2e,0x51,0x49,0xd9,0x39,0x73,0x10,
	0xec,0x1f,0xe2,0xa4,0xf8,0x8a,0xfb,0x07,
	0x02,0x48,0x30,0x45,0x02,0x21,0x00,0xce,
	0x94,0x6d,0x85,0x63,0x0e,0x5d,0x4b,0xb2,
	0x27,0xf6,0x96,0x78,0xfa,0xff,0x2d,0x39,
	0x7a,0xdb,0x25,0x61,0x7c,0xc9,0xb4,0xa5,
	0x32,0x3f,0x4b,0x76,0x01,0x21,0xe8,0x02,
	0x20,0x4d,0x05,0x0a,0x9c,0x08,0x83,0x30,
	0xf2,0x36,0xae,0xc8,0x80,0x75,0xde,0x5c,
	0xa6,0x31,0xa5,0xb7,0xa3,0xae,0xa6,0x92,
	0x37,0x1d,0x0f,0x99,0xc9,0x62,0xd7,0x8d,
	0x85,0x01,0x21,0x03,0x2c,0xc6,0x3e,0xba,
	0x4d,0x1f,0x35,0xb6,0x38,0xe7,0x36,0x86,
	0xf7,0x68,0xc2,0xc9,0x63,0x99,0xe6,0x9d,
	0xc7,0x3f,0xcc,0x39,0xfe,0x30,0xb5,0x43,
	0xed,0x85,0x9c,0x19,0x00,0x00,0x00,0x00,
}

	transactionTest = &Transaction {
		Hash:       []byte{
			0xff,0x85,0x7d,0x64,0xd0,0x31,0x82,0x39,
			0x16,0x20,0xa9,0x33,0xfc,0x84,0x77,0x37,
			0x2a,0xf4,0xba,0xc3,0xa3,0xf7,0x37,0x88,
			0x3a,0x57,0x88,0x2d,0x7c,0x7d,0x8f,0x36,
		},
		Version:    1,
		TxInCount:  1,
		TxIn:       []*TxInput{
			{
			Hash:            []byte{
				0x9b,0x34,0xe2,0x9c,0x56,0x2d,0xe2,0x0e,
				0xf3,0x26,0x68,0x2f,0xa5,0x82,0x67,0x98,
				0x47,0xa4,0xdc,0x4a,0xc7,0xd5,0x26,0x4d,
				0x32,0x28,0xa1,0x7c,0x8f,0xef,0x91,0x0c,
				},
			Index:           1,
			ScriptBytes:     0,
			SignatureScript: nil,
			Sequence:        4294967295,
			},
		},
		TxOutCount: 1,
		TxOut:      []*TxOutput{
			{
				Value:         1026582,
				PkScriptBytes: 22,
				PkScript: &PKScript{
					PkScript: []byte{
						0x00,0x14,0x52,0x15,0x83,0x08,0xca,0x2e,
						0x51,0x49,0xd9,0x39,0x73,0x10,0xec,0x1f,
						0xe2,0xa4,0xf8,0x8a,0xfb,0x07,
					},
					Pops: []*parsedOpcode{
						{
							Opcode: &opcodeArray[OP_0],
							Data: nil,
						},
						{
							Opcode: &opcodeArray[OP_DATA_20],
							Data: []byte{
								0x52,0x15,0x83,0x08,0xca,0x2e,0x51,0x49,
								0xd9,0x39,0x73,0x10,0xec,0x1f,0xe2,0xa4,
								0xf8,0x8a,0xfb,0x07,
							},
						},
					},
					Stype:     WitnessPubKeyHash,
					Addresses: []string{
						"bc1q2g2cxzx29eg5nkfewvgwc8lz5nug47c8ta5ene",
					},
				},
			},
		},
		Witness:    []*Witness{
			{
				WitnessStackCount: 2,
				WitnessScript: []*WitnessScript{
					{
						WitnessScriptBytes: 72,
						WitnessScript: []byte{
							0x30,0x45,0x02,0x21,0x00,0xce,0x94,0x6d,
							0x85,0x63,0x0e,0x5d,0x4b,0xb2,0x27,0xf6,
							0x96,0x78,0xfa,0xff,0x2d,0x39,0x7a,0xdb,
							0x25,0x61,0x7c,0xc9,0xb4,0xa5,0x32,0x3f,
							0x4b,0x76,0x01,0x21,0xe8,0x02,0x20,0x4d,
							0x05,0x0a,0x9c,0x08,0x83,0x30,0xf2,0x36,
							0xae,0xc8,0x80,0x75,0xde,0x5c,0xa6,0x31,
							0xa5,0xb7,0xa3,0xae,0xa6,0x92,0x37,0x1d,
							0x0f,0x99,0xc9,0x62,0xd7,0x8d,0x85,0x01,
						},
					},
					{
						WitnessScriptBytes: 33,
						WitnessScript: []byte{
							0x03,0x2c,0xc6,0x3e,0xba,0x4d,0x1f,0x35,
							0xb6,0x38,0xe7,0x36,0x86,0xf7,0x68,0xc2,
							0xc9,0x63,0x99,0xe6,0x9d,0xc7,0x3f,0xcc,
							0x39,0xfe,0x30,0xb5,0x43,0xed,0x85,0x9c,
							0x19},
					},
				},
			},
		},
		LockTime:   0,
	}


)
func TestDeserializeTransaction(t *testing.T) {
	tx, _, err := DeserializeTransaction(transactionTestBytes, &MainnetParams)
	if err != nil {
		t.Errorf("TestDesserializeTransaction error: %s", err)
	}
	//if size != len(transactionTestBytes) {
	//	t.Fatalf("size is not correct, we got: %d, but answer is: %d", size, len(transactionTestBytes))
	//}
	CompareTransaction(tx, transactionTest,t )
}


func CompareTransaction(got *Transaction, ans *Transaction, t *testing.T){
	if got.Version != ans.Version {
		t.Errorf("version is not correct, we got: %d, but answer is: %d", got.Version, ans.Version)
	}
	if bytes.Compare(got.Hash, ans.Hash) != 0 {
		t.Errorf("hash is not correct, we got: %x, but answer is: %x", got.Hash, ans.Hash)
	}
	if got.TxInCount != ans.TxInCount {
		t.Errorf("txInCount is not correct, we got: %d, but answer is: %d", got.TxInCount, ans.TxInCount)
	}
	if got.TxOutCount != ans.TxOutCount {
		t.Errorf("txOutCount is not correct, we got: %d, but answer is: %d", got.TxOutCount, ans.TxOutCount)
	}
	if len(got.TxIn) != len(ans.TxIn) {
		t.Errorf("txin length is not correct, we got: %d, but answer is: %d", len(got.TxIn), len(ans.TxIn))
	}
	if len(got.TxOut) != len(ans.TxOut) {
		t.Errorf("txout length is not correct, we got: %d, but answer is: %d", len(got.TxOut), len(ans.TxOut))
	}
	if len(got.Witness) != len(ans.Witness) {
		t.Fatalf("witness length is not correct, we got: %d, but answer is: %d", len(got.Witness), len(ans.Witness))
	}
	for i := 0; i < len(got.TxIn); i++ {
		CompareTxIn(got.TxIn[i], ans.TxIn[i], t)
	}
	for i := 0; i < len(got.TxOut); i++ {
		CompareTxOutput(got.TxOut[i], ans.TxOut[i], t)
	}
	for i := 0; i < len(got.Witness); i++ {
		CompareWitness(got.Witness[i], ans.Witness[i], t)
	}

}
