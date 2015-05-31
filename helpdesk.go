package helpdesk

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/fkse/helpdesk/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

// Globals
var db gorm.DB
var log *logrus.Logger

func Run(c *Configuration) {
	// Init logger
	log = logrus.New()
	log.Info("Initializing server..")
	// init database
	initDatabase(&c.Database)
	// get router
	router := NewRouter()
	// Start http server
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), router)
	// shutdown
}

func initDatabase(c *ConfigDatabase) {
	// make connection
	db, err := gorm.Open(c.Driver, c.ConnectionString())
	if err != nil {
		panic(err)
	}
	// Migration
	db.AutoMigrate(&model.User{})
}

// Create and populate router
func NewRouter() (r *mux.Router) {
	r = mux.NewRouter().StrictSlash(true)
	// init routes
	for _, route := range AppRoutes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(RequestLogger(RequestSetup(route.HandlerFunc), route.Name))
	}

	return r
}

// Request setupY
func RequestSetup(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set json content type
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// Server request
		h.ServeHTTP(w, r)
	})
}

// Request logging middleware
func RequestLogger(h http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Request time tracking
		start := time.Now()
		h.ServeHTTP(w, r)
		// Log request
		log.WithFields(logrus.Fields{
			"route": name,
			"time":  time.Since(start),
		}).Infof("%s\t%s", r.Method, r.RequestURI)
	})
}
