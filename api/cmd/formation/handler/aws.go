func Credentials(req *Request) *credentials.Credentials {
	if req != nil {
		if access, ok := req.ResourceProperties["AccessId"].(string); ok && access != "" {
			if secret, ok := req.ResourceProperties["SecretAccessKey"].(string); ok && secret != "" {
				return credentials.NewStaticCredentials(access, secret, "")
			}
		}
	}
	if os.Getenv("AWS_ACCESS") != "" {
		return credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS"), os.Getenv("AWS_SECRET"), "")
	}


	// return credentials.NewCredentials(&credentials.EC2RoleProvider{})
	return credentials.NewEnvCredentials()
}

func Region(req *Request) *string {
	if req != nil {
		if region, ok := req.ResourceProperties["Region"].(string); ok && region != "" {
			return aws.String(region)
		}
	}

	return aws.String(os.Getenv("AWS_REGION"))
}

func CloudFormation(req Request) *cloudformation.CloudFormation {
	return cloudformation.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func Lambda(req Request) *lambda.Lambda {
	return lambda.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func EC2(req Request) *ec2.EC2 {
	return ec2.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func ECR(req Request) *ecr.ECR {
	return ecr.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func ECS(req Request) *ecs.ECS {
	return ecs.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func KMS(req Request) *kms.KMS {
	return kms.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func S3(req Request) *s3.S3 {
	return s3.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func SNS(req Request) *sns.SNS {
	return sns.New(session.New(), &aws.Config{
		Credentials: Credentials(&req),
		Region:      Region(&req),
	})
}

func SQS() *sqs.SQS {
	return sqs.New(session.New(), &aws.Config{
		Credentials: Credentials(nil),
		Region:      Region(nil),
	})
}
