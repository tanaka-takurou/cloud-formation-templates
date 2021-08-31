# AWS CloudFormation Template dlq_for_lambda


## Dependence
- aws-lambda-go


## Requirements
- AWS (Lambda, SNS)
- aws-sam-cli
- golang environment


## Usage

### Edit Target email address
- open template.yml
- edit 'your.email@sample.com'


### Deploy
```bash
make clean build
AWS_PROFILE={profile} AWS_DEFAULT_REGION={region} make bucket={bucket} stack={stack name} deploy
```


### Check
```bash
aws lambda invoke --function-name ServerlessDLQLambdaFunction --invocation-type Event response.json
```
