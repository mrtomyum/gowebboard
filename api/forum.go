package api

// ForumController is an interface for Forum
type ForumController interface {
	Create(*ForumCreateRequest) (*ForumCreateResponse, error)
	Update(*ForumUpdateRequest) (*ForumUpdateResponse, error)
	List(*ForumListRequest) (*ForumListResponse, error)
	Delete(*ForumDeleteRequest) (*ForumDeleteResponse, error)
	Get(int) (*Forum, error)
}

// ForumCreateRequest is a Create request for ForumController interface
type ForumCreateRequest struct {
	Name string
}
// ForumCreateResponse is a Create response for ForumController interface
type ForumCreateResponse struct {

}

// ForumUpdateRequest is a Update request for ForumController interface
type ForumUpdateRequest struct {

}

// ForumUpdateResponse is a Update response for ForumController interface
type ForumUpdateResponse struct {

}

// ForumListRequest is a List request for ForumController interface
type ForumListRequest struct {

}

// ForumListResponse is a List response for ForumController interface
type ForumListResponse struct {

}

// ForumDeleteRequest is a Delete request for ForumController interface
type ForumDeleteRequest struct {

}

// ForumDeleteResponse is a Delete response for ForumController interface
type ForumDeleteResponse struct {

}

// Forum is a forum struct
type Forum struct {
	Name string
}