package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"todo/repo"
	"todo/service"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Error string `json:"error"`
}

type authRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type updateUserRequest struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

type updateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Deadline    *string `json:"deadline"`
	Status      *bool   `json:"status"`
}

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/register", registerHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", loginHandler).Methods(http.MethodPost)
	router.HandleFunc("/users/{userID:[0-9]+}", updateUser).Methods(http.MethodPatch)
	router.HandleFunc("/users/{userID:[0-9]+}", deleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/users/{userID:[0-9]+}/tasks", listTasks).Methods(http.MethodGet)
	router.HandleFunc("/users/{userID:[0-9]+}/tasks", createTaskHTTP).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{userID:[0-9]+}/{taskID:[0-9]+}", getTask).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{userID:[0-9]+}/{taskID:[0-9]+}", updateTaskHTTP).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{userID:[0-9]+}/{taskID:[0-9]+}", deleteTask).Methods(http.MethodDelete)
	return router
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.Login == "" || req.Password == "" || req.Name == "" {
		writeError(w, http.StatusBadRequest, fmt.Errorf("login, password and name are required"))
		return
	}
	if err := service.Register(req.Login, req.Password, req.Name); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	user, err := service.ReadUser(req.Login)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, sanitizeUser(user))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req authRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	user, err := service.Login(req.Login, req.Password)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err)
		return
	}

	writeJSON(w, http.StatusOK, sanitizeUser(user))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := pathInt(r, "userID")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	user, err := findUserByID(userID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	var req updateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if req.Login != nil && *req.Login != "" && *req.Login != user.Login {
		if _, err := service.ReadUser(*req.Login); err == nil {
			writeError(w, http.StatusBadRequest, fmt.Errorf("user with this login already exists"))
			return
		}
		service.ChangeLogin(&user, *req.Login)
	}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Password != nil {
		user.Password = *req.Password
	}

	service.UpdateUser(&user)
	writeJSON(w, http.StatusOK, sanitizeUser(user))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := pathInt(r, "userID")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	user, err := findUserByID(userID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	service.DeleteUser(&user)
	writeJSON(w, http.StatusOK, map[string]string{"message": "user deleted"})
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	userID, err := pathInt(r, "userID")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if _, err := findUserByID(userID); err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	writeJSON(w, http.StatusOK, service.ListTasks(userID))
}

func createTaskHTTP(w http.ResponseWriter, r *http.Request) {
	userID, err := pathInt(r, "userID")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if _, err := findUserByID(userID); err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	var req createTaskRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusBadRequest, fmt.Errorf("title is required"))
		return
	}

	writeJSON(w, http.StatusCreated, service.CreateTask(req.Title, req.Description, req.Deadline, userID))
}

func getTask(w http.ResponseWriter, r *http.Request) {
	userID, taskID, err := taskPathIDs(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task, err := service.ReadTask(userID, taskID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func updateTaskHTTP(w http.ResponseWriter, r *http.Request) {
	userID, taskID, err := taskPathIDs(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task, err := service.ReadTask(userID, taskID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	var req updateTaskRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Deadline != nil {
		task.Deadline = *req.Deadline
	}
	if req.Status != nil {
		task.Status = *req.Status
	}

	repo.SaveTask(task)
	writeJSON(w, http.StatusOK, task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	userID, taskID, err := taskPathIDs(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task, err := service.ReadTask(userID, taskID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	service.DeleteTask(&task)
	writeJSON(w, http.StatusOK, map[string]string{"message": "task deleted"})
}

func parseID(value, kind string) (int, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s id", kind)
	}
	return id, nil
}

func pathInt(r *http.Request, key string) (int, error) {
	return parseID(mux.Vars(r)[key], key)
}

func taskPathIDs(r *http.Request) (int, int, error) {
	userID, err := pathInt(r, "userID")
	if err != nil {
		return 0, 0, err
	}

	taskID, err := pathInt(r, "taskID")
	if err != nil {
		return 0, 0, err
	}

	return userID, taskID, nil
}

func decodeJSON(r *http.Request, dst any) error {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, errorResponse{Error: err.Error()})
}

func findUserByID(userID int) (repo.User, error) {
	for _, user := range repo.Users {
		if user.ID == userID {
			return user, nil
		}
	}
	return repo.User{}, fmt.Errorf("user not found")
}

func sanitizeUser(user repo.User) map[string]any {
	return map[string]any{
		"id":    user.ID,
		"login": user.Login,
		"name":  user.Name,
	}
}
