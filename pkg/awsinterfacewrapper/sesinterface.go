package awsinterfacewrapper

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type AwsSesV2ClientInterfaceWrapper interface {
	BatchGetMetricData(
		ctx context.Context,
		params *sesv2.BatchGetMetricDataInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.BatchGetMetricDataOutput, error)
	CreateConfigurationSet(
		ctx context.Context,
		params *sesv2.CreateConfigurationSetInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateConfigurationSetOutput, error)
	CreateConfigurationSetEventDestination(
		ctx context.Context,
		params *sesv2.CreateConfigurationSetEventDestinationInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateConfigurationSetEventDestinationOutput, error)
	CreateContact(
		ctx context.Context,
		params *sesv2.CreateContactInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateContactOutput, error)
	CreateContactList(
		ctx context.Context,
		params *sesv2.CreateContactListInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateContactListOutput, error)
	CreateCustomVerificationEmailTemplate(
		ctx context.Context,
		params *sesv2.CreateCustomVerificationEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateCustomVerificationEmailTemplateOutput, error)
	CreateDedicatedIpPool(
		ctx context.Context,
		params *sesv2.CreateDedicatedIpPoolInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateDedicatedIpPoolOutput, error)
	CreateDeliverabilityTestReport(
		ctx context.Context,
		params *sesv2.CreateDeliverabilityTestReportInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateDeliverabilityTestReportOutput, error)
	CreateEmailIdentity(
		ctx context.Context,
		params *sesv2.CreateEmailIdentityInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateEmailIdentityOutput, error)
	CreateEmailIdentityPolicy(
		ctx context.Context,
		params *sesv2.CreateEmailIdentityPolicyInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateEmailIdentityPolicyOutput, error)
	CreateEmailTemplate(
		ctx context.Context,
		params *sesv2.CreateEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateEmailTemplateOutput, error)
	CreateImportJob(
		ctx context.Context,
		params *sesv2.CreateImportJobInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.CreateImportJobOutput, error)
	DeleteConfigurationSet(
		ctx context.Context,
		params *sesv2.DeleteConfigurationSetInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetEventDestination(
		ctx context.Context,
		params *sesv2.DeleteConfigurationSetEventDestinationInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteConfigurationSetEventDestinationOutput, error)
	DeleteContact(
		ctx context.Context,
		params *sesv2.DeleteContactInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteContactOutput, error)
	DeleteContactList(
		ctx context.Context,
		params *sesv2.DeleteContactListInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteContactListOutput, error)
	DeleteCustomVerificationEmailTemplate(
		ctx context.Context,
		params *sesv2.DeleteCustomVerificationEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteCustomVerificationEmailTemplateOutput, error)
	DeleteDedicatedIpPool(
		ctx context.Context,
		params *sesv2.DeleteDedicatedIpPoolInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteDedicatedIpPoolOutput, error)
	DeleteEmailIdentity(
		ctx context.Context,
		params *sesv2.DeleteEmailIdentityInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteEmailIdentityOutput, error)
	DeleteEmailIdentityPolicy(
		ctx context.Context,
		params *sesv2.DeleteEmailIdentityPolicyInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteEmailIdentityPolicyOutput, error)
	DeleteEmailTemplate(
		ctx context.Context,
		params *sesv2.DeleteEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteEmailTemplateOutput, error)
	DeleteSuppressedDestination(
		ctx context.Context,
		params *sesv2.DeleteSuppressedDestinationInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.DeleteSuppressedDestinationOutput, error)
	GetAccount(
		ctx context.Context,
		params *sesv2.GetAccountInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetAccountOutput, error)
	GetBlacklistReports(
		ctx context.Context,
		params *sesv2.GetBlacklistReportsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetBlacklistReportsOutput, error)
	GetConfigurationSet(
		ctx context.Context,
		params *sesv2.GetConfigurationSetInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetConfigurationSetOutput, error)
	GetConfigurationSetEventDestinations(
		ctx context.Context,
		params *sesv2.GetConfigurationSetEventDestinationsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetConfigurationSetEventDestinationsOutput, error)
	GetContact(
		ctx context.Context,
		params *sesv2.GetContactInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetContactOutput, error)
	GetContactList(
		ctx context.Context,
		params *sesv2.GetContactListInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetContactListOutput, error)
	GetCustomVerificationEmailTemplate(
		ctx context.Context,
		params *sesv2.GetCustomVerificationEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetCustomVerificationEmailTemplateOutput, error)
	GetDedicatedIp(
		ctx context.Context,
		params *sesv2.GetDedicatedIpInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetDedicatedIpOutput, error)
	GetDedicatedIpPool(
		ctx context.Context,
		params *sesv2.GetDedicatedIpPoolInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetDedicatedIpPoolOutput, error)
	GetDedicatedIps(
		ctx context.Context,
		params *sesv2.GetDedicatedIpsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetDedicatedIpsOutput, error)
	GetDeliverabilityDashboardOptions(
		ctx context.Context,
		params *sesv2.GetDeliverabilityDashboardOptionsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error)
	GetDeliverabilityTestReport(
		ctx context.Context,
		params *sesv2.GetDeliverabilityTestReportInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetDeliverabilityTestReportOutput, error)
	GetDomainDeliverabilityCampaign(
		ctx context.Context,
		params *sesv2.GetDomainDeliverabilityCampaignInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetDomainDeliverabilityCampaignOutput, error)
	GetDomainStatisticsReport(
		ctx context.Context,
		params *sesv2.GetDomainStatisticsReportInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetDomainStatisticsReportOutput, error)
	GetEmailIdentity(
		ctx context.Context,
		params *sesv2.GetEmailIdentityInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetEmailIdentityOutput, error)
	GetEmailIdentityPolicies(
		ctx context.Context,
		params *sesv2.GetEmailIdentityPoliciesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetEmailIdentityPoliciesOutput, error)
	GetEmailTemplate(
		ctx context.Context,
		params *sesv2.GetEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetEmailTemplateOutput, error)
	GetImportJob(
		ctx context.Context,
		params *sesv2.GetImportJobInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetImportJobOutput, error)
	GetSuppressedDestination(
		ctx context.Context,
		params *sesv2.GetSuppressedDestinationInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.GetSuppressedDestinationOutput, error)
	ListConfigurationSets(
		ctx context.Context,
		params *sesv2.ListConfigurationSetsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListConfigurationSetsOutput, error)
	ListContactLists(
		ctx context.Context,
		params *sesv2.ListContactListsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListContactListsOutput, error)
	ListContacts(
		ctx context.Context,
		params *sesv2.ListContactsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListContactsOutput, error)
	ListCustomVerificationEmailTemplates(
		ctx context.Context,
		params *sesv2.ListCustomVerificationEmailTemplatesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error)
	ListDedicatedIpPools(
		ctx context.Context,
		params *sesv2.ListDedicatedIpPoolsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListDedicatedIpPoolsOutput, error)
	ListDeliverabilityTestReports(
		ctx context.Context,
		params *sesv2.ListDeliverabilityTestReportsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListDeliverabilityTestReportsOutput, error)
	ListDomainDeliverabilityCampaigns(
		ctx context.Context,
		params *sesv2.ListDomainDeliverabilityCampaignsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error)
	ListEmailIdentities(
		ctx context.Context,
		params *sesv2.ListEmailIdentitiesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListEmailIdentitiesOutput, error)
	ListEmailTemplates(
		ctx context.Context,
		params *sesv2.ListEmailTemplatesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListEmailTemplatesOutput, error)
	ListImportJobs(
		ctx context.Context,
		params *sesv2.ListImportJobsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListImportJobsOutput, error)
	ListRecommendations(
		ctx context.Context,
		params *sesv2.ListRecommendationsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListRecommendationsOutput, error)
	ListSuppressedDestinations(
		ctx context.Context,
		params *sesv2.ListSuppressedDestinationsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListSuppressedDestinationsOutput, error)
	ListTagsForResource(
		ctx context.Context,
		params *sesv2.ListTagsForResourceInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.ListTagsForResourceOutput, error)
	PutAccountDedicatedIpWarmupAttributes(
		ctx context.Context,
		params *sesv2.PutAccountDedicatedIpWarmupAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutAccountDedicatedIpWarmupAttributesOutput, error)
	PutAccountDetails(
		ctx context.Context,
		params *sesv2.PutAccountDetailsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutAccountDetailsOutput, error)
	PutAccountSendingAttributes(
		ctx context.Context,
		params *sesv2.PutAccountSendingAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutAccountSendingAttributesOutput, error)
	PutAccountSuppressionAttributes(
		ctx context.Context,
		params *sesv2.PutAccountSuppressionAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutAccountSuppressionAttributesOutput, error)
	PutAccountVdmAttributes(
		ctx context.Context,
		params *sesv2.PutAccountVdmAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutAccountVdmAttributesOutput, error)
	PutConfigurationSetDeliveryOptions(
		ctx context.Context,
		params *sesv2.PutConfigurationSetDeliveryOptionsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutConfigurationSetDeliveryOptionsOutput, error)
	PutConfigurationSetReputationOptions(
		ctx context.Context,
		params *sesv2.PutConfigurationSetReputationOptionsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutConfigurationSetReputationOptionsOutput, error)
	PutConfigurationSetSendingOptions(
		ctx context.Context,
		params *sesv2.PutConfigurationSetSendingOptionsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutConfigurationSetSendingOptionsOutput, error)
	PutConfigurationSetSuppressionOptions(
		ctx context.Context,
		params *sesv2.PutConfigurationSetSuppressionOptionsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutConfigurationSetSuppressionOptionsOutput, error)
	PutConfigurationSetTrackingOptions(
		ctx context.Context,
		params *sesv2.PutConfigurationSetTrackingOptionsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutConfigurationSetTrackingOptionsOutput, error)
	PutConfigurationSetVdmOptions(
		ctx context.Context,
		params *sesv2.PutConfigurationSetVdmOptionsInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutConfigurationSetVdmOptionsOutput, error)
	PutDedicatedIpInPool(
		ctx context.Context,
		params *sesv2.PutDedicatedIpInPoolInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutDedicatedIpInPoolOutput, error)
	PutDedicatedIpPoolScalingAttributes(
		ctx context.Context,
		params *sesv2.PutDedicatedIpPoolScalingAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutDedicatedIpPoolScalingAttributesOutput, error)
	PutDedicatedIpWarmupAttributes(
		ctx context.Context,
		params *sesv2.PutDedicatedIpWarmupAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutDedicatedIpWarmupAttributesOutput, error)
	PutDeliverabilityDashboardOption(
		ctx context.Context,
		params *sesv2.PutDeliverabilityDashboardOptionInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutDeliverabilityDashboardOptionOutput, error)
	PutEmailIdentityConfigurationSetAttributes(
		ctx context.Context,
		params *sesv2.PutEmailIdentityConfigurationSetAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutEmailIdentityConfigurationSetAttributesOutput, error)
	PutEmailIdentityDkimAttributes(
		ctx context.Context,
		params *sesv2.PutEmailIdentityDkimAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutEmailIdentityDkimAttributesOutput, error)
	PutEmailIdentityDkimSigningAttributes(
		ctx context.Context,
		params *sesv2.PutEmailIdentityDkimSigningAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutEmailIdentityDkimSigningAttributesOutput, error)
	PutEmailIdentityFeedbackAttributes(
		ctx context.Context,
		params *sesv2.PutEmailIdentityFeedbackAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutEmailIdentityFeedbackAttributesOutput, error)
	PutEmailIdentityMailFromAttributes(
		ctx context.Context,
		params *sesv2.PutEmailIdentityMailFromAttributesInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutEmailIdentityMailFromAttributesOutput, error)
	PutSuppressedDestination(
		ctx context.Context,
		params *sesv2.PutSuppressedDestinationInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.PutSuppressedDestinationOutput, error)
	SendBulkEmail(
		ctx context.Context,
		params *sesv2.SendBulkEmailInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.SendBulkEmailOutput, error)
	SendCustomVerificationEmail(
		ctx context.Context,
		params *sesv2.SendCustomVerificationEmailInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.SendCustomVerificationEmailOutput, error)
	SendEmail(
		ctx context.Context,
		params *sesv2.SendEmailInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.SendEmailOutput, error)
	TagResource(
		ctx context.Context,
		params *sesv2.TagResourceInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.TagResourceOutput, error)
	TestRenderEmailTemplate(
		ctx context.Context,
		params *sesv2.TestRenderEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.TestRenderEmailTemplateOutput, error)
	UntagResource(
		ctx context.Context,
		params *sesv2.UntagResourceInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.UntagResourceOutput, error)
	UpdateConfigurationSetEventDestination(
		ctx context.Context,
		params *sesv2.UpdateConfigurationSetEventDestinationInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.UpdateConfigurationSetEventDestinationOutput, error)
	UpdateContact(
		ctx context.Context,
		params *sesv2.UpdateContactInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.UpdateContactOutput, error)
	UpdateContactList(
		ctx context.Context,
		params *sesv2.UpdateContactListInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.UpdateContactListOutput, error)
	UpdateCustomVerificationEmailTemplate(
		ctx context.Context,
		params *sesv2.UpdateCustomVerificationEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.UpdateCustomVerificationEmailTemplateOutput, error)
	UpdateEmailIdentityPolicy(
		ctx context.Context,
		params *sesv2.UpdateEmailIdentityPolicyInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.UpdateEmailIdentityPolicyOutput, error)
	UpdateEmailTemplate(
		ctx context.Context,
		params *sesv2.UpdateEmailTemplateInput,
		optFns ...func(*sesv2.Options),
	) (*sesv2.UpdateEmailTemplateOutput, error)
}
