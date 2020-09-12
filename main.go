package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	// REGION ...
	REGION = "ap-northeast-1"
)

func main() {
	sess := session.Must(session.NewSession())
	svc := ec2.New(
		sess,
		aws.NewConfig().WithRegion(REGION))

	// 引数を使ってフィルタ設定を作成
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					// aws.String("running"),
					aws.String("stopped"),
				},
			},
		},
	}

	res, err := svc.DescribeInstances(params)
	if err != nil {
		log.Fatalln(err)
	}
	// 初期化した配列用意してfor内のprintしている結果を入れる
	//instanceInfo := make([][]string, 2)
	resultInstances := make([][]string, 2)
	for i, r := range res.Reservations {
		for _, ins := range r.Instances {
			var tagName string
			for _, t := range ins.Tags {
				if *t.Key == "Name" {
					tagName = *t.Value
				}
			}
			instanceInfo := []string{tagName, *ins.InstanceId, *ins.InstanceType, *ins.State.Name}
			fmt.Println(instanceInfo)
			resultInstances[i] = append(resultInstances[i], instanceInfo...)

			/*
				fmt.Println(
					tagName, "\t\t\t",
					*ins.InstanceId, "\t",
					*ins.InstanceType, "\t",
					*ins.State.Name, "\t",
				)
			*/
		}
	}
	fmt.Println(resultInstances)
}
