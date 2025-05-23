// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	"github.com/aws/smithy-go/tracing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson11_serializeOpGenerateDataSet struct {
}

func (*awsAwsjson11_serializeOpGenerateDataSet) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGenerateDataSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GenerateDataSetInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MarketplaceCommerceAnalytics20150701.GenerateDataSet")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGenerateDataSetInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpStartSupportDataExport struct {
}

func (*awsAwsjson11_serializeOpStartSupportDataExport) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpStartSupportDataExport) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	_, span := tracing.StartSpan(ctx, "OperationSerializer")
	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
	defer endTimer()
	defer span.End()
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartSupportDataExportInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MarketplaceCommerceAnalytics20150701.StartSupportDataExport")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentStartSupportDataExportInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	endTimer()
	span.End()
	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentCustomerDefinedValues(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentGenerateDataSetInput(v *GenerateDataSetInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CustomerDefinedValues != nil {
		ok := object.Key("customerDefinedValues")
		if err := awsAwsjson11_serializeDocumentCustomerDefinedValues(v.CustomerDefinedValues, ok); err != nil {
			return err
		}
	}

	if v.DataSetPublicationDate != nil {
		ok := object.Key("dataSetPublicationDate")
		ok.Double(smithytime.FormatEpochSeconds(*v.DataSetPublicationDate))
	}

	if len(v.DataSetType) > 0 {
		ok := object.Key("dataSetType")
		ok.String(string(v.DataSetType))
	}

	if v.DestinationS3BucketName != nil {
		ok := object.Key("destinationS3BucketName")
		ok.String(*v.DestinationS3BucketName)
	}

	if v.DestinationS3Prefix != nil {
		ok := object.Key("destinationS3Prefix")
		ok.String(*v.DestinationS3Prefix)
	}

	if v.RoleNameArn != nil {
		ok := object.Key("roleNameArn")
		ok.String(*v.RoleNameArn)
	}

	if v.SnsTopicArn != nil {
		ok := object.Key("snsTopicArn")
		ok.String(*v.SnsTopicArn)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentStartSupportDataExportInput(v *StartSupportDataExportInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CustomerDefinedValues != nil {
		ok := object.Key("customerDefinedValues")
		if err := awsAwsjson11_serializeDocumentCustomerDefinedValues(v.CustomerDefinedValues, ok); err != nil {
			return err
		}
	}

	if len(v.DataSetType) > 0 {
		ok := object.Key("dataSetType")
		ok.String(string(v.DataSetType))
	}

	if v.DestinationS3BucketName != nil {
		ok := object.Key("destinationS3BucketName")
		ok.String(*v.DestinationS3BucketName)
	}

	if v.DestinationS3Prefix != nil {
		ok := object.Key("destinationS3Prefix")
		ok.String(*v.DestinationS3Prefix)
	}

	if v.FromDate != nil {
		ok := object.Key("fromDate")
		ok.Double(smithytime.FormatEpochSeconds(*v.FromDate))
	}

	if v.RoleNameArn != nil {
		ok := object.Key("roleNameArn")
		ok.String(*v.RoleNameArn)
	}

	if v.SnsTopicArn != nil {
		ok := object.Key("snsTopicArn")
		ok.String(*v.SnsTopicArn)
	}

	return nil
}
