package engine

import (
	"fmt"
	"log"
	"time"
	"math/rand"

	"github.com/jbattistella/special-interests/database"
)

const For = "for"
const Against = "against"

type AppResponse struct {
	Group string
	Cause string
}

func Engine() (AppResponse, string, error){
	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to DB")
	}

	DB.AutoMigrate(&database.QC_Prompts{})


	max:=7
	min:=1

	var groupPrompt database.QC_Prompts
	randomNumber := generateRandomNumber(min, max)
	_ = DB.Where("id = ?", randomNumber).Find(&groupPrompt)

	var causePrompt database.QC_Prompts
	randomNumber2 := generateRandomNumber(min, max)
    _ = DB.Where("id = ?", randomNumber2).Find(&causePrompt)

	var preposition string
	preposition = getPreposition(min, max)

	appRes := AppResponse{
        Group:       groupPrompt.Group,
        Cause:       causePrompt.Cause,
    }

	fmt.Println(groupPrompt.Group)

	return appRes,preposition, err
}

func generateRandomNumber(min, max int) int {
    // Seed the random number generator using the current time
    rand.Seed(time.Now().UnixNano())

    // Generate a random number in the range [min, max]
    return rand.Intn(max-min+1) + min
}

func getPreposition(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min+1) 
	
	if result%2 == 0 {
		return For
	} else {
		return Against
	}
}