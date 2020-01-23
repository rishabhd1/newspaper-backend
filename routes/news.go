package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"newspaper-backend/config"
	"newspaper-backend/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveNewsPayload : Payload containing user's email and news' mongoID
type SaveNewsPayload struct {
	Email   string `json:"email" bson:"email"`
	MongoID string `json:"mongoID" bson:"mongoID"`
}

// GetHosts : Returns all Host names
func GetHosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var finalResponse models.FinalResponse

	field, e := collection.Distinct(ctx, "host", bson.D{{}})
	if e != nil {
		log.Println(e.Error())
		return
	}

	finalResponse.Status = "success"
	finalResponse.Body = field

	json.NewEncoder(w).Encode(finalResponse)
}

// AllNews : Returns all news
func AllNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var allNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, bson.M{})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allNews = append(allNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = allNews

	json.NewEncoder(w).Encode(finalResponse)
	return
}

// BusinessNews : Returns all business news
func BusinessNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var allBusinessNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, models.News{Category: "business"})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allBusinessNews = append(allBusinessNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = allBusinessNews

	json.NewEncoder(w).Encode(finalResponse)
	return
}

// SportsNews : Returns all sports news
func SportsNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var allSportsNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, models.News{Category: "sports"})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allSportsNews = append(allSportsNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = allSportsNews

	json.NewEncoder(w).Encode(finalResponse)
	return
}

// EntertainmentNews : Returns all entertainment news
func EntertainmentNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var allEntertainmentNews []models.News
	var finalResponse models.FinalResponse

	cursor, e := collection.Find(ctx, models.News{Category: "entertainment"})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		allEntertainmentNews = append(allEntertainmentNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = allEntertainmentNews

	json.NewEncoder(w).Encode(finalResponse)
	return
}

// ClickCount : Increase the clickCount by 1
func ClickCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	newsID := vars["news_id"]
	newsObjectID, e := primitive.ObjectIDFromHex(newsID)
	if e != nil {
		log.Println(e)
		return
	}

	filter := bson.M{"_id": bson.M{"$eq": newsObjectID}}
	update := bson.M{"$inc": bson.M{"clickCount": 1}}
	result, e := collection.UpdateOne(ctx, filter, update)
	if e != nil {
		log.Println(e)
		return
	}

	var finalResponse models.FinalResponse

	finalResponse.Status = "success"
	finalResponse.Body = result

	json.NewEncoder(w).Encode(finalResponse)
	return
}

// MostViewedNews : Returns the news with most clickCount (Descending to Ascending)
func MostViewedNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	collection := config.Client.Database("newspaper").Collection("news")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var mostViewedNews []models.News
	var finalResponse models.FinalResponse

	opts := options.Find()
	opts.SetSort(bson.D{{"clickCount", -1}})
	cursor, e := collection.Find(ctx, bson.D{}, opts)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + e.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var news models.News
		cursor.Decode(&news)
		mostViewedNews = append(mostViewedNews, news)
	}

	finalResponse.Status = "success"
	finalResponse.Body = mostViewedNews

	json.NewEncoder(w).Encode(finalResponse)
	return
}

// SaveNews : Save the news in user's profile
func SaveNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json/")

	var payload SaveNewsPayload
	var finalResponse models.FinalResponse

	e := json.NewDecoder(r.Body).Decode(&payload)
	if e != nil {
		log.Println(e.Error())
		return
	}

	mongoID, e := primitive.ObjectIDFromHex(payload.MongoID)
	if e != nil {
		log.Println(e.Error())
		return
	}

	collection := config.Client.Database("newspaper").Collection("profile")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": payload.Email}
	update := bson.M{"$push": bson.M{"news": mongoID}}
	result, e := collection.UpdateOne(ctx, filter, update)
	if e != nil {
		log.Println(e.Error())
		return
	}

	finalResponse.Status = "success"
	finalResponse.Body = result

	json.NewEncoder(w).Encode(finalResponse)
}
