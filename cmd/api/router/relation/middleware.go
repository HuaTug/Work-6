// Code generated by hertz generator.

package relation

import (
	"HuaTug.com/cmd/api/router/authfunc"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followserviceMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followservicepageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationserviceMw() []app.HandlerFunc {
	// your code...
	return authfunc.Auth()
}

func _relationservicepageMw() []app.HandlerFunc {
	// your code...
	return authfunc.Auth()
}