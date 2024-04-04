package handler

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/ko44d/go-api-sample/entity"
	"github.com/ko44d/go-api-sample/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddTask(t *testing.T) {
	t.Parallel()
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/addtask/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/addtask/ok_rsp.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/addtask/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/addtask/bad_req_rsp.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(testutil.LoadFile(t, tt.reqFile)))
			moq := &AddTaskServiceMock{}
			moq.AddTaskFunc = func(ctx context.Context, title string) (*entity.Task, error) {
				if tt.want.status == http.StatusOK {
					return &entity.Task{ID: 1}, nil
				}
				return nil, errors.New("error from mock")
			}
			sut := AddTask{
				Service:   moq,
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			rsp := w.Result()
			testutil.AssertResponse(t, rsp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
		})
	}
}
