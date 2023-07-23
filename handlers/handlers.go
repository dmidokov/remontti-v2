package handlers

import (
	"github.com/dmidokov/remontti-v2/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	_ "image/jpeg"
	"net/http"
)

// Переменная уровня пакета, используется для
// передачи конфигурации в функции обработчики

type Model struct {
	DB          *pgxpool.Pool
	Config      *config.Configuration
	CookieStore *sessions.CookieStore
	Logger      *logrus.Logger
}

type response struct {
	Status  string   `json:"status"`
	Errors  []string `json:"errors"`
	Message string   `json:"message" `
}

func New(db *pgxpool.Pool, cfg *config.Configuration, cookieStore *sessions.CookieStore, log *logrus.Logger) *Model {
	return &Model{
		DB:          db,
		Config:      cfg,
		CookieStore: cookieStore,
		Logger:      log,
	}
}

// Router Возвращает *mux.Router, c handlerFunction
// А также файловые сервер для выдачи статики css/js/jpg/...
func (hm *Model) Router(corsEnable bool) (*mux.Router, error) {

	router := mux.NewRouter()

	router.HandleFunc("/", hm.auth(handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist/", ""))).Methods(http.MethodGet)

	router.HandleFunc("/assets/{file}", handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", "")).Methods(http.MethodGet)

	router.HandleFunc("/login/", handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", "")).Methods(http.MethodGet)
	router.HandleFunc("/{folder}/", hm.auth(handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", ""))).Methods(http.MethodGet)

	router.HandleFunc("/{file}.{file}", handleFileServer(hm.Config.ROOT_PATH+"/web/vueui/remontti-ui/dist", "")).Methods(http.MethodGet)

	var methods []string
	methods = append(methods, http.MethodPost)
	methods = append(methods, http.MethodGet)
	if corsEnable {
		methods = append(methods, http.MethodOptions)
	}

	router.HandleFunc("/login", hm.loginPOST).Methods(http.MethodPost)
	router.HandleFunc("/logout", hm.logout).Methods(http.MethodGet)
	//TODO: подумать как иначе обобщить обращения к апи, чтобы не плодить миллион хендлеров можно подумать в сторону обработки
	// общих эндпоинтов типа
	// /api/v1/translations/
	// /api/v1/navigation
	// /api/v1/companies/
	// и внутри разбирать запросы или даже на уровне версии апи
	router.HandleFunc("/api/v1/translations/{pages}", hm.getTranslationsApi).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/navigation", hm.auth(hm.getNavigationApi)).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/companies", hm.auth(hm.getCompanies)).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/companies/{id}", hm.auth(hm.updateCompany)).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/companies/id/{id}", hm.auth(hm.getCompanyById)).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/companies/getCurrentCompanyName", hm.auth(hm.getCurrentCompanyName)).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/companies", hm.auth(hm.addCompany)).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/companies/delete", hm.auth(hm.deleteCompany)).Methods(http.MethodDelete)

	return router, nil
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
	http.Redirect(w, r, target, http.StatusMovedPermanently)
}

func handleFileServer(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodOptions {
			setCorsHeaders(&w, req)
			return
		}
		realHandler(w, req)
	}
}

func setCorsHeaders(w *http.ResponseWriter, _ *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
