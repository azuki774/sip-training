package digest

import "testing"

func TestWWWAuthenticate_ComputeResponse(t *testing.T) {
	type fields struct {
		AuthenticateScheme string
		UserName           string
		Realm              string
		NonceValue         string
		OpaqueValue        string
		Algorithm          string
		QOP                string
		URI                string
		NonceCount         string
		CNonse             string
		Response           string
	}
	type args struct {
		method   string
		password string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse string
	}{
		{
			name: "sample", // Ref. https://qiita.com/sgtrducky/items/38cd5aa47dec743b2358
			fields: fields{
				AuthenticateScheme: "digest",
				UserName:           "digest_user",
				Realm:              "digest",
				NonceValue:         "jeqPGvfFBQA=9e2e2d4c04ea2ad4758a1129484608c9e8abea85",
				OpaqueValue:        "",
				Algorithm:          "MD5",
				QOP:                "auth",
				URI:                "/digest/",
				NonceCount:         "00000001",
				CNonse:             "10057899296509bc",
				Response:           "",
			},
			args: args{
				method:   "GET",
				password: "password",
			},
			wantResponse: "e6af4f81fa87ccde2bc4769ee237eea2",
		},
		{
			name: "real_data1",
			fields: fields{
				AuthenticateScheme: "Digest",
				UserName:           "6002",
				Realm:              "asterisk",
				NonceValue:         "1717774189/360b1348b61842106c5bd646dbf2c9be",
				OpaqueValue:        "70210e76169d4618",
				Algorithm:          "MD5",
				QOP:                "auth",
				URI:                "sip:100.121.131.130;transport=UDP",
				NonceCount:         "00000001",
				CNonse:             "1b043ccbff85cf2b8cac03898ebb6267",
				Response:           "",
			},
			args: args{
				method:   "REGISTER",
				password: "unsecurepassword",
			},
			wantResponse: "f55e2f1f0b8c5d18abb597aa7d687a6b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WWWAuthenticate{
				AuthenticateScheme: tt.fields.AuthenticateScheme,
				UserName:           tt.fields.UserName,
				Realm:              tt.fields.Realm,
				NonceValue:         tt.fields.NonceValue,
				OpaqueValue:        tt.fields.OpaqueValue,
				Algorithm:          tt.fields.Algorithm,
				QOP:                tt.fields.QOP,
				URI:                tt.fields.URI,
				NonceCount:         tt.fields.NonceCount,
				CNonse:             tt.fields.CNonse,
				Response:           tt.fields.Response,
			}
			if gotResponse := w.ComputeResponse(tt.args.method, tt.args.password); gotResponse != tt.wantResponse {
				t.Errorf("WWWAuthenticate.ComputeResponse() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
