package domain

//
//import (
//	"awesomeProject/common"
//	errs "awesomeProject/errors"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//)
//
//type MentorRole struct {
//	ID                 primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty`
//	Domain             string               `bson:"domain"`
//	Subdomain          string               `bson:"subdomain"`
//	AnnouncementLetter string               `bson:"announcement_letter"`
//	AboutMe            string               `bson:"about_me"`
//	Courses            []primitive.ObjectID `bson:"courses"`
//	MenteeCount        common.AccountType   `bson:"mentee_count"`
//	AverageRating      string               `bson:"average_rating"`
//	CreatedAt          string               `bson:"password"`
//	User               primitive.ObjectID   `bson:"user"`
//}
//
//type MentorRoleRequest struct {
//	Domain             string             `json:"domain"`
//	Subdomain          string             `json:"subdomain"`
//	AnnouncementLetter string             `json:"announcement_letter"`
//	AboutMe            string             `json:"about_me"`
//	User               primitive.ObjectID `json:"user"`
//}
//
//type MentorRoleResponse struct {
//	Domain             string               `json:"domain"`
//	Subdomain          string               `json:"subdomain"`
//	AnnouncementLetter string               `json:"announcement_letter"`
//	AboutMe            string               `json:"about_me"`
//	Courses            []primitive.ObjectID `json:"courses"`
//	MenteeCount        common.AccountType   `json:"mentee_count"`
//	AverageRating      string               `json:"average_rating"`
//	CreatedAt          string               `json:"password"`
//	User               User                 `json:"user"`
//}
//
//func (mr MentorRole) ToResponse() MentorRoleResponse {
//	return MentorRoleResponse{
//		Domain:             mr.Domain,
//		Subdomain:          mr.Subdomain,
//		AnnouncementLetter: mr.AnnouncementLetter,
//		AboutMe:            mr.AboutMe,
//		Courses:            mr.Courses,
//		MenteeCount:        mr.MenteeCount,
//		AverageRating:      mr.AverageRating,
//		CreatedAt:          mr.CreatedAt,
//		User:               User,
//	}
//}
//
//type MentorRoleRepository interface {
//	Insert(MentorRole) (*MentorRole, *errs.AppError)
//}
