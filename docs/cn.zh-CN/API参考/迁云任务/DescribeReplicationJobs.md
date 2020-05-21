# DescribeReplicationJobs {#doc_api_smc_DescribeReplicationJobs .reference}

查询一个或多个迁移任务的详细信息。

## 接口说明 {#description .section}

请求参数的作用类似于一个过滤器，过滤器为逻辑与（AND）关系。如果某一参数为空，则过滤器不起作用。

## 调试 {#api_explorer .section}

[您可以在OpenAPI Explorer中直接运行该接口，免去您计算签名的困扰。运行成功后，OpenAPI Explorer可以自动生成SDK代码示例。](https://api.aliyun.com/#product=smc&api=DescribeReplicationJobs&type=RPC&version=2019-06-01)

## 请求参数 {#parameters .section}

|名称|类型|是否必选|示例值|描述|
|--|--|----|---|--|
|Action|String|否|DescribeReplicationJobs|系统规定参数。取值：DescribeReplicationJobs<br/><br/> |
|BusinessStatus|String|否|Preparing|迁移任务的业务状态，取值范围：<br/> -   Preparing（准备中）<br/>-   Syncing（同步中）<br/>-   Processing（处理中）<br/>-   Cleaning（清除中）<br/><br/> |
|JobId.N|RepeatList|否|j-xxxxxxxxxxxxxxx|迁移任务ID列表。N的最大值：50。<br/> |
|Name|String|否|MyMigrationTask|迁移任务的名称。<br/>|
|PageNumber|Integer|否|1|返回结果中，迁移任务列表的页码。起始值：1。<br/>默认值：1。<br/>|
|PageSize|Integer|否|10|分页查询时设置的每页行数。最大值：50。<br/>默认值：10。<br/>|
|RegionId|String|否|cn-hangzhou|迁移源需迁入的目标阿里云地域ID。<br/>例如，您需要迁移源服务器至上海，则RegionId为`cn-shanghai`。您可以调用[DescribeRegions](~~25609~~)查看最新的阿里云地域列表。<br/>|
|SourceId.N|RepeatList|否|s-xxxxxxxxxxxxxxx|迁移源ID列表。N的最大值：50。<br/>|
|Status|String|否|Ready|迁移任务的主状态。取值范围：<br/>-   Ready（未开始）<br/>-   Running（运行中）<br/>-   Stopped（已暂停）<br/>-   InError（出错）<br/>-   Finished（已完成）<br/>-   Expired（已过期）<br/>-   Deleting（删除中）<br/>|

## 返回数据 {#resultMapping .section}

|名称|类型|示例值|描述|
|--|--|---|--|
|PageNumber|Integer|1|迁移任务列表的页码。<br/>|
|PageSize|Integer|10|每页行数。<br/>|
|ReplicationJobs| | |迁移任务详情集合。<br/>|
|BusinessStatus|String|Preparing|迁移任务的业务状态。<br/>|
|CreationTime|String|2014-07-24T13:00:52Z|迁移任务的创建时间。<br/>|
|DataDisks| | |目标阿里云服务器ECS的数据盘。<br/>|
|Index|Integer|1|数据盘顺序。<br/>|
|Size|Integer|40|数据盘大小，单位为GiB。<br/>|
|Description|String|This is my migration task.|迁移任务的描述。<br/>|
|EndTime|String|2019-06-04T16:00:52Z|迁移任务的完成时间。<br/>|
|ErrorCode|String|InternalError|迁移任务的错误码。<br/>|
|ImageId|String|xxxxxxxxxxxxxxxx|迁移任务交付的目标镜像ID。<br/>|
|ImageName|String|MyAliCloudImage|迁移任务交付的目标镜像名称。<br/>|
|InstanceId|String|i-xxxxxxxxxxxx|目标实例ID。<br/>|
|JobId|String|j-xxxxxxxxxxxx|迁移任务ID。<br/>|
|Name|String|MyMigrationTask|迁移任务名称。<br/>|
|NetMode|Integer|0|迁移时使用的网络类型。<br/>|
|Progress|Float|55.45|迁移任务总进度。<br/>|
|RegionId|String|cn-hangzhou|迁移源需迁入的目标阿里云地域ID。<br/>|
|ReplicationParameters|String|BandWidthLimit:0|复制驱动器参数信息。<br/>|
|ScheduledStartTime|String|2019-06-04T 13:35:00Z|迁移任务的执行时间。<br/>|
|SourceId|String|s-xxxxxxxxxxx|迁移源ID。<br/>|
|StartTime|String|2019-06-04T14:40:52Z|迁移任务的开始时间。<br/>|
|Status|String|Running|迁移任务的主状态。<br/>|
|StatusInfo|String|statusinfo|迁移状态的详细信息。<br/>|
|SystemDiskSize|Integer|40|目标阿里云服务器ECS的系统盘大小。<br/>|
|TargetType|String|image|迁移交付的目标类型。<br/>|
|TransitionInstanceId|String|instanceid2|迁移中转实例ID。<br/>|
|VSwitchId|String|xxxxxxxxxxxxxxx|指定VPC下的虚拟交换机ID。<br/>|
|ValidTime|String|2019-06-08T14:40:52Z|迁移任务的过期时间。<br/>|
|VpcId|String|xxxxxxxxxxxxxxx|已配置高速通道服务或者VPN网关的VPC ID。<br/>|
|RequestId|String|473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E|请求ID。<br/>|
|TotalCount|Integer|5|迁移任务总个数。<br/>|

## 示例 {#demo .section}

请求示例

``` {#request_demo}

http(s)://smc.aliyuncs.com/?Action=DescribeReplicationJobs
&<公共请求参数>

```

正常返回示例

`XML` 格式

``` {#xml_return_success_demo}
<DescribeReplicationJobsResponse>
	  <PageNumber>1</PageNumber>
	  <PageSize>10</PageSize>
	  <ReplicationJobs>
		    <ReplicationJob>
			      <BusinessStatus>Preparing</BusinessStatus>
			      <CreationTime>2014-07-24T13:00:52Z</CreationTime>
			      <DataDisks>
				        <DataDisk>
					          <Index>1</Index>
					          <Size>40</Size>
				        </DataDisk>
			      </DataDisks>
			      <Description>This is my migration task.</Description>
			      <EndTime>2019-06-04T16:00:52Z</EndTime>
			      <ErrorCode>InternalError</ErrorCode>
			      <ImageId>xxxxxxxxxxxxxxx</ImageId>
			      <ImageName>MyAilCloudImage</ImageName>
			      <InstanceId>i-xxxxxxxxxxxxxxx</InstanceId>
			      <JobId>j-xxxxxxxxxxxxxxx</JobId>
			      <Name>MyMiagrationTask</Name>
			      <NetMode>0</NetMode>
			      <Progress>55.45</Progress>
			      <RegionId>cn-hangzhou</RegionId>
			      <ReplicationParameters>BandWidthLimit:0</ReplicationParameters>
			      <ScheduledStartTime>2019-06-04T13:35:00Z</ScheduledStartTime>
			      <SourceId>s-xxxxxxxxxxxxxxx</SourceId>
			      <StartTime>2019-06-04T14:40:52Z</StartTime>
			      <Status>Running</Status>
			      <StatusInfo>statusinfo</StatusInfo>
			      <SystemDiskSize>40</SystemDiskSize>
			      <TargetType>image</TargetType>
			      <TransitionInstanceId>instanceid2</TransitionInstanceId>
			      <VSwitchId>xxxxxxxxxxxxxxx</VSwitchId>
			      <VaildTime>30</VaildTime>
			      <Vpcld>xxxxxxxxxxxxxxx</Vpcld>
		    </ReplicationJob>
	  </ReplicationJobs>
	  <RequestId>473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E</RequestId>
	  <TotalCount>5</TotalCount>
</DescribeReplicationJobsResponse>
```

`JSON` 格式

``` {#json_return_success_demo}
{
	"PageNumber":1,
	"TotalCount":"5",
	"PageSize":10,
	"RequestId":"473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E",
	"ReplicationJobs":{
		"ReplicationJob":[
			{
				"Vpcld":"xxxxxxxxxxxxxxx",
				"ImageId":"xxxxxxxxxxxxxxx",
				"InstanceId":"i-xxxxxxxxxxxxxxx",
				"VSwitchId":"xxxxxxxxxxxxxxx",
				"StatusInfo":"statusinfo",
				"ImageName":"MyAilCloudImage",
				"ScheduledStartTime":"2019-06-04T13:35:00Z",
				"BusinessStatus":"Preparing",
				"JobId":"j-xxxxxxxxxxxxxxx",
				"StartTime":"2019-06-04T14:40:52Z",
				"Description":"This is my migration task.",
				"SystemDiskSize":"40",
				"TargetType":"image",
				"VaildTime":"30",
				"DataDisks":{
					"DataDisk":[
						{
							"Index":"1",
							"Size":"40"
						}
					]
				},
				"ReplicationParameters":"BandWidthLimit:0",
				"NetMode":"0",
				"SourceId":"s-xxxxxxxxxxxxxxx",
				"Progress":"55.45",
				"Name":"MyMiagrationTask",
				"CreationTime":"2014-07-24T13:00:52Z",
				"Status":"Running",
				"RegionId":"cn-hangzhou",
				"TransitionInstanceId":"instanceid2",
				"EndTime":"2019-06-04T16:00:52Z",
				"ErrorCode":"InternalError"
			}
		]
	}
}
```

## 错误码 { .section}

访问[错误中心](https://error-center.aliyun.com/status/product/smc)查看更多错误码。

