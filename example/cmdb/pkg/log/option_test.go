package log

import (
	"cmdb/pkg/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
	"time"
)

func TestNewXlog(t *testing.T) {
	type args struct {
		fs []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Xlog
		wantErr bool
	}{
		{
			name:    "echo",
			args:    args{fs: []Option{WithTrace(zapcore.DebugLevel)}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := NewXlog(tt.args.fs...)
				got.TraceLog(trace.GetTr(), "log test").Info("log hello", zap.String("an", "bo"))
				time.Sleep(3 * time.Second)

				if (err != nil) != tt.wantErr {
					t.Errorf("NewXlog() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewXlog() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
