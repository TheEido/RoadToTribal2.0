package api

import (
	"RoadToTribal2.0/internal/models"
	//"RoadToTribal2.0/internal/models"
	//transaction "RoadToTribal2.0/internal/repositories"
	"RoadToTribal2.0/internal/services/Transaction"
	"encoding/json"
	//"fmt"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	//"strconv"
)

type TransactionController struct {
	log            *zap.SugaredLogger
	validate       *validator.Validate
	transactionSvc *Transaction.DefaultTransactionService
}

func NewTransactionController(server *HTTPServer, logger *zap.SugaredLogger, v *validator.Validate, ts *Transaction.DefaultTransactionService) *TransactionController {
	b := &TransactionController{
		log:            logger,
		validate:       v,
		transactionSvc: ts,
	}

	// Load routes
	server.Router.Group(func(r chi.Router) {
		/*r.Use(JwtVerifyMiddleware(server.JwtService))*/
		r.Get("/transactions", b.handleGetAll)
		r.Get("/transactions/{id}", b.handleGetOne)
		r.Post("/transactions", b.handleCreate)
	})

	return b
}

func (b *TransactionController) handleGetAll(w http.ResponseWriter, r *http.Request) {
	result, ok := b.transactionSvc.FindAllTransactions(r.Context())

	if !ok {
		b.log.Errorf("There is no transactions")
		return
	}

	RenderJSON(r.Context(), w, http.StatusOK, result)
}

func (b *TransactionController) handleGetOne(w http.ResponseWriter, r *http.Request) {
	beneficiaryID := chi.URLParam(r, "id")

	if len(beneficiaryID) == 0 {
		b.log.Errorf("Beneficiary Id is required")
		return
	}

	result, ok := b.transactionSvc.FindTransactionDetails(r.Context(), beneficiaryID)

	if !ok {
		b.log.Errorf("Transaction is not exist")
		return
	}

	RenderJSON(r.Context(), w, http.StatusOK, result)
}

func (b *TransactionController) handleCreate(w http.ResponseWriter, r *http.Request) {
	var body *models.CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		b.log.Errorf("Malformed body: %s", err)
		return
	}

	err = b.validate.Struct(body)
	if err != nil {
		b.log.Errorf("Validation error: %s", err)
		return
	}
	// Call the service
	result, status := b.transactionSvc.AddTransaction(r.Context(), body)

	if status == false {
		b.log.Errorf("Add transaction failed. %v", err)
		return
	}

	RenderJSON(r.Context(), w, http.StatusOK, result)
}

/*func getTransaction(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	if ID == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	trans := &models.Transaction{}
	err, trans := transaction.GetTransactionById(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Contact with ID: %s not found", ID), 404)
		return
	}
	json.NewEncoder(w).Encode(trans)
	w.WriteHeader(200)
}*/

/*func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	err, transactions := transaction.GetAllTransaction()
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(200)
}

func addTransaction(w http.ResponseWriter, r *http.Request) {
	existingtransaction := &models.Transaction{}
	var trans models.Transaction
	json.NewDecoder(r.Body).Decode(&trans)
	trans.CreatedAt = "2022-07-19T13:02:01.440618Z"
	err, existingtransaction := transaction.GetTransactionById(strconv.Itoa(existingtransaction.ID))
	if err == nil {
		http.Error(w, fmt.Sprintf("Contact with ID: %s already exist", strconv.Itoa(trans.ID)), 400)
		return
	}
	err, _ = transaction.Addtransction(&trans)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Transaction created successfully"))
	w.WriteHeader(201)
}*/

/*func RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", getAllTransactions) //GET /transactions
		r.Get("/{ID}", getTransaction) //GET /transactions/0147344454
		r.Post("/", addTransaction)    //POST /transactions
		r.Put("/{ID}", updateTransaction)    //PUT /transactions/0147344454
		r.Delete("/{ID}", deleteTransaction) //DELETE /transactions/0147344454
	})
	return r
}*/
