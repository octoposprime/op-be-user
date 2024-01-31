package presentation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/user"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserFilter is a struct that represents the filter dto of a user.
type UserFilter struct {
	proto *pb.UserFilter
}

// NewUserFilter creates a new *UserFilter.
func NewUserFilter(pb *pb.UserFilter) *UserFilter {
	return &UserFilter{
		proto: pb,
	}
}

// String returns a string representation of the UserFilter.
func (s *UserFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
		"UserName: %v, "+
		"Email: %v, "+
		"UserType: %v, "+
		"UserStatus: %v, "+
		"Tags: %v, "+
		"FirstName: %v, "+
		"LastName: %v, "+
		"CreatedAtFrom: %v, "+
		"CreatedAtTo: %v, "+
		"UpdatedAtFrom: %v, "+
		"UpdatedAtTo: %v, "+
		"SearchText: %v, "+
		"SortType: %v, "+
		"SortField: %v, "+
		"Limit: %v, "+
		"Offset: %v",
		s.proto.Id,
		s.proto.Username,
		s.proto.Email,
		s.proto.UserType,
		s.proto.UserStatus,
		s.proto.Tags,
		s.proto.FirstName,
		s.proto.LastName,
		s.proto.CreatedAtFrom,
		s.proto.CreatedAtTo,
		s.proto.UpdatedAtFrom,
		s.proto.UpdatedAtTo,
		s.proto.SearchText,
		s.proto.SortType,
		s.proto.SortField,
		s.proto.Limit,
		s.proto.Offset)
}

// NewUserFilterFromEntity creates a new *UserFilter from entity.
func NewUserFilterFromEntity(entity *me.UserFilter) *UserFilter {
	id := entity.Id.String()
	userName := entity.UserName
	email := entity.Email
	userType := pb.UserType(entity.UserType)
	userStatus := pb.UserStatus(entity.UserStatus)
	tags := entity.Tags
	firstName := entity.FirstName
	lastName := entity.LastName
	createdAtFrom := timestamppb.New(entity.CreatedAtFrom)
	createdAtTo := timestamppb.New(entity.CreatedAtTo)
	updatedAtFrom := timestamppb.New(entity.UpdatedAtFrom)
	updatedAtTo := timestamppb.New(entity.UpdatedAtTo)
	searchText := entity.SearchText
	sortType := entity.SortType
	sortField := pb.UserSortField(entity.SortField)
	limit := int32(entity.Limit)
	offset := int32(entity.Offset)
	return &UserFilter{
		&pb.UserFilter{
			Id:            &id,
			Username:      &userName,
			Email:         &email,
			UserType:      &userType,
			UserStatus:    &userStatus,
			Tags:          tags,
			FirstName:     &firstName,
			LastName:      &lastName,
			CreatedAtFrom: createdAtFrom,
			CreatedAtTo:   createdAtTo,
			UpdatedAtFrom: updatedAtFrom,
			UpdatedAtTo:   updatedAtTo,
			SearchText:    &searchText,
			SortType:      &sortType,
			SortField:     &sortField,
			Limit:         &limit,
			Offset:        &offset,
		},
	}
}

// ToEntity returns a entity representation of the UserFilter.
func (s *UserFilter) ToEntity() *me.UserFilter {
	id := uuid.UUID{}
	if s.proto.Id != nil {
		id = tuuid.FromString(*s.proto.Id)
	}
	userName := ""
	if s.proto.Username != nil {
		userName = string(*s.proto.Username)
	}
	email := ""
	if s.proto.Email != nil {
		email = string(*s.proto.Email)
	}
	userType := 0
	if s.proto.UserType != nil {
		userType = int(*s.proto.UserType)
	}
	userStatus := 0
	if s.proto.UserStatus != nil {
		userStatus = int(*s.proto.UserStatus)
	}
	tags := []string{}
	if s.proto.Tags != nil {
		tags = s.proto.Tags
	}
	firstName := ""
	if s.proto.FirstName != nil {
		firstName = string(*s.proto.FirstName)
	}
	lastName := ""
	if s.proto.LastName != nil {
		lastName = string(*s.proto.LastName)
	}
	createdAtFrom := time.Time{}
	if s.proto.CreatedAtFrom != nil {
		createdAtFrom = s.proto.CreatedAtFrom.AsTime()
	}
	createdAtTo := time.Time{}
	if s.proto.CreatedAtTo != nil {
		createdAtTo = s.proto.CreatedAtTo.AsTime()
	}
	updatedAtFrom := time.Time{}
	if s.proto.UpdatedAtFrom != nil {
		updatedAtFrom = s.proto.UpdatedAtFrom.AsTime()
	}
	updatedAtTo := time.Time{}
	if s.proto.UpdatedAtTo != nil {
		updatedAtTo = s.proto.UpdatedAtTo.AsTime()
	}
	searchText := ""
	if s.proto.SearchText != nil {
		searchText = string(*s.proto.SearchText)
	}
	sortType := ""
	if s.proto.SortType != nil {
		sortType = string(*s.proto.SortType)
	}
	sortField := 0
	if s.proto.SortField != nil {
		sortField = int(*s.proto.SortField)
	}
	limit := 0
	if s.proto.Limit != nil {
		limit = int(*s.proto.Limit)
	}
	offset := 0
	if s.proto.Offset != nil {
		offset = int(*s.proto.Offset)
	}
	return &me.UserFilter{
		Id:            id,
		UserName:      userName,
		Email:         email,
		UserType:      mo.UserType(userType),
		UserStatus:    mo.UserStatus(userStatus),
		Tags:          tags,
		FirstName:     firstName,
		LastName:      lastName,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     mo.UserSortField(sortField),
		Limit:         limit,
		Offset:        offset,
	}
}
