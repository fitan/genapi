package log

import (
	"cmdb/pkg/trace"
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
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
			args: args{fs: []Option{WithTrace(trace.GetTr(), zapcore.DebugLevel)}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := NewXlog(tt.args.fs...)
				got.TraceLog(context.Background(), "log test").Info("log hello", zap.String("an", "bo"))
				got.Sync()
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
