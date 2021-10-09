package appointy

import (
	"context"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var err error
var lock sync.Mutex

// Schema for User Data Type(Document)
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"Name" bson:"Name"`
	Email    string             `json:"Email" bson:"Email"`
	Password string             `json:"Password" bson:"Password"`
}

// Schema for Post Data Type(Document)
type Post struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string              `json:"Caption" bson:"Caption"`
	ImageURL  string              `json:"ImageURL" bson:"ImageURL"`
	Timestamp primitive.Timestamp `json:"Timestamp,omitempty" bson:"Timestamp,omitempty"`
	UserID    primitive.ObjectID  `json:"UserID" bson:"UserID"`
}

// Default Route Handler
func handler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Fprintf(w, "HTTP route Working!")
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
}

// Handles POST Requests to create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	if r.Method != "POST" {
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check Input Validity
	if u.Email == "" || u.Name == "" || u.Password == "" {
		http.Error(w, "Parameters missing in JSON", http.StatusBadRequest)
		return
	}
	Users := client.Database("Appointy").Collection("Users")
	// To hash the password
	HashedPassword := sha512.New()
	HashedPassword.Write([]byte(u.Password))
	u.Password = string(HashedPassword.Sum(nil))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User: %+v", u)
	fmt.Printf("\n\n User: %+v", u)
	result, err := Users.InsertOne(
		context.Background(),
		u)
	if err != nil {
		fmt.Printf("\n\n Database Insertion Error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result != nil {
		return
	}
}

// Handles GET Requests to fetch user using Id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	if r.Method != "GET" {
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}
	UserID := r.URL.Query().Get("id")
	if UserID == "" {
		fmt.Printf("\n\n User ID not mentioned")
		http.Error(w, "UserID not Mentioned", http.StatusBadRequest)
		return
	}
	fmt.Printf(UserID)
	Users := client.Database("Appointy").Collection("Users")
	var RequestedUser bson.M
	UserOID, err := primitive.ObjectIDFromHex(UserID)
	err = Users.FindOne(context.Background(), bson.D{{"_id", UserOID}}).Decode(&RequestedUser)
	if err != nil {
		fmt.Printf("\n\n User not found")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(RequestedUser)
	json.NewEncoder(w).Encode(RequestedUser)
}

// Handles POST Request to create new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf(" Creating Post")
	if r.Method != "POST" {
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
	var p Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check Input Validity
	if p.Caption == "" || p.ImageURL == "" {
		http.Error(w, "Parameters missing in JSON", http.StatusBadRequest)
		return
	}
	p.Timestamp = primitive.Timestamp{T: uint32(time.Now().Unix())}
	Posts := client.Database("Appointy").Collection("Posts")
	fmt.Fprintf(w, "Post: %+v", p)
	fmt.Printf("\n\n Post: %+v", p)
	result, err := Posts.InsertOne(
		context.Background(),
		p)
	if err != nil {
		fmt.Printf("\n\n Database Insertion Error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result != nil {
		return
	}
}

// Handles GET Request to retrieve post using Id
func GetPostById(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	if r.Method != "GET" {
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}
	PostID := r.URL.Query().Get("id")
	if PostID == "" {
		fmt.Printf("\n\n Post ID is Missing")
		http.Error(w, "Post ID not Mentioned", http.StatusBadRequest)
		return
	}
	Posts := client.Database("Appointy").Collection("Posts")
	var RequestedPost bson.M
	PostOID, err := primitive.ObjectIDFromHex(PostID)
	err = Posts.FindOne(context.Background(), bson.D{{"_id", PostOID}}).Decode(&RequestedPost)
	if err != nil {
		fmt.Printf("\n\n Post not found")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(RequestedPost)
	json.NewEncoder(w).Encode(RequestedPost)
}

// Handles GET Request to retrieve all posts by an user using User Id
func ListPostsByUser(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	UserID := r.URL.Query().Get("id")
	page := r.URL.Query().Get("page")
	Posts := client.Database("Appointy").Collection("Posts")
	UserOID, err := primitive.ObjectIDFromHex(UserID)
	// For Pagination
	if page != "" {
		currentpage, err := strconv.Atoi(page)
		posts, err := Posts.CountDocuments(context.Background(), bson.D{{"UserID", UserOID}})
		if err != nil {
			fmt.Printf("\n\n No Posts found")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pages := posts / 4
		if posts%4 != 0 {
			pages += 1
		}
		if currentpage > int(pages) {
			http.Error(w, "Invalid Page No.", http.StatusBadRequest)
		}
		options := options.Find()
		skip := (currentpage - 1) * 4
		options.SetSkip(int64(skip))
		options.SetLimit(4)
		cursor, err1 := Posts.Find(context.Background(), bson.D{{"UserID", UserOID}}, options)
		if err1 != nil {
			fmt.Printf("\n\n No Posts found")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer cursor.Close(context.Background())
		var post bson.M
		for cursor.Next(context.Background()) {
			if err1 = cursor.Decode(&post); err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(post)
		}
	} else { // No Pagination
		cursor, err1 := Posts.Find(context.Background(), bson.D{{"UserID", UserOID}})
		if err1 != nil {
			fmt.Printf("\n\n No Posts found")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer cursor.Close(context.Background())
		var post bson.M
		for cursor.Next(context.Background()) {
			if err1 = cursor.Decode(&post); err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(post)
		}
	}
}

func RunServer() {
	// Connect to MongoDB Atlas
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Appointy:MuUzw9t9YeqyahZ@appointy.gs1hd.mongodb.net/Appointy?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	// HTTP routes and their handlers are defined below
	http.HandleFunc("/users/", GetUserById)
	http.HandleFunc("/users", CreateUser)
	http.HandleFunc("/posts/", GetPostById)
	http.HandleFunc("/posts", CreatePost)
	http.HandleFunc("/posts/users/", ListPostsByUser)
	http.HandleFunc("/", handler)
	//Create HTTP Server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
