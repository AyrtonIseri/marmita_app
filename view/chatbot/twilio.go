package whatsapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"marmita/types"
	"net/http"
	"net/http/httputil"
)

// aqui vamos formatar as mensagens para retornar para o twilio

func ReadIncomingMessage(w http.ResponseWriter, r *http.Request) error {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nNew request received. Information: \n %s \n", string(dump))

	// printing the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	// print raw body
	fmt.Println("\nRaw Body:", string(body))

	// Finally save it into a proper type
	var Response types.TwilioResponse

	err = json.NewDecoder(r.Body).Decode(&Response)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nFinal response WAID: ", Response.WppUser)
	fmt.Println("\nFinal response TwilioNumber: ", Response.TwilioWpp)
	fmt.Println("\nFinal response Body: ", Response.Body)

	return nil
}
