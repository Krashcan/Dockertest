package resources

import(
	"encoding/json"
	"net/http"

	"github.com/krashcan/dockertest/model"
	"github.com/krashcan/dockertest/util"
)

func ReverseString(w http.ResponseWriter,r *http.Request){
	input := &model.InputBody{}
	err := json.NewDecoder(r.Body).Decode(input)

	if util.ErrorExists(err, http.StatusBadRequest, w) {
		return
	}

	output := &model.OutputBody{util.Reverse(input.Input)}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(output)

	if util.ErrorExists(err, http.StatusInternalServerError, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
}