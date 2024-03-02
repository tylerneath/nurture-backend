package common

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	return zap.Must(zap.NewDevelopment())

}
