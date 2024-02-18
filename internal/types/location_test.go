package types

import "testing"

func Test_buildAddressGroup_KoreanAddresses(t *testing.T) {
	type args struct {
		sdName  string
		sggName string
		emdName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "All parts of address group are present",
			args: args{
				sdName:  "충청남도",
				sggName: "천안시 동남구",
				emdName: "북면",
			},
			want: "충청남도 천안시 동남구 북면",
		},
		{
			name: "EMD name ends with 동",
			args: args{
				sdName:  "서울특별시",
				sggName: "강남구",
				emdName: "역삼1동",
			},
			want: "서울특별시 강남구",
		},
		{
			name: "One part of address group is missing",
			args: args{
				sdName:  "충청남도",
				sggName: "",
				emdName: "북면",
			},
			want: "충청남도 북면",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildSdSggEm(tt.args.sdName, tt.args.sggName, tt.args.emdName); got != tt.want {
				t.Errorf("buildSdSggEm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildRoadNameGroup_KoreanRoads(t *testing.T) {
	type args struct {
		roadName           string
		buildingMainNumber string
		buildingSubNumber  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Building sub number is not empty and not zero",
			args: args{
				roadName:           "강남대로",
				buildingMainNumber: "123",
				buildingSubNumber:  "4",
			},
			want: "강남대로 123-4",
		},
		{
			name: "Building sub number is empty",
			args: args{
				roadName:           "강남대로",
				buildingMainNumber: "123",
				buildingSubNumber:  "",
			},
			want: "강남대로 123",
		},
		{
			name: "Building sub number is zero",
			args: args{
				roadName:           "강남대로",
				buildingMainNumber: "123",
				buildingSubNumber:  "0",
			},
			want: "강남대로 123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildAddrDetail(tt.args.roadName, tt.args.buildingMainNumber, tt.args.buildingSubNumber); got != tt.want {
				t.Errorf("buildAddrDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
