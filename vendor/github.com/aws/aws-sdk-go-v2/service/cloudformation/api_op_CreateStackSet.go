// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a stack set.
func (c *Client) CreateStackSet(ctx context.Context, params *CreateStackSetInput, optFns ...func(*Options)) (*CreateStackSetOutput, error) {
	if params == nil {
		params = &CreateStackSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStackSet", params, optFns, c.addOperationCreateStackSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStackSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStackSetInput struct {

	// The name to associate with the stack set. The name must be unique in the Region
	// where you create your stack set.
	//
	// A stack name can contain only alphanumeric characters (case-sensitive) and
	// hyphens. It must start with an alphabetic character and can't be longer than 128
	// characters.
	//
	// This member is required.
	StackSetName *string

	// The Amazon Resource Name (ARN) of the IAM role to use to create this stack set.
	//
	// Specify an IAM role only if you are using customized administrator roles to
	// control which users or groups can manage specific stack sets within the same
	// administrator account. For more information, see [Prerequisites: Granting Permissions for Stack Set Operations]in the CloudFormation User
	// Guide.
	//
	// [Prerequisites: Granting Permissions for Stack Set Operations]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html
	AdministrationRoleARN *string

	// Describes whether StackSets automatically deploys to Organizations accounts
	// that are added to the target organization or organizational unit (OU). Specify
	// only if PermissionModel is SERVICE_MANAGED .
	AutoDeployment *types.AutoDeployment

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account.
	//
	// By default, SELF is specified. Use SELF for stack sets with self-managed
	// permissions.
	//
	//   - To create a stack set with service-managed permissions while signed in to
	//   the management account, specify SELF .
	//
	//   - To create a stack set with service-managed permissions while signed in to a
	//   delegated administrator account, specify DELEGATED_ADMIN .
	//
	// Your Amazon Web Services account must be registered as a delegated admin in the
	//   management account. For more information, see [Register a delegated administrator]in the CloudFormation User
	//   Guide.
	//
	// Stack sets with service-managed permissions are created in the management
	// account, including stack sets that are created by delegated administrators.
	//
	// [Register a delegated administrator]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html
	CallAs types.CallAs

	// In some cases, you must explicitly acknowledge that your stack set template
	// contains certain capabilities in order for CloudFormation to create the stack
	// set and related stack instances.
	//
	//   - CAPABILITY_IAM and CAPABILITY_NAMED_IAM
	//
	// Some stack templates might include resources that can affect permissions in
	//   your Amazon Web Services account; for example, by creating new Identity and
	//   Access Management (IAM) users. For those stack sets, you must explicitly
	//   acknowledge this by specifying one of these capabilities.
	//
	// The following IAM resources require you to specify either the CAPABILITY_IAM or
	//   CAPABILITY_NAMED_IAM capability.
	//
	//   - If you have IAM resources, you can specify either capability.
	//
	//   - If you have IAM resources with custom names, you must specify
	//   CAPABILITY_NAMED_IAM .
	//
	//   - If you don't specify either of these capabilities, CloudFormation returns
	//   an InsufficientCapabilities error.
	//
	// If your stack template contains these resources, we recommend that you review
	//   all permissions associated with them and edit their permissions if necessary.
	//
	// [AWS::IAM::AccessKey]
	//
	// [AWS::IAM::Group]
	//
	// [AWS::IAM::InstanceProfile]
	//
	// [AWS::IAM::Policy]
	//
	// [AWS::IAM::Role]
	//
	// [AWS::IAM::User]
	//
	// [AWS::IAM::UserToGroupAddition]
	//
	// For more information, see [Acknowledging IAM Resources in CloudFormation Templates].
	//
	//   - CAPABILITY_AUTO_EXPAND
	//
	// Some templates reference macros. If your stack set template references one or
	//   more macros, you must create the stack set directly from the processed template,
	//   without first reviewing the resulting changes in a change set. To create the
	//   stack set directly, you must acknowledge this capability. For more information,
	//   see [Using CloudFormation Macros to Perform Custom Processing on Templates].
	//
	// Stack sets with service-managed permissions don't currently support the use of
	//   macros in templates. (This includes the [AWS::Include]and [AWS::Serverless]transforms, which are macros
	//   hosted by CloudFormation.) Even if you specify this capability for a stack set
	//   with service-managed permissions, if you reference a macro in your template the
	//   stack set operation will fail.
	//
	// [AWS::IAM::AccessKey]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
	// [AWS::Include]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html
	// [AWS::IAM::User]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
	// [AWS::IAM::InstanceProfile]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
	// [AWS::IAM::Policy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
	// [AWS::IAM::Group]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
	// [AWS::IAM::UserToGroupAddition]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
	// [AWS::IAM::Role]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
	// [Using CloudFormation Macros to Perform Custom Processing on Templates]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html
	// [AWS::Serverless]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html
	// [Acknowledging IAM Resources in CloudFormation Templates]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities
	Capabilities []types.Capability

	// A unique identifier for this CreateStackSet request. Specify this token if you
	// plan to retry requests so that CloudFormation knows that you're not attempting
	// to create another stack set with the same name. You might retry CreateStackSet
	// requests to ensure that CloudFormation successfully received them.
	//
	// If you don't specify an operation ID, the SDK generates one automatically.
	ClientRequestToken *string

	// A description of the stack set. You can use the description to identify the
	// stack set's purpose or other important information.
	Description *string

	// The name of the IAM execution role to use to create the stack set. If you do
	// not specify an execution role, CloudFormation uses the
	// AWSCloudFormationStackSetExecutionRole role for the stack set operation.
	//
	// Specify an IAM role only if you are using customized execution roles to control
	// which stack resources users and groups can include in their stack sets.
	ExecutionRoleName *string

	// Describes whether StackSets performs non-conflicting operations concurrently
	// and queues conflicting operations.
	ManagedExecution *types.ManagedExecution

	// The input parameters for the stack set template.
	Parameters []types.Parameter

	// Describes how the IAM roles required for stack set operations are created. By
	// default, SELF-MANAGED is specified.
	//
	//   - With self-managed permissions, you must create the administrator and
	//   execution roles required to deploy to target accounts. For more information, see
	//   [Grant Self-Managed Stack Set Permissions].
	//
	//   - With service-managed permissions, StackSets automatically creates the IAM
	//   roles required to deploy to accounts managed by Organizations. For more
	//   information, see [Grant Service-Managed Stack Set Permissions].
	//
	// [Grant Self-Managed Stack Set Permissions]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html
	// [Grant Service-Managed Stack Set Permissions]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html
	PermissionModel types.PermissionModels

	// The stack ID you are importing into a new stack set. Specify the Amazon
	// Resource Name (ARN) of the stack.
	StackId *string

	// The key-value pairs to associate with this stack set and the stacks created
	// from it. CloudFormation also propagates these tags to supported resources that
	// are created in the stacks. A maximum number of 50 tags can be specified.
	//
	// If you specify tags as part of a CreateStackSet action, CloudFormation checks
	// to see if you have the required IAM permission to tag resources. If you don't,
	// the entire CreateStackSet action fails with an access denied error, and the
	// stack set is not created.
	Tags []types.Tag

	// The structure that contains the template body, with a minimum length of 1 byte
	// and a maximum length of 51,200 bytes. For more information, see [Template Anatomy]in the
	// CloudFormation User Guide.
	//
	// Conditional: You must specify either the TemplateBody or the TemplateURL
	// parameter, but not both.
	//
	// [Template Anatomy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
	TemplateBody *string

	// The location of the file that contains the template body. The URL must point to
	// a template (maximum size: 460,800 bytes) that's located in an Amazon S3 bucket
	// or a Systems Manager document. For more information, see [Template Anatomy]in the CloudFormation
	// User Guide.
	//
	// Conditional: You must specify either the TemplateBody or the TemplateURL
	// parameter, but not both.
	//
	// [Template Anatomy]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
	TemplateURL *string

	noSmithyDocumentSerde
}

type CreateStackSetOutput struct {

	// The ID of the stack set that you're creating.
	StackSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStackSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateStackSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateStackSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateStackSet"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateStackSetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateStackSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStackSet(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

type idempotencyToken_initializeOpCreateStackSet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateStackSet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateStackSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateStackSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateStackSetInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateStackSetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateStackSet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateStackSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateStackSet",
	}
}