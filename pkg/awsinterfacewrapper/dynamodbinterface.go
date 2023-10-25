package awsinterfacewrapper

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DynamoDbAwsInterface - interface to organize AWS SDK functions better for easier testing.
// This is updated as of Fri Jun 30, 2023.
type DynamoDbAwsInterface interface {
	BatchExecuteStatement(
		ctx context.Context,
		params *dynamodb.BatchExecuteStatementInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.BatchExecuteStatementOutput, error)
	BatchGetItem(
		ctx context.Context,
		params *dynamodb.BatchGetItemInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.BatchGetItemOutput, error)
	BatchWriteItem(
		ctx context.Context,
		params *dynamodb.BatchWriteItemInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.BatchWriteItemOutput, error)
	CreateBackup(
		ctx context.Context,
		params *dynamodb.CreateBackupInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.CreateBackupOutput, error)
	CreateGlobalTable(
		ctx context.Context,
		params *dynamodb.CreateGlobalTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.CreateGlobalTableOutput, error)
	CreateTable(
		ctx context.Context,
		params *dynamodb.CreateTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.CreateTableOutput, error)
	DeleteBackup(
		ctx context.Context,
		params *dynamodb.DeleteBackupInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DeleteBackupOutput, error)
	DeleteItem(
		ctx context.Context,
		params *dynamodb.DeleteItemInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DeleteItemOutput, error)
	DeleteTable(
		ctx context.Context,
		params *dynamodb.DeleteTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DeleteTableOutput, error)
	DescribeBackup(
		ctx context.Context,
		params *dynamodb.DescribeBackupInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeBackupOutput, error)
	DescribeContinuousBackups(
		ctx context.Context,
		params *dynamodb.DescribeContinuousBackupsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeContinuousBackupsOutput, error)
	DescribeContributorInsights(
		ctx context.Context,
		params *dynamodb.DescribeContributorInsightsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeContributorInsightsOutput, error)
	DescribeEndpoints(
		ctx context.Context,
		params *dynamodb.DescribeEndpointsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeEndpointsOutput, error)
	DescribeExport(
		ctx context.Context,
		params *dynamodb.DescribeExportInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeExportOutput, error)
	DescribeGlobalTable(
		ctx context.Context,
		params *dynamodb.DescribeGlobalTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeGlobalTableOutput, error)
	DescribeGlobalTableSettings(
		ctx context.Context,
		params *dynamodb.DescribeGlobalTableSettingsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	DescribeImport(
		ctx context.Context,
		params *dynamodb.DescribeImportInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeImportOutput, error)
	DescribeKinesisStreamingDestination(
		ctx context.Context,
		params *dynamodb.DescribeKinesisStreamingDestinationInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error)
	DescribeLimits(
		ctx context.Context,
		params *dynamodb.DescribeLimitsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeLimitsOutput, error)
	DescribeTable(
		ctx context.Context,
		params *dynamodb.DescribeTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeTableOutput, error)
	DescribeTableReplicaAutoScaling(
		ctx context.Context,
		params *dynamodb.DescribeTableReplicaAutoScalingInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeTimeToLive(
		ctx context.Context,
		params *dynamodb.DescribeTimeToLiveInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DescribeTimeToLiveOutput, error)
	DisableKinesisStreamingDestination(
		ctx context.Context,
		params *dynamodb.DisableKinesisStreamingDestinationInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.DisableKinesisStreamingDestinationOutput, error)
	EnableKinesisStreamingDestination(
		ctx context.Context,
		params *dynamodb.EnableKinesisStreamingDestinationInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.EnableKinesisStreamingDestinationOutput, error)
	ExecuteStatement(
		ctx context.Context,
		params *dynamodb.ExecuteStatementInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ExecuteStatementOutput, error)
	ExecuteTransaction(
		ctx context.Context,
		params *dynamodb.ExecuteTransactionInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ExecuteTransactionOutput, error)
	ExportTableToPointInTime(
		ctx context.Context,
		params *dynamodb.ExportTableToPointInTimeInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ExportTableToPointInTimeOutput, error)
	GetItem(
		ctx context.Context,
		params *dynamodb.GetItemInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.GetItemOutput, error)
	ImportTable(
		ctx context.Context,
		params *dynamodb.ImportTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ImportTableOutput, error)
	ListBackups(
		ctx context.Context,
		params *dynamodb.ListBackupsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ListBackupsOutput, error)
	ListContributorInsights(
		ctx context.Context,
		params *dynamodb.ListContributorInsightsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ListContributorInsightsOutput, error)
	ListExports(
		ctx context.Context,
		params *dynamodb.ListExportsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ListExportsOutput, error)
	ListGlobalTables(
		ctx context.Context,
		params *dynamodb.ListGlobalTablesInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ListGlobalTablesOutput, error)
	ListImports(
		ctx context.Context,
		params *dynamodb.ListImportsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ListImportsOutput, error)
	ListTables(
		ctx context.Context,
		params *dynamodb.ListTablesInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ListTablesOutput, error)
	ListTagsOfResource(
		ctx context.Context,
		params *dynamodb.ListTagsOfResourceInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ListTagsOfResourceOutput, error)
	PutItem(
		ctx context.Context,
		params *dynamodb.PutItemInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.PutItemOutput, error)
	Query(
		ctx context.Context,
		params *dynamodb.QueryInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.QueryOutput, error)
	RestoreTableFromBackup(
		ctx context.Context,
		params *dynamodb.RestoreTableFromBackupInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.RestoreTableFromBackupOutput, error)
	RestoreTableToPointInTime(
		ctx context.Context,
		params *dynamodb.RestoreTableToPointInTimeInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	Scan(
		ctx context.Context,
		params *dynamodb.ScanInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.ScanOutput, error)
	TagResource(
		ctx context.Context,
		params *dynamodb.TagResourceInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.TagResourceOutput, error)
	TransactGetItems(
		ctx context.Context,
		params *dynamodb.TransactGetItemsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.TransactGetItemsOutput, error)
	TransactWriteItems(
		ctx context.Context,
		params *dynamodb.TransactWriteItemsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.TransactWriteItemsOutput, error)
	UntagResource(
		ctx context.Context,
		params *dynamodb.UntagResourceInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UntagResourceOutput, error)
	UpdateContinuousBackups(
		ctx context.Context,
		params *dynamodb.UpdateContinuousBackupsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateContinuousBackupsOutput, error)
	UpdateContributorInsights(
		ctx context.Context,
		params *dynamodb.UpdateContributorInsightsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateContributorInsightsOutput, error)
	UpdateGlobalTable(
		ctx context.Context,
		params *dynamodb.UpdateGlobalTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateGlobalTableOutput, error)
	UpdateGlobalTableSettings(
		ctx context.Context,
		params *dynamodb.UpdateGlobalTableSettingsInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	UpdateItem(
		ctx context.Context,
		params *dynamodb.UpdateItemInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateItemOutput, error)
	UpdateTable(
		ctx context.Context,
		params *dynamodb.UpdateTableInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateTableOutput, error)
	UpdateTableReplicaAutoScaling(
		ctx context.Context,
		params *dynamodb.UpdateTableReplicaAutoScalingInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	UpdateTimeToLive(
		ctx context.Context,
		params *dynamodb.UpdateTimeToLiveInput,
		optFns ...func(*dynamodb.Options),
	) (*dynamodb.UpdateTimeToLiveOutput, error)
}
