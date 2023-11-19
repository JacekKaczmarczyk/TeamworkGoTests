package customerimporter

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

type mockReader struct {
	Data []byte
	err  error
}

func (r *mockReader) Open(fileName string) (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(r.Data)), r.err
}

func TestCountAndSortEmailDomains(t *testing.T) {
	tests := []struct {
		name       string
		want       []Domain
		wantErr    bool
		mockReader mockReader
	}{
		{
			name: "Test_1",
			want: []Domain{
				{
					DomainName: "ezinearticles.com",
					Count:      1,
				},
				{
					DomainName: "othertest.com",
					Count:      4,
				},
				{
					DomainName: "test.com",
					Count:      2,
				},
			},
			wantErr: false,
			mockReader: mockReader{
				Data: []byte(`Norma,Allen,nallen8@test.com,Female,168.67.162.1
				Lillian,Lawrence,llawrence9@test.com,Female,190.106.124.105
				Irene,Crawford,icrawforda@test.com,Female,156.30.64.85
				Shirley,Alvarez,salvarezb@othertest.com,Female,233.224.134.184
				Patricia,Sims,psimsc@othertest.com,Female,235.115.22.151
				Joyce,Sanchez,jsancheze@ezinearticles.com,Female,246.105.133.101
				Joe,Washington,jwashingtonf@othertest.com,Male,60.176.60.134
				Anna,Rivera,ariverag@othertest.com,Female,105.158.80.2`),
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountAndSortEmailDomains(&tt.mockReader)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountAndSortEmailDomains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountAndSortEmailDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertDomainsMapToStruct(t *testing.T) {

	tests := []struct {
		name      string
		domainMap map[string]int
		want      []Domain
	}{
		{
			name: "Test_1",
			domainMap: map[string]int{
				"test.com":      2,
				"otherTest.com": 10,
			},
			want: []Domain{
				{
					DomainName: "test.com",
					Count:      2,
				},
				{
					DomainName: "otherTest.com",
					Count:      10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDomainsMapToStruct(tt.domainMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDomainsMapToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "TestCorrectEmail",
			email: "test@gmail.com",
			want:  true,
		},
		{
			name:  "TestIncorrectEmail",
			email: "testgmail.com",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEmail(tt.email); got != tt.want {
				t.Errorf("isValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractDomain(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{
			name:  "Test_1",
			email: "someMail@wp.pl",
			want:  "wp.pl",
		},
		{
			name:  "Test_2",
			email: "testMail@go.com",
			want:  "go.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractDomain(tt.email); got != tt.want {
				t.Errorf("extractDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
