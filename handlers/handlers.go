package handlers

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/http"

	"github.com/dmidokov/remontti-v2/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
)

// Переменная уровня пакета, используется для
// передачи конфигурации в функции обработчики

type HandlersModel struct {
	DB          *pgx.Conn
	Config      *config.Configuration
	CookieStore *sessions.CookieStore
}

type response struct {
	Status  string   `json:"status"`
	Errors  []string `json:"errors"`
	Message string   `json:"message" `
}

func New(db *pgx.Conn, cfg *config.Configuration, cookieStore *sessions.CookieStore) *HandlersModel {
	return &HandlersModel{
		DB:          db,
		Config:      cfg,
		CookieStore: cookieStore,
	}
}

// Возвращает *mux.Router, c handlerFunction
// А также файловые сервер для выдачи статики css/js/jpg/...
func (hm *HandlersModel) Router(corsEnable bool) (*mux.Router, error) {

	router := mux.NewRouter()

	router.HandleFunc("/", hm.auth1(handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist/", ""))).Methods(http.MethodGet)

	router.HandleFunc("/assets/{file}", handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", "")).Methods(http.MethodGet)

	router.HandleFunc("/{folder}/", hm.auth1(handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", ""))).Methods(http.MethodGet)

	router.HandleFunc("/{file}.{file}", handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", "")).Methods(http.MethodGet)

	router.HandleFunc("/login/", handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", "")).Methods(http.MethodGet)

	var methods []string
	methods = append(methods, http.MethodPost)
	methods = append(methods, http.MethodGet)
	if corsEnable {
		methods = append(methods, http.MethodOptions)
	}

	router.HandleFunc("/login", hm.loginPOST).Methods(methods...)
	router.HandleFunc("/logout", hm.logout).Methods(methods...)
	router.HandleFunc("/api/v1/translations/get", hm.getTranslationsApi).Methods(methods...)

	return router, nil
}

func settings(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Request was processed settings")
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
	http.Redirect(w, r, target, http.StatusMovedPermanently)
}

func handleFileServer(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		realHandler(w, req)
	}
}

func setCorsHeaders(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
