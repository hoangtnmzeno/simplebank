package api

// import (
// 	"net/http/httptest"
// 	mockdb "simplebank/db/mock"
// 	db "simplebank/db/sqlc"
// 	"simplebank/util"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// )

// func TestTransferAPI(t *testing.T) {
// 	amount := int64(10)

// 	user1, _ := randomUser(t)
// 	user2, _ := randomUser(t)
// 	user3, _ := randomUser(t)

// 	account1 := randomAccount(user1.Username)
// 	account2 := randomAccount(user2.Username)
// 	account3 := randomAccount(user3.Username)

// 	account1.Currency = util.USD
// 	account2.Currency = util.USD
// 	account3.Currency = util.EUR

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        util.USD,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(db.Account{})
// 			},
// 			checkResponse: func(recoder *httptest.ResponseRecorder) {},
// 		},
// 	}
// }
