package app

import (
	"flag"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/GalichAnton/go_final_project/internal/config"
	"github.com/GalichAnton/go_final_project/internal/handlers/done"
	"github.com/GalichAnton/go_final_project/internal/handlers/next_date"
	"github.com/GalichAnton/go_final_project/internal/handlers/sign"
	"github.com/GalichAnton/go_final_project/internal/handlers/tasks"
	"github.com/GalichAnton/go_final_project/internal/logger"
	"github.com/GalichAnton/go_final_project/internal/middleware/auth"
	logMiddlaware "github.com/GalichAnton/go_final_project/internal/middleware/log"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

const readHeaderTimeout = 5 * time.Second

var configPath string

const (
	taskPath     = "/api/task"
	tasksPath    = "/api/tasks"
	nextDatePath = "/api/nextdate"
	taskDonePath = "/api/task/done"
	singPath     = "/api/signin"
)

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig() error {
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer() error {
	webDir := "./web"
	fs := http.FileServer(http.Dir(webDir))
	logger.Init(logger.Core(logger.AtomicLevel(a.serviceProvider.LogConfig().Level())))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	pass := a.serviceProvider.pass

	mux.HandleFunc(nextDatePath, next_date.Handle)
	mux.HandleFunc(taskPath, auth.Auth(a.serviceProvider.TaskHandler(), pass))
	mux.HandleFunc(tasksPath, auth.Auth(tasks.New(a.serviceProvider.TaskService()).Handle, pass))
	mux.HandleFunc(taskDonePath, auth.Auth(done.New(a.serviceProvider.TaskService()).Handle, pass))
	mux.HandleFunc(singPath, sign.New(a.serviceProvider.TaskService()).Handle)

	muxWithLog := logMiddlaware.HttpLogInterceptor(mux)

	a.httpServer = &http.Server{
		Addr:              a.serviceProvider.HTTPConfig().Address(),
		Handler:           muxWithLog,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Info("HTTP server is running on ", zap.String("path", a.serviceProvider.HTTPConfig().Address()))
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
