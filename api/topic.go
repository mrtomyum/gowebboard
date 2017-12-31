package api

// TopicController is an interface for Topic
type TopicController interface {
	Create(*TopicCreateRequest) (*TopicCreateResponse, error)
	Update(*TopicUpdateRequest) (*TopicUpdateResponse, error)
	List(*TopicListRequest) (*TopicListResponse, error)
	Delete(*TopicDeleteRequest) (*TopicDeleteResponse, error)
	Get(int) (*User, error)
}

// TopicCreateRequest is create request for Topic Controller
type TopicCreateRequest struct {
	Name string
	Body  string
}

// TopicCreateResponse is create response for Topic Controller
type TopicCreateResponse struct {
	TopicID int64
}

// TopicUpdateRequest is Update request for Topic Controller
type TopicUpdateRequest struct {
	TopicID int64
	Name   string
	Body    string
}

// TopicUpdateResponse is Update response for Topic Controller
type TopicUpdateResponse struct {

}

// TopicListRequest is List request for Topic Controller
type TopicListRequest struct {
	Key string
}

// TopicListResponse is List response for Topic Controller
type TopicListResponse struct {

}

// TopicDeleteRequest is Delete request for Topic Controller
type TopicDeleteRequest struct {
	TopicID int64
}

// TopicDeleteResponse is Delete response for Topic Controller
type TopicDeleteResponse struct {

}

// Topic is topic struct
type Topic struct {
	Name string
}

// Create method
func (*Topic) Create(*TopicCreateRequest) (*TopicCreateResponse, error) {
	return nil,nil
}
