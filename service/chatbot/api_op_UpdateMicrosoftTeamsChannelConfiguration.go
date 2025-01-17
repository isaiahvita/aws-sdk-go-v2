// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chatbot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates MS Teams Channel Configuration
func (c *Client) UpdateMicrosoftTeamsChannelConfiguration(ctx context.Context, params *UpdateMicrosoftTeamsChannelConfigurationInput, optFns ...func(*Options)) (*UpdateMicrosoftTeamsChannelConfigurationOutput, error) {
	if params == nil {
		params = &UpdateMicrosoftTeamsChannelConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMicrosoftTeamsChannelConfiguration", params, optFns, c.addOperationUpdateMicrosoftTeamsChannelConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMicrosoftTeamsChannelConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMicrosoftTeamsChannelConfigurationInput struct {

	// The ID of the Microsoft Teams channel.
	//
	// This member is required.
	ChannelId *string

	// The ARN of the MicrosoftTeamsChannelConfiguration to update.
	//
	// This member is required.
	ChatConfigurationArn *string

	// The name of the Microsoft Teams channel.
	ChannelName *string

	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS
	// managed 'AdministratorAccess' policy is applied by default if this is not set.
	GuardrailPolicyArns []string

	// The ARN of the IAM role that defines the permissions for AWS Chatbot. This is a
	// user-defined role that AWS Chatbot will assume. This is not the service-linked
	// role. For more information, see IAM Policies for AWS Chatbot.
	IamRoleArn *string

	// Logging levels include ERROR, INFO, or NONE.
	LoggingLevel *string

	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot.
	SnsTopicArns []string

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired *bool

	noSmithyDocumentSerde
}

type UpdateMicrosoftTeamsChannelConfigurationOutput struct {

	// The configuration for a Microsoft Teams channel configured with AWS Chatbot.
	ChannelConfiguration *types.TeamsChannelConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMicrosoftTeamsChannelConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMicrosoftTeamsChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMicrosoftTeamsChannelConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMicrosoftTeamsChannelConfiguration"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateMicrosoftTeamsChannelConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMicrosoftTeamsChannelConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateMicrosoftTeamsChannelConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMicrosoftTeamsChannelConfiguration",
	}
}
