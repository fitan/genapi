package log

import (
	"cmdb/pkg/trace"
	"fmt"
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
				tl := got.TraceLog(trace.GetTrCxt(), "log test", trace.GetTp())
				tl.Info("log hello", zap.String("an", "bo"))
				err = tl.Sync()
				if err != nil {
					fmt.Println(err)
				}
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
