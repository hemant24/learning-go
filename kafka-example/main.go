package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	hostPort = "3000"
	topicName = "kafka-example-topic"
)

func main() {
	// Seed for fake skill score
	rand.Seed(time.Now().Unix())

	ds := &mockDataStore{}
	addFakeData(ds)


	producer, err := createKafkaProducer("127.0.0.1:9092")

	if err != nil {
		log.Fatal("Failed to connect to Kafka")
	}

	//Ensures that the topic has been created in kafka
	producer.Input() <- &sarama.ProducerMessage{
		Key: sarama.StringEncoder("init"),
		Topic: topicName,
		Timestamp: time.Now(),
	}

	log.Println("Creating Topic...")
	time.Sleep(1 * time.Second)
	/*
		This quick sleep allows the async producer some processing time to fire off the init message above to
		Kafka, since Go concurrency works on scheduling not threads.  Hmm... that may be another good topic for
		a blog post.
	*/


	go func(){
		consumeMessages("127.0.0.1:2181", msgHandler(ds))
	}()
	
	http.Handle("/api/skills", httpHandlers(ds, producer.Input()))
	log.Println("Listening on", hostPort)
	http.ListenAndServe(fmt.Sprintf(":%s", hostPort), nil)
}

func addFakeData(ds *mockDataStore) {
	user1 := make(map[string]skillScore)
	user1["Golang"] = skillScore{
		SkillName: "Golang",
		Score: 100,
		LastScored: time.Now().Add(-24 * time.Hour),
	}


	user1["Kafka"] = skillScore{
		SkillName: "Kafka",
		Score: 100,
		LastScored: time.Now().Add(-24 * time.Hour),
	}

	ds.data = make(map[string]map[string]skillScore)
	ds.data["user1"] = user1

}

func httpHandlers(ds *mockDataStore, c chan <- *sarama.ProducerMessage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getHandler(ds, w, r)
		} else if r.Method == http.MethodPost {
			postHandler(c, w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}


func postHandler(c chan <- *sarama.ProducerMessage, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	reqBody := &requestBody{}
	e := decoder.Decode(reqBody)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	for _, skill := range reqBody.Skills {
		processSkill(c, reqBody.ID, skill)
	}

	w.WriteHeader(http.StatusAccepted)
}

func processSkill(c chan <- *sarama.ProducerMessage, userID string, skillName string) {
	body := skillScoreMessage{
		SkillName: skillName,
		ProfileID: userID,
	}

	bodyBytes, _ := json.Marshal(body)

	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Key: sarama.StringEncoder(userID),
		Timestamp: time.Now(),
		Value: sarama.ByteEncoder(bodyBytes),
	}

	c <- msg
}

func getHandler(ds *mockDataStore, w http.ResponseWriter, r * http.Request) {
	userId := r.URL.Query().Get("userID")
	skill := r.URL.Query().Get("skill")

	if userId == "" || skill == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userSkills, ok := ds.ReadData(userId, skill)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		encoder.Encode(userSkills)
	}
}

func msgHandler(ds *mockDataStore) func(m *sarama.ConsumerMessage) error {
	return func(m *sarama.ConsumerMessage) error {
		// Empty body means it is an init message
		if len(m.Value) == 0 {
			return nil
		}

		skillMsg := &skillScoreMessage{}
		e := json.Unmarshal(m.Value, skillMsg)

		if e != nil {
			return e
		}

		//Simulate processing time
		time.Sleep(1 * time.Second)
		log.Printf("Adding skill %s to user %s", skillMsg.SkillName, skillMsg.ProfileID)

		score := skillScore{
			SkillName:  skillMsg.SkillName,
			Score:      rand.Float32() * 100,
			LastScored: time.Now(),
		}

		ds.WriteData(skillMsg.ProfileID, score)
		return nil
	}
}