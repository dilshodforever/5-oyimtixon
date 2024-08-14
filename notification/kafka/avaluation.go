package kafka

import (
	"encoding/json"
	"log"

	pb "github.com/dilshodforever/5-oyimtixon/model"
	"github.com/dilshodforever/5-oyimtixon/service"
)

func StartLevel(rootService *service.NotificationService) func(message []byte) {
	return func(message []byte) {
		var app pb.Send
		if err := json.Unmarshal(message, &app); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}
		err := rootService.CreateNotification(app)
		if err != nil {
			log.Printf("Cannot create evaluation via Kafka: %v", err)
			return
		}
		log.Print("Created evaluation")
	}
}

// func EvaluationUpdateHandler(evalService *service.EvaluationService) func(message []byte) {
// 	return func(message []byte) {
// 		var eval pb.EvaluationCreate
// 		if err := json.Unmarshal(message, &eval); err != nil {
// 			log.Printf("Cannot unmarshal JSON: %v", err)
// 			return
// 		}

// 		respEval, err := evalService.(context.Background(), &eval)
// 		if err != nil {
// 			log.Printf("Cannot create evaluation via Kafka: %v", err)
// 			return
// 		}
// 		log.Printf("Created evaluation: %+v", respEval)
// 	}
// }
