// parser_test.go
package serv

import (
	"testing"
)

// Just for emulating field access, so it will not throw "evaluated but not used"
func nothing(_ ...interface{}) {}

func BenchmarkDecode(b *testing.B) {
	data := NewTransferRequest().ToJSON()
	for decoderName, decodeFunc := range DecodesList {
		b.Run(decoderName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tr, _ := decodeFunc(data)
				nothing(tr)
			}
		})
	}
}
