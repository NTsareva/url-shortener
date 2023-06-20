package handlers

import "testing"

func TestHandlers_PostBodyHandler(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
		body        string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				statusCode:  201,
				contentType: "text/plain",
				body:        "XVlBzgba",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//дописать тело запроса
		})
	}
}
