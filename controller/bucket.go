package controller

import (
	"encoding/json"
	"net/http"
	"s3manager/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Server) BucketsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resp, err := s.getAllBuckets()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		respJSON, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
		return
	}
}

func (s *Server) getAllBuckets() ([]model.Bucket, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		return nil, err
	}
	input := &s3.ListBucketsInput{}
	svc := s3.New(sess)
	result, err := svc.ListBuckets(input)
	if err != nil {
		return nil, err
	}
	output := []model.Bucket{}
	for _, b := range result.Buckets {
		bckt := model.Bucket{
			Name: b.Name,
		}
		output = append(output, bckt)
	}
	return output, nil
}
