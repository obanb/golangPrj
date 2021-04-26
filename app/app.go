package app

import (
	"awesomeProject/domain"
	"awesomeProject/logger"
	"awesomeProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	//router := mux.NewRouter()

	gin := gin.Default()


	mySqlClient := getMySqlClient("root:root@tcp(localhost:3306)/banking")
	mongoClient := getMongoClient("mongodb://localhost:27017")

	// local db/service repositories
	//customerRepositoryDb := domain.NewCustomerRepositoryDb(mySqlClient)
	//accountRepositoryDb := domain.NewAccountRepositoryDb(mySqlClient)
	issueRepositorySql := domain.NewIssueRepositorySql(mySqlClient, mongoClient)
	dbReportRepository := domain.NewDbReportRepository(mySqlClient, mongoClient)
	// remote db repositories

	// rest handlers
	//ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	//ah := AccountHandlers{service: service.NewAccountService(accountRepositoryDb)}
	ih := IssueHandlers{service: service.NewIssueService(issueRepositorySql)}
	dbh := DbReportHandlers{service: service.NewDbReportService(dbReportRepository)}
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//// rest routes
	//router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)
	//router.
	//	HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
	//	Methods(http.MethodPost).
	//	Name("NewTransaction")


	gin.GET("/issues", ih.getAllIssues)


	gin.POST("/issue", ih.CreateIssue)
	gin.POST("/issuemongo", ih.CreateIssueMongo)

	gin.POST("/issue/create_many", ih.CreateIssues)

	gin.POST("/dbreport", dbh.CreateDbReport)

	// starting server
	fmt.Print(mongoClient.Database("localhost").Collection("ucty").Name())
	logger.Info("Http server start..")
	//log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", "localhost", "8000"), router))
	log.Fatal(gin.Run())
}
