package combinatorics

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCnk(t *testing.T) {
	type args struct {
		n int64
		k int64
	}
	tests := []struct {
		name string
		args args
		want big.Int
	}{
		{
			name: `Big int N=100 K=10`,
			args: args{
				n: 1000000,
				k: 3,
			},
			want: *big.NewInt(161700),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cnk(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf(`Cnk() = %v, want %v`, got, tt.want)
			}
		})
	}
}

func TestCnkUint(t *testing.T) {
	type args struct {
		n int64
		k int64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: `Uint64 int N=10 K=2`,
			args: args{
				n: 10,
				k: 2,
			},
			want: 45,
		},
		{
			name: `Uint64 int N=100 K=3`,
			args: args{
				n: 100,
				k: 3,
			},
			want: 161700,
		},
		{
			name: `Uint64 int N=1000 K=3`,
			args: args{
				n: 1000,
				k: 3,
			},
			want: 166167000,
		},
		{
			name: `Uint64 int N=1000 K=2`,
			args: args{
				n: 1000,
				k: 2,
			},
			want: 499500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CnkUint(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf(`CnkUint() = %v, want %v`, got, tt.want)
			}
		})
	}
}

func TestCnkStr(t *testing.T) {
	type args struct {
		n int64
		k int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `String int N=1000 K=2`,
			args: args{
				n: 1000,
				k: 2,
			},
			want: `499500`,
		},
		{
			name: `String int N=10000 K=20`,
			args: args{
				n: 10000,
				k: 20,
			},
			want: `40329089689367321847583694286698379982154793605060337426312000`,
		},
		{
			name: `String int N=100000 K=200`,
			args: args{
				n: 100000,
				k: 200,
			},
			want: `10390327301407696296947585866200244344914126174604513814671137403995324122598099169243332024144139660488824536654448170343406378688490194282395601942733016530226681727974437549686413900512750677435710654556383182308686247957628417543710122692563997357035473610983279620757224908931720735629527614329223690402631251342155503152924803442233970219747981636674700509124521982439882726578149452272271406746658972193494039610440206754298278357388358035606418515896063771696839079196169415672834770862969864687192330140796264138657972077457246422785949473767013072645685742067168551474887930288982986001358650432185971424047336212000`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CnkStr(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf(`CnkStr() = %v, want %v`, got, tt.want)
			}
		})
	}
}
