package model

import "testing"

//func TestBingoCard_GenerateCard(t *testing.T) {
//	type args struct {
//		bingoSeed int64
//	}
//	tests := []struct {
//		name string
//		args args
//		want [5][5]uint
//	}{
//		{
//			name: "success1",
//			args: args{
//				bingoSeed: 1,
//			},
//			want: [5][5]uint{
//				{12, 10, 14, 15, 3},
//				{19, 20, 16, 23, 26},
//				{45, 44, 31, 37, 33},
//				{60, 53, 52, 51, 54},
//				{66, 63, 70, 74, 73},
//			},
//		},
//		{
//			name: "success2",
//			args: args{
//				bingoSeed: 1000000,
//			},
//			want: [5][5]uint{
//				{6, 3, 12, 14, 13},
//				{24, 25, 22, 17, 28},
//				{32, 40, 36, 39, 34},
//				{55, 49, 48, 56, 57},
//				{71, 72, 73, 70, 66},
//			},
//		},
//		{
//			name: "success3",
//			args: args{
//				bingoSeed: 6964835249063,
//			},
//			want: [5][5]uint{
//				{1, 10, 15, 11, 4},
//				{22, 20, 21, 25, 24},
//				{41, 38, 36, 43, 44},
//				{48, 51, 55, 53, 57},
//				{68, 67, 73, 62, 74},
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			bc := RebuildBingoCard(tt.args.bingoSeed)
//			if got := bc.GenerateCard(); got != tt.want {
//				t.Errorf("GenerateCard() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestBingoCard_CheckBingo(t *testing.T) {
	type args struct {
		bingoSeed    int64
		bingoNumbers []uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "HorizontalBingo1",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					12, 10, 14, 15, 3,
				},
			},
			want: true,
		},
		{
			name: "HorizontalBingo2",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					19, 20, 16, 23, 26,
				},
			},
			want: true,
		},
		{
			name: "HorizontalBingo3",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					45, 44, 31, 37, 33,
				},
			},
			want: true,
		},
		{
			name: "HorizontalBingo4",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					60, 53, 52, 51, 54,
				},
			},
			want: true,
		},
		{
			name: "HorizontalBingo5",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					66, 63, 70, 74, 73,
				},
			},
			want: true,
		},
		{
			name: "NotHorizontalBingo",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					66, 63, 34, 54, 74, 73,
				},
			},
			want: false,
		},
		{
			name: "VerticalBingo1",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					12, 19, 45, 60, 66,
				},
			},
			want: true,
		},
		{
			name: "VerticalBingo2",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					10, 20, 44, 53, 63,
				},
			},
			want: true,
		},
		{
			name: "VerticalBingo3",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					14, 16, 31, 52, 70,
				},
			},
			want: true,
		},
		{
			name: "VerticalBingo4",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					15, 23, 37, 51, 74,
				},
			},
			want: true,
		},
		{
			name: "VerticalBingo5",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					3, 26, 33, 54, 73,
				},
			},
			want: true,
		},
		{
			name: "NotVerticalBingo",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					3, 26, 1, 54, 73,
				},
			},
			want: false,
		},
		{
			name: "DiagonalBingo1",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					12, 20, 51, 73,
				},
			},
			want: true,
		},
		{
			name: "DiagonalBingo2",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					3, 23, 53, 66,
				},
			},
			want: true,
		},
		{
			name: "NotDiagonalBingo",
			args: args{
				bingoSeed: 1,
				bingoNumbers: []uint{
					3, 2, 44, 51, 66,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := RebuildBingoCard(tt.args.bingoSeed, tt.args.bingoNumbers)
			if got := bc.CheckBingo(); got != tt.want {
				t.Errorf("CheckBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}
