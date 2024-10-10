package hash

import (
	"testing"
)

func TestHashService(t *testing.T) {
	t.Parallel()

	secretTest1 := "Pass12345"
	hashTest1 := Hash(secretTest1)

	secretTest2 := "AnotherPassword"

	type args struct {
		password    string
		hashedValue string
		shouldPass  bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "should hash password correctly",
			args: args{
				password:    secretTest1,
				hashedValue: hashTest1,
				shouldPass:  true,
			},
		},
		{
			name: "should fail verification with incorrect password",
			args: args{
				password:    secretTest2,
				hashedValue: "diffHash",
				shouldPass:  false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			verifyResult := Verify(tt.args.password, tt.args.hashedValue)
			if verifyResult != tt.args.shouldPass {
				t.Errorf("hash verify = %v, want %v", verifyResult, tt.args.shouldPass)
			}
		})
	}
}
