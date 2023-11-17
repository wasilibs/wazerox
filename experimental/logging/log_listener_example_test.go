package logging_test

import (
	"context"
	_ "embed"
	"log"
	"os"

	wazero "github.com/wasilibs/wazerox"
	"github.com/wasilibs/wazerox/experimental"
	"github.com/wasilibs/wazerox/experimental/logging"
	"github.com/wasilibs/wazerox/imports/wasi_snapshot_preview1"
)

// listenerWasm was generated by the following:
//
//	cd testdata; wat2wasm --debug-names listener.wat
//
//go:embed testdata/listener.wasm
var listenerWasm []byte

// This is a very basic integration of listener. The main goal is to show how
// it is configured.
func Example_newHostLoggingListenerFactory() {
	// Set context to one that has an experimental listener that logs all host functions.
	ctx := context.WithValue(context.Background(), experimental.FunctionListenerFactoryKey{},
		logging.NewHostLoggingListenerFactory(os.Stdout, logging.LogScopeAll))

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.InstantiateWithConfig(ctx, listenerWasm,
		wazero.NewModuleConfig().WithStdout(os.Stdout))
	if err != nil {
		log.Panicln(err)
	}

	_, err = mod.ExportedFunction("rand").Call(ctx, 4)
	if err != nil {
		log.Panicln(err)
	}

	// We should see the same function called twice: directly and indirectly.

	// Output:
	//--> listener.rand(len=4)
	//	==> wasi_snapshot_preview1.random_get(buf=4,buf_len=4)
	//	<== errno=ESUCCESS
	//	==> wasi_snapshot_preview1.random_get(buf=8,buf_len=4)
	//	<== errno=ESUCCESS
	//<--
}

// This example shows how to see all function calls, including between host
// functions.
func Example_newLoggingListenerFactory() {
	// Set context to one that has an experimental listener
	ctx := context.WithValue(context.Background(), experimental.FunctionListenerFactoryKey{}, logging.NewLoggingListenerFactory(os.Stdout))

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.InstantiateWithConfig(ctx, listenerWasm,
		wazero.NewModuleConfig().WithStdout(os.Stdout))
	if err != nil {
		log.Panicln(err)
	}

	_, err = mod.ExportedFunction("rand").Call(ctx, 4)
	if err != nil {
		log.Panicln(err)
	}

	// We should see the same function called twice: directly and indirectly.

	// Output:
	//--> listener.rand(len=4)
	//	--> listener.wasi_rand(len=4)
	//		==> wasi_snapshot_preview1.random_get(buf=4,buf_len=4)
	//		<== errno=ESUCCESS
	//		==> wasi_snapshot_preview1.random_get(buf=8,buf_len=4)
	//		<== errno=ESUCCESS
	//	<--
	//<--
}
